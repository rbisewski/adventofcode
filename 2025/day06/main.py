import re

def partOne(filename):
    lines = []

    with open(filename) as fh:
        for line in fh:
            ln = re.sub(r'\s+',' ',line.strip()).split(" ")
            lines.append(ln)
    
    total = 0
    for i in range(0, len(lines[0])):
        first = int(lines[0][i])
        second = int(lines[1][i])
        third = int(lines[2][i])
        fourth = int(lines[3][i])
        operator = lines[4][i]

        if operator == "*":
            total += (first * second * third * fourth)
        elif operator == "+":
            total += (first + second + third + fourth)
        
    print(f"Part One: {total}")

def partTwo(filename):
    lines = []
    operators = []

    with open(filename) as fh:
        for line in fh:
            if "*" in line:
                operators = re.sub(r'\s+',' ',line.strip()).split(" ")
            else:
                lines.append(list(line.rstrip('\n')))

    numbers = {}
    i = 0
    for ln in lines:
        for i in range(0,len(ln)):
            if not numbers.get(i):
                numbers[i] = ln[i]
            else:
                numbers[i] = f"{numbers[i]}{ln[i]}".strip()

    i += 1
    numbers[i] = ' '

    current_operator = 0
    temp_array = []
    total = 0
    for n in numbers:
        num = numbers[n]
        if num == ' ' or num == '':
            op = operators[current_operator]
            running_tally = 0

            for e in temp_array:
                if running_tally == 0:
                    running_tally = int(e)
                elif op == "+":
                    running_tally += int(e)
                elif op == "*":
                    running_tally *= int(e)

            temp_array = []
            current_operator += 1

            total += running_tally
        else:
            temp_array.append(num)

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")