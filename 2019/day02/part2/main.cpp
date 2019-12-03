#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>

#include "main.hpp"

using namespace std;

int run_program(vector<int> v, int noun, int verb) {

    int first_address = 0;
    int second_address = 0;
    int first = 0;
    int second = 0;
    int output_address = 0;

    vector<int> vect = v;
    vect[1] = noun;
    vect[2] = verb;

    // cycle thru the program
    for (int i = 0;;) {

        // handle the opcode 1 case
        if (vect[i] == 1) {

            first_address = vect[i+1];
            second_address = vect[i+2];
            output_address = vect[i+3];

            first = vect[first_address];
            second = vect[second_address];

            vect[output_address] = first + second;

            i += 4;
            continue;
        }

        // handle the opcode 2 case
        if (vect[i] == 2) {

            first_address = vect[i+1];
            second_address = vect[i+2];
            output_address = vect[i+3];

            first = vect[first_address];
            second = vect[second_address];

            vect[output_address] = first * second;

            i += 4;
            continue;
        }

        // handle the opcode 99 case
        if (vect[i] == 99) {
            break;
        }
    }

    return vect[0];
}

int main() {

    string fileName = "input1.txt";
    string str;
    vector<int> vect;
    bool program_done = false;

    int i = 0;
    int j = 0;

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

    for (i = 0; i < 1000; i++) {
        for (j = 0; j < 1000; j++) {
            int result = run_program(vect, i, j);
            if (result == 19690720) {
                program_done = true;
                break;
            }
        }
        if (program_done) {
            break;
        }
    }

    cout << "Noun is: " << i << endl;
    cout << "Verb is: " << j << endl;

    return 0;
}
