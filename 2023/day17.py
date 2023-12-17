from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import math



def reconstruct_path(came_from, current):
  path = [current]
  while current in came_from:
    current = came_from[current].pos
    path.append(current)
  return path[::-1]

def dir_repeated(came_from, current, delta: tuple[int, int], n: tuple[int, int], max_repeats) -> bool:
  start = current
  count = 0
  while current in came_from:
    curr_delta = came_from[current].delta
    if curr_delta == delta:
      count += 1
    else:
      return False
    if count > max_repeats:
      # print(f'skipping repeated {count} times for delta {delta} at {current}:{n}')
      return True

    current = came_from[current].pos

  return False

def valid_neighbor(came_from, current: tuple[int, int], n: tuple[int, int], delta: tuple[int, int], max_repeats=3) -> bool:
  # can't be where you just came from
  # can't reverse 
  # should this be curr or n??
  # n has not been put there so that doesn't quite make sense

  came_from_v = came_from[current] if current in came_from else None
  if came_from_v is not None:
    pass
    # not same direction, not previous node
    # this is not doing anything (yet??)
    if came_from_v.pos == n:
      return False
  
    # if delta == came_from_v.delta:
      # return False
  
  # no repeats
  return not dir_repeated(came_from, current, delta, n, max_repeats)

def distance(g:Grid, came_from, start, current, n, path):
  # path = reconstruct_path(came_from, current)
  return sum(g.get_int(v[0], v[1]) for v in path if v != start) + g.get_int(n[0], n[1])
  
@dataclass
class StarNode:
  pos: tuple[int, int]
  delta: tuple[int, int]

# return
def a_star(g: Grid, start: tuple[int, int], end_pos: tuple[int, int]):
  open_list = []
  open_list.append(start)
  came_from = {}
  
  g_score = defaultdict(lambda: math.inf)
  # g_score = defaultdict(int)
  g_score[start] = 0
  
  # wow this inf matters???
  f_score = defaultdict(lambda: math.inf)
  # f_score = defaultdict(int)
  # how to determine h(start)??
  f_score[start] = 0
  seen = set()
  
  while len(open_list) > 0:
    curr = open_list[0]
    curr_idx = 0
    
    for idx, item in enumerate(open_list):
      if f_score[idx] < f_score[curr_idx]:
        curr = f_score[idx]
        curr_idx = idx

    open_list.pop(curr_idx)
    seen.add(curr)
    
    path = reconstruct_path(came_from, curr)
    if curr == end_pos:
      # print(came_from)
      return path
      # return reconstruct_path(came_from, curr)
    

    for n in g.neighbors(curr[0], curr[1]):
      # if n in seen:
      #   print(f'seen {n}. skipping')
      #   continue

      g1 = g_score[curr] + distance(g, came_from, start, curr, n, path)
      # g1 = g_score[curr] + g.get_int(n[0], n[1])
      delta = (n[0] - curr[0], n[1] - curr[1])
      if g1 < g_score[n] and valid_neighbor(came_from, curr, n, delta):
        # print(g1, g_)

        # changing this to try and incorporate delta
        # came_from[n] = curr
        came_from[n] = StarNode(curr, delta)
        g_score[n] = g1
        f_score[n] = g1 + g.get_int(n[0], n[1])
        if n not in open_list:
          open_list.append(n)
  
  return None

async def part1():
  g = await Grid.grid_from_file("day17input")
  
  start = (0,0)
  end = (g.height - 1, g.width - 1)

  ret = a_star(g, start, end)
  print(ret)
  print(sum(g.get_int(v[0], v[1]) for v in ret if v != start))


async def part2():
  async for line in read_file("day17input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
