from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def process_beam(g: Grid, d: str, curr: tuple[int, int], energized: set[tuple[int, int]]):
  dirs = {
    'R': (0, 1),
    'L': (0, -1),
    'U': (-1, 0),
    'D': (1, 0),
  }
  left_mirror = {  # '/' mirror
      'R': 'U',
      'U': 'R',
      'L': 'D',
      'D': 'L',
  }
  right_mirror = {  # '/' mirror
      'R': 'D',
      'D': 'R',
      'L': 'U',
      'U': 'L',
  }

  while True:
    r, c = curr
    if r < 0 or r >= g.height or c < 0 or c >= g.width:
      break

    if (r, c, d) in energized:
      break

    energized.add((r, c, d))    

    v = g.get_value(r, c)

    match v:
      case '.':
        # continues in same direction
        pass

      case '/':
        d = left_mirror[d]

      case '\\':
        d = right_mirror[d]

      case '|':
        # flat side of splitter, split
        if d not in ['U', 'D']:
          process_beam(g, 'U', (r, c), energized)
          process_beam(g, 'D', (r, c), energized)
          break

      case '-':
        # flat side of splitter, split
        if d not in ['R', 'L']:
          process_beam(g, 'R', (r, c), energized)
          process_beam(g, 'L', (r, c), energized)
          break
      case _:
        raise ValueError("unknown value {v} at {r}, {c}")

    delta = dirs[d]
    r2, c2 = r + delta[0], c + delta[1]

    curr = (r2, c2)


def do_work(g: Grid, s: tuple[int, int], d: s)-> int:
  energized = set()
  process_beam(g, d, s, energized)
  
  coords = set((r,c) for (r, c, d) in energized)
  return len(coords)

def debug_print(g: Grid, energized: set[tuple[int, int]]):
  g2 = Grid.grid(g.width, g.height)
  for r in range(g.height):
    for c in range(g.width):
      if (r, c) in energized:
        g2.set(r, c, '#')
      else:
        g2.set(r, c, '.')
  g2.print()
  
async def part1():
  g = await Grid.grid_from_file("day16input")

  print(do_work(g, (0, 0), 'R'))


async def part2():
  g = await Grid.grid_from_file("day16input")

  vals = []
  for r in range(g.height):
    vals.append(do_work(g, (0, r), 'D'))
    vals.append(do_work(g, (g.width-1, r), 'U'))

  for c in range(g.width):
    vals.append(do_work(g, (c, 0), 'R'))
    vals.append(do_work(g, (g.height-1, r), 'L'))
    
  print(max(vals))


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
