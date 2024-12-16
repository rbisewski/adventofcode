from heapq import heappop, heappush

compass = {
                   "n": (-1, 0),
    "w": (0, -1),                 "e": (0, 1),
                   "s": (1, 0)
}

ninetyDegrees = {
    "e": ["n", "s"],
    "s": ["e", "w"],
    "w": ["s", "n"],
    "n": ["w", "e"],
}

def getLowestScore(start, end, grid):
    sx, sy = start
    pq = [(0, sx, sy, 'e')]
    visited = set()

    while pq:
        cost, x, y, facing = heappop(pq)

        if (x, y, facing) in visited:
            continue

        visited.add((x, y, facing))

        if (x,y) == (end[0], end[1]):
            return cost - 1000

        dx, dy = compass[facing]
        nx = x + dx
        ny = y + dy

        if grid[ny][nx] != "#":
            heappush(pq, (cost + 1, nx, ny, facing))

        for turn in ninetyDegrees[facing]:
            heappush(pq, (cost + 1000, x, y, turn))

    return 0

def getBestTiles(start, end, grid):
    sx, sy = start
    ex, ey = end
    scores = {}
    tiles = set()

    # set this to some very high number
    best = 100000000000000000000000000

    pq = [(0, (sx, sy), 'e', [(sx, sy)])]

    while pq:
        cost, (x, y), facing, path = heappop(pq)

        if ((x, y), facing) in scores and scores[((x, y), facing)] < cost:
            continue

        if (x, y) == (ex, ey) and cost <= best:
            best = cost
            tiles |= set(path)

        scores[((x, y), facing)] = cost

        dx, dy = compass[facing]
        nx = x + dx
        ny = y + dy

        if grid[ny][nx] != "#":
            heappush(pq, (cost + 1, (nx, ny), facing, path[:] + [(nx, ny)]))

        for turn in ninetyDegrees[facing]:
            heappush(pq, (cost + 1000, (x, y), turn, path))

    return len(tiles)

def partOne(filename):
    total = 0
    grid = []
    height = 0
    sx, sy = 0, 0
    ex, ey = 0, 0

    with open(filename) as fh:
        for line in fh:
            row = list(line.strip())
            if 'S' in row:
                sx = row.index("S")
                sy = height
            elif 'E' in row:
                ex = row.index("E")
                ey = height
            grid.append(row)
            height += 1

    total = getLowestScore((sx, sy), (ex, ey), grid)

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    grid = []
    height = 0
    sx, sy = 0, 0
    ex, ey = 0, 0

    with open(filename) as fh:
        for line in fh:
            row = list(line.strip())
            if 'S' in row:
                sx = row.index("S")
                sy = height
            elif 'E' in row:
                ex = row.index("E")
                ey = height
            grid.append(row)
            height += 1

    total = getBestTiles((sx, sy), (ex, ey), grid)

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")