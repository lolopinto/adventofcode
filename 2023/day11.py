from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
from itertools import combinations


async def build_grid(file: str):
  lines = []
  async for line in read_file("day11input"):
    no_galaxies = True
    for c in line:
      if c == '#':
        no_galaxies = False
        break

    lines.append(line)
    if no_galaxies:
      # print('row ', len(lines), ' has no galaxies')
      lines.append(line)
      
  cols = set()
  for c in range(len(lines[0])):
    no_galaxies = True
    for r in range(len(lines)):
      if lines[r][c] == '#':
        no_galaxies = False
        break
    if no_galaxies:
      cols.add(c)
      # print('column ', c, ' has no galaxies')

  outputs = []
  for line in lines:
    line2 = []
    for i in range(len(line)):
      line2.append(line[i])
      if i in cols:
        line2.append(line[i])
    outputs.append("".join(line2))

  g = Grid(len(outputs[0]), len(outputs))
  for r in range(len(outputs)):
    for c in range(len(outputs[0])):
      g.set(r, c, outputs[r][c])
  # print(len(outputs)), print(len(outputs[0]))
  return g

def distance(p1: tuple[int, int], p2: tuple[int, int]) -> int:
  return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

async def part1():
  g = await build_grid("day11input")
  # g.print()
  galaxies = []
  for r in range(g.height):
    for c in range(g.width):
      if g.get_value(r, c) == '#':
        galaxies.append((r, c))

  shortest = []
  # need to account for places that have been seen before
  # e.g. if you get to 5 and 6, then 6 and 5 should be 0
  
  # cache = defaultdict(factory=lambda: defaultdict(int))
  # cache = {}
  # def seen_min_before(start: tuple[int, int], end: tuple[int, int]) -> -1:
  #   if start in cache and end in cache[start]:
  #     return cache[start][end]
  #   return -1
  
  
  # we want a cache from any neighbors to any other neighbors
  # not just the points 

  ct = 0
  
  seen_ct = defaultdict(int)
  for p1, p2 in combinations(galaxies, 2):
    seen_ct[p2] += 1
    # if ct % 1000 == 0:
    #   print(ct)
    ct += 1
    # print(p1, p2)
    # short = 0
    # g2 = g.clone()
    # short = g2.dijkstra2(p1, p2)
    # short = g.dijkstra2(p1, p2)
    # cache[p1] = cache.get(p1, {})
    # cache[p1][p2] = short
    # # cache[p1][p2] = short
    # # cache[p2][p1] = short
    # cache[p2] = cache.get(p2, {})
    # cache[p2][p1] = short
    shortest.append(distance(p1, p2))

  print(sum(shortest))
  # print(seen_ct)

async def part2():
  async for line in read_file("day11input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
