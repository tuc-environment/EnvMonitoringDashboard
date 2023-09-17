import pandas as pd
import numpy as np
from sklearn.ensemble import RandomForestRegressor

TRAIN_DATA_PATH = "./model/train.xlsx"

class Model:
    def __init__(self) -> None:
        train = pd.read_excel(TRAIN_DATA_PATH, "Sheet1")[2:].to_numpy()
        label = pd.read_excel(TRAIN_DATA_PATH, "Sheet2")[2:].to_numpy()

        self.clf = RandomForestRegressor(random_state=3)
        self.clf.fit(train, label)
        print("Model is initialized.")

    def predict(self, args: np.ndarray) -> np.ndarray:
        return self.clf.predict(args).ravel()
