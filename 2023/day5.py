from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re

@dataclass
class Group:
  source: str
  destination: str
  items: list[list[int]]
  
  def get_next(self, val: int, map: dict[str, Group]) -> int:
    for v in self.items:
      assert len(v) == 3
      dest_start, source_start, length = v
      if val >= source_start and val < source_start + length:
        return dest_start + (val - source_start)
      
    return val

async def part1():
  seeds = []
  maps = {}
  curr = None
  
  async for line in read_file("day5input"):
    if line.startswith("seeds: "):
      seeds = ints(line.split(": ")[1])
    elif len(line) == 0:
      # reset here
      continue
    elif line.endswith(":"):
      match = re.match(r"(.+)-to-(.+) map:", line)
      assert match is not None
      source = match.group(1)
      dest = match.group(2)
      curr = Group(source, dest, [])
    else:
      assert curr is not None
      l = ints(line)
      assert len(l) == 3
      curr.items.append(ints(line))
      maps[curr.source] = curr

  locations = []
  for seed in seeds:
    curr_num = seed
    curr_val = "seed"
    while curr_val != "location":
      g = maps[curr_val]
      curr_num = g.get_next(curr_num, maps)
      curr_val = g.destination
    assert curr_num != -1
    locations.append(curr_num)      
    
  print(min(locations))

async def part2():
  async for line in read_file("day5input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
