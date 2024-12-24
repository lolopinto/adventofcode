from __future__ import annotations
from utils import read_file_groups, read_file, ints, get_file_groups
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math

@dataclass
class Equation:
  lhs: str
  rhs: str
  op: Literal["XOR", "OR", "AND"]
  result: str
  
  @staticmethod
  def parse(line: str) -> Equation:
    parts = line.split(" ")
    assert len(parts) == 5
    
    return Equation(
      lhs=parts[0],
      op=parts[1],
      rhs=parts[2],
      result=parts[4],
    )
    
  def can_solve(self, known: dict[str, int]) -> bool:
    return self.lhs in known and self.rhs in known
  
  def solve(self, known: dict[str, int]):
    left = known[self.lhs]
    right = known[self.rhs]
    
    match self.op:
      case 'AND':
        ret = left & right
      case 'OR':
        ret = left | right
      case 'XOR':
        ret = left ^ right
      case _:
        raise ValueError(f"unknown op {self.op}")
    # if self.result in known:
    #   print("{self.result} already in known" )
    known[self.result] = ret


async def part1():
  groups = await get_file_groups("day24input")
  assert len(groups) == 2
  
  known: dict[str, int] = dict()

  for line in groups[0]:
    parts = line.split(": ")
    assert len(parts) == 2
    known[parts[0]] = int(parts[1])

  equations = []    
  for line in groups[1]:
    eq = Equation.parse(line)
    equations.append(eq)

  while len(equations) > 0:
    todo = []
    for eq in equations:
      if eq.can_solve(known):
        eq.solve(known)
      else:
        todo.append(eq)
    equations = todo
    
  zes = []
  for k, v in known.items():
    if k[0] == 'z':
      zes.append((k, v))

  s = ""
  # 0 is least significant bit... so have to reverse
  zes = sorted(zes, key=lambda t: t[0], reverse=True)
  for _, v in zes:
    s += f"{v}"

  print(s, int(s, 2))
      
    
    


async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
