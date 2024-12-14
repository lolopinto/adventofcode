from __future__ import annotations
from utils import read_file_groups, read_file, ints, ints_list
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math

# example
# WIDTH = 11
# HEIGHT = 7

WIDTH = 101
HEIGHT = 103

@dataclass
class Point:
  x: int 
  y: int
  vx: int
  vy: int
  

  def move(self):
    self.x += self.vx
      
    if self.x < 0:
      self.x = WIDTH + self.x
    if self.x >= WIDTH:  
      self.x = self.x - WIDTH
      
    self.y += self.vy


    if self.y < 0:
      self.y = HEIGHT + self.y
    if self.y >= HEIGHT:
      self.y = self.y - HEIGHT
    
async def part1():

  points = []
  async for line in read_file("day14input"):
    m = re.search("p=(.+) v=(.+)", line)
    assert m is not None, line
    
    assert len(m.groups()) == 2, line
    
    pos = ints_list(m.group(1).split(","))
    vel = ints_list(m.group(2).split(","))
    
    points.append(Point(x=pos[0], y=pos[1], vx=vel[0], vy=vel[1]))

  for i in range(100):
    for p in points:
      p.move()

  q1 = []
  q2 = []
  q3 = []
  q4 = []
    
  mid_x = math.floor(WIDTH / 2)
  mid_y = math.floor(HEIGHT / 2)
  
  for p in points:
    if p.x < mid_x and p.y < mid_y:
      q1.append(p)
    elif p.x > mid_x and p.y < mid_y:  
      q2.append(p)  
    elif p.x < mid_x and p.y > mid_y:
      q3.append(p)
    elif p.x > mid_x and p.y > mid_y:
      q4.append(p)


  print(len(q1) * len(q2)* len(q3) * len(q4))
  

async def part2():
  async for line in read_file("day14input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
