from __future__ import annotations
from utils import read_file_groups, read_file, ints, read_file_lines
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math

@dataclass
class Stone:
  mark: int
  
  def blink(self):
    if self.mark == 0:
      return [Stone(mark=1)]
    
    s = str(self.mark)
    if len(s) % 2 == 0:
      
      mid = len(s) // 2
      left = s[:mid]
      right = s[mid:]
      
      return [Stone(mark=int(left)), Stone(mark=int(right))]
    
    return [Stone(mark=self.mark * 2024)]
    

async def part1():
  lines = await read_file_lines("day11input")
  assert len(lines) == 1
  
  stones = []
  for stone in ints(lines[0]):
    stones.append(Stone(mark=stone))
    
  # print(len(stones))

  for i in range(25):
    new_stones = []
    for stone in stones:
      new_stones.extend(stone.blink())
    stones = new_stones
    # print(stones)
  print(len(stones))


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
