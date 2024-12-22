from __future__ import annotations
from utils import read_file_groups, read_file, ints, get_file_groups
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math
from functools import cache


def is_possible(towels: list[str], line: str) -> bool:
  if len(line) == 0:
    return True

  for towel in towels:
    if line.startswith(towel):
      if is_possible(towels, line[len(towel):]):
        return True

  return False


@cache
def is_possible_count(towels: tuple[str], line: str) -> int:
  if len(line) == 0:
    return 1

  ret = 0
  for towel in towels:
    if line.startswith(towel):
      ret += is_possible_count(towels, line[len(towel):])

  return ret

async def part1():
  groups = await get_file_groups("day19input")
  assert len(groups) == 2
  
  assert len(groups[0]) == 1
  towels = groups[0][0].split(", ")
  
  ct = 0
  for line in groups[1]:
    if is_possible(towels, line):
      ct += 1

  print(ct)


async def part2():
  groups = await get_file_groups("day19input")
  assert len(groups) == 2
  
  assert len(groups[0]) == 1
  towels = groups[0][0].split(", ")
  
  ct = 0
  for line in groups[1]:
    ct += is_possible_count(tuple(towels), line)

  print(ct)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
