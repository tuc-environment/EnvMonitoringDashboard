from fastapi import FastAPI
from model import Model
import numpy as np

model = Model()
app = FastAPI()


@app.get("/stations")
def stations(
    temp: float,
    humidity: float,
    barometric_pressure: float,
    soil_temp_shallow: float,
    soil_temp_deep: float,
    soil_water_content_shallow: float,
    soil_water_content_deep: float,
    soil_electrical_conductivity: float,
):
    args = np.array([
        temp,
        humidity,
        barometric_pressure,
        soil_temp_shallow,
        soil_temp_deep,
        soil_water_content_shallow,
        soil_water_content_deep,
        soil_electrical_conductivity,
    ]).reshape(1, -1)
    result = model.predict(args)
    if len(result) >= 12:
        return {
            "down_soil_temp_shallow": result[0],
            "down_soil_temp_deep": result[1],
            "down_soil_water_content_shallow": result[2],
            "down_soil_water_content_deep": result[3],
            "down_temp": result[4],
            "down_humidity": result[5],
            "middle_soil_temp_shallow": result[6],
            "middle_soil_temp_deep": result[7],
            "middle_soil_water_content_shallow": result[8],
            "middle_soil_water_content_deep": result[9],
            "middle_temp": result[10],
            "middle_humidity": result[11]
        }
    else:
        return {}