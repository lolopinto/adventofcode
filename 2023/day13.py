from __future__ import annotations
from utils import read_file_groups, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def get_diff_for_row(g: Grid, r1: int, r2: int) -> (int):
  ct = 0  
  for c in range(g.width):
    val1 = g.get_value(r1, c)
    val2 = g.get_value(r2, c)
    if val1 != val2:
      ct += 1
  return ct

def get_diff_for_col(g: Grid, c1: int, c2: int) -> int: 
  ct = 0  
  for r in range(g.height):
    val1 = g.get_value(r, c1)
    val2 = g.get_value(r, c2)
    if val1 != val2:
      ct += 1
  return ct

def find_reflection(g: Grid, smudge_factor: int = 0) -> int:
  
  # using grid instead of lines so that we can flip it vertially
  for (r1, r2) in itertools.pairwise(range(g.height)):
    total_diff = 0
    diff = get_diff_for_row(g, r1, r2)
    total_diff += diff
    i = 1
    while True:
      rr1 = r1 - i
      rr2 = r2 + i
      if rr1 < 0 or rr2 >= g.height:
        break
      diff = get_diff_for_row(g, rr1, rr2)
      total_diff += diff
      i += 1

    if total_diff == smudge_factor:
      return 100 * (r1 + 1)

  for (c1, c2) in itertools.pairwise(range(g.width)):
    total_diff = 0
    diff = get_diff_for_col(g, c1, c2)
    total_diff += diff
    i = 1
    while True:
      cc1 = c1 - i
      cc2 = c2 + i
      if cc1 < 0 or cc2 >= g.width:
        break
      diff = get_diff_for_col(g, cc1, cc2)
      total_diff += diff
      i += 1

    if total_diff == smudge_factor:
      return c1 + 1

  raise ValueError("no reflection found")

async def part1():
  result = 0
  async for lines in read_file_groups("day13input"):
    g = Grid.from_lines(lines)    
    result += find_reflection(g)
  print(result)


async def part2():
  result = 0
  # this was tougher than i expected. tried to do a complicated thing with
  # checking previous reflections and it wasn't working
  # brute force didn't work either. I wonder what I was doing wrong
  async for lines in read_file_groups("day13input"):
    g = Grid.from_lines(lines)    
    result += find_reflection(g, 1)
  print(result)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
