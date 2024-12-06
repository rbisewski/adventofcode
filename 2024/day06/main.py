import copy

def loadMap(filename):
    map = []
    guardPos = []

    length = 0
    height = 0

    with open(filename) as fh:
        for line in fh:
            row = list(line.strip())
            map.append(row)
            length = len(row)
            height += 1

    for x in range(length):
        for y in range(height):
            if map[y][x] == "^":
                guardPos = [x,y]
                map[y][x] = '.'
                break
        if guardPos != []:
            break

    return map, guardPos, length, height

def runMap(map, guardPos, length, height):
    guardPosX = guardPos[0]
    guardPosY = guardPos[1]
    y = guardPosY
    x = guardPosX

    loopLimit = 0
    positions = [[x,y]]
    direction = "N"
    while True:
        match direction:
            case "N":
                y -= 1
            case "E":
                x += 1
            case "S":
                y += 1
            case "W":
                x -= 1

        if x < 0 or y < 0 or x >= length or y >= height:
            break

        nextBlock = map[y][x]
        match nextBlock:
            case ".":
                guardPosY = y
                guardPosX = x
                if [x,y] not in positions:
                    positions.append([x,y])

                # incredibly hacky low-code way via brute force
                elif [x,y] in positions:
                    loopLimit += 1
                    if loopLimit > 1000:
                        return []
            case "#":
                match direction:
                    case "N":
                        y += 1
                        direction = "E"
                    case "E":
                        x -= 1
                        direction = "S"
                    case "S":
                        y -= 1
                        direction = "W"
                    case "W":
                        x += 1
                        direction = "N"

    return positions

def partOne(filename):
    map, guardPos, length, height = loadMap(filename)
    positions = runMap(map, guardPos, length, height)
    print(f"Part One: {len(positions)}")

def partTwo(filename):
    map, guardPos, length, height = loadMap(filename)
    positions = runMap(map, guardPos, length, height)

    count = 0
    for p in positions:
        x = p[0]
        y = p[1]
        if map[y][x] == ".":
            tempMap = copy.deepcopy(map)
            tempMap[y][x] = "#"
            res = runMap(tempMap, guardPos, length, height)
            if res == []:
                count += 1

    print(f"Part Two: {count}")

# ---
partOne("input1.txt")
partTwo("input1.txt")