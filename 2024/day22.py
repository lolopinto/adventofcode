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

def mix(value: int, secret: int) -> int:
  return value ^ secret

def prune(secret: int) -> int:
  return secret % 16777216

def next_secret(secret: int) -> int:
  n = secret * 64
  secret = mix(n, secret)
  secret = prune(secret)
  n2 = secret // 32
  secret = mix(n2, secret)
  secret = prune(secret)
  
  n3 = secret * 2048
  secret = mix(n3, secret)
  secret = prune(secret)
  return secret

def get_nth_secret(secret: int, n: int) -> int:
  last = None
  for secret in yield_nth_secret(secret, n):
    last = secret
  assert last is not None
  return last

def yield_nth_secret(secret: int, n: int) -> int:
  for i in range(n):
    secret = next_secret(secret)
    yield secret

async def part1():
  s = 0
  async for line in read_file("day22input"):
    s += get_nth_secret(int(line), 2000)
    
  print(s)


def get_deltas(secret: int, n: int) -> dict[tuple[int, int, int, int], int]: 
  last = secret % 10
  deltas = []
  
  m = dict()
  for secret in yield_nth_secret(secret, n):
    curr = secret % 10
    delta = curr - last
    last = curr
    deltas.append(delta)
    if len(deltas) >= 4:
      t = tuple(deltas[-4:])
      if t not in m:
        m[t] = curr
      # print(t, curr)

  return m

async def part2():
  deltas_map = dict()
  all_deltas = set()
  all_numbers = []
  async for line in read_file("day22input"):
    n = int(line)
    deltas = get_deltas(n, 2000)
    deltas_map[n] = deltas
    all_deltas.update(deltas.keys())
    all_numbers.append(n)
    

  curr_max = 0      
  for delta in all_deltas:
    s = 0
    for n in all_numbers:
      if delta in deltas_map[n]:
        s += deltas_map[n][delta]

    curr_max = max(curr_max, s)
    
  print(curr_max)
    

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
