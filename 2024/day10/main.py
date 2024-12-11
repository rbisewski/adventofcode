import copy

successfulHikes = []

def hike(grid, path, pos):
    if len(path["points"]) == 10:
        successfulHikes.append(path)
        return

    x0 = pos[0]
    y0 = pos[1]

    neighbours = [
                  [x0,y0-1],
        [x0-1,y0],          [x0+1,y0],
                  [x0,y0+1]
    ]

    for n in neighbours:
        x1 = n[0]
        y1 = n[1]
        if x1 < 0 or y1 < 0 or len(grid) <= y1 or len(grid[y1]) <= x1:
            continue

        t = grid[y1][x1]

        if t == len(path["points"]):
            newPath = copy.deepcopy(path)
            newPath["points"].append({"topology": t, "coord": [x1,y1]})
            hike(grid, newPath, [x1,y1])

def partOne(filename):
    global successfulHikes
    total = 0

    grid = []
    length = 0
    height = 0
    with open(filename) as fh:
        for line in fh:
            row = [int(x) for x in list(line.strip())]
            grid.append(row)
            length = len(row)
            height += 1

    trailHeadStart = []
    for y in range(height):
        for x in range(length):
            if grid[y][x] == 0:
                trailHeadStart.append([x,y])

    for pos in trailHeadStart:
        successfulHikes = []
        path = {"start": pos, "points": [{"topology": 0, "coord": pos}]}
        hike(grid, path, pos)

        listOfNines = []
        for s in successfulHikes:
            for p in s["points"]:
                if p["topology"] == 9:
                    if p["coord"] not in listOfNines:
                        listOfNines.append(p["coord"])

        total += len(listOfNines)

    print(f"Part One: {total}")

def partTwo(filename):
    global successfulHikes
    total = 0

    grid = []
    length = 0
    height = 0
    with open(filename) as fh:
        for line in fh:
            row = [int(x) for x in list(line.strip())]
            grid.append(row)
            length = len(row)
            height += 1

    trailHeadStart = []
    for y in range(height):
        for x in range(length):
            if grid[y][x] == 0:
                trailHeadStart.append([x,y])

    for pos in trailHeadStart:
        successfulHikes = []
        path = {"start": pos, "points": [{"topology": 0, "coord": pos}]}
        hike(grid, path, pos)
        total += len(successfulHikes)

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")