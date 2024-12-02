import threading
import queue
import time
import movement
import uuid
import scan
import iostate
import led
import standby
from config import config
from typing import Dict

job_queue = queue.Queue()
# maps job id to job output
job_outputs: Dict[str, str] = {}


def queue_home():
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "home",
            "job_id": job_id,
        }
    )
    return job_id


def queue_store(position: int, ioPosition: int):
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "store",
            "ioPosition": ioPosition,
            "position": position,
            "job_id": job_id,
        }
    )
    return job_id


def queue_pickup(position: int, ioPosition: int):
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "pickup",
            "ioPosition": ioPosition,
            "position": position,
            "job_id": job_id,
        }
    )
    return job_id


def queue_scan_iopos(ioPosition: int):
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "scan_iopos",
            "ioPosition": ioPosition,
            "job_id": job_id,
        }
    )
    return job_id


def queue_standby():
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "standby",
            "job_id": job_id,
        }
    )
    return job_id


def queue_wakeup():
    job_id = uuid.uuid4()
    job_queue.put(
        {
            "type": "wakeup",
            "job_id": job_id,
        }
    )
    return job_id


def is_job_done(job_id):
    for job in job_queue.queue:
        if job["job_id"] == job_id:
            return False
    return True


def wait_for_job(job_id):
    while not is_job_done(job_id):
        time.sleep(0.1)
    return job_id


def job_done(output: str = None):
    print("job done")
    job = job_queue.get()
    job_outputs[job["job_id"]] = output
    job_queue.task_done()


def work():
    print("looking for work")
    try:
        job = job_queue.queue[0]
        print(job)
        if job["type"] == "store":
            led.animate_no_block(led.store_animation)
            if iostate.is_io_pos_free(job["ioPosition"]):
                raise Exception("error: source io empty")
            movement.pickup_by_io_pos_id(job["ioPosition"])
            movement.store_by_posid(job["position"])
            led.animate_no_block(led.done_animation)
            job_done("success")

        elif job["type"] == "pickup":
            led.animate_no_block(led.pickup_animation)

            io_pos_id = job["ioPosition"]
            if io_pos_id == 0:
                if iostate.is_io_pos_free(1):
                    io_pos_id = 1
                elif iostate.is_io_pos_free(2):
                    io_pos_id = 2
                elif iostate.is_io_pos_free(3):
                    io_pos_id = 3
                else:
                    raise Exception("error: all io positions full")

            movement.pickup_by_posid(job["position"])


            movement.store_by_io_pos_id(io_pos_id)
            led.animate_no_block(led.done_animation)
            job_done("success")

        elif job["type"] == "scan_iopos":
            # if scanner is blocked and we are not trying to read from the scanner pos, error
            if not iostate.is_io_pos_free(1) and job["ioPosition"] != 1:
                raise Exception("error: scanner pos not free")
            # if source io is empty, error
            if iostate.is_io_pos_free(job["ioPosition"]):
                raise Exception("source io empty")

            # dont do the dance if we are already at the scanner
            if job["ioPosition"] == 1:
                position = config["ioPositions"][str(job["ioPosition"])].get()
                movement.go_to_and_wait(position["x"], position["y"] - 8, 100)
                movement.extend_carriage()
                # pick up box by rising y to normal level of box
                movement.go_to_and_wait(None, position["y"], None)
            else:
                movement.pickup_by_io_pos_id(job["ioPosition"])
            movement.go_to_scanner()
            box_id = scan.do_scan()
            # movement.extend_carriage()
            movement.store_by_io_pos_id(1)
            job_done(box_id)
            # movement.store_by_io_pos_id(job["ioPosition"])

        elif job["type"] == "home":
            led.fill(255, 0, 0)
            movement.home()
            led.fill(255, 255, 255)
            job_done("success")

        elif job["type"] == "standby":
            standby.activate()
            job_done("success")

        elif job["type"] == "wakeup":
            standby.wakeup()
            job_done("success")

    except IndexError:
        return
    except Exception as e:
        job_done("exception: " + str(e))
