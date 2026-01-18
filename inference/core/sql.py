import os
from datetime import date, datetime, timedelta
from sqlalchemy import ForeignKey, create_engine, URL, select
from sqlalchemy.orm import DeclarativeBase, Mapped, Session, mapped_column
from .models import Garage, CalendarDate

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


class Calendar(Base):
    __tablename__ = "calendar"
    calendar_date: Mapped[date] = mapped_column(primary_key=True)
    is_campus_closed: Mapped[bool]
    is_weekend: Mapped[bool]
    utc_start: Mapped[datetime]
    utc_end: Mapped[datetime]

    def __repr__(self) -> str:
        return f"date={self.calendar_date}, is_campus_closed={self.is_campus_closed}, is_weekend={self.is_weekend}"


class SQLAccessor:
    def __init__(self):
        self.engine = create_engine(url_object)
        # This most likely won't change while I'm here LoL - 12/17/26
        self.NUM_OF_GARAGES = 4
        # Sets how far the time window[timestamp, timestamp + PREDICTION_DAY_WINDOW] should be in table calendar
        self.PREDICTION_DAY_WINDOW = 1

    def most_recent_garage_records(self) -> list[Garage]:
        stmt = (
            select(
                GarageInfo.garage_name,
                GarageFullness.garage_id,
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
            for name, garage_id, timestamp, fullness in results:
                garages.append(
                    Garage(
                        garage_id=garage_id,
                        name=name,
                        utc_timestamp=timestamp,
                        fullness=fullness,
                    )
                )
            return garages

    def most_recent_record(self, garage_id: int) -> Garage | None:
        stmt = (
            select(
                GarageInfo.garage_name,
                GarageFullness.garage_id,
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
            return Garage(
                garage_id=result.garage_id,
                name=result.garage_name,
                utc_timestamp=result.utc_timestamp,
                fullness=result.fullness,
            )

    def calendar_info(self, timestamp: datetime) -> list[CalendarDate] | None:
        timestamp_end = timestamp + timedelta(days=self.PREDICTION_DAY_WINDOW)
        stmt = select(
            Calendar.calendar_date,
            Calendar.is_weekend,
            Calendar.is_campus_closed,
            Calendar.utc_start,
            Calendar.utc_end,
        ).where(Calendar.utc_start < timestamp_end, Calendar.utc_end > timestamp)
        with Session(self.engine) as sess:
            results = sess.execute(stmt)
            # NOTE: Edge case where the utc_end is beyond the window so only 1 record gets returned, find a way to verify how results can have multiple days.
            # This is pretty rare though because the calendar table goes to 2030 and who knows what I'll be doing then...
            if results is None:
                return None
            calendar_info = []
            for (
                calendar_date,
                is_campus_closed,
                is_weekend,
                utc_start,
                utc_end,
            ) in results:
                calendar_info.append(
                    CalendarDate(
                        calendar_date=calendar_date,
                        is_campus_closed=is_campus_closed,
                        is_weekend=is_weekend,
                        utc_start=utc_start,
                        utc_end=utc_end,
                    )
                )
            return calendar_info
