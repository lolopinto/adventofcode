from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math

@dataclass
class Hailstone:
  position: tuple[int, int, int]
  velocity: tuple[int, int, int]
  
  @staticmethod
  def parse_parts(part: str) -> tuple[int, int, int]:
    return tuple(int(v) for v in part.split(', '))

  @staticmethod
  def parse(line: str) -> Hailstone:
    parts = line.split(' @ ')
    assert len(parts) == 2
    position = Hailstone.parse_parts(parts[0])
    velocity = Hailstone.parse_parts(parts[1])
    return Hailstone(position, velocity)
  
  # def in_target(self, target: tuple[int, int], index: int) -> bool:
  #   xvalue = self.position[0] + index * self.velocity[0]
  #   yvalue = self.position[1] + index * self.velocity[1]

  #   ret = target[0] <= xvalue <= target[1] and target[0] <= yvalue <= target[1]
  #   # if ret:
  #   #   print(f"{self} {index} {xvalue} {yvalue}")
  #   return ret, (xvalue, yvalue)
    
# def hailstones_intersect(h1: Hailstone, h2: Hailstone, target: tuple[int, int]) -> tuple[int|None, int | None]:
#   x = helper(h1.position[0], h1.velocity[0], h2.position[0], h2.velocity[0])
#   y = helper(h1.position[1], h1.velocity[1], h2.position[1], h2.velocity[1]) 

#   if x is not None:
#     h1, h1coords = h1.in_target(target, x) 
#     h2, h2coords = h2.in_target(target, x)
#     ret = h1 and h2 and (h1coords[1] == h2coords[1] or h1coords[0] == h2coords[0])
#     if ret:
#       print(f"{h1} {h2} {x} {h1coords} {h2coords}")
#     return ret
#     # xvalue = h1.position[0] + int(x) * h1.velocity[0]
#     # yvalue = h2.position[0] + int(x) * h2.velocity[0]
#     # return xvalue <= target[0] <= target[1] and yvalue <= target[0] <= target[1]
  
#   if y is not None:
#     h1, h1coords = h1.in_target(target, y) 
#     h2, h2coords = h2.in_target(target, y)
#     ret = h1 and h2 and (h1coords[1] == h2coords[1] or h1coords[0] == h2coords[0])
#     if ret:
#       print(f"{h1} {h2} {y} {h1coords} {h2coords}")
#     return ret

#     # return h1.in_target(target, y) and h2.in_target(target, y)
#     # xvalue = h1.position[1] + int(y) * h1.velocity[1]
#     # yvalue = h2.position[1] + int(y) * h2.velocity[1]
#     # return xvalue <= target[0] <= target[1] and yvalue <= target[0] <= target[1]
  
#   return False
  
# def helper(a0, da, b0, db):
#   if da == db:
#     return 0 if a0 == b0 else None  # Sequences are identical or parallel
#   n = (b0 - a0) / (da - db)
#   if n.is_integer() and n >= 0:
#     return n
#   return None

def value(start:int, delta:int, index: int):
  return start + index * delta

def in_target(value: int, target: tuple[int, int]) -> bool:
  return target[0] <= value <= target[1]

def hailstones_intersect(h1: Hailstone, h2: Hailstone, target: tuple[int, int]) -> bool:
  
  if hailstones_intersect_helper(h1, h2, target, check_x=True):
    return True
  return hailstones_intersect_helper(h1, h2, target, check_x=False)
    

def hailstones_intersect_helper(h1: Hailstone, h2: Hailstone, target: tuple[int, int], *, check_x: bool) -> bool:
  # make this large enough
  low, hi = 0, 1_000_000_000_000_000_000_000_000_000_000_000_000_000_000
  
  # value shouldn't change based on how many times i try or what hte max index is?
  
  while low < hi:
    mid = (low + hi) // 2
    
    x1 = value(h1.position[0], h1.velocity[0], mid)
    y1 = value(h1.position[1], h1.velocity[1], mid)

    x2 = value(h2.position[0], h2.velocity[0], mid)
    y2 = value(h2.position[1], h2.velocity[1], mid)
    
    if x1 == x2 and in_target(x1, target) and in_target(y1, target) and in_target(y2, target):
      print(f"{h1} {h2} {x1} {y1} {x2} {y2}")
      return True

    if check_x:
      if x1 < x2:
        low = mid + 1
      else:
        hi = mid - 1
    else:
      if y1 < y2:
        low = mid + 1
      else:
        hi = mid - 1
    
  return False
 
 
def find_intersection(h1: Hailstone, h2: Hailstone, target: tuple[int, int]) -> tuple[int, int]:   
  t1, t2 = target
  a0 = h1.position[0]
  da = h1.velocity[0]
  b0 = h2.position[0]
  db = h2.velocity[0]
  if da == db:
    return False
    # return (a0, None) if a0 == b0 else (None, None)  # Same or parallel sequences
  x = (b0 - a0) / (da - db)
  y = a0 + da * x
  return True if t1 <= x <= t2 and x.is_integer() else False

# 89, 95 wrong
async def part1():
  hailstones = []
  
  # target = (7, 27)
  target = (200000000000000, 400000000000000)
  
  async for line in read_file("day24input"):
    h = Hailstone.parse(line)
    hailstones.append(h)

  count = 0
  i = 1
  for h1, h2 in itertools.combinations(hailstones, 2):
    i += 1

    if find_intersection(h1, h2, target):
      # print(f"{i} {h1} {h2}")
      count += 1

    # x, y = find_arithmetic_intersection(h1, h2)
    # if x is None or y is None:
    #   continue
    # print(x, y)
    # if (x <= target[0] and x <= target[1]) and (y <= target[0] and y <= target[1]):
    #   count += 1
    #   continue
      
    # h1coords = (h1.position[0], h1.position[1])
    # h2coords = (h2.position[0], h1.position[2])
    
    # yay = False
    # while True:
    #   h1coords = (h1coords[0] + h1.velocity[0], h1coords[1] + h1.velocity[1])
    #   h2coords = (h2coords[0] + h1.velocity[0], h2coords[1] + h2.velocity[1])
      
    #   print(h1coords, h2coords)

    #   candidates = [h1coords[0], h1coords[1], h2coords[0], h2coords[1]]
    #   if not all(target[0] <= v <= target[1] for v in candidates):
    #     break

    #   # print(h1coords, h2coords)
    #   if h1coords[0] == h2coords[0] or h1coords[1] == h2coords[1]:
    #     # print(h1, h2, h1coords, h2coords)
    #     yay = True
    #     break

    #   # valid = h1coords[0] <= target[0] <= target[1] and 
    #   # h1coords[1] <= 
    #   # # TODO this is not generic enough
    #   # if h1coords[0] > target[0] or h1coords[1] < target[1]:
    #   #   break


    # if yay:
    #   count += 1

  print(count)
    


async def part2():
  async for line in read_file("day24input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
