import re

def partOne(filename):
    total = 0
    with open(filename) as fh:
        for line in fh:
            muls = re.findall(r"mul\(([0-9]+),([0-9]+)\)",line)
            for seq in muls:
                first = int(seq[0])
                second = int(seq[1])
                total += first * second
    print(f"Part One: {total}")

def partTwo(filename):
    enabled = True
    total = 0
    with open(filename) as fh:
        for line in fh:
            matches = re.findall(r"(mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\))",line)
            for seq in matches:
                if seq == "do()":
                    enabled = True
                    continue
                elif seq == "don't()":
                    enabled = False
                    continue

                if enabled:
                    seq = seq.strip("mul(").strip(")").split(',')
                    first = int(seq[0])
                    second = int(seq[1])
                    total += first * second

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")