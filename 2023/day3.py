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

async def init_grid() -> Grid[str]:
  init = False
  g = None
  r = 0

  # TODO better way to just init from a grid if we know it's a grid
  # e.g. just do all this from input file
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

async def part1():
  g = await init_grid()

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

# 316970 too low
# 532580 too low
def is_symbol(g: Grid[str], r: int, c:int) -> bool:
  v = g.get_value(r, c)
  if v.isdigit():
    return False
  if v == ".":
    return False
  return True

async def part2():
  g = await init_grid()

  curr = None
  l = []
  sum = 0
  stars = {}
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
          l = stars.get(star_coord) or []
          l.append(num)
          stars[star_coord] = l
        curr = None
        l = []

  for k, v in stars.items():
    if len(v) == 2:
      sum += v[0] * v[1]
  print(sum)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
