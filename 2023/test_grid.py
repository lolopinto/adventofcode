from grid import Grid
import pytest

def test_rotate_right_square_grid():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_right()
  assert g2.current_lines() == ["369", "258", "147"]

def test_rotate_left_square_grid():
  lines = """
123
456
789""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_left()
  assert g2.current_lines() == ["147", "258", "369"]

def test_rotate_right_rect_grid():
  lines = """
123
456""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_right()
  assert g2.current_lines() == ["36", "25", "14"]
  
def test_rotate_left_rect_grid():
  lines = """
123
456""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_left()
  assert g2.current_lines() == ["14", "25", "36"]
  
def test_rotate_right_rect_grid2():
  lines = """
123
456
789
abc
def
ghi
""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_right()
  assert g2.current_lines() == ["369cfi", "258beh", "147adg"]
  
def test_rotate_left_rect_grid2():
  lines = """
123
456
789
abc
def
ghi
""".strip().splitlines()
  g = Grid.from_lines(lines)
  g2 = g.rotate_left()
  assert g2.current_lines() == ["147adg", "258beh", "369cfi"]