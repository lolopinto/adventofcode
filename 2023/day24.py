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

@dataclass
class Hailstone:
  position: tuple[int, int, int]
  velocity: tuple[int, int, int]
  
  @staticmethod
  def parse_parts(part: str) -> tuple[int, int, int]:
    return tuple(int(v) for v in part.split(', '))

  @staticmethod
  def parse(line: str) -> Hailstone:
    parts = line.split(' @ ')
    assert len(parts) == 2
    position = Hailstone.parse_parts(parts[0])
    velocity = Hailstone.parse_parts(parts[1])
    return Hailstone(position, velocity)
  
  # this is using algebra and zamore's solution https://github.com/fzamore/advent-of-code/blob/main/2023/day24.py
  # https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection?fbclid=IwAR1ZEsxYuSE7otIZDhvmLtslgkxTIK6T8OWHctD3pfmFEEutkAmkrGwNAfo#Given_two_points_on_each_line
  @staticmethod
  def intersect_within(h1: Hailstone, h2: Hailstone, target: tuple[int, int]) -> bool:
    x1, y1, _ = h1.position
    x2, y2, _ = h2.position
    dx1, dy1, _ = h1.velocity
    dx2, dy2, _ = h2.velocity
        
    # parallel    
    if dx1 * dy2 == dx2 * dy1:
      return False
  
    b1 = y1 - x1 * dy1 / dx1
    b2 = y2 - x2 * dy2 / dx2
    
    xi = (b2 - b1) / (dy1 / dx1 - dy2 / dx2)
    yi = xi * dy1 / dx1 + b1
    
    intersect = (xi, yi)
    def in_past(h: Hailstone) -> bool:
      x, y, _ = h.position
      dx, dy, _ = h.velocity
      return (xi - x) / dx < 0 or (yi - y) / dy < 0
    
    if in_past(h1) or in_past(h2):
      return False
    
    return target[0] <= xi <= target[1] and target[0] <= yi <= target[1]


async def part1():
  hailstones = []
  
  # target = (7, 27)
  target = (200000000000000, 400000000000000)
  
  async for line in read_file("day24input"):
    h = Hailstone.parse(line)
    hailstones.append(h)

  count = 0
  for h1, h2 in itertools.combinations(hailstones, 2):
    if Hailstone.intersect_within(h1, h2, target):
      count += 1

  print(count)
    


async def part2():
  async for line in read_file("day24input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
