# This python scripts converts the json file into a single column text file
# The txt file is then passed into the database server for marking which days are regarded as holidays


import json
import datetime

with open("./campus_close.json", "r") as f:
    yearly_entries = json.load(f)

with open("./campus_close.txt", "w") as fp:
    for entry in yearly_entries:
        year = entry["year"]
        instances = entry["instances"]
        for i in instances:
            start_month = i["startmonth"]
            start_day = i["startday"]
            end_month = i["endmonth"]
            end_day = i["endday"]
            start_date = datetime.date(year, start_month, start_day)
            end_date = datetime.date(year, end_month, end_day)
            day_diff = datetime.timedelta(days=1)
            while start_date <= end_date:
                fp.write(start_date.isoformat()+"\n")
                start_date += day_diff
