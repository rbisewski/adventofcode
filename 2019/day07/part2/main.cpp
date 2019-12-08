#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>

#include "main.hpp"

using namespace std;

struct entry {
    int first;
    int second;
    int third;
    int fourth;
    int fifth;
};

int run_program(vector<int> *vect, int input_signal, int phase_signal) {

    int first_address = 0;
    int second_address = 0;
    int first = 0;
    int second = 0;
    int output_address = 0;

    int opcode = 0;
    int first_param_mode = 0;
    int second_param_mode = 0;

    bool phase_signal_used = false;
    bool input_signal_used = false;
    int output = 0;

    // cycle thru the program
    for (int i = 0; i < int(vect->size());i++) {

        opcode = 0;
        first_param_mode = 0;
        second_param_mode = 0;

        string element = to_string(vect->at(i));

        switch (element.length()) {
            case 1:
                opcode = stoi(element.substr(0,1));
                break;
            case 2:
                opcode = stoi(element.substr(0,2));
                break;
            case 3:
                opcode = stoi(element.substr(1,2));
                first_param_mode = stoi(element.substr(0,1));
                break;
            case 4:
                opcode = stoi(element.substr(2,2));
                first_param_mode = stoi(element.substr(1,1)); 
                second_param_mode = stoi(element.substr(0,1));
                break;
            case 5:
                opcode = stoi(element.substr(3,2));
                first_param_mode = stoi(element.substr(2,1));
                second_param_mode = stoi(element.substr(1,1));
                break;
            default:
                break;
        }

        // ADD - handle the opcode 1 case
        if (opcode == 1) {

            output_address = vect->at(i+3);

            first = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(vect->at(i+1));
            } 

            second = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(vect->at(i+2));
            } 

            vect->at(output_address) = first + second;

            i += 4;
            i--;
            continue;
        }

        // MULTIPLY - handle the opcode 2 case
        if (opcode == 2) {

            output_address = vect->at(i+3);

            first_address = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(first_address);
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(second_address);
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect->at(output_address) = first * second;

            i += 4;
            i--;
            continue;
        }

        // INPUT - handle the opcode 3 case
        if (opcode == 3) {
            output_address = vect->at(i+1);
            if (!phase_signal_used) {
                vect->at(output_address) = phase_signal;
                phase_signal_used = true;
            } else if (!input_signal_used) {
                vect->at(output_address) = input_signal;
                input_signal_used = true;
            }
            i += 2;
            i--;
            continue;
        }

        // OUTPUT - handle the opcode 4 case
        if (opcode == 4) {

            first_address = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(first_address);
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            output = first;

            i += 2;
            i--;
            continue;
        }

        // JUMP-IF-TRUE - handle the opcode 5 case
        if (opcode == 5) {

            first = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(vect->at(i+1));
            } 

            second = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(vect->at(i+2));
            }

            if (first != 0) {
                vect->at(i) = second;
                i = second;
                i--;
                continue;
            }

            i += 3;
            i--;
            continue;
        }

        // JUMP-IF-FALSE - handle the opcode 6 case
        if (opcode == 6) {

            first_address = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(first_address);
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(second_address);
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            if (first == 0) {
                vect->at(i) = second;
                i = second;
                i--;
                continue;
            }

            i += 3;
            i--;
            continue;
        }

        // LESS-THAN - handle the opcode 7 case
        if (opcode == 7) {

            output_address = vect->at(i+3);

            first_address = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(first_address);
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(second_address);
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect->at(output_address) = 0;
            if (first < second) {
                vect->at(output_address) = 1;
            }

            i += 4;
            i--;
            continue;
        }

        // EQUALS - handle the opcode 8 case
        if (opcode == 8) {

            output_address = vect->at(i+3);

            first_address = vect->at(i+1);
            if (first_param_mode == 0) {
                first = vect->at(first_address);
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect->at(i+2);
            if (second_param_mode == 0) {
                second = vect->at(second_address);
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect->at(output_address) = 0;
            if (first == second) {
                vect->at(output_address) = 1;
            }

            i += 4;
            i--;
            continue;
        }

        // EXIT - handle the opcode 99 case
        if (opcode == 99) {
            break;
        }
    }
    return output;
}

int main() {

    string fileName = "input2.txt";
    string str;
    vector<int> vect;

    int input_signal = 0;
    int output = 0;

    int i = 0;

    int highest_value = 0;

    ifstream in(fileName.c_str());
    
    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    getline(in, str);
    in.close();

    if (str.size() <= 0) {
        cerr << "This input file is empty... " << fileName << endl;
    }

    stringstream ss(str);

    for (i = 0; ss >> i;) {
        vect.push_back(i);
        if (ss.peek() == ',') {
            ss.ignore();
        }
    }

    int first_value  = 0;
    int second_value = 0;
    int third_value  = 0;
    int fourth_value = 0;
    int fifth_value  = 0;
    bool discard = false;
    vector<entry> list;

    for (i = 56789; i <= 98765; i++) {

        string v = to_string(i);

        first_value  = int(v[0]) - 48;
        second_value = int(v[1]) - 48;
        third_value  = int(v[2]) - 48;
        fourth_value = int(v[3]) - 48;
        fifth_value  = int(v[4]) - 48;

        for (size_t j = 0; j < v.size(); j++) {
            if (v[j] == '0') {
                discard = true;
                break;
            }
            if (v[j] == '1') {
                discard = true;
                break;
            }
            if (v[j] == '2') {
                discard = true;
                break;
            }
            if (v[j] == '3') {
                discard = true;
                break;
            }
            if (v[j] == '4') {
                discard = true;
                break;
            }
            if (v.find_first_of(v[j]) != v.find_last_of(v[j])) {
                discard = true;
                break;
            }
        }

        if (discard) {
            discard = false;
            continue;
        }

        entry e;
        e.first = first_value;
        e.second = second_value;
        e.third = third_value;
        e.fourth = fourth_value;
        e.fifth = fifth_value;

        list.push_back(e);
    }

    for (size_t s = 0; s < list.size(); s++) {

        entry e = list[s];

        //cout << e.first << e.second << e.third << e.fourth << e.fifth << endl;

        output = run_program(&vect, input_signal, e.first);
        output = run_program(&vect, output, e.second);
        output = run_program(&vect, output, e.third);
        output = run_program(&vect, output, e.fourth);
        output = run_program(&vect, output, e.fifth);

        if (output > highest_value) {
            highest_value = output;
        }
    }

    cout << highest_value << endl;

    return 0;
}
