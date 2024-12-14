from __future__ import annotations
from dataclasses import dataclass
from typing import TypeVar, Generic
from utils import read_file
from collections import defaultdict

T = TypeVar("T")

@dataclass
class Data:
  visited: bool = False
  # value: T = field(default_factory=lambda: 0 if isinstance(T, int) else "")
  value: T = None

@dataclass()
class Grid(Generic[T]):
  width: int
  height: int
  data: list[list[Data[T]]]

  def __init__(self, width: int, height: int, default_value: T | None = None) -> str:
    self.width = width
    self.height = height
    self.data = [[Data(value=default_value) for _ in range(width)] for _ in range(height)]
    # print(self.data)

  @staticmethod
  def square_grid(length: int) -> Grid[T]:
    return Grid(length, length)

  @staticmethod
  def grid(width: int, height: int) -> Grid[T]:
    return Grid(width, height)
  
  @staticmethod
  async def square_grid_from_file(file: str) -> Grid[str]:
    init = False
    g = None
    r = 0

    async for line in read_file(file):
      if not init:
        l = len(line)
        g = Grid.square_grid(l)
        init = True
      assert g is not None
      for c in range(l):
        g.set(r, c, line[c])
      r += 1
    return g
  
  # TODO kill square_grid_from_file above since it may not actually be square in all usages and this is better
  @staticmethod
  async def grid_from_file(file: str) -> Grid[str]:
    lines = [line async for line in read_file(file)]

    return Grid.from_lines(lines)
  
  @staticmethod
  def from_lines(lines: list[str]) -> Grid[str]:
    width = len(lines[0])
    height = len(lines)

    g = Grid(width, height)
    for r in range(g.height):
      for c in range(g.width):
        g.set(r, c, lines[r][c])

    return g  

  def print(self, *, none_value = "."):
    for r in range(self.height):
      for c in range(self.width):
        val = self.data[r][c].value 
        val = val if val is not None else none_value
        print(val, end="")
      print()
    print()
    
  def walk(self):
    for r in range(self.height):
      for c in range(self.width):
        yield (r, c)

  def set(self, r: int, c: int, val: T):
    self.data[r][c].value = val
    
  def get(self, r: int, c: int) -> Data[T]:
    return self.data[r][c]
  
  def get_value(self, r: int, c: int) -> T:
    return self.data[r][c].value
  
  def get_int(self, r: int, c: int) -> int:
    return int(self.data[r][c].value)

  def visit(self, r: int, c: int):
    self.data[r][c].visited = True
    
  def visited(self, r: int, c: int) -> bool:
    return self.data[r][c].visited
  
  def find(self, v: Any) -> tuple[int, int] | None:
    for r in range(self.height):
      for c in range(self.width):
        if self.get_value(r, c) == v:
          return (r, c)
    return None

  def current_lines(self) -> list[str]:
    result = []
    for r in range(self.height):
      line = ""
      for c in range(self.width):
        line += self.get_value(r, c)
      result.append(line)
    return result
  
  def clone(self) -> Grid[T]:
    g = Grid(self.width, self.height)
    for r in range(self.height):
      for c in range(self.width):
        g.set(r, c, self.get_value(r, c))
    return g
  
  def neighbors(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    if r > 0:
      neighbors.append((r - 1, c))
    if r < self.height - 1:
      neighbors.append((r + 1, c))
    if c > 0:
      neighbors.append((r, c - 1))
    if c < self.width - 1:
      neighbors.append((r, c + 1))
    return neighbors
  
  def right_and_down(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    if r < self.height - 1:
      neighbors.append((r + 1, c))
    if c < self.width - 1:
      neighbors.append((r, c + 1))
    return neighbors
  
  def left_and_down(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    if r < self.height - 1:
      neighbors.append((r + 1, c))
    if c > 0:
      neighbors.append((r, c - 1))
    return neighbors
  
  def left_and_up(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    if r > 0:
      neighbors.append((r - 1, c))
    if c > 0:
      neighbors.append((r, c - 1))
    return neighbors
  
  def right_and_up(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    if r > 0:
      neighbors.append((r - 1, c))
    if c < self.width - 1:
      neighbors.append((r, c + 1))
    return neighbors
  
  def top(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = r - 1
    while i >= 0:
      result.append((i, c))
      i -= 1
    return result
  
  def bottom(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = r + 1
    while i < self.height:
      result.append((i, c))
      i += 1
    return result
  
  def left(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = c - 1
    while i >= 0:
      result.append((r, i))
      i -= 1
    return result
  
  def right(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = c + 1
    while i < self.width:
      result.append((r, i))
      i += 1
    return result
  
  def neighbors8(self, r: int, c: int) -> list[tuple[int, int]]:
    neighbors = []
    for i in range(-1, 2):
      for j in range(-1, 2):
        if i == 0 and j == 0:
          continue
        r2 = r + i
        c2 = c + j
        if r2 >= 0 and r2 < self.height and c2 >= 0 and c2 < self.width:
          neighbors.append((r2, c2))
    return neighbors

  def visit_neighbors(self, r: int, c: int):
    for n in self.neighbors(r, c):
      self.visit(n[0], n[1])
      
  def rotate_left(self) -> Grid:
    # saw this on the internet based on 2023 day 13

    # lines = [*zip(*self.current_lines())]
    # g = Grid.from_lines(lines)
    # return g

    # this is easier to read I think lol
    # flipped
    g = Grid(self.height, self.width)
    for r in range(self.height):
      for c in range(self.width):
        r2 = c
        c2 = r
        g.set(r2, c2, self.get_value(r, c))
        
    return g
    
  def rotate_right(self) -> Grid:
    # flipped
    g = Grid(self.height, self.width)
    for r in range(self.height):
      for c in range(self.width):
        r2 = self.width - c - 1
        c2 = r
        g.set(r2, c2, self.get_value(r, c))
        
    return g
    
  # dijkstra with no mins. just result
  def dijkstra(self, start: tuple[int, int], end: tuple[int, int], seen_before: Optional[Callable] = None) -> int:
    visited = set()
    q = []
    q.append((start, 0))
    while len(q) > 0:
      curr, dist = q.pop(0)
      if curr == end:
        return dist
      visited.add(curr)
      for n in self.neighbors(curr[0], curr[1]):
        if seen_before is not None:
          seen = seen_before(curr, n)
          if seen != -1:
            print('using seen', seen)
            return dist + seen
        if n not in visited:
          q.append((n, dist + 1))
    return -1
  
  # translating 2023 day 15  
  def dijkstra2(self, start: tuple[int, int], end: tuple[int, int]) -> int:
    q = set()
    mins = defaultdict(int)
    unvisitedmins = set()
    for r in range(self.height):
      for c in range(self.width):
        q.add((r, c))

    mins[start] = 0
    
    curr = start
    while len(q) > 0:
      curr_val = mins[curr]
      
      for n in self.neighbors(curr[0], curr[1]):
        if self.visited(n[0], n[1]):
          continue
        
        new_min = curr_val + 1
        neighbor_min = mins[n]
        
        if neighbor_min == 0 or new_min < neighbor_min:
          mins[n] = new_min
          unvisitedmins.add(n)

      if curr in q:          
        q.remove(curr)
      self.visit(curr[0], curr[1])
      if curr in unvisitedmins:
        unvisitedmins.remove(curr)
      
      if curr == end:
        break
      
      min_so_far = None
      new_curr = None
      # print('unvisitedmins', unvisitedmins)
      assert len(unvisitedmins) > 0
      for n in unvisitedmins:
        v = mins[n]
        if v == 0 or self.visited(n[0], n[1]):
          continue
        if min_so_far is None or v < min_so_far:
          min_so_far = v
          new_curr = n
          
      if new_curr is not None:
        curr = new_curr
        continue

    return mins[end]
      
      
# TODO lines to get the lines
# TODO rotate