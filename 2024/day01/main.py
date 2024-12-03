def partOne(filename):
    array0 = []
    array1 = []

    length = 0
    with open(filename) as fh:
        for line in fh:
            entries = line.strip().split("   ")
            array0.append(int(entries[0]))
            array1.append(int(entries[1]))
            length += 1

    array0.sort()
    array1.sort()

    total = 0
    for i in range(length):
        total += abs(array0[i] - array1[i])

    print(f"Part One: {total}")

def partTwo(filename):
    array0 = []
    array1 = []

    with open(filename) as fh:
        for line in fh:
            entries = line.strip().split("   ")
            array0.append(int(entries[0]))
            array1.append(int(entries[1]))

    total = 0
    for a in array0:
        similarity = 0
        for b in array1:
            if a == b:
                similarity += 1
        total += (a * similarity)

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")