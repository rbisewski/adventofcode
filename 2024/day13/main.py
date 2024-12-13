import re

input = []

def partOne(filename):
    total = 0

    with open(filename) as fh:
        buttonA = (0,0)
        buttonB = (0,0)
        prize = (0,0)
        for line in fh:
            aMatches = re.findall(r"Button A: X\+([0-9]+), Y\+([0-9]+)",line)
            if aMatches:
                x = int(aMatches[0][0])
                y = int(aMatches[0][1])
                buttonA = (x,y)
                continue

            bMatches = re.findall(r"Button B: X\+([0-9]+), Y\+([0-9]+)",line)
            if bMatches:
                x = int(bMatches[0][0])
                y = int(bMatches[0][1])
                buttonB = (x,y)
                continue

            pMatches = re.findall(r"Prize: X\=([0-9]+), Y\=([0-9]+)",line)
            if pMatches:
                x = int(pMatches[0][0])
                y = int(pMatches[0][1])
                prize = (x,y)
                input.append({
                    "Button A": buttonA,
                    "Button B": buttonB,
                    "Prize": prize
                })
                continue

    # Button A costs 3 tokens
    # Button B costs 1 token
    allPrizes = []
    for i in input:
        ax = i["Button A"][0]
        ay = i["Button A"][1]

        bx = i["Button B"][0]
        by = i["Button B"][1]

        px = i["Prize"][0]
        py = i["Prize"][1]

        cheapest = {
            "A Pushes": -1,
            "B Pushes": -1,
            "Cost": -1
        }

        first = True
        for apushes in range(100):
            for bpushes in range(100):
                x = (ax * apushes) + (bx * bpushes)
                y = (ay * apushes) + (by * bpushes)

                if x == px and y == py:
                    cost = (apushes * 3) + (bpushes)

                    if first:
                        cheapest["A Pushes"] = apushes
                        cheapest["B Pushes"] = bpushes
                        cheapest["Cost"] = cost
                        first = False

                    elif cost < cheapest["Cost"]:
                        cheapest["A Pushes"] = apushes
                        cheapest["B Pushes"] = bpushes
                        cheapest["Cost"] = cost

        if cheapest["Cost"] != -1:
            allPrizes.append(cheapest)

    for p in allPrizes:
        total += p["Cost"]

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    grid = []
    length = 0
    height = 0

    with open(filename) as fh:
        for line in fh:
            row = list(line.strip())
            length = len(row)
            height += 1
            grid.append(row)

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input0.txt")