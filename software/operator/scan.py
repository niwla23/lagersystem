import cv2
import zxingcpp
import time
from picamera import PiCamera
from gpiozero import LED
import atexit
import led

print("loaded modules")

camera = PiCamera()
camera.resolution = (640, 480)

flash = LED(26)

def do_scan():
    led.fill(0, 0, 0)
    time.sleep(0.1)
    flash.on()
    led.fill(255, 255, 255)
    time.sleep(0.1)
    camera.capture('qrcode.jpg')
    flash.off()
    img = cv2.imread('qrcode.jpg')
    results = zxingcpp.read_barcodes(img)
    if len(results) == 0:
        return
    return results[0].text

def exit_handler():
    camera.close()
    print("camera closed")

atexit.register(exit_handler)


if __name__ == "__main__":
  try:
    while True:
      input("pres enter to scan...")
      print(do_scan())

  except:
    camera.close()
