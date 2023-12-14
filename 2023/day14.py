from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

async def part1():
  g = await Grid.grid_from_file("day14input")
  g.print()
  
  # for each column
  for c in range(g.width):
    for r in range(g.height):
      # only swapping .
      if g.get_value(r, c) != '.': 
        continue

      for r2, c2 in g.bottom(r, c):
        # look for the next rock
        match g.get_value(r2, c2):
          case '#':
            break
          case 'O':
            # found a rock, swap 
            g.set(r2, c2, '.')
            g.set(r, c, 'O')
            break
  
  g.print()

  sum = 0  
  for r in range(g.height):
    row = g.height - r
    for c in range(g.width):
      if g.get_value(r, c) == 'O':
        sum += row

  print(sum)


@dataclass
class Info:
  direction: str

async def part2():
  g = await Grid.grid_from_file("day14input")
  # g.print()
  # for each column
  
  lines = defaultdict(list)
  # lines = {}
  
  cycle = 1
  # cycle_key = None
  cycle_length = None
  while True:
    # N W S E
    directions = [(g.bottom, False, 'N'), (g.right, False, 'W'), (g.top, True, 'S'), (g.left, True, 'E')]
    for direction, rev, dir_str in directions:
      # print(f"going {dir_str}")
      
      width = range(g.width)
      height = range(g.height)
      if rev:
        width = list(reversed(width))
        height = list(reversed(height))
      for c in width:
        for r in height:
          # if dir_str == 'S':
          #   print(r,c)
          # if dir_str == 'S' and ((r,c) == (4,2) or (r,c) == (3,2)):
            
          #   print(r,c)

          # only swapping .
          if g.get_value(r, c) != '.': 
            continue

          rows = direction(r, c)
          # if rev:
          #   rows.reverse()
          # if dir_str == 'S' and ((r,c) == (4,2) or (r,c) == (3,2)):
          #   print(r,c, rows)
          for r2, c2 in rows:
              
            # look for the next rock
            match g.get_value(r2, c2):
              case '#':
                break
              case 'O':
                # found a rock, swap 
                g.set(r2, c2, '.')
                g.set(r, c, 'O')
                break
      # post direction 

    # post cycle
    curr_grid = tuple(g.current_lines())
    key = curr_grid
    # pri
    # key = f"{dir_str}{curr_grid}"
    # if key in lines:
    #   print(f"cycle. seen @ {cycle} {lines[key]}")
    # if key in lines:
    #   # cycle_key = key
    #   cycle_length = cycle - lines[key]
    #   # print('cycle detected')
    # lines[key] = cycle
    
    # if cycle_length is not None and cycle > cycle_length * 2:
    #   print('breaking cycle')
    #   break

    lines[key].append(cycle)
    # TODO 2
    if len(lines[key]) == 2:
      print('cycle detected')
      cycle_key = key
      cycle_length = lines[key][1] - lines[key][0]
      # print(f"cycle @ {cycle} {lines[key]}")
      break
    
    # if cycle > 
    cycle += 1
    
  # assert cycle_key is not None
  # cycle_length = lines[cycle_key][1] - lines[cycle_key][0]
  # cycle_start = lines[cycle_key][1]

  target_idx = 1_000_000_000 % cycle_length  
  print(cycle, cycle_length, target_idx)

  # remaining = 1_000_000_000 - cycle
  # cycles_remaining = remaining // cycle_length
  # cycle_delta = remaining % cycle_length
  # print(lines.values())
  # remaining = ((1_000_000 - cycle) % cycle_length )
  
  # print(cycle_length, cycle_start, cycle, remaining, cycles_remaining, cycle_delta)
  # answer = {k: v for k, v in lines.items() if v[0] == target_idx}
  answer = {k: v for k, v in lines.items() if v[-1] % cycle_length == target_idx}
  print(len(answer))
  # assert len(answer) == 1

  # print(list(answer.keys())[0])
  # 95160 too low when using the first key
  # 97241 correct answer using last key
  g2 = Grid.from_lines(list(answer.keys())[-1])
  # print(lines)
    # g.print()
    
  # g.print()

  sum = 0  
  for r in range(g2.height):
    row = g2.height - r
    for c in range(g2.width):
      # had this line wrong. had g instead of g2
      if g2.get_value(r, c) == 'O':
        sum += row

  print(sum)

if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())
