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
  for i in range(n):
    secret = next_secret(secret)
  return secret

async def part1():
  s = 0
  async for line in read_file("day22input"):
    s += get_nth_secret(int(line), 2000)
    
  print(s)


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
