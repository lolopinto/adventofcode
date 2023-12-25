from __future__ import annotations
from utils import read_file_groups, read_file, ints
from typing import Iterable
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math
import sys
from functools import cache
from copy import deepcopy
sys.setrecursionlimit(10000)

@dataclass
class Node:
  name: str
  linked: set[str]

def connected_groups(node_keys: Iterable[str], nodes: dict[str, Node], combos: Iterable[str]) -> list[set[str]]:
  
  @cache
  def connected_groups_helper(node_keys: Iterable[str], combos: Iterable[str], ) -> list[set[str]]:
    connected = set()
    disconnected = set()

    keys = list(node_keys)
    first = keys[0]
    
    to_process = [first]
    while len(to_process) > 0:
      node = to_process.pop()
      connected.add(node)
      for key in nodes[node].linked:
        if key not in connected:
          to_process.append(key)

    disconnected = tuple(set(keys) - connected)

    return [connected, disconnected]        

  return connected_groups_helper(node_keys, combos)

async def part1():
  nodes = {}

  async for line in read_file("day25input"):
    left, right = line.split(": ")
    if left not in nodes:
      nodes[left] = Node(left, set())
    for right in right.split(" "):
      if right not in nodes:
        nodes[right] = Node(right, set())
      nodes[right].linked.add(left)
      nodes[left].linked.add(right)

  node_keys = tuple(nodes.keys())
  for combos in itertools.combinations(node_keys, 6):

    valid = set()
    unique_items = set()

    for combo in itertools.combinations(combos, 2):

        l, r = combo
        left = nodes[l]
        right = nodes[r]

        if l in right.linked and r in left.linked:
          valid.add(combo)
          unique_items.add(l)
          unique_items.add(r)

    if len(valid) != 3 and len(unique_items) != 6:
      continue
    
    # need one more combination of these ones to get the valid thing
    # seems like too many combinations
    for valid in itertools.combinations(valid, 3):

      nodes2 = deepcopy(nodes)
      for group in valid:
        l, r = group
        nodes2[l].linked.remove(r)
        nodes2[r].linked.remove(l)
        
      # now go through and group them
      ret = connected_groups(node_keys, nodes2, node_keys)
    
      if len(ret) == 2 and len(ret[1]) != 0:
        print(len(ret[0]) * len(ret[1]))
        return


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
