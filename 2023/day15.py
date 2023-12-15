from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

@dataclass
class Info:
  label: str
  dash: bool
  equal: bool
  focal: int

def parse(s: str) -> Info:
  idx = s.find('=')
  if idx != -1:
    return Info(s[:idx], False, True, int(s[idx+1:]))
  else:
    idx = s.index('-')
    return Info(s[:idx], True, False, 0)

def hash(part: str) -> int:
  curr = 0
  for c in part:
    curr += ord(c)
    curr *= 17
    curr = curr % 256
  return curr

async def part1():
  async for line in read_file("day15input"):
    parts = line.split(",")
    print(sum([hash(part) for part in parts]))

async def part2():
  boxes = defaultdict(list)
  async for line in read_file("day15input"):
    parts = line.split(",")
    for part in parts:
      info = parse(part)

      box = hash(info.label)
      
      if info.dash:
        for idx in range(len(boxes[box])):
          item = boxes[box][idx]
          if item.label == info.label:
            boxes[box].pop(idx)
            break
      else:
        found = False
        for idx in range(len(boxes[box])):
          item = boxes[box][idx]
          if item.label == info.label:
            found = True
            # replace
            boxes[box][idx] = info
            break

        if not found:
          boxes[box].append(info)

  res = 0
  for k, l in boxes.items():
    for idx in range(len(l)):
      res += ((k + 1) * (idx + 1) * l[idx].focal)
  print(res)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
