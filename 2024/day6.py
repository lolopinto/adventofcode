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

async def part1():
  g = await Grid.grid_from_file("day6input")
  curr = None
  
  curr = g.find("^")
  assert curr is not None

  g.visit(curr[0], curr[1])
  
  direction = "up"

  while True:
    match direction:
      case "up":
        new_pos = curr[0] - 1, curr[1]
        if new_pos[0] < 0:
          break
        if g.get_value(new_pos[0], new_pos[1]) == "#":
          # go right
          direction = "right"
        else:
          curr = new_pos
          g.visit(curr[0], curr[1])
          
      case "right":
        new_pos = curr[0], curr[1] + 1
        if new_pos[1] >= g.width:
          break
        if g.get_value(new_pos[0], new_pos[1]) == "#":
          # go down
          direction = "down"
        else:
          curr = new_pos
          g.visit(curr[0], curr[1])
          
      case "down":
        new_pos = curr[0] + 1, curr[1]
        if new_pos[0] >= g.height:
          break
        if g.get_value(new_pos[0], new_pos[1]) == "#":
          # go left
          direction = "left"
        else:
          curr = new_pos
          g.visit(curr[0], curr[1])
          
      case "left":
        new_pos = curr[0], curr[1] - 1
        if new_pos[1] < 0:
          break
        if g.get_value(new_pos[0], new_pos[1]) == "#":
          # go up
          direction = "up"
        else:
          curr = new_pos
          g.visit(curr[0], curr[1])
          
        
  print(g.count_visited())


async def part2():
  async for line in read_file("day6input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
