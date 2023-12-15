from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def process(g: Grid, dir: str):
  dirs = {
    'N': g.bottom,
    'W': g.right,
    'S': g.top,
    'E': g.left
  }
  reverse = {
    'S': True,
    'E': True,
    'N': False,
    'W': False,
  }
  
  width = range(g.width)
  height = range(g.height)
  if reverse[dir]:
    width = list(reversed(width))
    height = list(reversed(height))

  # for each column
  for c in width:
    for r in height:
      # only swapping .
      if g.get_value(r, c) != '.': 
        continue

      for r2, c2 in dirs[dir](r, c):
        # look for the next rock
        match g.get_value(r2, c2):
          case '#':
            break
          case 'O':
            # found a rock, swap 
            g.set(r2, c2, '.')
            g.set(r, c, 'O')
            break
  
def result(g: Grid) -> int:
  res = 0  
  for r in range(g.height):
    row = g.height - r
    for c in range(g.width):
      if g.get_value(r, c) == 'O':
        res += row
  return res

async def part1():
  g = await Grid.grid_from_file("day14input")
  
  process(g, 'N')

  print(result(g))


async def part2():
  g = await Grid.grid_from_file("day14input")
  
  lines = defaultdict(list)
  
  cycle = 1
  cycle_length = None
  while True:
    for dir in 'NWSE':
      process(g, dir)

    # post cycle
    key = tuple(g.current_lines())

    lines[key].append(cycle)
    if len(lines[key]) == 2:
      cycle_length = lines[key][1] - lines[key][0]
      break
    
    cycle += 1
    
  target_idx = 1_000_000_000 % cycle_length  
  # print(cycle, cycle_length, target_idx)

  answer = {k: v for k, v in lines.items() if v[-1] % cycle_length == target_idx}

  g2 = Grid.from_lines(list(answer.keys())[-1])

  print(result(g2))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
