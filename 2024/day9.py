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
class File:
  file_id: int
  blocks: int

@dataclass
class FreeSpace:
  amount: int
 

async def part1():
  lines = [line async for line in read_file("day9input")]
  assert len(lines) == 1
  
  items = []
  file_id = 0
  is_file = True
  
  l = []

  for c in lines[0]:
    v = int(c)
    if is_file:
      l.extend([file_id] * v)
      file_id += 1
    else:
      l.extend(["."] * v)
      
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
  lines = [line async for line in read_file("day9input")]
  assert len(lines) == 1
  
  items = []
  file_id = 0
  is_file = True
  last_file_id = None
  
  for c in lines[0]:
    v = int(c)
    if is_file:
      items.append(File(blocks=v, file_id=file_id))
      last_file_id = file_id
      file_id += 1
    elif v > 0:
      items.append(FreeSpace(amount=v))
    is_file = not is_file
  
  
  def print_items():
    line = ""
    for item in items:
      if isinstance(item, FreeSpace):
        line += ("." * item.amount)
      else:
        line += (f"{item.file_id}" * item.blocks)
    print(line)
  
  def swap(items, free_space_indices, file_indices) -> tuple[bool, list | None]:
    nonlocal last_file_id

    for file_idx in range(len(file_indices)-1, -1, -1):

      i = file_indices[file_idx]
      item = items[i]
      if isinstance(item, FreeSpace):
        continue
      
      if item.file_id > last_file_id:
        continue
      
      last_file_id = item.file_id

      for free_idx in range(len(free_space_indices)):

        
        item_idx_swap = free_space_indices[free_idx]
        if item_idx_swap >= i:
          break
        
        free_space = items[item_idx_swap]
        
        can_swap = free_space.amount >= item.blocks

        if can_swap:

          new_list = []

          def add_to_list(item):
            if not isinstance(item, FreeSpace) or len(new_list) == 0:
              new_list.append(item)
              return
            
            last = new_list[-1]
            if not isinstance(last, FreeSpace):
              new_list.append(item)
              return
            
            last.amount += item.amount
            
          for ii, v in enumerate(items):
            
            if ii == i:
              # add free space in new replacement spot
              add_to_list(FreeSpace(amount=item.blocks))
              continue

            if ii != item_idx_swap:
              add_to_list(v)
              continue
            
            add_to_list(item)
            if free_space.amount > item.blocks:
              add_to_list(FreeSpace(amount=free_space.amount - item.blocks))
              
          return True, new_list

    return False, None
  
  while True:
    free_space_indices = []
    file_indices = []
    
    for idx, v in enumerate(items):
      if isinstance(v, FreeSpace):
        free_space_indices.append(idx)
      else:
        file_indices.append(idx)

    swapped, new_l = swap(items, free_space_indices, file_indices)
    # print()
    # print(swapped, new_l)
    
    if not swapped:
      break
    
    items = new_l

  idx = 0
  ss = 0
  for item in items:
    if isinstance(item, FreeSpace):
      idx += item.amount
      continue
    
    for j in range(item.blocks):
      ss += (item.file_id * idx)
      idx += 1
    
  print(ss)
    
  
if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
