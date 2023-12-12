from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import itertools

def replace_chars(s: str):
  # chat gpt helped write this :)
  char_to_replace = '?'
  replacements = ['#', '.']
  positions = [i for i, c in enumerate(s) if c == char_to_replace]
  combos = itertools.product(replacements, repeat=len(positions))

  result = []  
  for combo in combos:
    new_s = list(s)
    for pos, replacement in zip(positions, combo):
      new_s[pos] = replacement
    yield ''.join(new_s)
  return result

def is_valid(s: str, damaged: list[int]) -> bool:
  cts = []
  ct = 0
  for c in s:
    match c:
      case '#':
        ct += 1
      case '.':
        if ct > 0:
          cts.append(ct)
        ct = 0
      case _:
        assert False
  if ct > 0:
    cts.append(ct)

  return cts == damaged

async def part1():
  valids = []
  async for line in read_file("day12input"):
    parts = line.split()
    assert len(parts) == 2
    damaged = [int(v) for v in parts[1].split(",")]
    springs = parts[0]
    replacements = replace_chars(springs)
    valid = sum([1 for replacement in replacements if is_valid(replacement, damaged)])
    valids.append(valid)

  print(sum(valids))


async def part2():
  async for line in read_file("day12input"):
    pass






if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
