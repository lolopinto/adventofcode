from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass, field
import re
from grid import Grid
import itertools
import enum
import math

class PulseType(enum.Enum):
  HIGH = 'high'
  LOW = 'low'

class PulseState(enum.Enum):
  ON = 1
  OFF = 2
  
@dataclass
class Module:
  name: str
  conjunction: bool
  flip_flop: bool
  destinations: list[str]
  
  high_pulses: int = 0
  low_pulses: int = 0
  # last_pulse: Pulse | None = None
  pulse_state: PulseState = PulseState.OFF
  received: defaultdict[str, PulseType ] = field(default_factory= lambda: defaultdict(lambda x: PulseType.LOW))
  
  def receive_pulse(self, pulse: PulseType, modules: dict[str, Module], source: str):
    pulse_to_send = pulse
    if self.flip_flop:
      match pulse:
        case PulseType.HIGH:
          # nothing to do
          pulse_to_send = None

        case PulseType.LOW:
          
          match self.pulse_state:
            case PulseState.ON:
              self.pulse_state = PulseState.OFF
              pulse_to_send = PulseType.LOW

            case PulseState.OFF:
              self.pulse_state = PulseState.ON
              pulse_to_send = PulseType.HIGH              

    if self.conjunction:
      self.received[source] = pulse
      if all(v == PulseType.HIGH for v in self.received.values()):
        pulse_to_send = PulseType.LOW
      else:
        pulse_to_send = PulseType.HIGH
        
    return pulse_to_send;
          
  @staticmethod
  def press_button(modules: dict[str, Module], *, find_high_pulse_for=None):
    to_process = []
    # source, destination, pulse
    to_process.append(('button', 'broadcaster', PulseType.LOW))

    low = 0
    high = 0
    while len(to_process) > 0:
      source, destination, pulse_type = to_process.pop(0)

      if pulse_type == PulseType.HIGH:
        high += 1
      else:
        low += 1

      if destination not in modules:
        continue

      m = modules[destination]
      pulse_to_send = m.receive_pulse(pulse_type, modules, source)

      
      if find_high_pulse_for is not None and \
        m.name in find_high_pulse_for and \
        pulse_to_send == PulseType.HIGH:
        # print(f'source {source} {m.name} {find_high_pulse_for}')
        return m.name

      if pulse_to_send is None:
        continue
      

      for d in m.destinations:
        to_process.append((m.name, d, pulse_to_send))

    return (low, high)      
  
  @staticmethod
  def parse(line: str):
    parts = line.split(' -> ')
    assert len(parts) == 2

    left = parts[0]
    flip_flop = False
    conjunction = False

    match left[0]:
      case '%':
        name = left[1:]
        flip_flop = True
      case '&':
        name = left[1:]
        conjunction = True
      case _:
        name = left

    return Module(name, conjunction, flip_flop, parts[1].split(', '))        
        
async def parse_input():
  modules = {}

  conjunctions = set()  
  async for line in read_file("day20input"):
    m = Module.parse(line)
    if m.conjunction:
      conjunctions.add(m.name)
    modules[m.name] = m
    
  # set up conjunctions correctly
  for name in conjunctions:
    m = modules[name]
    for m2 in modules.values():
      if name in m2.destinations:
        m.received[m2.name] = PulseType.LOW
       
  return modules, conjunctions 

async def part1():
  modules, _ = await parse_input()

  total_high = 0
  total_low = 0
  
  for i in range(1000):
    low, high = Module.press_button(modules)
    total_low += low
    total_high += high
    
  print(total_low * total_high)
  

async def part2():
  modules, conjunctions = await parse_input()

  # this one randomly depends on the fact that the input 
  # seems to be 4 conjunctions -> 1 conjunction -> output so we can
  # depend on the fact that a conjunction sends low when it's received all highs
  # so we'll look for when each conjunction which is the first in the chain receives
  # it's high and then assume no overlaps and then find the lcm of the number of button
  # presses it takes to get each for the first time as the time they all converge and have all 
  # high which means they have all sent a high to rs which then sends a low to rx 
  # i was on the right track for the first day i attempted this but gave up :|
  # too many tricks this year
  
  # rs -> rx
  rx_received = [m.name for m in modules.values() if 'rx' in m.destinations]
  assert len(rx_received) == 1
  assert modules[rx_received[0]].conjunction is True
  # bt, dl, fr, rv
  keys = modules[rx_received[0]].received.keys()
  # print(keys)
 
  # all conjunctions
  assert all(modules[k].conjunction for k in keys)

  values = {}
  count = 1
    
  while True:
    # not sure why i'm seeing repeateds in the way i wrote this 
    # but i had to change this from a list to a map so it only works
    # for the first time it sees each one
    ret = Module.press_button(modules, find_high_pulse_for=keys)
    if isinstance(ret, str) and ret not in values:
      values[ret] = count
    count += 1

    if len(values) == len(keys):
      break

  print(math.lcm(*values.values()))
  

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
