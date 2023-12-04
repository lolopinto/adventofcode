from utils import read_file
from dataclasses import dataclass
from collections import defaultdict
import asyncio
import re

async def part1():
  sum = 0
  async for line in read_file("day4input"):
    parts = line.split(": ")
    parts2 = parts[1].split(" | ")
    winners = set(match.group(0) for match in re.finditer(r"\d+", parts2[0]))
    mine = set(match.group(0) for match in re.finditer(r"\d+", parts2[1]))
    if len(winners.intersection(mine)) > 0:
      sum += 2 ** (len(winners.intersection(mine)) -1)

  print(sum)


@dataclass
class Card():
  number: int
  wins: int 
  
@dataclass
class Deck():
  counts: defaultdict[int, int]
  
async def part2():
  all_cards = {}
  deck = Deck(defaultdict(int))
  async for line in read_file("day4input"):
    parts = line.split(": ")
    match = re.match(r"Card +(\d+)", parts[0])
    assert match is not None
    card_number = int(match.group(1))
    parts2 = parts[1].split(" | ")
    winners = set(match.group(0) for match in re.finditer(r"\d+", parts2[0]))
    mine = set(match.group(0) for match in re.finditer(r"\d+", parts2[1]))
    
    wins = len(winners.intersection(mine))
    card = Card(card_number, wins)
    all_cards[card_number] = card
    deck.counts[card_number] = 1
    
  for card_number in all_cards.keys():
    card = all_cards[card_number]
    
    multiplier = deck.counts[card_number]

    for i in range(1, card.wins + 1):
      new_number = card_number + i
      deck.counts[new_number] += multiplier
    
  print(sum(count for count in deck.counts.values()))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
