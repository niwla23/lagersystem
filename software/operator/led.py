# SPDX-FileCopyrightText: 2021 ladyada for Adafruit Industries
# SPDX-License-Identifier: MIT

# Simple test for NeoPixels on Raspberry Pi
import time
import board
import neopixel
import multiprocessing

# Choose an open pin connected to the Data In of the NeoPixel strip, i.e. board.D18
# NeoPixels must be connected to D10, D12, D18 or D21 to work.
pixel_pin = board.D12

# The number of NeoPixels
num_pixels = 11

# The order of the pixel colors - RGB or GRB. Some NeoPixels have red and green reversed!
# For RGBW NeoPixels, simply change the ORDER to RGBW or GRBW.

# broken. [blue,red,green]
ORDER = neopixel.RGB

pixels = neopixel.NeoPixel(
    pixel_pin, num_pixels, brightness=1, auto_write=False, pixel_order=ORDER
)


def rgb(r, g, b):
    return (b, r, g)


def wheel(pos):
    # Input a value 0 to 255 to get a color value.
    # The colours are a transition r - g - b - back to r.
    if pos < 0 or pos > 255:
        r = g = b = 0
    elif pos < 85:
        r = int(pos * 3)
        g = int(255 - pos * 3)
        b = 0
    elif pos < 170:
        pos -= 85
        r = int(255 - pos * 3)
        g = 0
        b = int(pos * 3)
    else:
        pos -= 170
        r = 0
        g = int(pos * 3)
        b = int(255 - pos * 3)
    return (r, g, b) if ORDER in (neopixel.RGB, neopixel.GRB) else (r, g, b, 0)


def rainbow_cycle(wait):
    for j in range(255):
        for i in range(num_pixels):
            pixel_index = (i * 256 // num_pixels) + j
            pixels[i] = wheel(pixel_index & 255)
        pixels.show()
        time.sleep(wait)


def rainbow_cycle_reverse(wait):
    for j in reversed(range(255)):
        for i in range(num_pixels):
            pixel_index = (i * 256 // num_pixels) + j
            pixels[i] = wheel(pixel_index & 255)
        pixels.show()
        time.sleep(wait)


def done_animation():
    for i in range(num_pixels):
        pixels[i] = rgb(0, 255, 0)
        pixels.show()
        time.sleep(0.1)

    time.sleep(5)
    # fade to white
    for i in range(255):
        pixels.fill(rgb(i, 255, i))
        pixels.show()
        time.sleep(0.01)
    # for i in range(num_pixels):
    #     pixels[i] = rgb(0, 0, 0)
    #     pixels.show()
    #     time.sleep(0.1)


def pickup_animation():
    shift = 0
    while True:
        for i in range(num_pixels):
            is_on = (i + shift) % 4 == 0
            pixels[i] = rgb(255, 0, 255) if is_on else rgb(0, 0, 0)
        pixels.show()
        shift += 1
        time.sleep(0.2)


def store_animation():
    while True:
        rainbow_cycle_reverse(0.001)


def rainbow_forward_animation():
    while True:
        rainbow_cycle(0.001)


def rainbow_reverse_animation():
    while True:
        rainbow_cycle_reverse(0.001)


animation_thread = multiprocessing.Process()


def animate_no_block(animation):
    global animation_thread
    try:
        animation_thread.kill()
    except AttributeError:
        pass
    animation_thread = multiprocessing.Process(target=animation)
    animation_thread.start()


def fill(r, g, b):
    try:
        animation_thread.kill()
    except AttributeError:
        pass
    pixels.fill(rgb(r, g, b))
    pixels.show()


# while True:
#     # Comment this line out if you have RGBW/GRBW NeoPixels
#     pixels.fill((255, 0, 0))
#     # Uncomment this line if you have RGBW/GRBW NeoPixels
#     # pixels.fill((255, 0, 0, 0))
#     pixels.show()
#     time.sleep(1)

#     # Comment this line out if you have RGBW/GRBW NeoPixels
#     pixels.fill((0, 255, 0))
#     # Uncomment this line if you have RGBW/GRBW NeoPixels
#     # pixels.fill((0, 255, 0, 0))
#     pixels.show()
#     time.sleep(1)

#     # Comment this line out if you have RGBW/GRBW NeoPixels
#     pixels.fill((0, 0, 255))
#     # Uncomment this line if you have RGBW/GRBW NeoPixels
#     # pixels.fill((0, 0, 255, 0))
#     pixels.show()
#     time.sleep(1)

#     for i in range(20):
#         rainbow_cycle(0.001)  # rainbow cycle with 1ms delay per step
