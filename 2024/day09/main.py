import copy

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
    files = {}
    for i in range(len(fs)):
        id = fs[i]
        if files.get(id):
            files[id].append(i)
        else:
            files[id] = [i]

    cursor = len(fs)-1
    numbers = []
    while cursor >= 0:
        for f in files:
            if f == '.':
                continue

            elif cursor in files[f]:
                filesize = len(files[f])
                for i in files['.']:
                    if numbers == []:
                        numbers.append(i)

                    elif filesize == len(numbers):
                        oldFileLocation = copy.deepcopy(files[f])
                        files[f] = numbers
                        files[f].sort()
                        files['.'] = files['.'] + oldFileLocation
                        files['.'].sort()
                        for n in numbers:
                            files['.'].remove(n)
                        break

                    elif i == numbers[len(numbers)-1] + 1:
                        numbers.append(i)

                numbers = []
                filesize = 0

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