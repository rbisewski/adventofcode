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

    int orbits = 0;
    for (auto const& [key, val] : bodies) {
        string current = key;
        while (true) {
            if (current.compare("COM") == 0) {
                break;
            }
            current = bodies[current];
            orbits++;
        }
    }

    cout << orbits << endl;

    return 0;
}
