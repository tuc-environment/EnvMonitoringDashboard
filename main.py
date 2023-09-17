from fastapi import FastAPI
from model import Model

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
    args = [
        temp,
        humidity,
        barometric_pressure,
        soil_temp_shallow,
        soil_temp_deep,
        soil_water_content_shallow,
        soil_water_content_deep,
        soil_electrical_conductivity,
    ]
    result = model.predict(args)

    return {
        "temp": result[0],
        "humidity": result[1],
        "barometric_pressure": result[2],
        "soil_temp_shallow": result[3],
        "soil_temp_deep": result[4],
        "soil_water_content_shallow": result[5],
        "soil_water_content_deep": result[6],
        "soil_electrical_conductivity": result[7],
    }
