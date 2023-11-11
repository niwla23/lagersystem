from gpiozero import Button
import led
import worker
import standby
import config
import requests
import movement


class Buttons:
    light = Button(0)
    silent_mode = Button(5)
    clear_io = Button(6)
    home = Button(13)
    standby = Button(19)


def toggle_standby(x):
    if standby.active:
        worker.queue_wakeup()
    else:
        worker.queue_standby()


silent_mode_on = False


def silent_mode(x):
    global silent_mode_on
    movement.set_silentmode(not silent_mode_on)
    silent_mode_on = not silent_mode_on


def clear_io(x):
    base_url = config.config["managerBaseUrl"].get(str)
    requests.post(f"{base_url}/warehouses/1/clearIO")


light_mode = 0


def light(x):
    global light_mode
    if light_mode == 9:
        light_mode = 0
    else:
        light_mode += 1

    if light_mode == 0:
        led.fill(0, 0, 0)
    elif light_mode == 1:
        led.animate_no_block(led.rainbow_reverse_animation)
    elif light_mode == 2:
        led.animate_no_block(led.rainbow_forward_animation)
    elif light_mode == 3:
        led.fill(255, 0, 0)
    elif light_mode == 4:
        led.fill(0, 255, 0)
    elif light_mode == 5:
        led.fill(0, 0, 255)
    elif light_mode == 6:
        led.fill(255, 255, 0)
    elif light_mode == 7:
        led.fill(255, 0, 255)
    elif light_mode == 8:
        led.fill(0, 255, 255)
    elif light_mode == 9:
        led.fill(255, 255, 255)


Buttons.standby.when_pressed = toggle_standby
Buttons.silent_mode.when_pressed = silent_mode
Buttons.home.when_pressed = worker.queue_home
Buttons.light.when_pressed = light
Buttons.clear_io.when_pressed = clear_io

print("button thing")
