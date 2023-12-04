from collections import defaultdict
from utils import read_file
from grid import Grid
import asyncio

def check(g: Grid, num: int, coords: list[tuple[int, int]]) -> (bool, str | None, tuple[int, int] | None):
  for r, c in coords:
    neighbors = g.neighbors8(r, c)
    found = False
    for r2, c2 in neighbors:
      if is_symbol(g, r2, c2):
        return True, (r2, c2), g.get_value(r2, c2)
  return False, None, None


def is_symbol(g: Grid[str], r: int, c:int) -> bool:
  v = g.get_value(r, c)
  if v.isdigit():
    return False
  if v == ".":
    return False
  return True


async def part1():
  g = await Grid.square_grid_from_file("day3input")

  curr = None
  l = []
  sum = 0
  for r in range(g.height):
    for c in range(g.width):
      v = g.get_value(r, c)
      assert v is not None
      if v.isdigit():
        if curr is None:
          curr = []
        curr.append(v)
        l.append((r, c))
      elif curr is not None:
        num = int("".join(curr))
        is_part, _, _ = check(g, num, l)
        if is_part:
          sum += num
          
        curr = None
        l = []

  print(sum)


async def part2():
  g = await Grid.square_grid_from_file("day3input")

  curr = None
  l = []
  sum = 0
  stars = defaultdict(list)
  for r in range(g.height):
    for c in range(g.width):
      v = g.get_value(r, c)
      assert v is not None
      if v.isdigit():
        if curr is None:
          curr = []
        curr.append(v)
        l.append((r, c))
      elif curr is not None:
        num = int("".join(curr))
        is_part, star_coord, symbol = check(g, num, l)
        if is_part and symbol == '*':
          stars[star_coord].append(num)
        curr = None
        l = []

  for v in stars.values():
    if len(v) == 2:
      sum += v[0] * v[1]
  print(sum)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
