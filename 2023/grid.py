from __future__ import annotations
from dataclasses import dataclass
from typing import TypeVar, Generic

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
    print(self.data)

  @staticmethod
  def square_grid(length: int) -> Grid[T]:
    return Grid(length, length)

  @staticmethod
  def grid(width: int, height: int) -> Grid[T]:
    return Grid(width, height)
  
  def set(self, r: int, c: int, val: T):
    self.data[c][c].value = val
    
  def get(self, r: int, c: int) -> Data[T]:
    return self.data[r][c]
  
  def visit(self, r: int, c: int):
    self.data[r][c].visited = True
    
  def visited(self, r: int, c: int) -> bool:
    return self.data[r][c].visited
  
  def get_value(self, r: int, c: int) -> T:
    return self.data[r][c].value
  
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
    for i in range(r - 1, r + 2):
      for j in range(c - 1, c + 2):
        if i == r and j == c:
          continue
        if i < 0 or i >= self.height:
          continue
        if j < 0 or j >= self.width:
          continue
        neighbors.append((i, j))

  def visit_neighbors(self, r: int, c: int):
    for n in self.neighbors(r, c):
      self.visit(n[0], n[1])