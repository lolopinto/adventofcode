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
class Button:
  x: int
  y: int
  
  @staticmethod
  def parse(s: str) -> Button:
    m = re.search("Button .: X\+(.+), Y\+(.+)", s)
    assert m is not None
    
    return Button(int(m.group(1)), int(m.group(2)))

@dataclass
class Prize:
  x: int
  y: int
  
  @staticmethod
  def parse(s: str) -> Prize:
    m = re.search("Prize: X=(.+), Y=(.+)", s)
    assert m is not None
    
    return Prize(int(m.group(1)), int(m.group(2)))

@dataclass
class ClawMachine:
  a: Button
  b: Button
  prize: Prize
  
  def tokens(self) -> int | None:
    costs = []
    for i in range(0, 100):
      for j in range(0, 100):
        if (i * self.a.x + j * self.b.x == self.prize.x) and (i * self.a.y + j * self.b.y == self.prize.y):
          costs.append(i * 3 + j)

    if len(costs) == 0:
      return None
    return min(costs)

async def part1():
  
  s = 0
  async for group in read_file_groups("day13input"):
    assert len(group) == 3
    
    a = Button.parse(group[0])
    b = Button.parse(group[1])
    prize = Prize.parse(group[2])
    
    cm = ClawMachine(a, b, prize)
    if cost := cm.tokens():
      s += cost

  print(s)


async def part2():
  async for line in read_file("day13input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
