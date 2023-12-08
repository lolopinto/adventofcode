from __future__ import annotations
from typing import Callable
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import functools
import math

@dataclass
class Node:
  left: str
  right: str


def find_steps(start: str, directions: str, nodes: dict[str, Node], is_end: Callable[[str], bool]) -> str:
  next = start
  curr = nodes[next]
  steps = 0
  while True:
    idx = steps % len(directions)
    next = curr.left if directions[idx] == "L" else curr.right
    curr = nodes[next]
    steps += 1
    if is_end(next):
      break

  return steps

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

  print(find_steps("AAA", directions, nodes, lambda x: x == "ZZZ"))


async def part2():
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

  starts = [key for key in nodes.keys() if key[-1] == "A"]
  
  z_steps = []
  for start in starts:
    z_steps.append(find_steps(start, directions, nodes, lambda x: x[-1]== "Z"))
    
  print(math.lcm(*z_steps))


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
