from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

def process_beam(g: Grid, d: str, curr: tuple[int, int], energized: set[tuple[int, int]]):
  # r, c = curr
  # if r < 0 or r >= g.height or c < 0 or c >= g.width:
  #   print(f'discarding beam @ {r,c} going in dir {d}')
  #   return
  # if curr in energized:
  #   print(f'already energized @ {curr}')
  #   return

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
      print(f'discarding beam @ {r,c} going in dir {d}')
      break

    if (r, c, d) in energized:
      print(f'already energized @ {r, c, d}')
      break

    energized.add((r, c, d))    

    v = g.get_value(r, c)
    # energize new spot 
    # energized.add((r2, c2))

    match v:
      case '.':
        # continues in same direction
        # print('continuing. hit empty space')
        # energized.add((r2, c2, d))
        pass
        # continue
      case '/':
        print(f'left mirror, changing direction from {d} -> {left_mirror[d]}')
        d = left_mirror[d]
        # energized.add((r2, c2))
        # pass
      case '\\':
        print(f'right mirror, changing direction from {d} -> {right_mirror[d]}')

        d = right_mirror[d]
        # energized.add((r2, c2))
        # continue
      case '|':
        # pointy end of splitter, continue
        if d in ['U', 'D']:
          print('pointy head of | splitter. continue')
          # energized.add((r2, c2))
          # continue
        else:
        # otherwise, split
          print(f'flat side of | splitter, splitting into U,D @ {r, c}')
          # energized.add((r2, c2, d))
          process_beam(g, 'U', (r, c), energized)
          process_beam(g, 'D', (r, c), energized)
          break

      case '-':
        if d in ['R', 'L']:
          print('pointy head of - splitter. continue')
          # energized.add((r2, c2))
          # continue
        else:
          print(f'flat side of - splitter, splitting into R,L @ {r, c}')

          # otherwise, split
          # if not (r2, c2) in energized:
          # energized.add((r2, c2, d))
          process_beam(g, 'R', (r, c), energized)
          process_beam(g, 'L', (r, c), energized)
          break
      case _:
        raise ValueError("unknown value {v} at {r}, {c}")

    delta = dirs[d]
    r2, c2 = r + delta[0], c + delta[1]

    print(f"curr {curr}, dir: {d}, delta: {delta}, potential:{(r2, c2)}")

    curr = (r2, c2)

async def part1():
  g = await Grid.grid_from_file("day16input")
  s = (0, 0)
  d = 'R'
  energized = set()
  process_beam(g, d, s, energized)
  
  print(len(energized))
  coords = set((r,c) for (r, c, d) in energized)
  # 125 too low
  # debug_print(g, coords)
  print(len(coords))
  
def debug_print(g: Grid, energized: set[tuple[int, int]]):
  g2 = Grid.grid(g.width, g.height)
  for r in range(g.height):
    for c in range(g.width):
      if (r, c) in energized:
        g2.set(r, c, '#')
      else:
        g2.set(r, c, '.')
  g2.print()
  

async def part2():
  async for line in read_file("day16input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())

