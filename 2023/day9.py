from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from itertools import pairwise
import math


async def part1():
  # ints fail, abs fail
  # and then i was returning the last item instead of the sum of it all. epic fail all over the place here
  
  candidates = []
  async for line in read_file("day9input"):
    l = ints(line)
    lists = [l]
    while True:
      next = []
      s = set()
      for a, b in pairwise(l):
        diff = b - a
        s.add(diff)
        next.append(diff)
      l = next
      lists.append(l)
      if len(s) == 1:
        break
    
    lasts = [l[-1] for l in lists]
    candidates.append(sum(lasts))

  print(sum(candidates))
    


async def part2():
  candidates = []
  async for line in read_file("day9input"):
    l = ints(line)
    lists = [l]
    while True:
      next = []
      s = set()
      for a, b in pairwise(l):
        diff = b - a
        s.add(diff)
        next.append(diff)
      l = next
      lists.append(l)
      if len(s) == 1:
        break
    
    ans = 0
    lists.reverse()
    for l in lists:
      ans = l[0] - ans 
    candidates.append(ans)
      
  print(sum(candidates))


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
