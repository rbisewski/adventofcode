positionsRecorded = set()
regions = {}

def getNeighbours(pos):
    x0 = pos[0]
    y0 = pos[1]
    return [
                  [x0,y0-1],
        [x0-1,y0],          [x0+1,y0],
                  [x0,y0+1]
    ]

def evaluateNeighbours(x1, y1, letter, grid, length, height):
    valid = []

    for nx, ny in getNeighbours([x1, y1]):
        if 0 > ny or ny >= height or 0 > nx or nx >= length:
            continue

        if grid[ny][nx] != letter:
            continue

        if (nx, ny) not in positionsRecorded:
            valid.append((nx, ny))

    return valid

def mapOutRegion(x, y, letter, grid, length, height):
    pointsToEvaluate = [(x, y)]

    tiles = []
    while pointsToEvaluate:
        x1, y1 = pointsToEvaluate.pop()
        if (x1, y1) in positionsRecorded:
            continue

        positionsRecorded.add((x1, y1))
        tiles.append((y1, x1))

        neighbours = evaluateNeighbours(x1, y1, letter, grid, length, height)
        pointsToEvaluate += neighbours

    return tiles

def resetGlobals():
    global positionsRecorded, regions
    positionsRecorded = set()
    regions = {}

def partOne(filename):
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

    incrementer = 0
    for y in range(height):
        for x in range(length):
            if (x, y) not in positionsRecorded:
                letter = grid[y][x]
                regions[f"{letter}-{incrementer}"] = mapOutRegion(x, y, letter, grid, length, height)
                incrementer += 1

    for r in regions:
        area = len(regions[r])
        perimeter = 0

        for x, y in regions[r]:
            for n in getNeighbours([x,y]):
                nx = n[0]
                ny = n[1]
                if (nx, ny) not in regions[r]:
                    perimeter += 1

        total += area * perimeter

    print(f"Part One: {total}")

def partTwo(filename):
    resetGlobals()
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

    incrementer = 0
    for y in range(height):
        for x in range(length):
            if (x, y) not in positionsRecorded:
                letter = grid[y][x]
                regions[f"{letter}-{incrementer}"] = mapOutRegion(x, y, letter, grid, length, height)
                incrementer += 1

    for r in regions:
        area = len(regions[r])

        perimeter = set()
        for x, y in regions[r]:
            for n in getNeighbours([x,y]):
                nx = n[0]
                ny = n[1]
                if (nx, ny) not in regions[r]:
                    perimeter.add(((x,y),(nx,ny)))

        sides = set()
        for pos1, pos2 in perimeter:
            x1 = pos1[0]
            y1 = pos1[1]
            x2 = pos2[0]
            y2 = pos2[1]
            if ((x1+1, y1),(x2+1, y2)) in perimeter:
                continue
            elif ((x1, y1+1),(x2, y2+1)) in perimeter:
                continue
            else:
                sides.add((pos1, pos2))

        total += area * len(sides)

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")