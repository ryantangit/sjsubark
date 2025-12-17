import os
from sqlalchemy import String, create_engine, URL, select
from sqlalchemy.orm import DeclarativeBase, Mapped, Session, mapped_column
from datetime import datetime
from dotenv import load_dotenv

load_dotenv()

url_object = URL.create(
    "postgresql+psycopg",
    username=os.getenv("SJSUBARK_PSQL_USER"),
    password=os.getenv("SJSUBARK_PSQL_PASSWORD"), 
    host=os.getenv("SJSUBARK_PSQL_HOST"),
    database=os.getenv("SJSUBARK_PSQL_DB"),
    port= int(os.getenv("SJSUBARK_PSQL_PORT", 5432)),
)

class Base(DeclarativeBase):
    pass

class GarageRecord(Base):
    __tablename__ = "garage_fullness"
    transaction_id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str] = mapped_column(String(30))
    utc_timestamp: Mapped[datetime]  
    second: Mapped[int]
    minute: Mapped[int]
    hour: Mapped[int]
    day: Mapped[int]
    month: Mapped[int]
    year: Mapped[int]
    weekday: Mapped[int]
    is_weekend: Mapped[bool]
    is_campus_closed: Mapped[bool]
    fullness: Mapped[int]
    
    def __repr__(self) -> str:
        return f"name={self.name}, utc_timestamp={self.utc_timestamp}, fullness={self.fullness}"

class SQLAccessor:
    def __init__(self):
        self.engine = create_engine(url_object)
    def most_recent_garage(self):
        stmt = select(GarageRecord).order_by(GarageRecord.utc_timestamp.desc()).limit(4)
        with Session(self.engine) as sess:
            return sess.execute(stmt).scalars().all()

