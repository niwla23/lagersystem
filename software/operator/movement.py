import serial
import time
import led
from config import config
import standby

ser = serial.Serial("/dev/ttyACM0", 115200)
# absolute positioning
ser.write(b"G90\n")


def read_all_lines():
    data = ser.read_all().decode()
    return data.split("\n")


def all_equal(iterator):
    iterator = iter(iterator)
    try:
        first = next(iterator)
    except StopIteration:
        return True
    return all(first == x for x in iterator)


def handle_alarm():
    print("system is in alarm state, exiting")
    exit(1)


def jog(x, y, z):
    ser.write(f"$J=G91 F20000 X{x} Y{y} Z{z}\n".encode())


def turn_off():
    time.sleep(0.1)
    ser.write(b"$X\n")
    ser.write(b"$1=0\n")
    time.sleep(0.1)
    ser.write(b"G91\n G01 F1000 X0.1\n G90\n")
    time.sleep(0.5)
    ser.write(b"$1=255\n")
    time.sleep(0.1)


def restart_motors():
    turn_off()
    time.sleep(5)
    # power motors
    ser.write(b"G91\n G01 F1000 X0.1\n G90\n")
    time.sleep(0.5)


def home():
    ser.read_all()
    ser.write(b"$H\n")
    time.sleep(4)
    while True:
        time.sleep(0.1)
        lines = ser.read_all().decode()
        if "ok" in lines:
            offset = config["globalOffset"].get(float)
            ser.write(f"G92 X{offset} Y{offset} Z{offset}\n".encode())
            ser.write(b"G01 F1000 X0 Y0 Z0\n")
            return


def get_position():
    read_all_lines()
    while True:
        ser.write(b"?\n")
        lines = read_all_lines()
        for line in lines:
            try:
                if line.split("|")[0] == "Alarm":
                    handle_alarm()

                work_pos_str = line.split("|")[1].split(":")[1].split(",")
                work_pos = [float(coord) for coord in work_pos_str]
                if len(work_pos) < 3:
                    continue
                return work_pos
            except (IndexError, ValueError):
                continue

    work_pos_str = line.split("|")[1].split(":")[1].split(",")
    work_pos = [float(coord) for coord in work_pos_str]
    return work_pos


def go_to_and_wait(x, y, z):
    standby.last_action = time.time()
    if standby.active:
        standby.wakeup()

    x = x + config["positionOffset"]["x"].get(float) if x is not None else None
    y = y + config["positionOffset"]["y"].get(float) if y is not None else None
    z = z + config["positionOffset"]["z"].get(float) if z is not None else None
    command = "G01"
    if x is not None:
        command += f" X{x}"
    if y is not None:
        command += f" Y{y}"
    if z is not None:
        command += f" Z{z}"
    command += " F20000\n"
    ser.write(bytes(command, "utf-8"))

    buffer = []
    while True:
        standby.last_action = time.time()
        work_pos = get_position()
        buffer.append(work_pos)

        deltas = [abs(a - b)
                  for a, b in zip(work_pos, [x, y, z]) if b is not None]

        if all(x < 0.01 for x in deltas) or not any([x is not None for x in [x, y, z]]):
            return

        if all_equal(buffer[-100:]) and len(buffer) > 100:
            raise Exception("the thingy is not moving")

        time.sleep(0.1)


def extend_carriage():
    go_to_and_wait(None, None, 135)


def retract_carriage():
    go_to_and_wait(None, None, 0)


def pickup_by_xy(x: float, y: float):
    # retract_carriage()
    # go to the box, but lower so we can grab it from the bottom
    go_to_and_wait(x, y - 8, 100)
    extend_carriage()
    # pick up box by rising y to normal level of box
    go_to_and_wait(None, y, None)
    # go_to_and_wait(None, y+12, 0)
    retract_carriage()


def store_by_xy(x: float, y: float):
    # retract_carriage()
    # go to position
    go_to_and_wait(x, y, None)
    extend_carriage()
    # lower carriage
    go_to_and_wait(None, y - 8, 135)
    # back out
    go_to_and_wait(None, None, 100)


def get_xy_for_posid(posId: int):
    position = config["positions"][str(posId)].get()
    return position


def go_to_posid(posId: int):
    position = get_xy_for_posid(posId)
    go_to_and_wait(position["x"], position["y"], None)


def go_to_io_pos(ioPosId: int):
    position = config["ioPositions"][str(ioPosId)].get()
    go_to_and_wait(position["x"], position["y"], 0)


def go_to_scanner():
    position = config["scanPosition"].get()
    go_to_and_wait(position["x"], position["y"], None)
    go_to_and_wait(None, None, position["z"])


def pickup_by_posid(posId: int):
    position = get_xy_for_posid(posId)
    pickup_by_xy(position["x"], position["y"])


def pickup_by_io_pos_id(ioPosId: int):
    position = config["ioPositions"][str(ioPosId)].get()
    pickup_by_xy(position["x"], position["y"])


def store_by_posid(posId: int):
    position = get_xy_for_posid(posId)
    store_by_xy(position["x"], position["y"])


def store_by_io_pos_id(ioPosId: int):
    position = config["ioPositions"][str(ioPosId)].get()
    store_by_xy(position["x"], position["y"])


def set_silentmode(on: bool):
    print("silent mode on:", on)
    if on:
        led.fill(0, 0, 255)
        ser.write(b"F650")
    else:
        led.fill(255, 255, 255)
        ser.write(b"F20000")

# home()
# print("homed.")
# time.sleep(1)
# pick_up_by_posid(2)
# time.sleep(5)
# store_by_posid(2)
