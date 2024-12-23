from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math


def check(b, c, connected):
  return b in connected[c]

def normalize(a, b, c):
  return tuple(sorted((a, b, c)))

def normalize_list(t, i):
  return tuple(sorted((t + (i,))))

def get_lan3(connected):
  candidates = set()
  for k, v in connected.items():
    # i initially used pairwise instead of permutations here :(
    for item1, item2 in itertools.permutations(v, 2):
      if check(item1, item2, connected):
        candidates.add(normalize(k, item1, item2))
  
  return candidates

async def part1():
  connected = defaultdict(set)
  async for line in read_file("day23input"):
    p0, p1 = line.split("-")
    connected[p0].add(p1)
    connected[p1].add(p0)

  candidates = get_lan3(connected)

  s = len([v for v in candidates if any(vv.startswith("t") for vv in v)])  
  
  print(s)

async def part2():
  connected = defaultdict(set)
  async for line in read_file("day23input"):
    p0, p1 = line.split("-")
    connected[p0].add(p1)
    connected[p1].add(p0)

  # while we can add, keep looping
  # start from the answer of part 1 and keep looping to add items if possible
  
  # there's probably a smarter way to do this lol
  # also networkx again will work here
  
  # brute force. part 1 + part 2 took 175s to run
  candidates = get_lan3(connected)

  while True:

    to_add = set()
    for k in connected.keys():
      for candidate in candidates:
        if k in candidate:
          continue

        all_connected = True
        for item in candidate:
          if not check(k, item, connected):
            all_connected = False
            break

        if all_connected:
          new_val = normalize_list(candidate, k)
          if new_val not in candidates and new_val not in to_add:
            to_add.add(new_val)

    if not to_add:
      break
    for item in to_add:
      candidates.add(item)
      
  biggest = None
  for k in candidates:
    if biggest is None:
      biggest = k
    elif len(k) > len(biggest):
      biggest = k
      
  print(",".join(biggest))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
