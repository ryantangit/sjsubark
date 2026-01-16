from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from contextlib import asynccontextmanager
from core.trainedmodel import TrainedModel
from core.sql import SQLAccessor

ml_model = TrainedModel("./gradient_boost_pipeline.joblib")
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

'''
@app.get("/predict")
def predict(garage_id: int, increment: int = 0) -> None:
    # Retrieve the most recent record
    record = database.most_recent_record(garage_id)
'''
