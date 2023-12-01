from utils import read_file
import asyncio

async def main():
    
    possible = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "1": 1,
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
    }
    sum = 0
    async for line in read_file("day1input"):
        digits = {}
        for digit in possible.keys():
            if digit in line:
                idx = line.index(digit)
                digits[idx] = digit
        assert len(digits) > 0


        sorted_keys = sorted(digits.keys())
        # print(sorted_keys)
        first = possible[digits[sorted_keys[0]]]
        last = possible[digits[sorted_keys[-1]]]

        assert first is not None 
        assert last is not None
        val = (first * 10) + last
        print(line, first, last, val)

        sum += val
            
    print(sum)

if __name__ == "__main__":
    asyncio.run(main())
    
    # 54629 too high