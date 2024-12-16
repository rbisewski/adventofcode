def partOne(filename):
    total = 0
    length = 0
    height = 0
    robot = (0,0)

    with open(filename) as fh:
        map, moves = fh.read().split("\n\n")

    moves = moves.replace("\n", "")

    grid = []
    for line in map.split("\n"):
        newLine = list(line)
        if "@" in newLine:
            robot = (newLine.index("@"), height)
        grid.append(newLine)
        height += 1

    length = len(grid[0])

    for m in moves:
        x0 = robot[0]
        y0 = robot[1]

        x1 = x0
        y1 = y0
        direction = ""

        match m:
            case '^':
                y1 -= 1
                direction = "N"
            case 'v':
                y1 += 1
                direction = "S"
            case '>':
                x1 += 1
                direction = "E"
            case '<':
                x1 -= 1
                direction = "W"

        tile = grid[y1][x1]
        match tile:
            case '.':
                grid[y1][x1] = '@'
                grid[y0][x0] = '.'
                robot = (x1, y1)
                continue
            case '#':
                continue
            case 'O':
                pass

        x = x1
        y = y1
        while tile == 'O':
            match direction:
                case "N":
                    y -= 1
                case "S":
                    y += 1
                case "E":
                    x += 1
                case "W":
                    x -= 1
            tile = grid[y][x]

        match tile:
            case '#':
                continue
            case '.':
                grid[y0][x0] = '.'
                grid[y1][x1] = '@'
                grid[y][x] = 'O'
                robot = (x1, y1)

    for y in range(height):
        for x in range(length):
            tile = grid[y][x]
            if tile == 'O':
                total += (100 * y) + x

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    length = 0
    height = 0
    robot = (0,0)

    with open(filename) as fh:
        map, moves = fh.read().split("\n\n")

    moves = moves.replace("\n", "")

    grid = []
    for line in map.split("\n"):
        enlargedMapLine = line.replace("#", "##").replace("O", "[]").replace(".", "..").replace("@", "@.")
        newLine = list(enlargedMapLine)
        if "@" in newLine:
            robot = (newLine.index("@"), height)
        grid.append(newLine)
        height += 1

    length = len(grid[0])

    x = robot[0]
    y = robot[1]

    directions = {
                       "^": (-1,  0),
        "<": ( 0, -1),                ">": ( 0,  1),
                       "v": ( 1,  0)
    }

    for m in moves:
        dy, dx = directions[m]
        pos = [(y, x)]

        for y1, x1 in pos:
            ny = y1 + dy
            nx = x1 + dx
            if (ny, nx) not in pos:
                match grid[ny][nx]:
                    case "#":
                        break
                    case "[":
                        pos.extend([
                            (ny, nx),
                            (ny, nx + 1)
                        ])
                    case "]":
                        pos.extend([
                            (ny, nx),
                            (ny, nx - 1)
                        ])

        if grid[ny][nx] == "#":
            continue

        for i, j in pos[::-1]:
            a = grid[i + dy][j + dx]
            b = grid[i][j]
            grid[i][j] = a
            grid[i + dy][j + dx] = b

        y = y + dy
        x = x + dx

    for y in range(height):
        for x in range(length):
            tile = grid[y][x]
            if tile == "[":
                total += (100 * y) + x

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")