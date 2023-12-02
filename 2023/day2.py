from utils import read_file
import asyncio
import re
from dataclasses import dataclass

@dataclass
class Group:
  red: int
  blue: int
  green: int

@dataclass
class Game:
  id: int
  sets: list[Group]
  
  def possible(self) -> bool:
    for g in self.sets:
      if g.red > 12 or g.blue > 14 or g.green > 13:
        return False

    return True      
  
  # the language here was confusing and took me a while to understand what was being asked
  def power(self) -> int:
    reds = []
    greens = []
    blues = []
    for g in self.sets:
      reds.append(g.red)
      greens.append(g.green)
      blues.append(g.blue)
    return max(reds) * max(greens) * max(blues)

def parse_line(line: str) -> Game:
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
    if g.possible():
      sum += g.id
  print(sum)

async def part2():
  sum = 0
  async for line in read_file("day2input"):
    g = parse_line(line)
    sum += g.power()
  print(sum)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
