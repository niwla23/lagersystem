import movement
from collections import Counter
import random
import config

movement.home()

free_spot = 1
start_pos = 65
end_pos = 128

try:
  movement.pickup_by_posid(1)
  movement.store_by_io_pos_id(2)
  while True:
    full_positions = list(range(start_pos,free_spot)) + list(range(free_spot+1, end_pos+1))
    # print(full_positions)
    position_to_pick = random.choice(full_positions)
    print(f"moving {position_to_pick} to {free_spot}")
    movement.pickup_by_posid(position_to_pick)
    movement.store_by_posid(free_spot)
    free_spot = position_to_pick
    config.reload()
except Exception as e:
  print(e)
  movement.ser.write(0x18)
  exit(0)
