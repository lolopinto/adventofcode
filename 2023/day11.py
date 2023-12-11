from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
from itertools import combinations

@dataclass
class RemappedGrid:
  g: Grid[str]
  expanded_rows: set[int]
  expanded_cols: set[int]
  factor: int
  
  def remap_index(self, r: int, c: int) -> tuple[int, int]:
    add_r = 0
    for row in self.expanded_rows:
      if row < r:
        add_r += self.factor - 1
    add_c = 0
    for col in self.expanded_cols:
      if col < c:
        add_c += self.factor - 1    
    return (r + add_r, c + add_c)

# exploding the grid not working
# need to remap the points to the new grid
async def build_grid(file: str, expand=2):
  lines = []
  r = 0
  rows = set()
  async for line in read_file("day11input"):
    no_galaxies = True
    for c in line:
      if c == '#':
        no_galaxies = False
        break

    lines.append(line)
    if no_galaxies:
      rows.add(r)
    r += 1

  cols = set()
  for c in range(len(lines[0])):
    no_galaxies = True
    for r in range(len(lines)):
      if lines[r][c] == '#':
        no_galaxies = False
        break
    if no_galaxies:
      cols.add(c)

  g = Grid.square_grid(len(lines))
  for r in range(g.height):
    for c in range(g.width):
      g.set(r, c, lines[r][c])
  
  return RemappedGrid(g, rows, cols, expand)

def distance(p1: tuple[int, int], p2: tuple[int, int]) -> int:
  return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

async def part1():
  remapped = await build_grid("day11input")
  g = remapped.g
  galaxies = []
  for r in range(g.height):
    for c in range(g.width):
      if g.get_value(r, c) == '#':
        r2, c2 = remapped.remap_index(r, c)
        galaxies.append((r2, c2))

  shortest = []
  
  for p1, p2 in combinations(galaxies, 2):
    shortest.append(distance(p1, p2))

  print(sum(shortest))

async def part2():
  remapped = await build_grid("day11input", expand=1000000)
  g = remapped.g
  galaxies = []
  for r in range(g.height):
    for c in range(g.width):
      if g.get_value(r, c) == '#':
        r2, c2 = remapped.remap_index(r, c)
        galaxies.append((r2, c2))

  shortest = []
  
  for p1, p2 in combinations(galaxies, 2):
    shortest.append(distance(p1, p2))

  print(sum(shortest))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
