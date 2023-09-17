import numpy as np
from sklearn.ensemble import RandomForestRegressor
import csv

TRAIN_DATA_PATH = "./model/train.csv"
LABEL_DATA_PATH = "./model/label.csv"


def read_csv(path):
    with open(path, "r") as f:
        reader = csv.reader(f)
        data = list(reader)
        data = np.array(data, dtype="float64")
    return data


class Model:
    def __init__(self) -> None:
        pattern = read_csv(TRAIN_DATA_PATH)
        label = read_csv(LABEL_DATA_PATH)

        self.clf = RandomForestRegressor(random_state=3)
        self.clf.fit(pattern, label.ravel())
        print("Model is initialized.")

    def predict(self, args: np.ndarray) -> np.ndarray:
        return self.clf.predict(args)
