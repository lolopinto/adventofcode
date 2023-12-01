from utils import read_file
import asyncio

async def main():
    sum = 0
    async for line in read_file("day1input"):
        digits = []
        for c in line:
            if c.isdigit():
                digits.append(int(c))
        assert len(digits) > 0
        val = digits[0] * 10 + digits[-1]
        sum += val
            
    print(sum)

if __name__ == "__main__":
    asyncio.run(main())