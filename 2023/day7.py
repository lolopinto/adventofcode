from __future__ import annotations
from utils import read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
import enum
import functools

class HandType(enum.IntEnum):
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

hand_strength2 = {
  'A': 14,
  'K': 13,
  'Q': 12,
  'T': 10,
  '9': 9,
  '8': 8,
  '7': 7,
  '6': 6,
  '5': 5,
  '4': 4,
  '3': 3,
  '2': 2,
  'J': 1,
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
      
    counts_sorted = sorted(counts.values(), reverse = True)
    match counts_sorted:
      case [5]:
        return HandType.FIVE_OF_A_KIND
      case [4, 1]:
          return HandType.FOUR_OF_A_KIND
      case [3, 2]:
          return HandType.FULL_HOUSE
      case [3, 1, 1]:
          return HandType.THREE_OF_A_KIND
      case [2, 2, 1]:
          return HandType.TWO_PAIR
      case [2, 1, 1, 1]:
        return HandType.ONE_PAIR
      case _:
        return HandType.HIGH_CARD

  @staticmethod
  def determine_hand_type2(cards: str) -> HandType:
    if not 'J' in cards or cards == 'JJJJJ':
      return Hand.determine_hand_type(cards)
    
    counts = defaultdict(int)
    for card in cards:
      counts[card] += 1
      
    not_j = [key for key in counts.keys() if key != 'J']
    hand_types = []
    for key in not_j:
      hand_types.append(Hand.determine_hand_type(cards.replace('J', key)))

    return min(hand_types)

  @staticmethod
  def cmp(obj1: Hand, obj2: Hand, for_joker=False) -> int:
    if obj1.hand_type.value != obj2.hand_type.value:
      return obj1.hand_type.value - obj2.hand_type.value

    for c1, c2 in zip(obj1.cards, obj2.cards):
      if c1 != c2:
        if for_joker:
          return hand_strength2[c2] - hand_strength2[c1]
        return hand_strength[c2] - hand_strength[c1]

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
  hands = []
  async for line in read_file("day7input"):
    cards, bid = line.split()
    hand_type = Hand.determine_hand_type2(cards)
    card = Hand(cards, int(bid), hand_type)
    hands.append(card)
    
  key = functools.cmp_to_key(lambda obj1, obj2: Hand.cmp(obj1, obj2, True))
  sorted_hands = sorted(hands, key=key)
  
  result = 0
  for i in range(len(sorted_hands)):
    result += sorted_hands[i].bid * (len(sorted_hands) - i)

  print(result)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
