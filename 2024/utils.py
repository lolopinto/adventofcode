from aiofiles import open
from typing import AsyncGenerator
import re

async def read_file(file: str) -> AsyncGenerator[str, None]:
  async with open(file) as f:
    async for line in f:
      yield line.strip()
      
async def read_file_lines(file: str) -> list[str]:
  return [line async for line in read_file(file)]

async def read_file_groups(file: str) -> AsyncGenerator[list[str], None]:
  l = []
  async with open(file) as f:
    async for line in f:
      line = line.strip()
      if line == "":
        yield l
        l = []
      else:
        l.append(line)
  if l:
    yield l
    
async def get_file_groups(file: str)-> list[list[str]]:
  groups = []

  async for group in read_file_groups(file):
    groups.append(group)

  return groups

async def read_file_chunks(file: str, length: int) -> AsyncGenerator[list[str], None]:
  l = []
  next_empty = False
  async with open(file) as f:
    async for line in f:
      line = line.strip()
      if next_empty:
        assert line == ""
        next_empty = False
        continue
      
      l.append(line)
      if len(l) == length:
        yield l
        l = []
        next_empty = True


def ints(s: str, delimiter : str | None = None)-> list[int]:
  return [int(v) for v in (s.split(delimiter) if delimiter is not None else s.split())]

def ints_list(s: list[str])-> list[int]:
  return [int(v) for v in s]


# TODO grid implementation

# TODO graph implementation

# TODO cube implementation