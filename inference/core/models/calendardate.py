from datetime import date, datetime
from pydantic import BaseModel
from pandas import DataFrame


class CalendarDate(BaseModel):
    calendar_date: date
    is_campus_closed: bool
    is_weekend: bool
    utc_start: datetime
    utc_end: datetime

    def to_df(self) -> DataFrame:
        return DataFrame(
            {
                "is_campus_closed": [int(self.is_campus_closed)],
                "is_weekend": [int(self.is_weekend)],
            }
        )
