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
    winners = set(int(v) for v in parts2[0].split(" ") if v != "")
    mine = set(int(v) for v in parts2[1].split(" ") if v != "")
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
    cards_number = int(match.group(1))
    parts2 = parts[1].split(" | ")
    winners = set(int(v) for v in parts2[0].split(" ") if v != "")
    mine = set(int(v) for v in parts2[1].split(" ") if v != "")
    
    wins = len(winners.intersection(mine))
    card = Card(cards_number, wins)
    all_cards[cards_number] = card
    deck.counts[cards_number] = 1
    
  for card_number in all_cards.keys():
    card = all_cards[card_number]
    
    multiplier = deck.counts[card_number]

    for i in range(1, card.wins + 1):
      new_number = card_number + i
      deck.counts[new_number] += multiplier
    
  sum = 0
  for count in deck.counts.values():
    sum += count
  print(sum)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
