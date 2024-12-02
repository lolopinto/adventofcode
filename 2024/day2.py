from utils import read_file, ints
import asyncio
from collections import defaultdict


async def part1():
    ct = 0
    async for line in read_file("day2input"):
      l = ints(line)
      ascending = l[:]
      descending = l[:]
      
      ascending.sort()
      descending.sort(reverse=True)
      
      if l != ascending and l != descending:
        # not sorted in either direction, bye
        continue
      
      safe = True
      for i, item in enumerate(l):
        if i == 0:
          continue
        prev = l[i - 1]
        diff = abs(item - prev)
        if diff < 1 or diff > 3:
          safe = False
          break

      if safe:
        ct += 1
        
    print(ct)

async def part2():
  pass


if __name__ == "__main__":
    asyncio.run(part1())
    # asyncio.run(part2())