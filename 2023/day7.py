from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import enum
import functools

class HandType(enum.Enum):
  FIVE_OF_A_KIND = 1
  FOUR_OF_A_KIND = 2
  FULL_HOUSE = 3
  THREE_OF_A_KIND = 4
  TWO_PAIR = 5
  ONE_PAIR = 6
  HIGH_CARD = 7

hand_strength = {
  'A': 14,
  'K': 13,
  'Q': 12,
  'J': 11,
  'T': 10,
  '9': 9,
  '8': 8,
  '7': 7,
  '6': 6,
  '5': 5,
  '4': 4,
  '3': 3,
  '2': 2
}

@dataclass
class Hand:
  cards: str
  bid: int
  hand_type: HandType
  
  @staticmethod
  def determine_hand_type(cards: str) -> HandType:
    assert len(cards) == 5
    counts = defaultdict(int)
    for card in cards:
      counts[card] += 1
      
    match len(counts):
      case 1:
        return HandType.FIVE_OF_A_KIND
      case 2:
        if 4 in counts.values():
          return HandType.FOUR_OF_A_KIND
        else:
          return HandType.FULL_HOUSE
      case 3:
        if 3 in counts.values():
          return HandType.THREE_OF_A_KIND
        else:
          return HandType.TWO_PAIR
      case 4:
        return HandType.ONE_PAIR
      case _:
        return HandType.HIGH_CARD

  @staticmethod
  def cmp(obj1: Hand, obj2: Hand) -> int:
    if obj1.hand_type.value != obj2.hand_type.value:
      return obj1.hand_type.value - obj2.hand_type.value

    for c1, c2 in zip(obj1.cards, obj2.cards):
      if c1 != c2:
        # bigger is better?
        return hand_strength[c2] - hand_strength[c1]
        # return ord(c1) - ord(c2)

    return 0

async def part1():
  hands = []
  async for line in read_file("day7input"):
    cards, bid = line.split()
    hand_type = Hand.determine_hand_type(cards)
    card = Hand(cards, int(bid), hand_type)
    hands.append(card)
    
  key = functools.cmp_to_key(Hand.cmp)
  sorted_hands = sorted(hands, key=key)
  
  result = 0
  for i in range(len(sorted_hands)):
    result += sorted_hands[i].bid * (len(sorted_hands) - i)

  print(result)

async def part2():
  async for line in read_file("day7input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
