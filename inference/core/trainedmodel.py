import joblib
import pandas as pd
from typing import Optional
from sklearn.pipeline import Pipeline 

def drop_unnecessary_cols(df):
    cols_to_drop = ['utc_timestamp', "time_lag1_diff", "second"]
    # We only drop if they exist (to avoid errors during single-record prediction)
    return df.drop(columns=[c for c in cols_to_drop if c in df.columns])


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

    def predict(self, record):
        if not self.pipeline: return -1
        record = drop_unnecessary_cols(record)
        forecast = self.pipeline.predict(record)[0]
        return round(forecast)
