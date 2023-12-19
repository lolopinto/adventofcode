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
        return self.result(rule, rating, workflows)
      
  def condition(self, condition: str, rating: Rating)-> bool:
    parts1 = condition.split("<")
    parts2 = condition.split(">")
    if len(parts1) != 2 and len(parts2) != 2:
      raise ValueError(f"unknown condition {condition}")

    parts = parts1 if len(parts1) == 2 else parts2

    key = parts[0]
    rhs = int(parts[1])
    lhs = rating.m[key]      
    if len(parts1) == 2:
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


  def deps(self) -> list[str]:
    deps = []
    for rule in self.rules:
      parts = rule.split(":")
      if parts[-1] in 'AR':
        continue
      deps.append(parts[-1])
    return deps
  
  def count(self, workflows: dict[str, Workflow], results: dict[str, int]):
    result = 0
    
    curr_range = (1, 4001)
    curr_key = None

    ranges = []
    for rule in self.rules:
      parts = rule.split(":")
      
      rule_key = rule
      delta_range = None

      if len(parts) == 2:
        parts_less = parts[0].split("<")
        parts_right = parts[0].split(">")
        assert len(parts_less) == 2 or len(parts_right) == 2
        
        # passing these many times
        # less than 
        if len(parts_less) == 2:
          v = int(parts_less[1])
          curr_range = (1, v)
          delta_range = (v, 4001)
          curr_key = parts_less[0]
        else:
          v = int(parts_right[1])
          curr_range = (v + 1, 4001)
          delta_range = (1, v + 1)
          curr_key = parts_right[0]
          # print(key)

        # now evaluate the rest of the rules, how many times
        rule_key = parts[1]

        ranges.append((curr_key, curr_range))

      else:
        rule_key = rule


      assert curr_range is not None
      assert curr_key is not None
      
      # so wanna change the result to be dependent not what we have here
      # for in, we want (1, 1351) -> px
      # and rest -> qqz 
      
      res = self.count_part2(rule_key, results)
      
      keys = set()
      # l = len(range(curr_range[0], curr_range[1]))
      l = 1
      for r in ranges:
        keys.add(r[0])
        length = len(range(r[1][0], r[1][1]))
        l *= length

      # if self.id == 'rfg':
      print(ranges)      
      # TODO need to handle repeated ranges
      # assert len(keys) == len(ranges)
      print('adding', l * res)
      result += l * res
      
      if delta_range is not None:
        # remove last range, add delta range as new range
        ranges.pop(-1)
        ranges.append((curr_key, delta_range))

        # ranges.append((curr_key, curr_range))
        curr_range = delta_range


    return result    

  def count_part2(self, str: s, results: dict[str, int]) -> int:
    match str:
      case 'A':
        return 1
      case 'R':
        return 0
      case _:
        return results[str]


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

  deps_list = {}
  for line in groups[0]:
    w = Workflow.parse(line)    
    # print(w.id, w.deps())
    workflows[w.id] = w
    deps_list[w.id] = w.deps()

  # print(len(deps_list))
    
  done = set()

  deps_tree = defaultdict(list)
  while len(deps_list) > 0:
    deps_list2 = {}
    for k, v in deps_list.items():
      all_done = True
      for dep in v:
        if dep not in done:
          all_done = False
          break
      if all_done:
        done.add(k)
        for dep in v:
          deps_tree[k].append(dep)
          for dep2 in deps_tree[dep]:
            deps_tree[k].append(dep2)
      else:
        deps_list2[k] = v
    deps_list = deps_list2
  

  deps_tree_sorted = [(k, v) for k,v in deps_tree.items()]
  deps_tree_sorted = sorted(deps_tree_sorted, key=lambda x: len(x[1]))
  # last should be in
  assert deps_tree_sorted[-1][0] == "in"
  # print(deps_tree_sorted)
  # print(done)
  
  # lets start simple
  candidates = [x for x in deps_tree_sorted if len(x[1]) <= 10]
  
  # print('candidates', candidates)
  results = {}

  # order matters here
  # going in order of dependencies
  for k, _ in deps_tree_sorted:
    yay = False
    for cand in candidates:
      if k == cand[0]:
        yay = True
        break
    if not yay:
      continue
    print(f"processing {k}")
    w = workflows[k]
    ret = w.count(workflows, results)
    print(k, ret)
    results[k] = ret
    
  # this doesn't have distinct?!
  # back to ranges lol?
  print(results)


if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())
