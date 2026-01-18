from datetime import datetime
from zoneinfo import ZoneInfo
from pydantic import BaseModel
from pandas import DataFrame

tz_la = ZoneInfo("America/Los_Angeles")


class Garage(BaseModel):
    garage_id: int
    name: str
    utc_timestamp: datetime
    fullness: int

    def to_df(self) -> DataFrame:
        la_time = self.utc_timestamp.astimezone(tz_la)
        return DataFrame(
            {
                "garage_id": [self.garage_id],
                "fullness": [self.fullness],
                "hour": [la_time.hour],
                "min": [la_time.minute],
                "day": [la_time.day],
                "year": [la_time.year],
                "month": [la_time.month],
                "day_of_week": [la_time.weekday()],
            }
        )
