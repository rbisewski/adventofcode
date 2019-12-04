#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>
#include <cstdlib>

#include "main.hpp"

using namespace std;

int main() {

    int start = 284639;
    int end = 748758;
    bool valid = false;

    vector<int> partial_list;
    vector<int> list;
    vector<int> list_wo_nondoubles;

    int counts[10];

    //
    // obtain list of numbers where each digit is >= the previous
    //
    for (int i = start; i <= end; i++) {
        valid = true;
        string tmp = to_string(i);
        for (int j = 0; j < 6; j++) {
            if (j == 0) {
                continue;
            }
            if (int(tmp[j]) < int(tmp[j-1])) {
                valid = false;
                break;
            }
        }
        if (!valid) {
            continue;
        }
        partial_list.push_back(i);
    }

    //
    // obtain list of numbers which have multiples of a given digit
    //
    for (size_t i = 0; i < partial_list.size(); i++) {
        valid = false;
        string tmp = to_string(partial_list[i]);
        for (int j = 0; j < 6; j++) {
            if (j == 0) {
                continue;
            }
            if (int(tmp[j]) == int(tmp[j-1])) {
                valid = true;
                break;
            }
        }
        if (!valid) {
            continue;
        }
        list.push_back(partial_list[i]);
    }

    //
    // remove instances without doubles
    //
    for (size_t i = 0; i < list.size(); i++) {
        valid = false;
        string tmp = to_string(list[i]);
        for (int j = 0; j < 6; j++) {
            int a = int(tmp[j]) - 48;
            counts[a]++;
        }
        for (int j = 0; j < 10; j++) {
            // only interested in doubles
            if (counts[j] == 2) {
                valid = true;
            }
            // reset this for the next pass-thru
            counts[j] = 0;
        }
        if (!valid) {
            continue;
        }
        list_wo_nondoubles.push_back(list[i]);
    }

    cout << list_wo_nondoubles.size() << endl;

    return 0;
}
