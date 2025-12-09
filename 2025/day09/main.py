import itertools

def calculate_line(p, q):
    ln = set()
    a1, a2 = min(p[0], q[0]), max(p[0], q[0])
    b1, b2 = min(p[1], q[1]), max(p[1], q[1])
    for a in range(a1, a2 + 1):
        for b in range(b1, b2 + 1):
            ln.add((a, b))
    return ln

def partOne(filename):
    lines = []

    with open(filename) as fh:
        for line in fh:
            ln = line.strip().split(",")
            x = int(ln[0])
            y = int(ln[1])
            lines.append([x,y])

    largest_area = 0
    i = 0
    while i < len(lines):
        ln = lines[i]
        x1 = ln[0]
        y1 = ln[1]

        for line in lines:
            if ln == line:
                continue
            x2 = line[0]
            y2 = line[1]

            area = (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)
            if area > largest_area:
                largest_area = area

        i += 1

    print(f"Part One: {largest_area}")

def partTwo(filename):
    tiles = []
    with open(filename) as fh:
        for ln in fh:
            tiles.append(tuple([int(x) for x in ln.split(',')]))

    areas = []
    for pair in itertools.combinations(tiles, 2):
        x1 = pair[0][0]
        y1 = pair[0][1]
        x2 = pair[1][0]
        y2 = pair[1][1]
        area = (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)
        areas.append((pair, area))

    # calculate the set of points that form the perimeter of the polygon
    perimeter = set()
    for i in range(1, len(tiles)):
        ln = calculate_line(tiles[i - 1], tiles[i])
        perimeter |= ln
    perimeter |= calculate_line(tiles[-1], tiles[0])

    # check if the sides of the rectangle are bounded by the perimeter
    largest_area = 0
    areas.sort(key=lambda x: x[1], reverse=True)
    for rectangle in areas:
        pair, area = rectangle
        x1 = pair[0][0]
        y1 = pair[0][1]
        x2 = pair[1][0]
        y2 = pair[1][1]

        contains_perimeter = True
        for s in perimeter:
            if min(x1, x2) < s[0] and s[0] < max(x1, x2):
                if min(y1, y2) < s[1] and s[1] < max(y1, y2):
                    contains_perimeter = False
                    break

        if contains_perimeter:
            largest_area = area
            break

    print(f"Part Two: {largest_area}")

# ----
partOne("input1.txt")
partTwo("input1.txt")