#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>

#include "main.hpp"

using namespace std;

int calculateFuel(double mass) {

    double fuel = mass;

    fuel /= 3;
    fuel = floor(fuel);
    fuel -= 2;

    return int(fuel);
}

int main() {

    string fileName = "input1.txt";
    int32_t totalFuelRequirement = 0;
    int32_t fuelForTheFuel = 0;
    double mass = 0.0;
    int32_t impulse = 0;
    string str;

    ifstream in(fileName.c_str());
    
    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    
    while (getline(in, str)) {

    	if(str.size() <= 0) {
               continue;
        }

        mass = stof(str);
        impulse = calculateFuel(mass);
        totalFuelRequirement += impulse;

        // calculate the fuel-for-the-fuel
        fuelForTheFuel = 0;
        for (;;) {
            impulse = calculateFuel(impulse);
            if (impulse <= 0) {
                break;
            }
            fuelForTheFuel += impulse;
        }

        totalFuelRequirement += fuelForTheFuel;
    }

    in.close();

    cout << "The total fuel required is: " << totalFuelRequirement << endl;

    return 0;
}
