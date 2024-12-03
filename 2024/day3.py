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
  s = 0
  enabled = True
  first = True
  async for line in read_file("day3input"):
    dont_matches = [v.start() for v in re.finditer("don't()", line)]
    do_matches = [v.start() for v in re.finditer("do\(\)", line)]
    
    l = []
    for v in dont_matches:
      l.append((v, False))
    for v in do_matches:  
      l.append((v, True))
    for item in re.finditer("mul\(\d+,\d+\)", line):
      l.append((item.start(), item))
    
    l.sort(key=lambda x: x[0])    
    if first:
      l.insert(0, (-1, True))
      first = False
    
    for _, item in l:
      if item is True:
        enabled = True
        continue
      if item is False:
        enabled = False
        continue
      if not enabled:
        continue
      v = re.findall("mul\((\d+,\d+)\)", item.group())[0].split(",")
      s += int(v[0]) * int(v[1])
      
  print(s)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
