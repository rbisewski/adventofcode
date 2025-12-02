def partOne(filename):
    moves = []
    with open(filename) as fh:
        for line in fh:
            moves.append(line)

    total = 0
    count = 50
    for move in moves:
        direction = move[0]
        steps = int(move[1:])

        if direction == 'L':
            step = 1
        elif direction == 'R':
            step = -1

        count = (count + (step * steps)) % 100
        if count == 0:
            total += 1

    print(f"Part One: {total}")

def partTwo(filename):
    moves = []
    with open(filename) as fh:
        for line in fh:
            moves.append(line)

    total = 0
    count = 50
    for move in moves:
        direction = move[0]
        steps = int(move[1:])

        if direction == 'L':
            step = 1
        elif direction == 'R':
            step = -1

        for _ in range(steps):
            count = (count + step) % 100
            if count == 0:
                total += 1

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")