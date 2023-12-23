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
import sys
sys.setrecursionlimit(10_000)

slopes = {
  '>': (0, 1),
  '<': (0, -1),
  '^': (-1, 0),
  'v': (1, 0),
}

async def longest_path(check_slopes: bool): 
  g = await Grid.grid_from_file("day23input")
  
  start = None  
  end = None
  for c in range(g.width):
    if (v := g.get_value(0, c)) == '.':
      start = (0, c)
    if (v := g.get_value(g.height-1, c)) == '.':
      end = (g.height -1, c)
      
  assert start is not None
  assert end is not None
  
  visited = set()
  visited.add(start)
  
  # tried to do a modified dijkstra first and wasn't working. had chatgpt 
  # help me modify this
  def dfs(r, c, length):
    nonlocal max_length
    
    if end[0] == r and end[1] == c:
      max_length = max(max_length, length)
      return
    
    for n in g.neighbors(r, c):
      # ignore
      v = g.get_value(n[0], n[1])
      if v == '#' or n in visited:
        continue
          
      delta = (n[0] - r, n[1] - c)

      # can only go in direction pointing to
      if check_slopes and v in slopes and delta != slopes[v]:
        # print(f'skipping {v}')
        continue
        
      visited.add(n)
      dfs(n[0], n[1], length + 1)
      visited.remove(n)

  max_length = 0
  dfs(start[0], start[1], 0)

  return max_length 

async def part1():
  print(await longest_path(True))


async def part2():
  pass
  # print(await longest_path(True))
  # g = await Grid.grid_from_file("day23input")
  
  # start = None  
  # end = None
  # for c in range(g.width):
  #   if (v := g.get_value(0, c)) == '.':
  #     start = (0, c)
  #   if (v := g.get_value(g.height-1, c)) == '.':
  #     end = (g.height -1, c)
      
  # assert start is not None
  # assert end is not None
  
  # dp = defaultdict(int)
  # dp[start] = 1
  
  # for (r, c) in g.walk():
  #   # print(r, c)
  #   v = g.get_value(r,c)
  #   if v == '#':
  #     continue
  #   if r > 0:
  #     dp[(r,c)] = max(dp[(r, c)], dp[(r-1, c)] +1)
  #   if c > 0:
  #     dp[(r,c)] = max(dp[(r,c)], dp[(r, c-1)] +1)
      
  # print(len(dp))
  
  # for c in range(g.width):
  #   print(dp[(1, c)], end = " ")
  # print()
  # for c in range(g.width):
  #   print(dp[(2, c)], end = " ")
  # print()
  # for c in range(g.width):
  #   print(dp[(3, c)], end = " ")
  # print()
  # for c in range(g.width):
  #   print(dp[(4, c)], end = " ")


  # print(dp[end])
      
    




if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
