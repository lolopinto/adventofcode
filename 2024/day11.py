from __future__ import annotations
from utils import read_file_groups, read_file, ints, read_file_lines
import asyncio
from collections import defaultdict, Counter
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math
from functools import lru_cache

def blink(stone):
  if stone == 0:
    return [1]
  
  s = str(stone)
  if len(s) % 2 == 0:
    
    mid = len(s) // 2
    left = s[:mid]
    right = s[mid:]
    
    return [int(left), int(right)]
  
  return [stone * 2024]

# does a count of each unique stone so that we only need to blink a stone once
# i hate aoc
def blinkCounter(stone_counts: dict[int, int]) -> dict[int, int]:
  result = Counter()
  
  for stone in stone_counts:
    for new_stone in blink(stone):
      result[new_stone] += stone_counts[stone]
  return result


async def part1():
  lines = await read_file_lines("day11input")
  assert len(lines) == 1
  
  stones = []
  for stone in ints(lines[0]):
    stones.append(stone)
  
  for i in range(25):
    new_stones = []
    for stone in stones:
      new_stones.extend(blink(stone))
    stones = new_stones
  print(len(stones))

async def part2():
  lines = await read_file_lines("day11input")
  assert len(lines) == 1
  
  stones = []
  for stone in ints(lines[0]):
    stones.append(stone)
  
  stone_counts = Counter(stones)
  for i in range(75):
    stone_counts = blinkCounter(stone_counts)

  print(sum(stone_counts.values()))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
