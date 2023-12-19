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
  
  # instead of counts, return valid ranges that are correct
  # e.g. x: 0-4000
  # m: 0-1716
  # a: 0-1338
  # s: 0-1338
  def count(self, workflows: dict[str, Workflow]):
    # res = 0
    # start = 4000

    # this is not accounting for different values of xmas
    # so the 4000 needs to account for that.
    # e.g. 4000 permuations of x, 4000 permutations of m and then multiply and add as necessary?

    result = defaultdict(list)
    
    curr_range = (1, 4001)
    curr_key = None

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
          # times = len(range(int(parts_less[1])))
          curr_range = (1, v)
          delta_range = (v, 4001)
          curr_key = parts_less[0]
        else:
          # times = len(range(int(parts_right[1]) +1, 4001))
          v = int(parts_right[1])
          curr_range = (v + 1, 4001)
          delta_range = (1, v + 1)
          curr_key = parts_right[0]
          # print(key)

        # now evaluate the rest of the rules, how many times
        rule_key = parts[1]
      else:
        rule_key = rule


      assert curr_range is not None
      assert curr_key is not None

      # so wanna change the result to be dependent not what we have here
      # for in, we want (1, 1351) -> px
      # and rest -> qqz 
      
      res = self.count_part2(rule_key, workflows)
      match res:
        case True: 
          result[curr_key].append(curr_range)
        case False:
          continue
        case _:
          # dict 
          # print(f"other dict {rule}", result, res)
          result[curr_key].append(curr_range)
          for k, v in res.items():
            result[k].extend(v)
      
      if delta_range is not None: 
        curr_range = delta_range
          
      # else:
      #   assert curr_range is not None
      #   assert curr_key is not None


      # # TODO 
      #   res = self.count_part2(rule, workflows)
      #   match res:
      #     case True: 
      #       result[curr_key].append(curr_range)
      #     case False:
      #       continue
      #     case _:
      #       # dict 
      #       # print(f"other dict {rule}", result, res)
      #       for k, v in res.items():
      #         result[k].extend(v)

    return result    

  def count_part2(self, str: s, workflows: dict[str, Workflow]) -> bool | defaultdict(list):
    match str:
      case 'A':
        return True
      case 'R':
        return False
      case _:
        return workflows[str].count(workflows)


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

  print(len(deps_list))
    
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
  print(deps_tree_sorted)
  print(done)
  
  # lets start simple
  candidates = [x for x in deps_tree_sorted if len(x[1]) < 100]
  # candidates = deps_tree_sorted
  print(candidates)

# 0
# lnx 4000
# pv 1716
# gd 0
# crn 1338

  for cand in candidates:
    w = workflows[cand[0]]
    print(w.id, w.count(workflows))

  # solution once things are working?
  # workflows["in"].count(workflows)

  # s = 0
  # for _ in groups[1]:
  #   w = workflows["in"]
  #   s += w.count(workflows)
  # print(s)


  # print(done)
  # do we actually need this
  # l = list(deps_tree.keys())
  # print(all(v in done for v in l))
  # print(len(deps_tree.keys()))

  # deps_tree = workflows.copy()
  # while len(deps_tree) > 0:
  # need to build dependency tree

  # s = 0
  # for line in groups[1]:
  #   r = Rating.parse(line)
  #   start = workflows["in"]
  #   # cond = start.conditions(r, workflows)
  #   # print(cond)
  #   # s += cond
    
    
  # print(s)

if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())
