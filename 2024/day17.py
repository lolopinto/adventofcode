from __future__ import annotations
from utils import read_file_groups, read_file, ints, ints_list
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math


def literal_value(registers, operand):
  return operand

def combo_value(registers, operand):
  match operand:
    case 0 | 1 | 2 | 3:
      return operand
    case 4:
      return registers["A"]
    case 5:
      return registers["B"]
    case 6:
      return registers["C"]
    case _:
      raise ValueError(f"Invalid register {operand}")

def perform_operator(registers, opcode, operand) -> dict:
  match opcode:
    case 0 | 6 | 7:
      numerator = registers["A"]
      denominator = 2 ** combo_value(registers, operand)
      result = numerator // denominator
      
      if opcode == 0:
        registers["A"] = result
      elif opcode == 6:
        registers["B"] = result
      else:
        registers["C"] = result
      
    case 1:
      registers["B"] = registers["B"] ^ literal_value(registers, operand)
      
    case 2:
      registers["B"] = combo_value(registers, operand) % 8
      
    case 3:
      if registers["A"] != 0:
        return {
          "jumps": literal_value(registers, operand)
        }
      
    case 4:
      registers["B"] = registers["B"] ^ registers["C"]
      
    case 5:
      value = combo_value(registers, operand) % 8
      return {
        "outputs": value,
      }
      
  return {}
    

def run_program(registers, program):
  inputs = ints(program, ",")

  i = 0  
  
  outputs = []
  while True:
    if i >= len(inputs):
      break
    
    opcode, operand = inputs[i], inputs[i + 1]
    
    result = perform_operator(registers, opcode, operand)
    if "jumps" in result:
      i = result["jumps"]
    else:
      i += 2

    if "outputs" in result:
      outputs.append(str(result["outputs"]))
      
  return ",".join(outputs)

async def load():
  groups = [group async for group in read_file_groups("day17input")]

  assert len(groups) == 2
  
  registers = {}
  
  for register in groups[0]:
    m = re.search("Register (.): (.+)", register)
    assert m is not None
    registers[m.group(1)] = int(m.group(2))

  assert len(registers) == 3
  assert len(groups[1]) == 1
  
  program = groups[1][0]
  program = program.split(": ")[1]
  return registers, program

async def part1():
  registers, program = await load()
  
  print(run_program(registers, program))

async def part2():
  registers, program = await load()

  a_start = 0
  while True:
    current = registers.copy()
    current["A"] = a_start
    
    if program == run_program(current, program):
      break
    a_start += 1
    if a_start % 1_000_000 == 0:
      print(a_start)
    
  print(a_start)


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
