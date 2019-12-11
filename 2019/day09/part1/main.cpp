#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <charconv>
#include <array>
#include <algorithm>
#include <chrono>

#define ADD 1
#define MULTIPLY 2
#define INPUT 3
#define OUTPUT 4
#define JUMP_IF_TRUE 5
#define JUMP_IF_FALSE 6
#define LESS_THAN 7
#define EQUAL 8
#define RELATIVE 9
#define HALT 99

using namespace std;

using instruction_set = vector<long long>;

class intcode_computer {
private:
	int instruction_pointer = 0, relative_base = 0;
	bool has_halted, debug_toggle;
	vector<long long> input_data, output_data;
	instruction_set instructions_immutable = {}, instructions_mutable = {};

public:
	void add_input(vector<long long> input) {
		input_data.insert(input_data.end(), input.begin(), input.end());
	}

	long long  get_input() {
		if (has_input()) {
			long long  input = input_data.front();
			input_data.erase(input_data.begin());
			return input;
		}
	}

	void clear_input() {
		input_data.clear();
	}

	bool has_input() {
		return !input_data.empty();
	}

	void add_output(vector<long long> output) {
		output_data.insert(output_data.end(), output.begin(), output.end());
	}

	long long get_output() {
		if (has_output()) {
			long long  output = output_data.front();
			output_data.erase(output_data.begin());
			return output;
		}
	}

	void clear_output() {
		output_data.clear();
	}

	bool has_output() {
		return !output_data.empty();
	}

	bool has_finished() {
		return has_halted;
	}


	void set_immutable(instruction_set instructions) {
		instructions_immutable = instructions;
	}

	void reset_mutable() {
		relative_base = 0;
		instruction_pointer = 0;
		has_halted = false;
		instructions_mutable = instructions_immutable;
	}

	void toggle_debug() {
		debug_toggle = !debug_toggle;
	}

	bool has_debug() {
		return debug_toggle;
	}

	void run_mutable() {
		string str, inst;
		vector<int> args;
		long long size = 0, * p1 = &size, * p2 = &size, * p3 = &size;
		while (instruction_pointer < instructions_mutable.size()) {
			str.clear();
			args.clear();
			inst.clear();

			str = to_string(instructions_mutable[instruction_pointer]);
			size = str.size();

			for (int k = size - 1; k >= 0; --k) {
				if (size == 1 || k > size - 3) {
					inst += str[k];
				}
				else if (size > 1 && k < size - 2) {
					args.push_back(str[k] - '0');
				}
			}

			reverse(inst.begin(), inst.end());

			if (args.size() == 0) {
				args = { false, false, false };
			}

			p1 = args[0] == 0 ? &instructions_mutable[instructions_mutable[instruction_pointer + 1]] : args[0] == 1 ? &instructions_mutable[instruction_pointer + 1] : &instructions_mutable[relative_base + instructions_mutable[instruction_pointer + 1]];
			p2 = args[1] == 0 ? &instructions_mutable[instructions_mutable[instruction_pointer + 2]] : args[1] == 1 ? &instructions_mutable[instruction_pointer + 2] : &instructions_mutable[relative_base + instructions_mutable[instruction_pointer + 2]];
			p3 = args.size() != 3 ? &instructions_mutable[instructions_mutable[instruction_pointer + 3]] : &instructions_mutable[relative_base + instructions_mutable[instruction_pointer + 3]];

			if (has_debug()) {
				cout << endl << "Step " << instruction_pointer << ":" << endl;
			}

			switch (stoi(inst)) {
			case INPUT:
				if (!has_input()) {
					return;
				}
				if (has_debug()) {
					cout << "\t" << *p1 << " set to " << input_data.front() << "." << endl;
				}
				*p1 = get_input();
				instruction_pointer += 2;
				break;
			case OUTPUT:
				if (has_debug()) {
					cout << "\t" << *p1 << " added to output." << endl;
				}
				add_output({ *p1 });
				instruction_pointer += 2;
				break;
			case JUMP_IF_TRUE:
				if (has_debug()) {
					cout << "\t" << "i" << instruction_pointer << " set to " << (*p1 != 0) << " " << (*p1 != 0 ? *p2 : instruction_pointer + 3) << "." << endl;
				}
				instruction_pointer = *p1 != 0 ? *p2 : instruction_pointer + 3;
				break;
			case JUMP_IF_FALSE:
				if (has_debug()) {
					cout << "\t" << "i" << instruction_pointer << " set to " << (*p1 == 0) << " " << (*p1 == 0 ? *p2 : instruction_pointer + 3) << "." << endl;
				}
				instruction_pointer = *p1 == 0 ? *p2 : instruction_pointer + 3;
				break;
			case ADD:
				if (has_debug()) {
					cout << "\t" << *p3 << " set to " << (*p1 + *p2) << "." << endl;
				}
				*p3 = *p1 + *p2;
				instruction_pointer += 4;
				break;
			case MULTIPLY:
				if (has_debug()) {
					cout << "\t" << *p3 << " set to " << *p1 * *p2 << "." << endl;
				}
				*p3 = *p1 * *p2;
				instruction_pointer += 4;
				break;
			case LESS_THAN:
				if (has_debug()) {
					cout << "\t" << *p3 << " set to " << (*p1 < *p2) << "." << endl;
				}
				*p3 = *p1 < *p2;
				instruction_pointer += 4;
				break;
			case EQUAL:
				if (has_debug()) {
					cout << "\t" << *p3 << " set to " << (*p1 == *p2) << "." << endl;
				}
				*p3 = *p1 == *p2;
				instruction_pointer += 4;
				break;
			case RELATIVE:
				if (has_debug()) {
					cout << "\t" << "r" << relative_base << " changed by " << *p1 << "." << endl;
				}
				relative_base += *p1;
				instruction_pointer += 2;
				break;
			case HALT:
				if (has_debug()) {
					cout << "\t" << "Virtual machine halted. " << endl << endl;
				}
				has_halted = true;
				return;
			}
		}
	}

	void allocate(int amount) {
		for (int i = amount; i > 0; --i) {
			instructions_immutable.push_back(0);
		}
	}

	static void add_input_sequentially(vector<intcode_computer>& intcode_computers, vector<int> sequences) {
		for (int i = 0; i < intcode_computers.size() && i < sequences.size(); ++i) {
			intcode_computers[i].add_input({ sequences[i] });
		}
	}

	static void set_immutable_sequentially(vector<intcode_computer>& intcode_computers, instruction_set basecodes) {
		for (int i = 0; i < intcode_computers.size(); ++i) {
			intcode_computers[i].set_immutable({ basecodes });
		}
	}

	static void reset_mutable_sequentially(vector<intcode_computer>& intcode_computers) {
		for (int i = 0; i < intcode_computers.size(); ++i) {
			intcode_computers[i].reset_mutable();
		}
	}
};

int main() {
	string line, buffer;
	vector<long long> instructions;
	ifstream("input1.txt") >> line;
	for (int i = 0; i <= line.size(); ++i) {
		if (line[i] != ',') {
			buffer += line[i];
		}
		else if (!isdigit(line[i])) {
			instructions.push_back(stoll(buffer));
			buffer.clear();
		}
	}

	intcode_computer virtual_machine;
	virtual_machine.set_immutable(instructions);
	virtual_machine.allocate(2048);

	virtual_machine.reset_mutable();
	virtual_machine.add_input({ 1 });
	virtual_machine.run_mutable();

	cout << "Part I: " << virtual_machine.get_output() << endl;

	virtual_machine.reset_mutable();
	virtual_machine.add_input({ 2 });
	virtual_machine.run_mutable();

	cout << "Part II: " << virtual_machine.get_output() << endl;
}