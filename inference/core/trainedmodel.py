import joblib
from typing import Optional
from sklearn.ensemble import HistGradientBoostingRegressor

class TrainedModel:
    def __init__(self, filepath: str):
        self.filepath: str = filepath
        self.model: Optional[HistGradientBoostingRegressor] = None

    def load_model(self):
        try:
            self.model = joblib.load(self.filepath)
            print("model loaded successfully")
        except:
            print("Failed to load model")
    
    def unload_model(self):
        self.model = None

    def predict(self, garage: str):
        pass
