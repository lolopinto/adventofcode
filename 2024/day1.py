from utils import read_file
import asyncio


async def part1():
  l1 = []
  l2 = []
  async for line in read_file("day1input"):
    parts = line.split()
    l1.append(int(parts[0]))
    l2.append(int(parts[1]))

  l1.sort()  
  l2.sort()
  
  ans = 0
  for i1, i2 in zip(l1, l2):
    ans += max(i1, i2) - min(i1, i2)
    
  print(ans)


if __name__ == "__main__":
    asyncio.run(part1())
    # asyncio.run(part2())