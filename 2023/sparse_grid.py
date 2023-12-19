from dataclasses import dataclass
from typing import Any
import math 

@dataclass
class SparseGrid:
  data: dict[tuple[int, int], int]
  
  def __init__(self):
    self.data = {}
    self.min_row = None
    self.max_row = None
    self.min_col = None
    self.max_col = None
    
  def set(self, pos: tuple[int, int], v: Any):
    self.data[pos] = v
    if self.min_row is None:
      self.min_row = pos[0]
      self.max_row = pos[0]
      self.min_col = pos[1]
      self.max_col = pos[1]
    else:
      self.min_row = min(self.min_row, pos[0])
      self.max_row = max(self.max_row, pos[0])
      self.min_col = min(self.min_col, pos[1])
      self.max_col = max(self.max_col, pos[1])
    
  def get_value(self, pos: tuple[int, int], default_value=None) -> Any:
    if pos in self.data:
      return self.data[pos]
    return default_value
    # return self.data.get(pos, default_value)
  
  def len(self) -> int:
    return len(self.data)
  
  # @property
  # def min_row(self) -> int:
  #   return min(r for r, _ in self.data.keys())
  
  # @property
  # def max_row(self) -> int:
  #   return max(r for r, _ in self.data.keys())
  
  # @property
  # def min_col(self) -> int:
  #   return min(c for _, c in self.data.keys())
  
  # @property
  # def max_col(self) -> int:
  #   return max(c for _, c in self.data.keys())
  
  # TODO need a walk that only walks existing values
  # using keys and values and sorting
  def walk(self) -> tuple[int, int]:
    for r in range(self.min_row, self.max_row + 1):
      for c in range(self.min_col, self.max_col + 1):
        yield (r, c)
        
  def in_bounds(self, pos: tuple[int, int]) -> bool:
    r, c = pos
    return r >= self.min_row and r <= self.max_row and c >= self.min_col and c <= self.max_col
  
  
  def print(self):
    for (r, c) in g.walk():
      if (r, c) in self.data:
        print('#', end='')
      else:
        print('.', end='')
      print()