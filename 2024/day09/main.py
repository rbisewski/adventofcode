def partOne(filename):
    checksum = 0

    diskmap = ""
    with open(filename) as fh:
        for line in fh:
            diskmap = [int(x) for x in list(line.strip())]

    id = 0
    isID = True
    blockState = []
    for d in diskmap:
        if isID:
            for i in range(d):
                blockState.append(id)
            id += 1
            isID = False
        elif not isID:
            if d > 0:
                for i in range(d):
                    blockState.append(".")
            isID = True

    fs = blockState
    for i in range(len(fs)):
        if fs[i] == '.':
            a = len(fs)-1
            while a > i:
                if fs[a] != '.':
                    first = fs[i]
                    second = fs[a]
                    fs[i] = second
                    fs[a] = first
                    break
                else:
                    a -= 1

    pos = 0
    for block in fs:
        if block != '.':
            checksum += (pos * int(block))
        pos += 1

    print(f"Part One: {checksum}")

def partTwo(filename):
    checksum = 0

    diskmap = ""
    with open(filename) as fh:
        for line in fh:
            diskmap = [int(x) for x in list(line.strip())]

    id = 0
    isID = True
    fs = []
    for d in diskmap:
        if isID:
            for i in range(d):
                fs.append(id)
            id += 1
            isID = False
        elif not isID:
            if d > 0:
                for i in range(d):
                    fs.append(".")
            isID = True

    blockCount = {}
    for f in fs:
        if blockCount.get(f):
            blockCount[f] += 1
        else:
            blockCount[f] = 1

    last = len(fs)-1
    cursor = fs[last]
    while cursor != 0:
        size = blockCount[cursor]
        previous = 0
        fileMoved = False
        while True:
            if fileMoved:
                break
            a = fs.index('.', previous + 1)
            if a > fs.index(cursor):
                break
            for i in range(size):
                if fs[a + i] != '.':
                    previous = a
                    break
                elif i == size-1:
                    for e in range(size):
                        fs[fs.index(cursor)] = '.'
                    for j in range(size):
                        fs[a + j] = cursor
                    fileMoved = True
        cursor -= 1

    pos = 0
    for block in fs:
        if block != '.':
            checksum += (pos * int(block))
        pos += 1

    print(f"Part Two: {checksum}")

# ---
partOne("input0.txt")
partTwo("input0.txt")