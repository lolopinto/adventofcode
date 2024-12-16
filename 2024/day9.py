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

# @dataclass
# class File:
#   file_id: int
#   blocks: int

# @dataclass
# class FreeSpace:
#   amount: int

async def part1():
  lines = [line async for line in read_file("day9input")]
  assert len(lines) == 1
  
  items = []
  file_id = 0
  is_file = True
  
  l = []
  free_space_indices = []

  for c in lines[0]:
    v = int(c)
    if is_file:
      l.extend([file_id] * v)
      # items.append(File(blocks=v, file_id=file_id))
      file_id += 1
    else:
      start_idx = len(l)
      l.extend(["."] * v)
      for i in range(start_idx, v+start_idx):
        free_space_indices.append(i)
      # items.append(FreeSpace(amount=v))
      
    is_file = not is_file
  
  def swap(l, free_space_indices):
    free_idx = 0
    for i in range(len(l) -1, -1, -1):
      v = l[i]
      if v == ".":
        continue
      
      l_idx_swap = free_space_indices[free_idx]

      if l_idx_swap >= i:
        break
      
      l[l_idx_swap] = v
      
      free_idx += 1
      l[i] = "."

      if free_idx == len(free_space_indices):
        break
    
  swap(l, free_space_indices)

  while True:
    free_space_indices = []
    last_number_idx = -1
    
    done = True
    for idx, v in enumerate(l):
      if v == ".":
        free_space_indices.append(idx)
      else:
        if idx - last_number_idx != 1:
          done = False
        last_number_idx = idx

    if done:
      break

    swap(l, free_space_indices)
    
  s = 0
  for idx, v in enumerate(l):
    if v != ".":
      s += v * idx

  print(s)
  

async def part2():
  async for line in read_file("day9input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
