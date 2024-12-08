def partOne(filename):
    map = []
    height = 0
    length = 0

    with open(filename) as fh:
        for line in fh:
            tiles = list(line.strip())
            map.append(tiles)
            length = len(tiles)
            height += 1

    antennas = []
    for y in range(height):
        for x in range(length):
            tile = map[y][x]
            if tile != '.':
                antennas.append([tile,x,y])

    antinodes = []
    for a in antennas:
        antenna = a[0]
        ax = a[1]
        ay = a[2]
        for y in range(height):
            for x in range(length):
                tile = map[y][x]
                if tile == antenna:
                    dx = ax - x
                    dy = ay - y
                    for p in [[ax+dx,ay+dy],[x+dx,y+dy]]:
                        if p != [x,y] and p != [ax,ay] and (p[0] >= 0) and (p[0] < length) and (p[1] >= 0) and (p[1] < height) and (p not in antinodes):
                            antinodes.append(p)

    print(f"Part One: {len(antinodes)}")

def partTwo(filename):
    map = []
    height = 0
    length = 0

    with open(filename) as fh:
        for line in fh:
            tiles = list(line.strip())
            map.append(tiles)
            length = len(tiles)
            height += 1

    antennas = []
    for y in range(height):
        for x in range(length):
            tile = map[y][x]
            if tile != '.':
                antennas.append([tile,x,y])

    antinodes = []
    for a in antennas:
        antenna = a[0]
        ax = a[1]
        ay = a[2]
        for y in range(height):
            for x in range(length):
                if ax == x and ay == y:
                    continue

                tile = map[y][x]
                if tile == antenna:
                    dx = ax - x
                    dy = ay - y

                    possibleAntinodes = []
                    for b in [[ax,ay],[x,y]]:
                        p = [b[0], b[1]]
                        while True:
                            p = [p[0]+dx,p[1]+dy]
                            if (p in antinodes):
                                continue
                            if (p[0] < 0) or (p[0] >= length) or (p[1] < 0) or (p[1] >= height):
                                break
                            possibleAntinodes.append(p)

                    for p in possibleAntinodes:
                        if (p[0] >= 0) and (p[0] < length) and (p[1] >= 0) and (p[1] < height) and (p not in antinodes):
                            antinodes.append(p)

    print(f"Part Two: {len(antinodes)}")

# ---
partOne("input1.txt")
partTwo("input1.txt")