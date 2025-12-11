from collections import deque

def min_xor_operations(target, numbers):
    nums_int = [int(x, 2) for x in numbers]
 
    queue = deque([(0, 0)])
    visited = {0}
    while queue:
        current, steps = queue.popleft()

        if current == int(target, 2):
            break

        for num in nums_int:
            next_state = current ^ num
            if next_state not in visited:
                visited.add(next_state)
                queue.append((next_state, steps + 1))

    return steps

def partOne(filename):
    least_pushes = {}

    i = 0
    with open(filename) as fh:
        for line in fh:
            clean_ln = line.strip()

            light = clean_ln.split("]")[0].strip('[').replace(".","0").replace("#","1")
            btns = list(filter(None,clean_ln.split("]")[1].split("{")[0].split(" ")))

            btns_as_binary = []
            for b in btns:
                init = ["0"] * len(light)
                ones = b.replace("(","").replace(")","").split(",")
                for o in ones:
                    num = int(o)
                    init[num] = '1'
                btns_as_binary.append("".join(init))

            least_pushes[i] = min_xor_operations(light, btns_as_binary)

            i+=1

    total = 0
    for key in least_pushes:
        total += least_pushes[key]

    print(f"Part One: {total}")

def partTwo(filename):
    total = 0
    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input0.txt")