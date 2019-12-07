#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>
#include <map>

#include "main.hpp"

using namespace std;

int main() {

    string fileName = "input1.txt";
    string delimiter = ")";
    string str;
    string current;

    map<string,string> bodies;

    ifstream in(fileName.c_str());
    
    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    while (getline(in, str)) {
        string parent = str.substr(0,str.find(delimiter));
        string child = str.substr(str.find(delimiter)+1,str.length()-1);
        bodies[child] = parent;
    }
    in.close();

    if (bodies.size() <= 0) {
        cerr << "This input file is empty... " << fileName << endl;
    }

    vector<string> you_path;
    int YOU_to_COM = 0;
    current = "YOU";
    while (true) {
        you_path.push_back(current);
        if (current.compare("COM") == 0) {
            break;
        }
        current = bodies[current];
        YOU_to_COM++;
    }

    vector<string> san_path;
    int SAN_to_COM = 0;
    current = "SAN";
    while (true) {
        san_path.push_back(current);
        if (current.compare("COM") == 0) {
            break;
        }
        current = bodies[current];
        SAN_to_COM++;
    }

    string element;
    bool element_found = false;
    for (size_t i = 0; i < you_path.size(); i++) {
        for (size_t j = 0; j < san_path.size(); j++) {
            if (you_path[i] == san_path[j]) {
                element_found = true;
                element = you_path[i];
                break;
            }
        }
        if (element_found) {
            break;
        }
    }

    //
    // calculate the distance between the mid-point
    //
    // each couter starts at -1 since we ignore the starting position
    //

    int YOU_to_element = -1;
    current = "YOU";
    while (true) {
        if (current.compare(element) == 0) {
            break;
        }
        current = bodies[current];
        YOU_to_element++;
    }

    int SAN_to_element = -1;
    current = "SAN";
    while (true) {
        if (current.compare(element) == 0) {
            break;
        }
        current = bodies[current];
        SAN_to_element++;
    }

    cout << (YOU_to_element + SAN_to_element) << endl;

    return 0;
}
