import standby
import socket
import threading
import buttons
import worker
import iostate
import movement
import config
import time
from flask import Flask
import led

led.fill(255, 0, 0)


print("loaded flask")

print("loaded config")

print("loaded movement")

print("loaded iostate")

print("loaded worker")

print("loaded input")

print("loaded threading")

print("loaded socket")


socket.setdefaulttimeout(120)  # seconds

print("loaded modules, homing...")
movement.restart_motors()
movement.home()
movement.set_silentmode(False)

app = Flask(__name__)


def background_task():
    while True:
        worker.work()
        standy_timeout_seconds = config.config["standyTimeoutSeconds"].get(int)
        if (
            standby.last_action + standy_timeout_seconds < time.time()
            and not standby.active
        ):
            print("turning off motors")
            # standby.enable()
            worker.queue_standby()
        time.sleep(1)


threading.Thread(target=background_task).start()


@app.route("/scanBoxId/<int:ioPosition>")
def scanBoxId(ioPosition: int):
    startTime = time.time()
    job_id = worker.wait_for_job(worker.queue_scan_iopos(ioPosition))
    box_id = worker.job_outputs[job_id]

    return {"boxId": box_id, "duration": time.time() - startTime}


@app.route("/getIOState")
def getIOState():
    return {"ioState": iostate.get_io_state()}


@app.route("/home")
def home():
    startTime = time.time()
    job_id = worker.wait_for_job(worker.queue_home())
    out = worker.job_outputs[job_id]
    return {"duration": time.time() - startTime, "message": out}


@app.route("/storeBox/<int:position>/<int:ioPosition>")
def storeBox(position: int, ioPosition: int):
    startTime = time.time()
    job_id = worker.wait_for_job(worker.queue_store(position, ioPosition))
    time.sleep(0.1)
    out = worker.job_outputs[job_id]
    status_code = 200 if out == "success" else 500

    return {"duration": time.time() - startTime, "message": out}, status_code


@app.route("/pickupBox/<int:position>/<int:ioPosition>")
def pickupBox(position: int, ioPosition: int):
    print(f"[request received] pickup request from {position} to {ioPosition}")
    startTime = time.time()
    job_id = worker.wait_for_job(worker.queue_pickup(position, ioPosition))
    out = worker.job_outputs[job_id]

    return {"duration": time.time() - startTime, "message": out}


@app.route("/getPositions")
def getPositions():
    return {
        "positions": config.config["positions"].get(),
    }


try:
    led.fill(255, 255, 255)
    app.run(host="0.0.0.0")
except KeyboardInterrupt:
    print("stopping")
    standby.activate()
    exit()
