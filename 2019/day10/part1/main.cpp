#include <cmath>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>
#include <map>
#include <set>

#include "main.hpp"

using namespace std;

int main() {

    string fileName = "input1.txt";
    string str;
    vector<string> input;

    ifstream in(fileName.c_str());

    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }
    while (getline(in, str)) {
        input.push_back(str);
    }
    in.close();

    if (str.length() <= 0) {
        cerr << "This input file is empty... " << fileName << endl;
    }

	map<int, int> asteroid_count;

	// Assuming a square map
	int dim = input.size();

	for (int y = 0; y < dim; y++) {
		for (int x = 0; x < dim; x++) {

			set<double> angle_set;
			if (input[y][x] == '#') {

				for (int y2 = 0; y2 < dim; y2++) {
					for (int x2 = 0; x2 < dim; x2++) {

						if (input[y2][x2] == '#' and !(y2 == y and x2 == x))
							angle_set.insert(atan2(y2 - y, x2 - x));

					}
				}
				asteroid_count[x * 100 + y] = angle_set.size();
			}
		}
	}

	auto max = max_element(asteroid_count.begin(), asteroid_count.end(), [](pair<int, int> a, pair<int, int> b) {
		return a.second < b.second;
	});

	int x = (max->first) / 100, y = (max->first) % 100;
	cout << "Number of asteriods? " << max->second << " (at " << x << ", " << y << ")\n";


	map<double, vector<int>> angle_map;
	for (int y2 = 0; y2 < dim; y2++) {
		for (int x2 = 0; x2 < dim; x2++) {
			if (input[y2][x2] == '#' and !(y2 == y and x2 == x)) {
				// Angle adjusted so that it runs from 0->2PI clockwise
				double angle = atan2(x - x2, y - y2);
				angle = angle <= 0 ? abs(angle) : 2 * M_PI - angle;
				angle_map[angle].push_back(x2 * 100 + y2);
			}
		}
	}


	for (auto a : angle_map) {
		sort(a.second.begin(), a.second.end(), [x, y](int a, int b) {
			double a_dist, b_dist;
			a_dist = sqrt(pow(a / 100 - x, 2) + pow(a % 100 - y, 2));
			b_dist = sqrt(pow(b / 100 - x, 2) + pow(b % 100 - y, 2));
			return b_dist < a_dist;
		});
	}


	vector<int> destroyed_asteroids;
	while (destroyed_asteroids.size() < 200) {
		for (auto a : angle_map) {
			if (a.second.size() > 0) {
				destroyed_asteroids.push_back(a.second.back());
				a.second.pop_back();
			}
		}
	}

	cout << "200th destroyed asteriod? " << destroyed_asteroids[199] << "\n";

	cout << "\n";

    return 0;
}
