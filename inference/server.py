from fastapi import FastAPI
from pydantic import BaseModel
from contextlib import asynccontextmanager
from core.trainedmodel import TrainedModel
from core.sql import SQLAccessor 

ml_model = TrainedModel("./gradient_boost_model.pkl")
database = SQLAccessor()

@asynccontextmanager
async def lifespan(app: FastAPI):
    ml_model.load_model()
    yield
    ml_model.unload_model()

app = FastAPI(lifespan=lifespan)

@app.get("/health")
def health_check():
    return {
            "status": "ok", 
            "model_loaded": ml_model.model != None
            }

class Prediction(BaseModel):
    name: str
    forecast: int

@app.get("/predict")
def predict(garage: str) -> list[Prediction]:
    return []

