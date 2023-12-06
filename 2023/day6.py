from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import math

@dataclass()
class Race:
  time: int
  distance: int
  
  def ways_to_win(self) -> int:
    ct = 0
    for i in range(0, self.time):
      time_left = self.time -i 
      distance_traveled = time_left * i 
      if distance_traveled > self.distance:
        ct += 1
    return ct

async def part1():
  races = []
  async for line in read_file("day6input"):
    parts = line.split(": ")
    match parts[0]:
      case "Time":
        times = ints(parts[1])
        for i in range(len(times)):
          races.append(Race(times[i], 0))
        
      case "Distance":
        distances = ints(parts[1])
        for i in range(len(distances)):
          race = races[i]
          assert race is not None
          race.distance = distances[i]
        
  # print(races)
  print(math.prod([race.ways_to_win() for race in races]))


async def part2():
  async for line in read_file("day6input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
