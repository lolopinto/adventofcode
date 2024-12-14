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
  groups = [group async for group in read_file_groups("day5input")]
  assert len(groups) == 2
  
  before_orderings = defaultdict(set)
  after_orderings = defaultdict(set)
  
  for line in groups[0]:
    before, after = ints(line, "|")
    before_orderings[before].add(after)
    after_orderings[after].add(before)
    
  s = 0
  for line in groups[1]:
    l = ints(line, ",")
    
    valid = True
    for i in range(len(l)):
      for j in range(i, len(l)):
        if i == j:
          continue

        first = l[i]
        second = l[j]
        
        if second in before_orderings[first]:
          continue
        
        if first in after_orderings[second]:
          continue

        valid = False
        break
      
      if not valid:
        break

    if valid:
      mid = len(l)// 2
      s += l[mid]

  print(s)

async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
