from __future__ import annotations
from utils import get_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools


@dataclass
class Workflow:
  id: str
  rules: list[str] # parts divided by 2
  
  @staticmethod
  def parse(line: s) -> Workflow:
    match = re.match(r"(\w+)\{(.+)\}", line)
    assert match is not None
    groups = match.groups()
    assert len(groups) == 2
    return Workflow(groups[0], groups[1].split(","))
  

  def eval(self, rating: Rating, workflows: dict[str, Workflow])-> bool:
    for rule in self.rules:
      parts = rule.split(":")
      if len(parts) == 2:
        if not self.condition(parts[0], rating):
          continue
        
        return self.result(parts[1], rating, workflows)
      else: 
        # print("ruke", rule, self.rules)
        return self.result(rule, rating, workflows)
      
  def condition(self, condition: str, rating: Rating)-> bool:
    parts = condition.split("<")
    if len(parts) == 2:
      key = parts[0]
      rhs = int(parts[1])
      lhs = rating.m[key]      
      return lhs < rhs

    parts = condition.split(">")
    if len(parts) == 2:
      key = parts[0]
      rhs = int(parts[1])
      lhs = rating.m[key]      
      return lhs > rhs

    raise ValueError(f"unknown condition {condition}")
      
  def result(self, key: string, rating: Rating, workflows: dict[str, Workflow]) -> bool:
    match key:
      case 'A':
        return True
      case 'R':
        return False
      case _:
        # print(key)
        return workflows[key].eval(rating, workflows)


@dataclass
class Rating:
  m: dict[str,int]
  
  @staticmethod
  def parse(line: s) -> Rating:
    match = re.match(r"\{(.+)\}", line)
    assert match is not None
    groups = match.groups()
    v = groups[0]
    parts = v.split(",")
    
    m = {}
    for part in parts:
      pp = part.split('=')
      assert len(pp) == 2
      
      m[pp[0]] = int(pp[1])
    return Rating(m)

  def sum(self) -> int:
    return sum(self.m.values())

async def part1():
  groups = await get_file_groups("day19input")
  assert len(groups) == 2
  
  workflows = {}
  for line in groups[0]:
    w = Workflow.parse(line)    
    workflows[w.id] = w

  s = 0
  for line in groups[1]:
    r = Rating.parse(line)
    start = workflows["in"]
    if start.eval(r, workflows):
      s += r.sum()

  print(s)



async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
