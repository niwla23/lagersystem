import movement
import json
import sys, tty, termios

def read_char():
  fd = sys.stdin.fileno()
  old_settings = termios.tcgetattr(fd)
  try:
    tty.setraw(sys.stdin.fileno())
    ch = sys.stdin.read(1)
  finally:
    termios.tcsetattr(fd, termios.TCSADRAIN, old_settings)
  return ch

def bround(x, base=0.5):
    return base * round(x/base)

#movement.home()

position_config = {}

pos_id = int(input("starting position id: "))
while True:
  print(f"configuring position {pos_id}")
  print("use wasd to move on x/y, er for z. To test pickup press I, for storing O. Q to exit")
  print("to re-home press h")
  while True:
    movement.get_position()
    key = read_char()
    if key == "q":
      print(json.dumps(position_config, indent=2))
      exit(0)
    elif key == "w":
      movement.jog(0,0.5,0)
      
    elif key == "W":
      movement.jog(0,46,0)
      
    elif key == "s":
      movement.jog(0,-0.5,0)
      
    elif key == "S":
      movement.jog(0,-46,0)
      
    elif key == "a":
      movement.jog(1,0,0)
      
    elif key == "A":
      movement.jog(73,0,0)
      
    elif key == "d":
      movement.jog(-0.5,0,0)
      
    elif key == "D":
      movement.jog(-73,0,0)
      
    elif key == "e":
      movement.jog(0,0,3)
      
    elif key == "E":
      movement.go_to_and_wait(None, None, 120)
      
    elif key == "r":
      movement.jog(0,0,-3)
      
    elif key == "R":
      movement.retract_carriage()
      
    elif key == "i":
      pos = movement.get_position()
      movement.pickup_by_xy(pos[0], pos[1]) 
      
    elif key == "o":
      pos = movement.get_position()
      movement.store_by_xy(pos[0], pos[1])
      
    elif key == "h":
      movement.home()

    elif key == "c":
      movement.ser.write(b"!")

    elif key == "\x7f":
      pos_id -=1
      break

    elif key == "\r":
      pos = movement.get_position()

      print()
      position_config[pos_id] = {
        "x": bround(pos[0]),
        "y": bround(pos[1])
      }
      with open("temp_config.yaml", 'w') as file:
        file.write(json.dumps(position_config, indent=2))
      pos_id +=1
      break

    new_pos = movement.get_position()
    print(f"\rX:{new_pos[0]} Y:{new_pos[1]} Z:{new_pos[2]}                                                                                       \r", end="")
