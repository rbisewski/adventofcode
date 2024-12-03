import copy

def isSafe(entries):
    array = copy.deepcopy(entries)
    array.sort()

    if entries != array and entries != array[::-1]:
        return False

    for i in range(len(entries)-1):
        a = int(entries[i])
        b = int(entries[i+1])

        if a == b:
            return False
        elif abs(a-b) > 3:
            return False

    return True

def partOne(filename):
    count = 0
    with open(filename) as fh:
        for line in fh:
            entries = [int(x) for x in line.strip().split(" ")]

            if isSafe(entries):
                count += 1

    print(f"Part One: {count} safe reports")

def partTwo(filename):
    count = 0
    with open(filename) as fh:
        for line in fh:
            entries = [int(x) for x in line.strip().split(" ")]

            if isSafe(entries):
                count += 1
            else:
                for i in range(len(entries)):
                    array = copy.deepcopy(entries)
                    del array[i]
                    if isSafe(array):
                        count += 1
                        break

    print(f"Part Two: {count} safe reports")

# ----
partOne("input1.txt")
partTwo("input1.txt")