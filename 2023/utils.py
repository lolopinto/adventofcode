from aiofiles import open

async def read_file(file: str):
  async with open(file) as f:
    async for line in f:
      yield line.strip()