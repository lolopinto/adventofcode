from __future__ import annotations
from dataclasses import dataclass
from typing import TypeVar, Generic
from utils import read_file

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

  def __init__(self, width: int, height: int) -> str:
    self.width = width
    self.height = height
    self.data = [[Data() for _ in range(width)] for _ in range(height)]
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

    async for line in read_file("day3input"):
      if not init:
        l = len(line)
        g = Grid.square_grid(l)
        init = True
      assert g is not None
      for c in range(l):
        g.set(r, c, line[c])
      r += 1
    return g
  
  def set(self, r: int, c: int, val: T):
    self.data[r][c].value = val
    
  def get(self, r: int, c: int) -> Data[T]:
    return self.data[r][c]
  
  def get_value(self, r: int, c: int) -> T:
    return self.data[r][c].value

  def visit(self, r: int, c: int):
    self.data[r][c].visited = True
    
  def visited(self, r: int, c: int) -> bool:
    return self.data[r][c].visited
  
  
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
      i -= 1
      result.append((i, c))
    return result
  
  def bottom(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = r + 1
    while i < self.height:
      i += 1
      result.append((i, c))
    return result
  
  def left(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = c - 1
    while i >= 0:
      i -= 1
      result.append((r, i))
    return result
  
  def right(self, r: int, c: int) -> list[tuple[int, int]]:
    result = []
    i = c + 1
    while i < self.width:
      i += 1
      result.append((r, i))
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