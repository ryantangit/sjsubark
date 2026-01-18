import joblib
import pandas as pd
from typing import Optional
from sklearn.pipeline import Pipeline

from core.models import Garage, CalendarDate


class TrainedModel:
    def __init__(self, filepath: str):
        self.filepath: str = filepath
        self.pipeline: Optional[Pipeline] = None

    def load_pipeline(self):
        try:
            self.pipeline = joblib.load(self.filepath)
            print("pipeline loaded successfully")
        except Exception as e:
            print("Failed to load pipeline")
            print(e)

    def unload_model(self):
        self.pipeline = None

    def predict(self, garage: Garage, calendar_date: CalendarDate):
        if not self.pipeline:
            return -1
        garage_df = garage.to_df()
        calendar_df = calendar_date.to_df()
        combined_df = pd.concat([garage_df, calendar_df], axis=1)
        return round(self.pipeline.predict(combined_df)[0])
