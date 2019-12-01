#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>

using namespace std;

int main() {

    string fileName = "input1.txt";
    int32_t totalFuelRequirement = 0;
    double current = 0.0;

    ifstream in(fileName.c_str());
    
    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    
    string str;
    while (getline(in, str)) {

    	if(str.size() <= 0) {
               continue;
        }

        current = stof(str);

        current /= 3;
        current = floor(current);
        current -= 2;

        totalFuelRequirement += int(current);
    }

    in.close();

    cout << "The total fuel is: " << totalFuelRequirement << endl;

    return 0;
}
