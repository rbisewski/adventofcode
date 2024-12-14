import copy
import re

def partOne(filename):
    robots = []

    with open(filename) as fh:
        for line in fh:
            matches = re.findall(r"p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)",line.strip())
            px = int(matches[0][0])
            py = int(matches[0][1])
            vx = int(matches[0][2])
            vy = int(matches[0][3])
            entry = {"position":(px,py), "velocity": (vx,vy)}
            robots.append(entry)

    wide = 0
    tall = 0
    if filename == "input0.txt":
        wide = 11
        tall = 7

    elif filename == "input1.txt":
        wide = 101
        tall = 103

    middleWideRow = (wide-1) // 2
    middleTallRow = (tall-1) // 2

    for s in range(100):
        for r in robots:
            vx = r["velocity"][0]
            vy = r["velocity"][1]

            px = r["position"][0]
            py = r["position"][1]

            r["position"] = ((px + vx) % wide, (py + vy) % tall)

    q = {
        "I":0,
        "II":0,
        "III":0,
        "IV":0
    }
    for r in robots:
        px = r["position"][0]
        py = r["position"][1]

        if px < middleWideRow and py < middleTallRow:
            q["II"] += 1
        elif px < middleWideRow and py > middleTallRow:
            q["I"] += 1
        elif px > middleWideRow and py < middleTallRow:
            q["III"] += 1
        elif px > middleWideRow and py > middleTallRow:
            q["IV"] += 1

    safetyFactor = q["I"] * q["II"] * q["III"] * q["IV"]

    print(f"Part One: {safetyFactor}")

def partTwo(filename):
    robots = []

    with open(filename) as fh:
        for line in fh:
            matches = re.findall(r"p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)",line.strip())
            px = int(matches[0][0])
            py = int(matches[0][1])
            vx = int(matches[0][2])
            vy = int(matches[0][3])
            entry = {"position":(px,py), "velocity": (vx,vy)}
            robots.append(entry)

    wide = 0
    tall = 0
    if filename == "input0.txt":
        wide = 11
        tall = 7

    elif filename == "input1.txt":
        wide = 101
        tall = 103

    middleWideRow = (wide-1) // 2
    middleTallRow = (tall-1) // 2

    grid = []
    for w in range(wide):
        row = []
        for t in range(tall):
            row.append(".")
        grid.append(row)

    #  To determine the exact number of seconds
    #
    #  0. initially set this to something large, like 10,000 or 50,000
    #
    #  1. python main.py > /tmp/frames
    #
    #  2. cat /tmp/frame | grep "'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O'"
    #
    #  3. open the file with a text editor to confirm this is the Christmas tree
    #
    seconds = 7916

    for s in range(seconds):
        newGrid = copy.deepcopy(grid)

        for r in robots:
            vx = r["velocity"][0]
            vy = r["velocity"][1]

            px = r["position"][0]
            py = r["position"][1]

            r["position"] = ((px + vx) % wide, (py + vy) % tall)

            newGrid[r["position"][0]][r["position"][1]] = "O"

        print(f"At second: {s}")
        for g in newGrid:
            print(g)
        print("---")

    # print out the frames
    print(f"Part Two: {seconds}")

# ---
partOne("input1.txt")
partTwo("input1.txt")