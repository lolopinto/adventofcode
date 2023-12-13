from __future__ import annotations
from utils import read_file_groups, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def get_lines_for_row(g: Grid, r1: int, r2: int) -> (str, str):
  line1 = []
  line2 = []
  
  for c in range(g.width):
    line1.append(g.get_value(r1, c))
    line2.append(g.get_value(r2, c))
  return ("".join(line1), "".join(line2))

def get_lines_for_col(g: Grid, c1: int, c2: int) -> (str, str):
  line1 = []
  line2 = []
  
  for r in range(g.height):
    line1.append(g.get_value(r, c1))
    line2.append(g.get_value(r, c2))
  return ("".join(line1), "".join(line2))

def find_reflection(g: Grid):
  
  # using grid instead of lines so that we can flip it vertially
  for (r1, r2) in itertools.pairwise(range(g.height)):
    line1, line2 = get_lines_for_row(g, r1, r2)
    if line1 != line2:
      continue
    i = 1
    equal = True
    while True:
      rr1 = r1 - i
      rr2 = r2 + i
      if rr1 < 0 or rr2 >= g.height:
        break
      linea, lineb = get_lines_for_row(g, rr1, rr2)
      if linea != lineb:
        equal = False
        break
      i += 1
    if equal:
      print(f"found reflection at row {r1}, {r2}")
      return 100 * (r1 + 1)

  for (c1, c2) in itertools.pairwise(range(g.width)):
    line1, line2 = get_lines_for_col(g, c1, c2)
    if line1 != line2:
      continue
    i = 1
    equal = True
    while True:
      cc1 = c1 - i
      cc2 = c2 + i
      if cc1 < 0 or cc2 >= g.width:
        break
      linea, lineb = get_lines_for_col(g, cc1, cc2)
      if linea != lineb:
        equal = False
        break
      i += 1
    if equal:
      print(f"found reflection at col {c1}, {c2}")
      return c1 + 1

  raise ValueError("no reflection found")

async def part1():
  result = 0
  async for lines in read_file_groups("day13input"):
    g = Grid.from_lines(lines)    
    result += find_reflection(g)
  print(result)


async def part2():
  async for line in read_file_groups("day13input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
