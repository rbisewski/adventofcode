#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>
#include <cstdlib>

#include "main.hpp"

using namespace std;

struct coord {
    int32_t x;
    int32_t y;
};

struct movement {
    char direction;
    int distance;
};

int main() {

    string fileName = "input1.txt";
    string str;
    int wire_num = 1;
    string wire_directions[2];
    vector<string> wire1;
    vector<string> wire2;

    vector<movement> refined_wire1;
    vector<movement> refined_wire2;

    vector<coord> path1;
    vector<coord> path2;

    vector<coord> intersecting_coords;

    ifstream in(fileName.c_str());
    
    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }

    //
    // obtain the contents
    //

    while (getline(in, str)) {

        if (wire_num == 1) {
            stringstream ss(str);
            while(ss.good()) {
                string substr;
                getline(ss, substr, ',');
                wire1.push_back(substr);
            }
        }

        if (wire_num == 2) {
            stringstream ss(str);
            while(ss.good()) {
                string substr;
                getline(ss, substr, ',');
                wire2.push_back(substr);
            }
        }

        wire_num++;
    }

    in.close();

    //
    // obtain the list of movements
    //

    for (size_t i = 0; i < wire1.size(); i++) {
        movement point;
        point.direction = wire1[i].c_str()[0];
        stringstream ss(wire1[i].substr(1));
        while(ss.good()) {
            string substr;
            getline(ss, substr, '\0');
            point.distance = stoi(substr);
        }
        refined_wire1.push_back(point);
    }

    for (size_t i = 0; i < wire2.size(); i++) {
        movement point;
        point.direction = wire2[i].c_str()[0];
        stringstream ss(wire2[i].substr(1));
        while(ss.good()) {
            string substr;
            getline(ss, substr, '\0');
            point.distance = stoi(substr);
        }
        refined_wire2.push_back(point);
    }

    coord tracker;
    tracker.x = 0;
    tracker.y = 0;
    path1.push_back(tracker);

    for (size_t i = 0; i < refined_wire1.size(); i++) {
        switch (refined_wire1[i].direction) {
            case 'U':
                for (int32_t j = 0; j < refined_wire1[i].distance; j++) {
                    tracker.y++;
                    path1.push_back(tracker);
                }
                break;
            case 'D':
                for (int32_t j = 0; j < refined_wire1[i].distance; j++) {
                    tracker.y--;
                    path1.push_back(tracker);
                }
                break;
            case 'R':
                for (int32_t j = 0; j < refined_wire1[i].distance; j++) {
                    tracker.x++;
                    path1.push_back(tracker);
                }
                break;
            case 'L':
                for (int32_t j = 0; j < refined_wire1[i].distance; j++) {
                    tracker.x--;
                    path1.push_back(tracker);
                }
                break;
            default:
                break;
        }
    }

    tracker.x = 0;
    tracker.y = 0;
    path2.push_back(tracker);

    for (size_t i = 0; i < refined_wire2.size(); i++) {
        switch (refined_wire2[i].direction) {
            case 'U':
                for (int32_t j = 0; j < refined_wire2[i].distance; j++) {
                    tracker.y++;
                    path2.push_back(tracker);
                }
                break;
            case 'D':
                for (int32_t j = 0; j < refined_wire2[i].distance; j++) {
                    tracker.y--;
                    path2.push_back(tracker);
                }
                break;
            case 'R':
                for (int32_t j = 0; j < refined_wire2[i].distance; j++) {
                    tracker.x++;
                    path2.push_back(tracker);
                }
                break;
            case 'L':
                for (int32_t j = 0; j < refined_wire2[i].distance; j++) {
                    tracker.x--;
                    path2.push_back(tracker);
                }
                break;
            default:
                break;
        }
    }

    //
    // check if the movements intersect
    //
    for (size_t i = 0; i < path1.size(); i++) {
        for (size_t j = 0; j < path2.size(); j++) {
            coord point1 = path1[i];
            coord point2 = path2[j];
            if (point1.x == point2.x && point1.y == point2.y) {
                intersecting_coords.push_back(point1);
            }
        }
    }

    // pick some random very large value
    int shortest_manhattan_distance = 100000000;

    for (size_t i = 0; i < intersecting_coords.size(); i++) {
        int current_manhattan_distance = abs(intersecting_coords[i].x) + abs(intersecting_coords[i].y);
        if (current_manhattan_distance == 0) {
            continue;
        }
        if (current_manhattan_distance < shortest_manhattan_distance) {
            shortest_manhattan_distance = current_manhattan_distance;
        }
    }

    cout << "The shortest Manhattan distance is: " << shortest_manhattan_distance << endl;

    return 0;
}
