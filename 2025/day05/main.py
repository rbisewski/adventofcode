def merge_ranges(ranges):
    ranges.sort(key=lambda x: x[0])

    merged = [ranges[0]]
    for current_start, current_end in ranges[1:]:
        last_start, last_end = merged[-1]

        # check if these overlap, and if so, merge them by taking the larger end-value
        if current_start <= last_end + 1:
            merged[-1] = (last_start, max(last_end, current_end))
            continue

        merged.append((current_start, current_end))
 
    return [[start,end] for start, end in merged]

def partOne(filename):
    fresh = []
    available = []

    with open(filename) as fh:
        for line in fh:
            ln = line.strip()
            if ln == '':
                continue

            if "-" in ln:
                fresh.append(ln)
                continue
            
            available.append(ln)

    total = 0
    for a in available:
        ingredient = int(a)
        for f in fresh:
            pieces = f.split("-")
            start = int(pieces[0])
            end = int(pieces[1])
            if ingredient >= start and ingredient <= end:
                total += 1
                break

    print(f"Part One: {total}")

def partTwo(filename):
    fresh = []

    with open(filename) as fh:
        for line in fh:
            ln = line.strip()
            if ln == '':
                continue

            if "-" in ln:
                fresh.append(ln)
                continue

    sets = []
    for f in fresh:
        pieces = f.split("-")
        start = int(pieces[0])
        end = int(pieces[1])
        sets.append((start,end))

    new_ranges = merge_ranges(sets)

    total = 0
    for r in new_ranges:
        start = r[0]
        end = r[1]
        total += (end - start + 1)

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")