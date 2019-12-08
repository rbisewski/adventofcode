#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>

#include "main.hpp"

using namespace std;

int run_program(vector<int> v, int input_signal, int phase_signal) {

    int first_address = 0;
    int second_address = 0;
    int first = 0;
    int second = 0;
    int output_address = 0;

    vector<int> vect = v;

    int opcode = 0;
    int first_param_mode = 0;
    int second_param_mode = 0;

    bool phase_signal_used = false;
    bool input_signal_used = false;
    int output = 0;

    // cycle thru the program
    for (int i = 0; i < int(vect.size());i++) {

        opcode = 0;
        first_param_mode = 0;
        second_param_mode = 0;

        switch (to_string(vect[i]).length()) {
            case 1:
                opcode = stoi(to_string(vect[i]).substr(0,1));
                break;
            case 2:
                opcode = stoi(to_string(vect[i]).substr(0,2));
                break;
            case 3:
                opcode = stoi(to_string(vect[i]).substr(1,2));
                first_param_mode = stoi(to_string(vect[i]).substr(0,1));
                break;
            case 4:
                opcode = stoi(to_string(vect[i]).substr(2,2));
                first_param_mode = stoi(to_string(vect[i]).substr(1,1)); 
                second_param_mode = stoi(to_string(vect[i]).substr(0,1));
                break;
            case 5:
                opcode = stoi(to_string(vect[i]).substr(3,2));
                first_param_mode = stoi(to_string(vect[i]).substr(2,1));
                second_param_mode = stoi(to_string(vect[i]).substr(1,1));
                break;
            default:
                break;
        }

        // ADD - handle the opcode 1 case
        if (opcode == 1) {

            output_address = vect[i+3];

            first = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[vect[i+1]];
            } 

            second = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[vect[i+2]];
            } 

            vect[output_address] = first + second;

            i += 4;
            i--;
            continue;
        }

        // MULTIPLY - handle the opcode 2 case
        if (opcode == 2) {

            output_address = vect[i+3];

            first_address = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[first_address];
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[second_address];
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect[output_address] = first * second;

            i += 4;
            i--;
            continue;
        }

        // INPUT - handle the opcode 3 case
        if (opcode == 3) {
            output_address = vect[i+1];
            if (!phase_signal_used) {
                vect[output_address] = phase_signal;
                phase_signal_used = true;
            } else if (!input_signal_used) {
                vect[output_address] = input_signal;
                input_signal_used = true;
            }
            i += 2;
            i--;
            continue;
        }

        // OUTPUT - handle the opcode 4 case
        if (opcode == 4) {

            first_address = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[first_address];
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

            first = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[vect[i+1]];
            } 

            second = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[vect[i+2]];
            }

            if (first != 0) {
                vect[i] = second;
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

            first_address = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[first_address];
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[second_address];
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            if (first == 0) {
                vect[i] = second;
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

            output_address = vect[i+3];

            first_address = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[first_address];
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[second_address];
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect[output_address] = 0;
            if (first < second) {
                vect[output_address] = 1;
            }

            i += 4;
            i--;
            continue;
        }

        // EQUALS - handle the opcode 8 case
        if (opcode == 8) {

            output_address = vect[i+3];

            first_address = vect[i+1];
            if (first_param_mode == 0) {
                first = vect[first_address];
            } else if (first_param_mode == 1) {
                first = first_address;
            }

            second_address = vect[i+2];
            if (second_param_mode == 0) {
                second = vect[second_address];
            } else if (second_param_mode == 1) {
                second = second_address;
            }

            vect[output_address] = 0;
            if (first == second) {
                vect[output_address] = 1;
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

    string fileName = "input1.txt";
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

    for (i = 1234; i < 44444; i++) {

        string v = to_string(i);

        switch (v.length()) {
            case 4:
                first_value  = 0;
                second_value = int(v[0]) - 48;
                third_value  = int(v[1]) - 48;
                fourth_value = int(v[2]) - 48;
                fifth_value  = int(v[3]) - 48;
                break;
            case 5:
                first_value  = int(v[0]) - 48;
                second_value = int(v[1]) - 48;
                third_value  = int(v[2]) - 48;
                fourth_value = int(v[3]) - 48;
                fifth_value  = int(v[4]) - 48;
                break;
            default:
                break;
        }

        for (size_t j = 0; j < v.size(); j++) {
            if (v.length() == 4 && (int(v[j])-48) == 0) {
                discard = true;
                break;
            }
            if (v[j] == '5') {
                discard = true;
                break;
            }
            if (v[j] == '6') {
                discard = true;
                break;
            }
            if (v[j] == '7') {
                discard = true;
                break;
            }
            if (v[j] == '8') {
                discard = true;
                break;
            }
            if (v[j] == '9') {
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

        output = run_program(vect, input_signal, first_value);
        output = run_program(vect, output, second_value);
        output = run_program(vect, output, third_value);
        output = run_program(vect, output, fourth_value);
        output = run_program(vect, output, fifth_value);

        if (output > highest_value) {
            highest_value = output;
        }
    }

    cout << highest_value << endl;

    return 0;
}
