from gpiozero import LineSensor
from config import config

# dynamic typing lets goo
positions = config["ioPositions"].get()
for i, pos in positions.items():
    pos["sensor"] = LineSensor(pos["sensorPin"])

def get_io_state():
    answer = {}

    for i, pos in positions.items():
        answer[i] = "free" if pos["sensor"].value == 1 else "occupied"

    return answer

def is_io_pos_free(io_pos_id: int):
    return positions[str(io_pos_id)]["sensor"].value == 1
