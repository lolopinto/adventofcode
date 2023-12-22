from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

@dataclass
class Brick:
  start: tuple[int, int, int]
  end: tuple[int, int, int]
  
  @staticmethod
  def parse(line: str) -> Brick:
    parts = line.split('~')
    assert len(parts) == 2
    start = tuple(int(v) for v in parts[0].split(','))
    end = tuple(int(v) for v in parts[1].split(','))
    assert len(start) == 3
    assert len(end) == 3
    assert end[0] >= start[0]
    assert end[1] >= start[1]
    assert end[2] >= start[2]
    return Brick(start, end)
  
  def clone(self) -> Brick: 
    return Brick(self.start, self.end)

  def positions(self) -> set[tuple[int, int, int]]:
    s = {self.start, self.end}
    if self.start[0] != self.end[0]:
      for x in range(self.start[0], self.end[0] + 1):
        s.add((x, self.start[1], self.start[2]))

    if self.start[1] != self.end[1]:
      for y in range(self.start[1], self.end[1] + 1):
        s.add((self.start[0], y, self.start[2]))

    if self.start[2] != self.end[2]:
      for z in range(self.start[2], self.end[2] + 1):
        s.add((self.start[0], self.start[1], z))

    return s
  
  def settle(self, positions: set[tuple[int, int, int]]):
    # get current positions
    our = self.positions()
    # z = self.start[0]
    
    start = self.start
    end = self.end
    
    # create dupe positions and remove current brick from being checked
    dupe = positions.copy()
    for p in our:
      dupe.remove(p)    

    # go as far as the ground
    prev_start = None
    prev_end = None
    while start[2] != 1:
      # print(f'brick {self} at start {start}')
      start = (start[0], start[1], start[2] - 1)
      end = (end[0], end[1] , end[2] - 1)
      
      b2 = Brick(start, end)
      # check every position in new possible brick to see 
      
      break_outer = False
      for p in b2.positions():
        if p in dupe:
          if prev_start is None:
            # print(f'pos {p} already taken so {self} not settling at {start} {end}')
            return
          break_outer = True
          break

      if break_outer:
        break
         
      if start[2] < 1:
        print(f'{self} broken {start} {end}')
        break
      
      prev_start = start
      prev_end = end
  
    # nothing to do here    
    if prev_start == self.start and prev_end == self.end or prev_start is None or prev_end is None:
      return
    
    # print(f'brick {self} settling at new locations: {prev_start} {prev_end}')
    
    for p in our:
      positions.remove(p)
      
    self.start = prev_start
    self.end = prev_end

    for p in self.positions():
      positions.add(p)


  def disintegrate(self, positions: set[tuple[int, int, int]]):
    for p in self.positions():
      positions.remove(p)

async def parse_and_sort():
  bricks = []

  positions = set()
  async for line in read_file("day22input"):
    b = Brick.parse(line)
    for p in b.positions():
      if p in positions:
        raise ValueError(f'already taken {p} for brick {b}')
      positions.add(p)

    bricks.append(b)
    

  # sort bricks
  bricks.sort(key=lambda b: b.start[2])
  return bricks, positions

async def part1():
  bricks, positions = await parse_and_sort()

  # sort bricks, then settle them
  for b in bricks:
    b.settle(positions)

  ct = 0
  for b in bricks:
    dupe = positions.copy()
    for p in b.positions():
      dupe.remove(p)
    changed = False
    
    # check other bricks and see if there's any settling after removing current
    for b2 in bricks:
      if b == b2:
        continue
      
      b3 = b2.clone()
      b3.settle(dupe)
      if b2.start != b3.start or b2.end != b3.end:
        changed = True
        break
    
    if not changed:
      ct += 1
      
  print(ct)    


async def part2():
  bricks, positions = await parse_and_sort()

  # sort bricks, then settle them
  for b in bricks:
    b.settle(positions)

  values = []
  for b in bricks:
    dupe = positions.copy()

    # disintegrate this brick
    b.disintegrate(dupe)
    
    count_changed = 0 
    for b2 in bricks:
      if b == b2:
        continue
        
      b3 = b2.clone()
      b3.settle(dupe)
      if b2.start != b3.start or b2.end != b3.end:
        count_changed += 1
    
    values.append(count_changed)
    
  print(sum(values))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
