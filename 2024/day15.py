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
  groups = [g async for g in read_file_groups("day15input")]
  assert len(groups) == 2
  
  g = Grid.from_lines(groups[0])
  
  directions = "".join(groups[1])
  curr = g.find("@")
  num_boxes = g.count("O")
  assert curr is not None
  
  # g.print()

  candidates = {
    "^": g.top,
    "<": g.left,
    ">": g.right,
    "v": g.bottom,
  }

  for dir_idx, curr_dir in enumerate(directions):
    positions = candidates[curr_dir](curr[0], curr[1])

    new_positions = []
    seen_ball = False
    last = "@"
    end_with_wall = False
    seen_empty = False

    for idx, pos in enumerate(positions):
      val = g.get_value(pos[0], pos[1])

      if val == "#":
        end_with_wall = True
        break
      
      new_positions.append((last, pos))

      if val == ".":
        seen_empty = True
        break

      last = val

    if len(new_positions) == 0:
      # print("exiting because nothing to move")
      continue
    
    if end_with_wall and not seen_empty:
      continue
    

    g.set(curr[0], curr[1], ".")
    for val, pos in new_positions:
      g.set(pos[0], pos[1], val)

    curr = new_positions[0][1]

  # g.print()
  
  s = 0
  for r, c in g.walk():
    v = g.get_value(r, c)
    if v == "O":
      s += 100 * r + c
      
  print(s)
      

async def part2():
  async for line in read_file("day15input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
