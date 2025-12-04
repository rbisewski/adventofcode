def check_surrounds(x,y,str,lines,height,width):
    if str == ".":
        return 0

    directions = [
        [x-1,y-1],[x-1,y ],[x-1,y+1],
        [x,  y-1],[      ],[x,  y+1],
        [x+1,y-1],[x+1,y ],[x+1,y+1]
    ]

    count = 0
    for d in directions:
        if d == []:
            continue

        x1 = d[0]
        y1 = d[1]

        if x1 < 0 or y1 < 0 or y1 == height or x1 == width:
            continue

        char = lines[x1][y1]
        if char == "@":
            count += 1

    # check if it can be accessed by the forklift
    if count < 4:
        return 1

    return 0

def partOne(filename):
    lines = []
    height = 0
    width = 0

    with open(filename) as fh:
        for line in fh:
            elements = list(line.strip())
            width = len(elements)
            lines.append(elements)

    height = len(lines)

    total = 0
    for x in range(0, len(lines)):
        for y in range(0, len(lines[x])):
            total += check_surrounds(x,y,lines[x][y], lines, height, width)

    print(f"Part One: {total}")

def partTwo(filename):
    lines = []
    height = 0
    width = 0

    with open(filename) as fh:
        for line in fh:
            elements = list(line.strip())
            width = len(elements)
            lines.append(elements)

    height = len(lines)

    total = 0
    previous_total = -1
    while (total != previous_total):
        previous_total = total
        for x in range(0, len(lines)):
            for y in range(0, len(lines[x])):
                outcome = check_surrounds(x,y,lines[x][y],lines,height,width)
                if outcome == 1:
                    lines[x][y] = '.'
                total += outcome

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")