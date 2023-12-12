from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import itertools
from functools import cache

# changed from all combinations to something that calculates as we're going through
# inspired by something on reddit
@cache
def count_solutions(s: str, damaged: tuple[int], num_in_group: int) -> int:
  if len(s) == 0:
    return not damaged and not num_in_group
  
  ct = 0
  
  possibilities = [s[0]]
  if s[0] == '?':
    possibilities = ['.', '#']
  
  for c in possibilities:
    match c:
      case '#':
        ct += count_solutions(s[1:], damaged, num_in_group + 1)
      case _:
        if num_in_group and damaged and damaged[0] == num_in_group:
            ct += count_solutions(s[1:], damaged[1:], 0)
        elif not num_in_group:
          ct += count_solutions(s[1:], damaged, 0)
  return ct

async def part1():
  valids = []
  async for line in read_file("day12input"):
    parts = line.split()
    assert len(parts) == 2
    damaged = tuple([int(v) for v in parts[1].split(",")])
    ct = count_solutions(parts[0] + "_", damaged, 0)
    valids.append(ct)

  print(sum(valids))


def transform_line(line: str) -> str:
  parts = line.split()
  assert len(parts) == 2
  damaged = [int(v) for v in parts[1].split(",")]
  springs = '?'.join([parts[0]] * 5)
  damages = ",".join([parts[1]] * 5)
  return f"{springs} {damages}"


async def part2():
  valids = []
  async for line in read_file("day12input"):
    line = transform_line(line)
    parts = line.split()
    assert len(parts) == 2
    damaged = tuple([int(v) for v in parts[1].split(",")])
    ct = count_solutions(parts[0] + "_", damaged, 0)
    valids.append(ct)

  print(sum(valids))


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
