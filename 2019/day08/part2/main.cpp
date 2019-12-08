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
    string str;
    ifstream in(fileName.c_str());

    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    while (getline(in, str)) {}
    in.close();

    if (str.length() <= 0) {
        cerr << "This input file is empty... " << fileName << endl;
    }

    int wide_counter = 0;
    int tall_counter = 0;
    int layer_counter = 0;

    for (size_t i = 0; i < str.length(); i++) {

        if (wide_counter == 24) {

            if (tall_counter == 5) {
                wide_counter = 0;
                tall_counter = 0;
                layer_counter++;
                continue;
            }

            wide_counter = 0;
            tall_counter++;
            continue;
        }

        wide_counter++;
    }

    int final_image[25][6] = {2};

    wide_counter = 0;
    tall_counter = 0;
    layer_counter = 0;

    int layer_with_fewest_0_digits = 0;
    int fewest_zero_digits = 10000;
    int ones = 0;
    int twos = 0;

    int zero_digits = 0;
    int one_digits = 0;
    int two_digits = 0;

    int current_digit = 0;

    for (size_t i = 0; i < str.length(); i++) {

        current_digit = int(str[i])-48;

        if (layer_counter == 0) {
            final_image[wide_counter][tall_counter] = current_digit;
        } else {
            if (final_image[wide_counter][tall_counter] == 2) {
                final_image[wide_counter][tall_counter] = current_digit;
            }
        }

        switch(current_digit) {
            case 0:
                zero_digits++;
                break;
            case 1:
                one_digits++;
                break;
            case 2:
                two_digits++;
                break;
            default:
                break;
        }

        if (wide_counter == 24) {

            if (tall_counter == 5) {
                wide_counter = 0;
                tall_counter = 0;

                if (zero_digits < fewest_zero_digits) {
                    fewest_zero_digits = zero_digits;
                    layer_with_fewest_0_digits = layer_counter;
                    ones = one_digits;
                    twos = two_digits;
                }

                zero_digits = 0;
                one_digits = 0;
                two_digits = 0;
                layer_counter++;
                continue;
            }

            wide_counter = 0;
            tall_counter++;
            continue;
        }

        wide_counter++;
    }

    cout << "Layer " << layer_with_fewest_0_digits << endl;
    cout << "-> Zero digits: " << fewest_zero_digits << endl;
    cout << "-> One  digits: " << ones << endl;
    cout << "-> Two  digits: " << twos << endl;
    cout << "---" << endl << ones * twos << endl;

    cout << endl << "Image produced: " << endl << endl;
    for (int i = 0; i < 6; i++) {
        for (int j = 0; j < 25; j++) {
            if (final_image[j][i] == 0) cout << " ";
            else cout << final_image[j][i];
        }
        cout << endl;
    }
    cout << endl;

    return 0;
}
