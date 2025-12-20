from fastapi import FastAPI
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
    return {
            "status": "ok", 
            "model_loaded": ml_model.pipeline != None
            }

class Prediction(BaseModel):
    name: str
    forecast: int

@app.get("/predict")
def predict(garage: str, increment: int = 0) -> Prediction:
    # Retrieve the most recent record
    record = database.most_recent_record(garage)
    if not len(record) == 0:
        #Exception - Server Error
        pass
    return Prediction(name=garage, forecast=ml_model.predict(record))

