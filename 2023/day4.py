from utils import read_file
import asyncio

async def part1():
  sum = 0
  async for line in read_file("day4input"):
    parts = line.split(": ")
    parts2 = parts[1].split(" | ")
    # print(parts2)
    winners = set(int(v.strip()) for v in parts2[0].split(" ") if v != "")
    mine = set(int(v.strip()) for v in parts2[1].split(" ") if v != "")
    # print(winners, mine, winners.intersection(mine))
    if len(winners.intersection(mine)) > 0:
      sum += 2 ** (len(winners.intersection(mine)) -1)

  print(sum)


async def part2():
  async for line in read_file("day4input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
