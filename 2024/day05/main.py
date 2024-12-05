import copy

def valid(array,rules):
    for r in rules:
        first = r[0]
        second = r[1]

        if first not in array or second not in array:
            continue

        if array.index(first) > array.index(second):
            return False

    return True

def partOne(filename):
    total = 0
    rules = []
    updates = []

    with open(filename) as fh:
        for line in fh:
            if "|" in line:
                row = [int(x) for x in line.strip("\n").split("|")]
                rules.append(row)
            elif "," in line:
                row = [int(x) for x in line.strip("\n").split(",")]
                updates.append(row)

    validUpdates = []
    for u in updates:
        if valid(u,rules):
            validUpdates.append(u)

    for v in validUpdates:
        middlePosition = int((len(v)-1)/2)
        total += v[middlePosition]

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    rules = []
    updates = []

    with open(filename) as fh:
        for line in fh:
            if "|" in line:
                row = [int(x) for x in line.strip("\n").split("|")]
                rules.append(row)
            elif "," in line:
                row = [int(x) for x in line.strip("\n").split(",")]
                updates.append(row)

    invalidUpdates = []
    for u in updates:
        for r in rules:
            first = r[0]
            second = r[1]

            if first not in u or second not in u:
                continue

            if u.index(first) > u.index(second):
                invalidUpdates.append(u)
                break
 
    fixedUpdates = []
    for i in invalidUpdates:
        iu = copy.deepcopy(i)

        while not valid(iu,rules):
            for r in rules:
                first = r[0]
                second = r[1]

                if first not in iu or second not in iu:
                    continue

                if iu.index(first) > iu.index(second):
                    iu[iu.index(first)] = second
                    iu[iu.index(second)] = first

        fixedUpdates.append(iu)

    for v in fixedUpdates:
        middlePosition = int((len(v)-1)/2)
        total += v[middlePosition]

    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")