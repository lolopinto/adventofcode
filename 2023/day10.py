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
      # find its fake neighbors
      fake_neighbors = [n for n in g.neighbors(r, c) if g.get_value(n[0], n[1]) != "."]
      neighbors = set()
      # for each of its fake neighbors, check their neighbors and confirm that S is one of them
      for n in fake_neighbors:
        n2 = get_neighbors(g, n[0], n[1])
        # check if S position is one of its neighbors' neighbors
        if (r,c) in n2:
          neighbors.add((n[0],n[1]))
      assert len(neighbors) == 2

    case '.':
      raise ValueError(f"visiting ground at {r}, {c}")
    case _:
      raise ValueError(f"unknown value {val} at {r}, {c}")
    
  if len(neighbors) != 2:
    print(neighbors, val, (r,c))
    raise ValueError(f"expected 2 neighbors, got {len(neighbors)} at {r}, {c}, value: {val}")

  return neighbors

def get_distances(g: Grid):
  start = g.find("S")
  assert start is not None
  
  q = []
  q.append(start)
  distances = defaultdict(int)
  distances[start] = 0
  g.visit(start[0], start[1])
  
  while len(q) > 0:
    (r, c) = q.pop(0)
    
    distance = distances[(r, c)]
    neighbors = get_neighbors(g, r, c)
    
    for r2,c2 in neighbors:
      if g.visited(r2, c2):
        continue
      g.visit(r2, c2)
      q.append((r2, c2))
      distances[(r2, c2)] = distance + 1
  return distances

async def part1():
  g = await Grid.square_grid_from_file("day10input")

  distances = get_distances(g)
  print(max(distances.values()))

# for part 2, i mostly had to use reddit and others solutions to get this to work
# uses https://en.wikipedia.org/wiki/Point_in_polygon#Ray_casting_algorithm
# https://www.reddit.com/r/adventofcode/comments/18evyu9/comment/kcqtow6/ amongst others was helpful in getting this to work
# builds on top of existing solution for part 1 (get_distances) 
async def part2():
  g = await Grid.square_grid_from_file("day10input")

  distances = get_distances(g)

  inside = 0
  for r in range(g.height):
    for c in range(g.width):
      if (r,c) in distances:
        continue

      r2, c2 = r, c
      crosses = 0

      while r2 < g.height and c2 < g.width:
        v = g.get_value(r2, c2)
        if (r2, c2) in distances and v != "L" and v != '7':
          crosses += 1
        r2 += 1
        c2 += 1

      if crosses % 2 == 1:
        inside += 1

  print(inside)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
