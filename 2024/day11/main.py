from functools import cache

@cache
def blink(n,remaining):
    if remaining == 0:
        return 1

    remaining -= 1

    if n == 0:
        return blink(1,remaining)
    elif len(f"{n}") % 2 == 0:
        n = f"{n}"
        firstHalf = blink(int(n[:len(n)//2]), remaining)
        secondHalf = blink(int(n[len(n)//2:]), remaining)
        return firstHalf + secondHalf
    else:
        return blink(n * 2024, remaining)

def partOne(filename):
    total = 0
    stones = []
    with open(filename) as fh:
        for line in fh:
            stones = [int(x) for x in line.strip().split()]

    for s in stones:
        total += blink(int(s), 25)

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    stones = []
    with open(filename) as fh:
        for line in fh:
            stones = [int(x) for x in line.strip().split()]

    for s in stones:
        total += blink(int(s), 75)

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")