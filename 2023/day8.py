from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re

@dataclass
class Node:
  left: str
  right: str

async def part1():
  directions = ""
  nodes = {}
  
  async for line in read_file("day8input"):
    if line == "":
      continue
    if directions == "":
      directions = line
      continue

    match = re.match(r"(.+) = \((.+), (.+)\)", line)
    assert match is not None
    node = Node(match.group(2), match.group(3))
    nodes[match.group(1)] = node

  steps = 0
  next = "AAA"
  curr = nodes[next]
  while True:
    idx = steps % len(directions)
    next = curr.left if directions[idx] == "L" else curr.right
    curr = nodes[next]
    steps += 1
    if next == "ZZZ":
      break
    
  print(steps)


async def part2():
  async for line in read_file("day8input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
