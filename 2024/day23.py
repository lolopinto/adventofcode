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


def check(b, c, connected):
  return b in connected[c]

def normalize(a, b, c):
  return tuple(sorted((a, b, c)))


async def part1():
  connected = defaultdict(set)
  async for line in read_file("day23input"):
    p0, p1 = line.split("-")
    connected[p0].add(p1)
    connected[p1].add(p0)

  candidates = set()
  for k, v in connected.items():
    # i initially used pairwise instead of permutations here :(
    for item1, item2 in itertools.permutations(v, 2):
      if check(item1, item2, connected):
        candidates.add(normalize(k, item1, item2))
  
  s = len([v for v in candidates if any(vv.startswith("t") for vv in v)])  
  
  print(s)

async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
