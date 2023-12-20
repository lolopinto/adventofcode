from __future__ import annotations
from utils import get_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import math

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
        return self.result(rule, rating, workflows)
      
  def condition(self, condition: str, rating: Rating)-> bool:    
    assert condition[1] in "<>"

    key = condition[0]
    rhs = int(condition[2:])
    lhs = rating.m[key]
    if condition[1] == '<':
      return lhs < rhs
    else:
      return lhs > rhs

      
  def result(self, key: string, rating: Rating, workflows: dict[str, Workflow]) -> bool:
    match key:
      case 'A':
        return True
      case 'R':
        return False
      case _:
        return workflows[key].eval(rating, workflows)

  
  @staticmethod
  def intersect_ranges(r1, r2):
    start = max(r1[0], r2[0])
    end = min(r1[1], r2[1])
    return None if start > end else (start, end)

  def count(self, workflows: dict[str, Workflow], ranges: dict[str, tuple[int, int]]):
    result = 0
    
    for rule in self.rules:
      parts = rule.split(":")
      
      delta_range = None

      if len(parts) == 2:
        condition = parts[0]
        assert condition[1] in "<>"

        v = int(condition[2:])
        curr_key = condition[0]
        # less than 
        if condition[1] == '<':
          curr_range = (1, v)
          delta_range = (v, 4001)
        else:
          curr_range = (v + 1, 4001)
          delta_range = (1, v + 1)

        char_range = ranges[curr_key]
        inter = Workflow.intersect_ranges(char_range, curr_range)
        if inter is not None:
          ranges_temp = ranges.copy()
          ranges_temp[curr_key] = inter

          # evaluate this workflow just for this rang
          result += self.count_part2(parts[1], workflows, ranges_temp)

        # update range for the next rule
        inter = Workflow.intersect_ranges(char_range, delta_range)
        if inter is not None:
          ranges[curr_key] = inter
      else:
        # this is usually the last one so can just handle it at the end...
        result += self.count_part2(rule, workflows, ranges)

    return result    

  def count_part2(self, str: s, workflows: dict[str, Workflow], ranges: dict[str, int]) -> bool | dict:
    match str:
      case 'A':
        return math.prod(len(range(ranges[c][0], ranges[c][1])) for c in ranges)
      case 'R':
        return 0
      case _:
        return workflows[str].count(workflows, ranges)


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
  groups = await get_file_groups("day19input")
  assert len(groups) == 2
  
  workflows = {}

  for line in groups[0]:
    w = Workflow.parse(line)    
    workflows[w.id] = w
    
  ranges = {
    'x': (1, 4001),
    'm': (1, 4001),
    'a': (1, 4001),
    's': (1, 4001),
  }
  print(workflows["in"].count(workflows, ranges))
  

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
