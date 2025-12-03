def partOne(filename):
    lines = []
    with open(filename) as fh:
        for line in fh:
            lines.append(line.strip())

    largest_joltage_per_line = []
    for ln in lines:
        largest_joltage = 0
        array = list(ln)

        while len(array) > 1:
            first = array.pop(0)
            for a in array:
                current_joltage = int(f"{first}{a}")
                if current_joltage > largest_joltage:
                    largest_joltage = current_joltage
        
        largest_joltage_per_line.append(largest_joltage)

    total = 0
    for j in largest_joltage_per_line:
        total += j

    print(f"Part One: {total}")

def partTwo(filename):
    lines = []
    with open(filename) as fh:
        for line in fh:
            lines.append(line.strip())

    def largest_joltage(line):
        k = 12
        stack = []
        to_remove = len(line) - k  # get the number of digits to remove

        for d in line:
            while to_remove > 0 and stack and stack[-1] < d:
                stack.pop()
                to_remove -= 1
            stack.append(d)

        if to_remove > 0:
            stack = stack[:-to_remove]

        return int(''.join(stack[:k]))

    largest_joltage_per_line = []
    for line in lines:
        largest_joltage_per_line.append(largest_joltage(line))

    total = 0
    for j in largest_joltage_per_line:
        total += j

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")