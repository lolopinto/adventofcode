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
import networkx as nx

@dataclass
class Node:
  name: str
  linked: set[str]

async def part1():
  # got this from someone in reddit
  # need to start using all these tools that all these folks use instead of reinventing the wheel
  G = nx.Graph()

  async for line in read_file("day25input"):
    left, right = line.split(": ")
    for right in right.split(" "):
      G.add_edge(left, right)

  G.remove_edges_from(nx.minimum_edge_cut(G))

  print(math.prod([len(c) for c in nx.connected_components(G)]))


if __name__ == "__main__":
    asyncio.run(part1())
