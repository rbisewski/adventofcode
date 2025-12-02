def partOne(filename):
    ranges = []
    with open(filename) as fh:
        for line in fh:
            ranges = line.split(",")
    
    invalid_ids = []
    for r in ranges:
        pieces = r.split("-")
        start = int(pieces[0])
        end = int(pieces[1])
        for i in range(start,end+1):
            id = f"{i}"
            middle = len(id) // 2 + len(id) % 2
            first_half = id[:middle]
            second_half = id[middle:]
            if first_half == second_half:
                invalid_ids.append(i)

    total = 0
    for i in invalid_ids:
        total += i

    print(f"Part One: {total}")

def partTwo(filename):
    ranges = []
    with open(filename) as fh:
        for line in fh:
            ranges = line.split(",")

    def split_string_n_ways(s, n):
        length = len(s)
        avg_len = length // n
        remainder = length % n
        parts = []
        start = 0

        for i in range(n):
            part_len = avg_len + (1 if i < remainder else 0)
            parts.append(s[start:start+part_len])
            start += part_len

        return parts
    
    invalid_ids = []
    for r in ranges:
        pieces = r.split("-")
        start = int(pieces[0])
        end = int(pieces[1])

        # for every individual range
        for i in range(start,end+1):
            id = f"{i}"
            length = len(id)

            # for every combo of a given ID
            for ln in range(2,length+1):
                parts = split_string_n_ways(id, ln)
                all_equal = all(x == parts[0] for x in parts)
                if all_equal:
                    invalid_ids.append(i)
                    break

    total = 0
    for i in invalid_ids:
        total += i

    print(f"Part Two: {total}")

# ----
partOne("input1.txt")
partTwo("input1.txt")