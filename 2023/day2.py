from utils import read_file
import asyncio
import re
from dataclasses  import dataclass
# from typinq import T, L


@dataclass
class Group:
  red: int
  blue: int
  green: int

@dataclass
class Game:
  id: int
  groups: list[Group]
  
  def possible(self):
    red = 0
    blue = 0
    green = 0
    for g in self.groups:
      red += g.red
      blue += g.blue
      green += g.green
      
    return red <= 12 and blue <= 14 and green <= 13
  
def parse_line(line: str) -> Game:
  # print(line)
  parts = line.split(":")
  assert len(parts) == 2

  g = r"Game (\d+)"
  match = re.match(g, parts[0])
  assert match is not None

  id = int(match.group(1))
  groups = []
  for part in parts[1].strip().split(";"):
    g = Group(red=0, blue=0, green=0)
    for group in part.strip().split(","):
      group = group.strip()
      parts = group.split(" ")
      count = int(parts[0])
      color = parts[1]
      assert color in ["red", "blue", "green"]  
      match color:
        case "red":
          g.red = count
        case "blue":
          g.blue = count
        case "green":
          g.green = count
        case _:
          raise Exception("Unknown color: " + color)
    groups.append(g)
  return Game(id, groups)
  

async def part1():
  sum = 0
  async for line in read_file("day2input"):
    g = parse_line(line)
    # print(g)
    if g.possible():
      sum += g.id
  print(sum)


async def part2():
  pass

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
