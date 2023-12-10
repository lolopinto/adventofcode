from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import math
from grid import Grid


def get_neighbors(g: Grid, r: int, c: int):
  val = g.get_value(r, c)
  # neighbors = g.neighbors(r, c)
  # neighbors = [n for n in neighbors if g.get_value(n[0], n[1]) != "."]

  # # assume nothing to do here
  # if len(neighbors) == 2:
  #   return neighbors

  match val:
    case '|':
      # same column
      neighbors = [n for n in g.neighbors(r, c) if n[1] == c]
    case '-':
      # same row
      neighbors = [n for n in g.neighbors(r, c) if n[0] == r]
    case 'L':
      # north and east 
      neighbors = [n for n in g.right_and_up(r, c) if g.get_value(n[0], n[1]) != "."]
    case 'J':
      # north and west      
      neighbors = [n for n in g.left_and_up(r, c) if g.get_value(n[0], n[1]) != "."]
    case '7':
      # south and west
      neighbors = [n for n in g.left_and_down(r, c) if g.get_value(n[0], n[1]) != "."]
    case 'F':
      neighbors = [n for n in g.right_and_down(r, c) if g.get_value(n[0], n[1]) != "."]
    case "S":
      # assume this is correct
      neighbors = [n for n in g.neighbors(r, c) if g.get_value(n[0], n[1]) != "."]
    case '.':
      raise ValueError(f"visiting ground at {r}, {c}")
    case _:
      raise ValueError(f"unknown value {val} at {r}, {c}")
    
  if len(neighbors) != 2:
    print(len(neighbors), neighbors, val)
    raise ValueError(f"expected 2 neighbors, got {len(neighbors)} at {r}, {c}")

  return neighbors

async def part1():
  g = await Grid.square_grid_from_file("day10input")

  # g.print()
  start = None

  for r in range(g.height):
    for c in range(g.width):
      if g.get_value(r, c) == "S":
        start = (r, c)
        break

  assert start is not None
  
  # distances = defaultdict(int)
  q = []
  q.append(start)
  # distances[start] = 0
  g2 = Grid.square_grid(g.width)
  g2.set(start[0], start[1], 0)
  g.visit(start[0], start[1])
  
  while len(q) > 0:
    (r, c) = q.pop(0)
    
    distance = g2.get_value(r, c) if g2.get_value(r, c) is not None else 0
    # neighbors without ground
    neighbors = get_neighbors(g, r, c)
    # neighbors = [n for n in neighbors if g.get_value(n[0], n[1]) != "."]
    # print(len(neighbors), neighbors, g.get_value(r, c))
    # if len(neighbors) !=2:
    #   print(len(neighbors), neighbors, g.get_value(r, c))
    
    for n in neighbors:
      r2, c2 = n
      if g.visited(r2, c2):
        continue
      g.visit(r2, c2)
      q.append(n)
      g2.set(r2, c2, distance + 1)
      # distances[n] = distance + 1

  # g2.print()
  print(max(g2.get_value(r,c) for r in range(g2.height) for c in range(g2.width) if g2.get_value(r, c) != None))

async def part2():
  async for line in read_file("day10input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
