def directionToPosition(d,i,j):
    match d:
        case 'NW':
            return [i-1,j-1]
        case 'N':
            return [i,  j-1]
        case 'NE':
            return [i+1,j-1]
        case 'W':
            return [i-1,j  ]
        case 'E':
            return [i+1,j  ]
        case 'SW':
            return [i-1,j+1]
        case 'S':
            return [i  ,j+1]
        case 'SE':
            return [i+1,j+1]
        case _:
            return []

def checkSurroundingLetters(letter,direction,i,j,array,length,height):
    nextLetter = ""
    match letter:
        case 'X':
            nextLetter = "M"
        case 'M':
            nextLetter = "A"
        case 'A':
            nextLetter = "S"
        case 'S':
            return 1

    pos = directionToPosition(direction,i,j)

    if pos == []:
        return 0

    x = pos[0]
    y = pos[1]

    if x < 0 or y < 0:
        return 0
    elif x >= length or y >= height:
        return 0

    if array[x][y] == nextLetter:
        return checkSurroundingLetters(nextLetter,direction,x,y,array,length,height)

    return 0

def checkForXmas(letter,i,j,array,length,height):
    a = 0
    valid = [False, False]

    for direction in [
        ["SE","NW"],
        ["NE","SW"],
    ]:
        d0 = direction[0]
        d1 = direction[1]

        pos0 = directionToPosition(d0,i,j)
        pos1 = directionToPosition(d1,i,j)

        if pos0 == [] or pos1 == []:
            return 0

        x0 = pos0[0]
        y0 = pos0[1]
        x1 = pos1[0]
        y1 = pos1[1]

        if x0 < 0 or y0 < 0 or x1 < 0 or y1 < 0:
            return 0
        elif x0 >= length or y0 >= height or x1 >= length or y1 >= height:
            return 0

        if array[x0][y0] == "M" and array[x1][y1] == "S":
            valid[a] = True
        elif array[x0][y0] == "S" and array[x1][y1] == "M":
            valid[a] = True

        a += 1

    if valid == [True, True]:
        return 1
    
    return 0

def partOne(filename):
    total = 0

    array = []
    length = 0
    height = 0
    with open(filename) as fh:
        for line in fh:
            ln = list(line.strip("\n"))
            array.append(ln)
            length = len(ln)
            height += 1

    for i in range(length):
        for j in range(height):
            letter = array[i][j]
            if letter == 'X':
                for d in [
                    "NW","N","NE",
                    "W",     "E",
                    "SW","S","SE"
                    ]:
                    total += checkSurroundingLetters(letter,d,i,j,array,length,height)

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0

    array = []
    length = 0
    height = 0
    with open(filename) as fh:
        for line in fh:
            ln = list(line.strip("\n"))
            array.append(ln)
            length = len(ln)
            height += 1

    for i in range(length):
        for j in range(height):
            letter = array[i][j]
            if letter == 'A':
                total += checkForXmas(letter,i,j,array,length,height)

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")