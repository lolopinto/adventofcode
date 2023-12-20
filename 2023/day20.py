from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass, field
import re
from grid import Grid
import itertools
import enum

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
  def press_button(modules: dict[str, Module], *, find_low_rx=False):
    to_process = []
    # source, destination, pulse
    to_process.append(('button', 'broadcaster', PulseType.LOW))

    low = 0
    high = 0
    while len(to_process) > 0:
      source, destination, pulse_type = to_process.pop(0)
      
      # print(f"Sending {pulse_type} from {source} to {destination}")
      # print(f"{source} -> {pulse_type.value} -> {destination}")

      if pulse_type == PulseType.HIGH:
        high += 1
      else:
        low += 1

      if destination not in modules:
        continue

      m = modules[destination]
      pulse_to_send = m.receive_pulse(pulse_type, modules, source)
      
      # if find_low_rx and pulse_to_send == PulseType.HIGH and m.name == 'rx':
      #   return True
      if find_low_rx and pulse_to_send == PulseType.LOW and m.name == 'rx':
        return True

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

  # rs -> rx  
  print(conjunctions)
  keys = modules['rs'].received.keys()

  count = 1
  while True:
    ret = Module.press_button(modules, find_low_rx=True)
    all_off = all(m.pulse_state == PulseState.OFF for m in modules.values())
    all_on = all(m.pulse_state == PulseState.ON for m in modules.values())
    if all_off:
      print(f"all off after {count} presses")

    if all_on:
      print(f"all on after {count} presses")

    # for k in keys:
    #   # print(f"{k} is {modules[k].name}")
    #   if modules[k].pulse_state == PulseState.ON:
    #     print(f"{k} is on")

    count += 1
    
    if count == 100_000:
      break

  print(count)

if __name__ == "__main__":
    # asyncio.run(part1())
    asyncio.run(part2())
