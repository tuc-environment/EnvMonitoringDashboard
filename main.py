from typing import Union

from fastapi import FastAPI
import uvicorn

app = FastAPI()


@app.get("/stations")
def stations():
    return {
        "temp": 10,
        "humidity": 10,
        "barometric_pressure": 10,
        "soil_temp_shallow": 10,
        "soil_temp_deep": 10,
        "soil_water_content_shallow": 10,
        "soil_water_content_deep": 10,
        "soil_electrical_conductivity": 10,
    }


if __name__ == "__main__":
    config = uvicorn.Config("main:app", port=5000, log_level="info")
    server = uvicorn.Server(config)
    server.run()
