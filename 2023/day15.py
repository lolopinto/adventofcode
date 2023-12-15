from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def process(part: str) -> int:
  curr = 0
  for c in part:
    curr += ord(c)
    curr *= 17
    curr = curr % 256
  return curr

async def part1():
  async for line in read_file("day15input"):
    parts = line.split(",")
    print(sum([process(part) for part in parts]))

async def part2():
  async for line in read_file("day15input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
