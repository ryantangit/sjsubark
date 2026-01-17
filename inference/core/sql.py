import os
from datetime import datetime
from pydantic import BaseModel
from sqlalchemy import ForeignKey, create_engine, URL, select
from sqlalchemy.orm import DeclarativeBase, Mapped, Session, mapped_column

from dotenv import load_dotenv

load_dotenv()


url_object = URL.create(
    "postgresql+psycopg",
    username=os.getenv("SJSUBARK_PSQL_USER"),
    password=os.getenv("SJSUBARK_PSQL_PASSWORD"),
    host=os.getenv("SJSUBARK_PSQL_HOST"),
    database=os.getenv("SJSUBARK_PSQL_DB"),
    port=int(os.getenv("SJSUBARK_PSQL_PORT", 5432)),
)


class Garage(BaseModel):
    name: str
    utc_timestamp: datetime
    fullness: int


class Base(DeclarativeBase):
    pass


class GarageFullness(Base):
    __tablename__ = "garage_fullness"
    transaction_id: Mapped[int] = mapped_column(primary_key=True)
    utc_timestamp: Mapped[datetime]
    garage_id: Mapped[int] = mapped_column(ForeignKey("garage_info.garage_id"))
    fullness: Mapped[int]

    def __repr__(self) -> str:
        return f"id={self.garage_id}, utc_timestamp={self.utc_timestamp}, fullness={self.fullness}"


class GarageInfo(Base):
    __tablename__ = "garage_info"
    garage_id: Mapped[int] = mapped_column(primary_key=True)
    garage_name: Mapped[str]
    address: Mapped[str]

    def __repr__(self) -> str:
        return f"id={self.garage_id}, name={self.garage_name}, address={self.address}"

class SQLAccessor:
    def __init__(self):
        self.engine = create_engine(url_object)
        self.NUM_OF_GARAGES = (
            4  # This most likely won't change while I'm here LoL - 12/17/26
        )

    def most_recent_garage_records(self) -> list[Garage]:
        stmt = (
            select(
                GarageInfo.garage_name,
                GarageFullness.utc_timestamp,
                GarageFullness.fullness,
            )
            .join(GarageInfo)
            .order_by(GarageFullness.utc_timestamp.desc())
            .limit(self.NUM_OF_GARAGES)
        )
        with Session(self.engine) as sess:
            results = sess.execute(stmt)
            garages = []
            for name, timestamp, fullness in results:
                garages.append(
                    Garage(name=name, utc_timestamp=timestamp, fullness=fullness)
                )
            return garages

    def most_recent_record(self, garage_id: int) -> Garage | None:
        stmt = (
            select(
                GarageInfo.garage_name,
                GarageFullness.utc_timestamp,
                GarageFullness.fullness,
            )
            .join(GarageInfo)
            .where(GarageInfo.garage_id == garage_id)
            .order_by(GarageFullness.utc_timestamp.desc())
            .limit(1)
            )

        with Session(self.engine) as sess:
            result = sess.execute(stmt).first()
            if result is None:
                return None
            return Garage(name=result.garage_name, utc_timestamp=result.utc_timestamp, fullness=result.fullness)

