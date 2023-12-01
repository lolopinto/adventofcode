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
        digits_first = {}
        digits_last = {}
        for digit in possible.keys():
            first_idx = line.find(digit)
            if first_idx > -1:
                digits_first[first_idx] = digit
            last_idx = line.rfind(digit)
            if last_idx > -1:
                digits_last[last_idx] = digit
        assert len(digits_first) > 0
        assert len(digits_last) > 0


        sorted_first = sorted(digits_first.keys())
        sorted_last = sorted(digits_last.keys())
        first = possible[digits_first[sorted_first[0]]]
        last = possible[digits_last[sorted_last[-1]]]

        assert first is not None 
        assert last is not None
        val = (first * 10) + last
        print(line, first, last, val)

        sum += val
            
    print(sum)

if __name__ == "__main__":
    asyncio.run(main())
    
    # 54629 too high