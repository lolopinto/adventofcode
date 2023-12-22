from __future__ import annotations
from dataclasses import dataclass
from typing import Any
import math 
from grid import Grid

@dataclass
class InfiniteGrid:
  data: dict[tuple[int, int], Any]
  initial_width: int
  initial_height: int
  
  def __init__(self, *, initial_width: int, initial_height: int):
    self.data = {}
    self.initial_height = initial_height
    self.initial_width = initial_width
    
  @staticmethod
  def from_grid(g: Grid) -> InfiniteGrid:
    ret = InfiniteGrid(initial_width=g.width, initial_height=g.height)
    for pos in g.walk():
      ret.set(pos, g.get_value(pos[0], pos[1]))
    return ret

  def set(self, pos: tuple[int, int], v: Any):
    self.data[pos] = v
      
  # neighbors in infinite grid go forever
  def neighbors(self, pos: tuple[int, int]) -> list[tuple[int, int]]:
    # gets fields that exist, if they don't, it sets it based on the initial 
    # 5 * 5 grid to start
    
    ret = []
    
    r, c = pos
    
    neighbors = [
      (r - 1, c),
      (r + 1, c),
      (r, c - 1),
      (r, c + 1),
    ]
    for i, n in enumerate(neighbors):
      r2, c2 = n
      
      change = False
      if r2 < 0 or r2 >= self.initial_width:
        r2 = r2 % self.initial_width
        change = True
      if c2 < 0 or c2 >= self.initial_height:
        c2 = c2 % self.initial_height
        change = True

      # update grid with new values that are being added 
      if change:
        self.data[n] = self.get_value((r2, c2))

    return neighbors
    
  def get_valuex(self, pos: tuple[int, int]) -> Any:
    return self.data[pos]

  def get_value(self, pos: tuple[int, int], default_value=None) -> Any:
    if pos in self.data:
      return self.data[pos]
    return default_value
  
  def len(self) -> int:
    return len(self.data)

  # TODO walk, print etc may not make sense, can take from sparse_grid as it does
