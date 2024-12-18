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

async def part1():
  g = await Grid.grid_from_file("day10input")
  
  trailheads = []
  endings = []
  for r, c in g.walk():
    if g.get_value(r, c) == "0":
      trailheads.append((r, c))
    elif g.get_value(r, c) == "9":
      endings.append((r, c))

  def valid(g, curr, next):
    return int(g.get_value(next[0], next[1])) - int(g.get_value(curr[0], curr[1])) == 1

  score = 0
  for th in trailheads:
    for e in endings:
      g2 = g.clone()
      try:
        g2.dijkstra2(th, e, valid)
        score += 1
      except AssertionError:
        pass

  print(score)

async def part2():
  g = await Grid.grid_from_file("day10input")
  
  trailheads = []
  endings = []
  for r, c in g.walk():
    if g.get_value(r, c) == "0":
      trailheads.append((r, c))
    elif g.get_value(r, c) == "9":
      endings.append((r, c))

  def valid(g, curr, next):
    return int(g.get_value(next[0], next[1])) - int(g.get_value(curr[0], curr[1])) == 1

  score = 0
  for th in trailheads:
    for e in endings:
      g2 = g.clone()
      try:
        all_paths = g2.all_paths(th, e, valid)
        score += len(all_paths)
      except AssertionError:
        pass

  print(score)

if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())
