def get_distance(junctions):
    num_junctions = len(junctions)
    distances = []
    
    for index_a in range(num_junctions):
        ax, ay, az = junctions[index_a]
        for index_b in range(index_a + 1, num_junctions):
            bx, by, bz = junctions[index_b]

            distance_sq = (ax - bx)**2 + (ay - by)**2 + (az - bz)**2
            distances.append((distance_sq, index_a, index_b))

    distances.sort()
    return distances

def build_circuit(junction_to_circuit_id, circuit_sets, index_a, index_b, current_circuit_id):
    circuit_a = junction_to_circuit_id[index_a]
    circuit_b = junction_to_circuit_id[index_b]
 
    relevant_circuit_id = None

    # neither are connected, so make a new circuit
    if circuit_a is None and circuit_b is None:
        circuit_sets[current_circuit_id] = {index_a, index_b}
        junction_to_circuit_id[index_a] = current_circuit_id
        junction_to_circuit_id[index_b] = current_circuit_id
        relevant_circuit_id = current_circuit_id
        current_circuit_id += 1

    # in the same circuit
    elif circuit_a == circuit_b and circuit_a is not None:
        relevant_circuit_id = circuit_a

    # B is in a circuit, but not A
    elif circuit_a is None:
        circuit_sets[circuit_b].add(index_a)
        junction_to_circuit_id[index_a] = circuit_b
        relevant_circuit_id = circuit_b

    # A is in a circuit, but not B
    elif circuit_b is None:
        circuit_sets[circuit_a].add(index_b)
        junction_to_circuit_id[index_b] = circuit_a
        relevant_circuit_id = circuit_a

    # merge A and B together
    else:
        union_set = circuit_sets[circuit_a] | circuit_sets[circuit_b]
        circuit_sets[circuit_a] = union_set
        del circuit_sets[circuit_b]
        
        # update each junction's circuit ID
        for junction_index in union_set:
            junction_to_circuit_id[junction_index] = circuit_a

        relevant_circuit_id = circuit_a
 
    return current_circuit_id, relevant_circuit_id

def partOne(filename):
    junctions = []
    with open(filename, "r") as data:
        for line in data:
            x, y, z = map(int, line.strip().split(","))
            junctions.append((x, y, z))

    sorted_list_of_distances = get_distance(junctions)

    # initially the junctions are not connected
    junction_to_circuit_id = {}
    for i in range(len(junctions)):
        junction_to_circuit_id[i] = None

    circuit_sets = {}
    current_circuit_id = 1

    threshold = sorted_list_of_distances[:999 + 1]

    for t in threshold:
        index_a = t[1]
        index_b = t[2]

        current_circuit_id, _ = build_circuit(
            junction_to_circuit_id, 
            circuit_sets, 
            index_a, 
            index_b, 
            current_circuit_id
        )
 
    circuit_sizes = []
    for s in circuit_sets.values():
        circuit_sizes.append(len(s))

    circuit_sizes.sort(reverse=True)

    result = circuit_sizes[0] * circuit_sizes[1] * circuit_sizes[2]

    print(f"Part One: {result}")

def partTwo(filename):
    junctions = []
    with open(filename, "r") as data:
        for line in data:
            x, y, z = map(int, line.strip().split(","))
            junctions.append((x, y, z))

    sorted_list_of_distances = get_distance(junctions)

    # initially the junctions are not connected
    junction_to_circuit_id = {}
    for i in range(len(junctions)):
        junction_to_circuit_id[i] = None

    circuit_sets = {}
    current_circuit_id = 1
 
    result = None
    for s in sorted_list_of_distances:
        index_a = s[1]
        index_b = s[2]

        current_circuit_id, relevant_circuit_id = build_circuit(
            junction_to_circuit_id, 
            circuit_sets, 
            index_a, 
            index_b, 
            current_circuit_id
        )
 
        if relevant_circuit_id in circuit_sets:
            if len(circuit_sets[relevant_circuit_id]) == len(junctions):
                ax = junctions[index_a][0]
                bx = junctions[index_b][0]
                result = ax * bx
                break
            
    print(f"Part Two: {result}")

# ----
partOne("input1.txt")
partTwo("input1.txt")