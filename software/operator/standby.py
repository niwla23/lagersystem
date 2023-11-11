import movement
import led
import time

active = False
last_action = time.time()

def activate():
    global active
    global last_action
    led.fill(0, 0, 50)
    movement.go_to_and_wait(None, None, 0)
    movement.go_to_and_wait(0, 0, 0)
    movement.turn_off()
    active = True
    last_action = time.time()

def wakeup():
    global active
    global last_action

    led.fill(255,0,0)
    print("system is on standby, waking up")
    active = False
    movement.home()
    led.fill(255,255,255)
    last_action = time.time()
    