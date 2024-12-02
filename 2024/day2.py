from utils import read_file, ints
import asyncio
from collections import defaultdict

def is_safe(l: list[int]) -> bool:
  ascending = l[:]
  descending = l[:]
  
  ascending.sort()
  descending.sort(reverse=True)
  
  if l != ascending and l != descending:
    # not sorted in either direction, bye
    return False
  
  for i, item in enumerate(l):
    if i == 0:
      continue
    prev = l[i - 1]
    diff = abs(item - prev)
    if diff < 1 or diff > 3:
      return False

  return True

async def part1():
    ct = 0
    async for line in read_file("day2input"):
      if is_safe(ints(line)):
        ct += 1
        
    print(ct)

async def part2():
    ct = 0
    async for line in read_file("day2input"):
      l = ints(line)
      to_pop = [None]
      if is_safe(l):
        ct += 1
        continue

      orig = l[:]
      for i in range(len(orig)):
        l = orig[:]
        l.pop(i)
        if is_safe(l):
          ct += 1
          break
    print(ct)


if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())