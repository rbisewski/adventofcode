def evaluate(filename,operatorTypes):
    total = 0
    with open(filename) as fh:
        for line in fh:
            pieces = line.strip().split(": ")
            answer = int(pieces[0])
            numbers = [int(x) for x in pieces[1].split(" ")]

            count = len(numbers)-1
            finalList = [[]]
            groups = [list(operatorTypes)] * count
            for i in groups:
                finalList = [x+[y] for x in finalList for y in i]

            permutations = [''.join(item) for item in finalList]

            for p in permutations:
                ops = list(p)
                i = 0
                runningTotal = numbers[i]
                for o in ops:
                    n1 = numbers[i+1]
                    match o:
                        case "*":
                            runningTotal = (runningTotal * n1)
                        case "+":
                            runningTotal = (runningTotal + n1)
                        case "c":
                            runningTotal = int(f"{runningTotal}{n1}")
                    i += 1
                
                if runningTotal == answer:
                    total += answer
                    break
    return total

def partOne(filename):
    total = evaluate(filename,"*+")
    print(f"Part One: {total}")

def partTwo(filename):
    total = evaluate(filename,"*+c")
    print(f"Part Two: {total}")

# ---
partOne("input1.txt")
partTwo("input1.txt")