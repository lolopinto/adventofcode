from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
# from grid import Grid
import itertools
import enum
import math

async def part1():
  s = 0
  async for line in read_file("day3input"):
    v = re.findall("mul\((\d+,\d+)\)", line)
    s += sum([int(x) * int(y) for x, y in [x.split(",") for x in v]])
    
  print(s)


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    # asyncio.run(part2())
