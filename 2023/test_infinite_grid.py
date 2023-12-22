from grid import Grid
from infinite_grid import InfiniteGrid
import pytest

def test_infinite_grid():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  inf = InfiniteGrid.from_grid(g)
  assert inf.initial_height == g.height
  assert inf.initial_width == g.width
  assert inf.len() == g.width * g.height
  
def test_simple_neighbors():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  inf = InfiniteGrid.from_grid(g)
  assert inf.initial_height == g.height
  assert inf.initial_width == g.width
  
  n = inf.neighbors((1, 1))
  assert len(n) == 4
  assert inf.len() == g.width * g.height

def test_simple_overlapping_neighbor():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  inf = InfiniteGrid.from_grid(g)
  assert inf.initial_height == g.height
  assert inf.initial_width == g.width
  
  # left corner neighbor has added 2
  n = inf.neighbors((0, 0))
  assert len(n) == 4
  assert inf.len() - (g.width * g.height) == 2
  
def test_simple_overlapping_neighbor2():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  inf = InfiniteGrid.from_grid(g)
  assert inf.initial_height == g.height
  assert inf.initial_width == g.width
  
  # left middle neighbor has added 1
  n = inf.neighbors((1, 0))
  assert len(n) == 4
  assert inf.len() - (g.width * g.height) == 1
  
  # g2 = g.rotate_right()
  # assert g2.current_lines() == ["369", "258", "147"]

# def test_rotate_left_square_grid():
#   lines = """
# 123
# 456
# 789""".strip().splitlines()
#   g = Grid.from_lines(lines)
#   g2 = g.rotate_left()
#   assert g2.current_lines() == ["147", "258", "369"]

# def test_rotate_right_rect_grid():
#   lines = """
# 123
# 456""".strip().splitlines()
#   g = Grid.from_lines(lines)
#   g2 = g.rotate_right()
#   assert g2.current_lines() == ["36", "25", "14"]
  
# def test_rotate_left_rect_grid():
#   lines = """
# 123
# 456""".strip().splitlines()
#   g = Grid.from_lines(lines)
#   g2 = g.rotate_left()
#   assert g2.current_lines() == ["14", "25", "36"]
  
# def test_rotate_right_rect_grid2():
#   lines = """
# 123
# 456
# 789
# abc
# def
# ghi
# """.strip().splitlines()
#   g = Grid.from_lines(lines)
#   g2 = g.rotate_right()
#   assert g2.current_lines() == ["369cfi", "258beh", "147adg"]
  
# def test_rotate_left_rect_grid2():
#   lines = """
# 123
# 456
# 789
# abc
# def
# ghi
# """.strip().splitlines()
#   g = Grid.from_lines(lines)
#   g2 = g.rotate_left()
#   assert g2.current_lines() == ["147adg", "258beh", "369cfi"]