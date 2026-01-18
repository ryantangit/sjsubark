from datetime import datetime, timedelta
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from contextlib import asynccontextmanager
from core import TrainedModel, SQLAccessor

ml_model = TrainedModel("./boost_pipelines.joblib")
database = SQLAccessor()


@asynccontextmanager
async def lifespan(app: FastAPI):
    ml_model.load_pipeline()
    yield
    ml_model.unload_model()


app = FastAPI(lifespan=lifespan)


@app.get("/health")
def health_check():
    return {"status": "ok", "model_loaded": ml_model.pipeline != None}


@app.get("/recent")
def recent():
    recent_records = database.most_recent_garage_records()
    return recent_records


class Prediction(BaseModel):
    name: str
    forecast: int


@app.get("/predict")
def predict(garage_id: int, increment: int = 0) -> Prediction:
    # Retrieve the most recent record
    garage = database.most_recent_record(garage_id)
    if not garage:
        raise HTTPException(status_code=404, detail="Garage not found.")
    calendar_info = database.calendar_info(garage.utc_timestamp)
    if not calendar_info:
        raise HTTPException(
            status_code=404, detail="Record UTC time not found in calendar."
        )
    forecast = -1
    for _ in range(0, 1 + increment):
        calendar_date = calendar_info[0]
        for date in calendar_info:
            if (
                date.utc_start >= garage.utc_timestamp
                and date.utc_end < garage.utc_timestamp
            ):
                calendar_date = date
                break
        forecast = ml_model.predict(garage, calendar_date)
        garage.fullness = forecast
        garage.utc_timestamp += timedelta(minutes=10)
    return Prediction(name=garage.name, forecast=forecast)
