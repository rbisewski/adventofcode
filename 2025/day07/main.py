worlds = {}
def evaluate_world_line(lines, row, col):
    height = len(lines)
    width = len(lines[0])

    if row < 0 or row >= height:
        return 1

    elif col < 0 or col >= width:
        return 1

    elif (row, col) in worlds:
        pass

    elif lines[row][col] == "^":
        worlds[(row, col)] = evaluate_world_line(lines, row + 1, col - 1) + evaluate_world_line(lines, row + 1, col + 1)

    elif lines[row][col] == ".":
        worlds[(row, col)] = evaluate_world_line(lines, row + 1, col)
    
    return worlds[(row, col)]

def partOne(filename):
    lines = []

    with open(filename) as fh:
        for line in fh:
            lines.append(list(line.strip()))

    i = 0
    for c in lines[0]:
        if c == "S":
            break
        i += 1

    total = 0
    first_split = False 
    previous_line = []
    for ln in lines:
        if "S" in ln:
            first_split = True
            continue
        elif first_split:
            ln[i] = "|"
            first_split = False
        elif "^" in ln:
            beam_positions = [i for i, ch in enumerate(previous_line) if ch == '|']
            carot_positions = [i for i, ch in enumerate(ln) if ch == '^']

            characters_replaced = []
            for c in carot_positions:
                if c not in characters_replaced:
                    characters_replaced.append(c)

                if c in beam_positions:
                    total += 1
                    ln[c-1] = '|'
                    ln[c+1] = '|'

            beam_positions = [i for i, ch in enumerate(previous_line) if ch == '|']
            for b in beam_positions:
                if b in characters_replaced:
                    continue
                ln[b] = '|'
        else:
            beam_positions = [i for i, ch in enumerate(previous_line) if ch == '|']
            for b in beam_positions:
                ln[b] = '|'

        previous_line = ln
    
    print(f"Part One: {total}")

def partTwo(filename):
    lines = []

    with open(filename) as fh:
        for line in fh:
            lines.append(list(line.strip()))

    s_location = lines[0].index('S')
    total = evaluate_world_line(lines, 1, s_location)

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")