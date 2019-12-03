#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>

#include "main.hpp"

using namespace std;

int main() {

    string fileName = "input1.txt";
    string str;
    vector<int> vect;

    int i = 0;

    int first_address = 0;
    int second_address = 0;
    int first = 0;
    int second = 0;
    int output_address = 0;

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

    // replace position 1 with the value `12`
    vect[1] = 12;

    // replace position 2 with the value `2`
    vect[2] = 2;

    // cycle thru the program
    for (i = 0;;) {

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

    cout << "Value at position zero is: " << vect[0] << endl;

    return 0;
}
