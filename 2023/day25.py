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
  
  # def clone(self):
  #   return Node(self.name, self.linked.copy())
  
  # def connected(self, other: str, nodes: dict[str, Node])-> bool:
  #   if other in self.linked:
  #     return True

  #   seen = set()
  #   for node in self.linked:
  #     seen.add(node)
      
  #     for node2 in nodes[node].linked:
  #     if other in nodes[node].linked:
  #       return True
      
  #     if nodes[node].connected(other, nodes):
  #       return True

  #   return False

def connected_groups(node_keys: Iterable[str], nodes: dict[str, Node], combos: Iterable[str]) -> list[set[str]]:
  
  @cache
  def connected_groups_helper(node_keys: Iterable[str], combos: Iterable[str], ) -> list[set[str]]:
    connected = set()
    disconnected = set()

    keys = list(node_keys)
    first = keys[0]
    # connected.add(first)
    
    to_process = [first]
    seen = set()
    # seen.add(first)
    # print('first', first)
    while len(to_process) > 0:
      node = to_process.pop()
      seen.add(node)
      connected.add(node)
      for key in nodes[node].linked:
        if key not in seen:
          to_process.append(key)
          # connected.add(key)

    disconnected = tuple(set(keys) - connected)
    # for key in keys[1:]:
      
    #   if nodes[first].connected(key, nodes):
    #     connected.add(key)
    #   else:
    #     disconnected.add(key)    

    # if len(connected) == 0 or len(disconnected) == 0:
    return [connected, disconnected]
    
    # other, other2 = connected_groups(node_keys, nodes, disconnected)
    # if len(other2) == 0:
    #   return [connected, disconnected]
    # else:
    #   return [connected, other, other2]
    

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

  # print(nodes['hfx'])
  # print(nodes['pzl'])
  # print(nodes['nvd'])
  # print(nodes['jqt'])
  # print(nodes['bvb'])
  # print(nodes['cmg'])
  # return

  node_keys = tuple(nodes.keys())
  for combos in itertools.combinations(node_keys, 6):
    # if not ('hfx' in combos and 'pzl' in combos and 'nvd' in combos and 'jqt' in combos and 'bvb' in combos and 'cmg' in combos):
    #   continue
    # print(combos)
    # all_linked = True
    # ct = 0
    valid = set()
    unique_items = set()

    for combo in itertools.combinations(combos, 2):

        l, r = combo
        left = nodes[l]
        right = nodes[r]
        # print(l, r)
        # if not l in right.linked and not r in left.linked:
        #   all_linked = False
        #   break
        if l in right.linked and r in left.linked:
          # ct += 1
          valid.add(combo)
          unique_items.add(l)
          unique_items.add(r)
        # print(combo)

    # print('valid', len(valid), len(unique_items), valid, unique_items)
    if len(valid) != 3 and len(unique_items) != 6:
      continue
    
    # need one more comvination of these ones to get the valid thing
    # sems like too many combinations
    for valid in itertools.combinations(valid, 3):

    # print(valid, combos)
    # continue
    
    # print(valid)

      expected = [('hfx', 'pzl'), ('bvb', 'cmg'), ('nvd', 'jqt')]
    
    # try a winning solution
    # if not all((v in valid or (v[1], v[0]) in valid) for v in expected):
    #   continue
      # continue
    
    # print(valid)

      nodes2 = deepcopy(nodes)
      for group in valid:

        l, r = group
        nodes2[l].linked.remove(r)
        nodes2[r].linked.remove(l)
        
        # how go through and group them

        # print(l, r)
        # print(nodes[l].linked)
        # print(nodes2[l].linked)
        
        # print(nodes[r].linked)
        # print(nodes2[r].linked)
        # never find the 9 6...
      ret = connected_groups(node_keys, nodes2, node_keys)
    
      if len(ret) == 2 and len(ret[1]) != 0:
        # print(ret, valid)
        print(len(ret[0]) * len(ret[1]))
        return
      # print(len(ret[0]), len(ret[1]))
    # bug here finding the 14 and 1
    # if len(ret) == 2 and len(ret[1]) != 0 and len(ret[1]) != 1 and len(ret[0]) != 0 and len(ret[0]) != 1:
      # print(ret, valid)
      # print(len(ret[0]), len(ret[1]))
      # return
      # for k in node_keys:
      #   if k in nodes2:
      #     nodes2[k].linked = set()

    # now count if we can uniquely group them
    # check if this is what 
      # print(groups)
      # for node in groups:
      #   print(node, nodes[node].linked)
      # print()
    # nodes2 = nodes.copy()

    # try this combo
    # print(combos)

  # for node, node
  # print(len(nodes))


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
