from __future__ import annotations
from utils import read_file_groups, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def get_diff(g: Grid, r1: int, r2: int) -> (int):
  return sum([1 for c in range(g.width) if g.get_value(r1, c) != g.get_value(r2, c)])

def find_reflection_row(g: Grid, smudge_factor: int = 0) -> int | None:
  for (r1, r2) in itertools.pairwise(range(g.height)):
    total_diff = 0
    diff = get_diff(g, r1, r2)
    total_diff += diff
    i = 1
    while True:
      rr1 = r1 - i
      rr2 = r2 + i
      if rr1 < 0 or rr2 >= g.height:
        break
      diff = get_diff(g, rr1, rr2)
      total_diff += diff
      i += 1

    if total_diff == smudge_factor:
      return r1
  return None
async def part1():
  result = 0
  async for lines in read_file_groups("day13input"):
    g = Grid.from_lines(lines)
    if (h := find_reflection_row(g)) != None:
      result += (100 * (h + 1))
      continue
    g2 = g.rotate_left()
    if (h := find_reflection_row(g2)) != None:
      result += (h + 1)
      continue
    raise ValueError("no reflection found")
  print(result)


async def part2():
  result = 0
  async for lines in read_file_groups("day13input"):
    g = Grid.from_lines(lines)
    # walrus operator with a check tricky!
    if (h := find_reflection_row(g, 1)) != None:
      result += (100 * (h + 1))
      continue
    g2 = g.rotate_left()
    if (h := find_reflection_row(g2, 1)) != None:
      result += (h + 1)
      continue
    raise ValueError("no reflection found")
  print(result)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
