#include <cmath>
#include <cstdlib>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>
#include <sstream>
#include <map>
#include <numeric>

#include "main.hpp"

using namespace std;

struct coord {
    int id;
    int t;
    int px;
    int py;
    int pz;
    int vx;
    int vy;
    int vz;
};

int main() {

    string fileName = "input1.txt";
    int last = 1000;
    string str;

    vector<coord> coords;
    vector<coord> timesteps_1;
    vector<coord> timesteps_2;
    vector<coord> timesteps_3;
    vector<coord> timesteps_4;

    ifstream in(fileName.c_str());

    if(!in) {
    	cerr << "Error: Unable to open the following file... " << fileName << endl;
    	return 1;
    }

    int id = 1;
    while (getline(in, str)) {

        coord a = {
            .id = id, 
            .t = 0,
            .px = 0,
            .py = 0,
            .pz = 0,
            .vx = 0,
            .vy = 0,
            .vz = 0
            };
        int comma_counter = 0;
        string tmp_string = "";

        for (size_t i = 0; i < str.length(); i++) {
        
            switch(str[i]) {
                case ',':
                    comma_counter++;

                    if (comma_counter == 1) {
                        a.px = stoi(tmp_string);
                        tmp_string = "";

                    } else if (comma_counter == 2) {
                        a.py = stoi(tmp_string);
                        tmp_string = "";

                    }
                    break;

                case '<':
                case '=':
                case ' ':
                case 'x':
                case 'y':
                case 'z':
                    break;

                case '>':
                    if (comma_counter == 2) {
                        a.pz = stoi(tmp_string);
                        tmp_string = "";
                    }
                    break;

                default:
                    tmp_string.push_back(str[i]);
                    break;
            }
        }

        coords.push_back(a);
    }

    in.close();

    if (coords.size() <= 0) {
        cerr << "This input file is empty... " << fileName << endl;
    }

    timesteps_1.push_back(coords[0]);
    timesteps_2.push_back(coords[1]);
    timesteps_3.push_back(coords[2]);
    timesteps_4.push_back(coords[3]);

    int t = 1;
    for (t = 1; t <= last; t++) {

        coord one = {
            .id = 1, 
            .t = t,
            .px = timesteps_1[t-1].px,
            .py = timesteps_1[t-1].py,
            .pz = timesteps_1[t-1].pz,
            .vx = timesteps_1[t-1].vx,
            .vy = timesteps_1[t-1].vy,
            .vz = timesteps_1[t-1].vz
            };

        coord two = {
            .id = 2, 
            .t = t,
            .px = timesteps_2[t-1].px,
            .py = timesteps_2[t-1].py,
            .pz = timesteps_2[t-1].pz,
            .vx = timesteps_2[t-1].vx,
            .vy = timesteps_2[t-1].vy,
            .vz = timesteps_2[t-1].vz
            };

        coord three = {
            .id = 3, 
            .t = t,
            .px = timesteps_3[t-1].px,
            .py = timesteps_3[t-1].py,
            .pz = timesteps_3[t-1].pz,
            .vx = timesteps_3[t-1].vx,
            .vy = timesteps_3[t-1].vy,
            .vz = timesteps_3[t-1].vz
            };

        coord four = {
            .id = 4, 
            .t = t,
            .px = timesteps_4[t-1].px,
            .py = timesteps_4[t-1].py,
            .pz = timesteps_4[t-1].pz,
            .vx = timesteps_4[t-1].vx,
            .vy = timesteps_4[t-1].vy,
            .vz = timesteps_4[t-1].vz
            };

        // add the elements to the timestep
        timesteps_1.push_back(one);
        timesteps_2.push_back(two);
        timesteps_3.push_back(three);
        timesteps_4.push_back(four);

        // apply gravity
        for (int i = 1; i <= 4; i++) {
            switch (i) {
                case 1:
                    // compare 1 to 2
                    if (one.px < two.px) {
                        timesteps_1[t].vx++;
                        timesteps_2[t].vx--;
                    } else if (one.px > two.px) {
                        timesteps_1[t].vx--;
                        timesteps_2[t].vx++;
                    }

                    if (one.py < two.py) {
                        timesteps_1[t].vy++;
                        timesteps_2[t].vy--;
                    } else if (one.py > two.py) {
                        timesteps_1[t].vy--;
                        timesteps_2[t].vy++;
                    }

                    if (one.pz < two.pz) {
                        timesteps_1[t].vz++;
                        timesteps_2[t].vz--;
                    } else if (one.pz > two.pz) {
                        timesteps_1[t].vz--;
                        timesteps_2[t].vz++;
                    }

                    // compare 1 to 3
                    if (one.px < three.px) {
                        timesteps_1[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (one.px > three.px) {
                        timesteps_1[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (one.py < three.py) {
                        timesteps_1[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (one.py > three.py) {
                        timesteps_1[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (one.pz < three.pz) {
                        timesteps_1[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (one.pz > three.pz) {
                        timesteps_1[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 1 to 4
                    if (one.px < four.px) {
                        timesteps_1[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (one.px > four.px) {
                        timesteps_1[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (one.py < four.py) {
                        timesteps_1[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (one.py > four.py) {
                        timesteps_1[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (one.pz < four.pz) {
                        timesteps_1[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (one.pz > four.pz) {
                        timesteps_1[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;

                case 2:
                    // compare 2 to 3
                    if (two.px < three.px) {
                        timesteps_2[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (two.px > three.px) {
                        timesteps_2[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (two.py < three.py) {
                        timesteps_2[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (two.py > three.py) {
                        timesteps_2[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (two.pz < three.pz) {
                        timesteps_2[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (two.pz > three.pz) {
                        timesteps_2[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 2 to 4
                    if (two.px < four.px) {
                        timesteps_2[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (two.px > four.px) {
                        timesteps_2[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (two.py < four.py) {
                        timesteps_2[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (two.py > four.py) {
                        timesteps_2[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (two.pz < four.pz) {
                        timesteps_2[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (two.pz > four.pz) {
                        timesteps_2[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 3:
                    // compare 3 to 4
                    if (three.px < four.px) {
                        timesteps_3[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (three.px > four.px) {
                        timesteps_3[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (three.py < four.py) {
                        timesteps_3[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (three.py > four.py) {
                        timesteps_3[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (three.pz < four.pz) {
                        timesteps_3[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (three.pz > four.pz) {
                        timesteps_3[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 4:
                    break;
                default:
                    break;
            }
        }

        // apply velocity
        timesteps_1[t].px += timesteps_1[t].vx;
        timesteps_1[t].py += timesteps_1[t].vy;
        timesteps_1[t].pz += timesteps_1[t].vz;

        timesteps_2[t].px += timesteps_2[t].vx;
        timesteps_2[t].py += timesteps_2[t].vy;
        timesteps_2[t].pz += timesteps_2[t].vz;

        timesteps_3[t].px += timesteps_3[t].vx;
        timesteps_3[t].py += timesteps_3[t].vy;
        timesteps_3[t].pz += timesteps_3[t].vz;

        timesteps_4[t].px += timesteps_4[t].vx;
        timesteps_4[t].py += timesteps_4[t].vy;
        timesteps_4[t].pz += timesteps_4[t].vz;

        //cout << "After " << t << " steps:" << endl;

        //cout << "pos=<x=" << timesteps_1[t].px << ",y=" << timesteps_1[t].py << ",z=" << timesteps_1[t].pz << ">,";
        //cout << "vel=<x=" << timesteps_1[t].vx << ",y=" << timesteps_1[t].vy << ",z=" << timesteps_1[t].vz << ">" << endl;

        //cout << "pos=<x=" << timesteps_2[t].px << ",y=" << timesteps_2[t].py << ",z=" << timesteps_2[t].pz << ">,";
        //cout << "vel=<x=" << timesteps_2[t].vx << ",y=" << timesteps_2[t].vy << ",z=" << timesteps_2[t].vz << ">" << endl;

        //cout << "pos=<x=" << timesteps_3[t].px << ",y=" << timesteps_3[t].py << ",z=" << timesteps_3[t].pz << ">,";
        //cout << "vel=<x=" << timesteps_3[t].vx << ",y=" << timesteps_3[t].vy << ",z=" << timesteps_3[t].vz << ">" << endl;

        //cout << "pos=<x=" << timesteps_4[t].px << ",y=" << timesteps_4[t].py << ",z=" << timesteps_4[t].pz << ">,";
        //cout << "vel=<x=" << timesteps_4[t].vx << ",y=" << timesteps_4[t].vy << ",z=" << timesteps_4[t].vz << ">" << endl;

        //cout << endl;
    }

    int potential_1 = abs(timesteps_1[last].px) + abs(timesteps_1[last].py) + abs(timesteps_1[last].pz);
    int potential_2 = abs(timesteps_2[last].px) + abs(timesteps_2[last].py) + abs(timesteps_2[last].pz);
    int potential_3 = abs(timesteps_3[last].px) + abs(timesteps_3[last].py) + abs(timesteps_3[last].pz);
    int potential_4 = abs(timesteps_4[last].px) + abs(timesteps_4[last].py) + abs(timesteps_4[last].pz);

    int kinetic_1 = abs(timesteps_1[last].vx) + abs(timesteps_1[last].vy) + abs(timesteps_1[last].vz);
    int kinetic_2 = abs(timesteps_2[last].vx) + abs(timesteps_2[last].vy) + abs(timesteps_2[last].vz);
    int kinetic_3 = abs(timesteps_3[last].vx) + abs(timesteps_3[last].vy) + abs(timesteps_3[last].vz);
    int kinetic_4 = abs(timesteps_4[last].vx) + abs(timesteps_4[last].vy) + abs(timesteps_4[last].vz);

    int energy_1 = potential_1 * kinetic_1;
    int energy_2 = potential_2 * kinetic_2;
    int energy_3 = potential_3 * kinetic_3;
    int energy_4 = potential_4 * kinetic_4;

    int total_energy = energy_1 + energy_2 + energy_3 + energy_4;

    cout << "Total energy is: " << total_energy << endl;

    //
    // part 2, find x
    //
    timesteps_1.clear();
    timesteps_2.clear();
    timesteps_3.clear();
    timesteps_4.clear();

    timesteps_1.push_back(coords[0]);
    timesteps_2.push_back(coords[1]);
    timesteps_3.push_back(coords[2]);
    timesteps_4.push_back(coords[3]);

    t = 1;
    int repeat_x = 0;
    for (t = 1; ; t++) {

        coord one = {
            .id = 1, 
            .t = t,
            .px = timesteps_1[t-1].px,
            .py = timesteps_1[t-1].py,
            .pz = timesteps_1[t-1].pz,
            .vx = timesteps_1[t-1].vx,
            .vy = timesteps_1[t-1].vy,
            .vz = timesteps_1[t-1].vz
            };

        coord two = {
            .id = 2, 
            .t = t,
            .px = timesteps_2[t-1].px,
            .py = timesteps_2[t-1].py,
            .pz = timesteps_2[t-1].pz,
            .vx = timesteps_2[t-1].vx,
            .vy = timesteps_2[t-1].vy,
            .vz = timesteps_2[t-1].vz
            };

        coord three = {
            .id = 3, 
            .t = t,
            .px = timesteps_3[t-1].px,
            .py = timesteps_3[t-1].py,
            .pz = timesteps_3[t-1].pz,
            .vx = timesteps_3[t-1].vx,
            .vy = timesteps_3[t-1].vy,
            .vz = timesteps_3[t-1].vz
            };

        coord four = {
            .id = 4, 
            .t = t,
            .px = timesteps_4[t-1].px,
            .py = timesteps_4[t-1].py,
            .pz = timesteps_4[t-1].pz,
            .vx = timesteps_4[t-1].vx,
            .vy = timesteps_4[t-1].vy,
            .vz = timesteps_4[t-1].vz
            };

        // add the elements to the timestep
        timesteps_1.push_back(one);
        timesteps_2.push_back(two);
        timesteps_3.push_back(three);
        timesteps_4.push_back(four);

        // apply gravity
        for (int i = 1; i <= 4; i++) {
            switch (i) {
                case 1:
                    // compare 1 to 2
                    if (one.px < two.px) {
                        timesteps_1[t].vx++;
                        timesteps_2[t].vx--;
                    } else if (one.px > two.px) {
                        timesteps_1[t].vx--;
                        timesteps_2[t].vx++;
                    }

                    if (one.py < two.py) {
                        timesteps_1[t].vy++;
                        timesteps_2[t].vy--;
                    } else if (one.py > two.py) {
                        timesteps_1[t].vy--;
                        timesteps_2[t].vy++;
                    }

                    if (one.pz < two.pz) {
                        timesteps_1[t].vz++;
                        timesteps_2[t].vz--;
                    } else if (one.pz > two.pz) {
                        timesteps_1[t].vz--;
                        timesteps_2[t].vz++;
                    }

                    // compare 1 to 3
                    if (one.px < three.px) {
                        timesteps_1[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (one.px > three.px) {
                        timesteps_1[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (one.py < three.py) {
                        timesteps_1[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (one.py > three.py) {
                        timesteps_1[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (one.pz < three.pz) {
                        timesteps_1[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (one.pz > three.pz) {
                        timesteps_1[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 1 to 4
                    if (one.px < four.px) {
                        timesteps_1[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (one.px > four.px) {
                        timesteps_1[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (one.py < four.py) {
                        timesteps_1[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (one.py > four.py) {
                        timesteps_1[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (one.pz < four.pz) {
                        timesteps_1[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (one.pz > four.pz) {
                        timesteps_1[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;

                case 2:
                    // compare 2 to 3
                    if (two.px < three.px) {
                        timesteps_2[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (two.px > three.px) {
                        timesteps_2[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (two.py < three.py) {
                        timesteps_2[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (two.py > three.py) {
                        timesteps_2[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (two.pz < three.pz) {
                        timesteps_2[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (two.pz > three.pz) {
                        timesteps_2[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 2 to 4
                    if (two.px < four.px) {
                        timesteps_2[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (two.px > four.px) {
                        timesteps_2[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (two.py < four.py) {
                        timesteps_2[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (two.py > four.py) {
                        timesteps_2[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (two.pz < four.pz) {
                        timesteps_2[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (two.pz > four.pz) {
                        timesteps_2[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 3:
                    // compare 3 to 4
                    if (three.px < four.px) {
                        timesteps_3[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (three.px > four.px) {
                        timesteps_3[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (three.py < four.py) {
                        timesteps_3[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (three.py > four.py) {
                        timesteps_3[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (three.pz < four.pz) {
                        timesteps_3[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (three.pz > four.pz) {
                        timesteps_3[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 4:
                    break;
                default:
                    break;
            }
        }

        // apply velocity
        timesteps_1[t].px += timesteps_1[t].vx;
        timesteps_1[t].py += timesteps_1[t].vy;
        timesteps_1[t].pz += timesteps_1[t].vz;

        timesteps_2[t].px += timesteps_2[t].vx;
        timesteps_2[t].py += timesteps_2[t].vy;
        timesteps_2[t].pz += timesteps_2[t].vz;

        timesteps_3[t].px += timesteps_3[t].vx;
        timesteps_3[t].py += timesteps_3[t].vy;
        timesteps_3[t].pz += timesteps_3[t].vz;

        timesteps_4[t].px += timesteps_4[t].vx;
        timesteps_4[t].py += timesteps_4[t].vy;
        timesteps_4[t].pz += timesteps_4[t].vz;

        //
        // check to see if a given state ever existed before
        //
        bool previous_state_detected = false;
        for (int u = 0; u < t; u++) {
            if (timesteps_1[t].px == timesteps_1[u].px
            && timesteps_1[t].vx == 0
            && timesteps_2[t].px == timesteps_2[u].px
            && timesteps_2[t].vx == 0
            && timesteps_3[t].px == timesteps_3[u].px
            && timesteps_3[t].vx == 0
            && timesteps_4[t].px == timesteps_4[u].px
            && timesteps_4[t].vx == 0) {

                previous_state_detected = true;
                break;
            }
        }

        if (previous_state_detected) {
            break;
        }

        repeat_x++;
    }

    //
    // part 2, find y
    //
    timesteps_1.clear();
    timesteps_2.clear();
    timesteps_3.clear();
    timesteps_4.clear();

    timesteps_1.push_back(coords[0]);
    timesteps_2.push_back(coords[1]);
    timesteps_3.push_back(coords[2]);
    timesteps_4.push_back(coords[3]);

    t = 1;
    int repeat_y = 0;
    for (t = 1; ; t++) {

        coord one = {
            .id = 1, 
            .t = t,
            .px = timesteps_1[t-1].px,
            .py = timesteps_1[t-1].py,
            .pz = timesteps_1[t-1].pz,
            .vx = timesteps_1[t-1].vx,
            .vy = timesteps_1[t-1].vy,
            .vz = timesteps_1[t-1].vz
            };

        coord two = {
            .id = 2, 
            .t = t,
            .px = timesteps_2[t-1].px,
            .py = timesteps_2[t-1].py,
            .pz = timesteps_2[t-1].pz,
            .vx = timesteps_2[t-1].vx,
            .vy = timesteps_2[t-1].vy,
            .vz = timesteps_2[t-1].vz
            };

        coord three = {
            .id = 3, 
            .t = t,
            .px = timesteps_3[t-1].px,
            .py = timesteps_3[t-1].py,
            .pz = timesteps_3[t-1].pz,
            .vx = timesteps_3[t-1].vx,
            .vy = timesteps_3[t-1].vy,
            .vz = timesteps_3[t-1].vz
            };

        coord four = {
            .id = 4, 
            .t = t,
            .px = timesteps_4[t-1].px,
            .py = timesteps_4[t-1].py,
            .pz = timesteps_4[t-1].pz,
            .vx = timesteps_4[t-1].vx,
            .vy = timesteps_4[t-1].vy,
            .vz = timesteps_4[t-1].vz
            };

        // add the elements to the timestep
        timesteps_1.push_back(one);
        timesteps_2.push_back(two);
        timesteps_3.push_back(three);
        timesteps_4.push_back(four);

        // apply gravity
        for (int i = 1; i <= 4; i++) {
            switch (i) {
                case 1:
                    // compare 1 to 2
                    if (one.px < two.px) {
                        timesteps_1[t].vx++;
                        timesteps_2[t].vx--;
                    } else if (one.px > two.px) {
                        timesteps_1[t].vx--;
                        timesteps_2[t].vx++;
                    }

                    if (one.py < two.py) {
                        timesteps_1[t].vy++;
                        timesteps_2[t].vy--;
                    } else if (one.py > two.py) {
                        timesteps_1[t].vy--;
                        timesteps_2[t].vy++;
                    }

                    if (one.pz < two.pz) {
                        timesteps_1[t].vz++;
                        timesteps_2[t].vz--;
                    } else if (one.pz > two.pz) {
                        timesteps_1[t].vz--;
                        timesteps_2[t].vz++;
                    }

                    // compare 1 to 3
                    if (one.px < three.px) {
                        timesteps_1[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (one.px > three.px) {
                        timesteps_1[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (one.py < three.py) {
                        timesteps_1[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (one.py > three.py) {
                        timesteps_1[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (one.pz < three.pz) {
                        timesteps_1[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (one.pz > three.pz) {
                        timesteps_1[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 1 to 4
                    if (one.px < four.px) {
                        timesteps_1[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (one.px > four.px) {
                        timesteps_1[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (one.py < four.py) {
                        timesteps_1[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (one.py > four.py) {
                        timesteps_1[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (one.pz < four.pz) {
                        timesteps_1[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (one.pz > four.pz) {
                        timesteps_1[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;

                case 2:
                    // compare 2 to 3
                    if (two.px < three.px) {
                        timesteps_2[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (two.px > three.px) {
                        timesteps_2[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (two.py < three.py) {
                        timesteps_2[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (two.py > three.py) {
                        timesteps_2[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (two.pz < three.pz) {
                        timesteps_2[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (two.pz > three.pz) {
                        timesteps_2[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 2 to 4
                    if (two.px < four.px) {
                        timesteps_2[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (two.px > four.px) {
                        timesteps_2[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (two.py < four.py) {
                        timesteps_2[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (two.py > four.py) {
                        timesteps_2[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (two.pz < four.pz) {
                        timesteps_2[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (two.pz > four.pz) {
                        timesteps_2[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 3:
                    // compare 3 to 4
                    if (three.px < four.px) {
                        timesteps_3[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (three.px > four.px) {
                        timesteps_3[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (three.py < four.py) {
                        timesteps_3[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (three.py > four.py) {
                        timesteps_3[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (three.pz < four.pz) {
                        timesteps_3[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (three.pz > four.pz) {
                        timesteps_3[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 4:
                    break;
                default:
                    break;
            }
        }

        // apply velocity
        timesteps_1[t].px += timesteps_1[t].vx;
        timesteps_1[t].py += timesteps_1[t].vy;
        timesteps_1[t].pz += timesteps_1[t].vz;

        timesteps_2[t].px += timesteps_2[t].vx;
        timesteps_2[t].py += timesteps_2[t].vy;
        timesteps_2[t].pz += timesteps_2[t].vz;

        timesteps_3[t].px += timesteps_3[t].vx;
        timesteps_3[t].py += timesteps_3[t].vy;
        timesteps_3[t].pz += timesteps_3[t].vz;

        timesteps_4[t].px += timesteps_4[t].vx;
        timesteps_4[t].py += timesteps_4[t].vy;
        timesteps_4[t].pz += timesteps_4[t].vz;

        //
        // check to see if a given state ever existed before
        //
        bool previous_state_detected = false;
        for (int u = 0; u < t; u++) {
            if (timesteps_1[t].py == timesteps_1[u].py
             && timesteps_1[t].vy == 0
             && timesteps_2[t].py == timesteps_2[u].py
             && timesteps_2[t].vy == 0
             && timesteps_3[t].py == timesteps_3[u].py
             && timesteps_3[t].vy == 0
             && timesteps_4[t].py == timesteps_4[u].py
             && timesteps_4[t].vy == 0) {

                previous_state_detected = true;
                break;
            }
        }

        if (previous_state_detected) {
            break;
        }

        repeat_y++;
    }

     //
    // part 2, find y
    //
    timesteps_1.clear();
    timesteps_2.clear();
    timesteps_3.clear();
    timesteps_4.clear();

    timesteps_1.push_back(coords[0]);
    timesteps_2.push_back(coords[1]);
    timesteps_3.push_back(coords[2]);
    timesteps_4.push_back(coords[3]);

    t = 1;
    int repeat_z = 0;
    for (t = 1; ; t++) {

        coord one = {
            .id = 1, 
            .t = t,
            .px = timesteps_1[t-1].px,
            .py = timesteps_1[t-1].py,
            .pz = timesteps_1[t-1].pz,
            .vx = timesteps_1[t-1].vx,
            .vy = timesteps_1[t-1].vy,
            .vz = timesteps_1[t-1].vz
            };

        coord two = {
            .id = 2, 
            .t = t,
            .px = timesteps_2[t-1].px,
            .py = timesteps_2[t-1].py,
            .pz = timesteps_2[t-1].pz,
            .vx = timesteps_2[t-1].vx,
            .vy = timesteps_2[t-1].vy,
            .vz = timesteps_2[t-1].vz
            };

        coord three = {
            .id = 3, 
            .t = t,
            .px = timesteps_3[t-1].px,
            .py = timesteps_3[t-1].py,
            .pz = timesteps_3[t-1].pz,
            .vx = timesteps_3[t-1].vx,
            .vy = timesteps_3[t-1].vy,
            .vz = timesteps_3[t-1].vz
            };

        coord four = {
            .id = 4, 
            .t = t,
            .px = timesteps_4[t-1].px,
            .py = timesteps_4[t-1].py,
            .pz = timesteps_4[t-1].pz,
            .vx = timesteps_4[t-1].vx,
            .vy = timesteps_4[t-1].vy,
            .vz = timesteps_4[t-1].vz
            };

        // add the elements to the timestep
        timesteps_1.push_back(one);
        timesteps_2.push_back(two);
        timesteps_3.push_back(three);
        timesteps_4.push_back(four);

        // apply gravity
        for (int i = 1; i <= 4; i++) {
            switch (i) {
                case 1:
                    // compare 1 to 2
                    if (one.px < two.px) {
                        timesteps_1[t].vx++;
                        timesteps_2[t].vx--;
                    } else if (one.px > two.px) {
                        timesteps_1[t].vx--;
                        timesteps_2[t].vx++;
                    }

                    if (one.py < two.py) {
                        timesteps_1[t].vy++;
                        timesteps_2[t].vy--;
                    } else if (one.py > two.py) {
                        timesteps_1[t].vy--;
                        timesteps_2[t].vy++;
                    }

                    if (one.pz < two.pz) {
                        timesteps_1[t].vz++;
                        timesteps_2[t].vz--;
                    } else if (one.pz > two.pz) {
                        timesteps_1[t].vz--;
                        timesteps_2[t].vz++;
                    }

                    // compare 1 to 3
                    if (one.px < three.px) {
                        timesteps_1[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (one.px > three.px) {
                        timesteps_1[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (one.py < three.py) {
                        timesteps_1[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (one.py > three.py) {
                        timesteps_1[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (one.pz < three.pz) {
                        timesteps_1[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (one.pz > three.pz) {
                        timesteps_1[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 1 to 4
                    if (one.px < four.px) {
                        timesteps_1[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (one.px > four.px) {
                        timesteps_1[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (one.py < four.py) {
                        timesteps_1[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (one.py > four.py) {
                        timesteps_1[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (one.pz < four.pz) {
                        timesteps_1[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (one.pz > four.pz) {
                        timesteps_1[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;

                case 2:
                    // compare 2 to 3
                    if (two.px < three.px) {
                        timesteps_2[t].vx++;
                        timesteps_3[t].vx--;
                    } else if (two.px > three.px) {
                        timesteps_2[t].vx--;
                        timesteps_3[t].vx++;
                    }

                    if (two.py < three.py) {
                        timesteps_2[t].vy++;
                        timesteps_3[t].vy--;
                    } else if (two.py > three.py) {
                        timesteps_2[t].vy--;
                        timesteps_3[t].vy++;
                    }

                    if (two.pz < three.pz) {
                        timesteps_2[t].vz++;
                        timesteps_3[t].vz--;
                    } else if (two.pz > three.pz) {
                        timesteps_2[t].vz--;
                        timesteps_3[t].vz++;
                    }

                    // compare 2 to 4
                    if (two.px < four.px) {
                        timesteps_2[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (two.px > four.px) {
                        timesteps_2[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (two.py < four.py) {
                        timesteps_2[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (two.py > four.py) {
                        timesteps_2[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (two.pz < four.pz) {
                        timesteps_2[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (two.pz > four.pz) {
                        timesteps_2[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 3:
                    // compare 3 to 4
                    if (three.px < four.px) {
                        timesteps_3[t].vx++;
                        timesteps_4[t].vx--;
                    } else if (three.px > four.px) {
                        timesteps_3[t].vx--;
                        timesteps_4[t].vx++;
                    }

                    if (three.py < four.py) {
                        timesteps_3[t].vy++;
                        timesteps_4[t].vy--;
                    } else if (three.py > four.py) {
                        timesteps_3[t].vy--;
                        timesteps_4[t].vy++;
                    }

                    if (three.pz < four.pz) {
                        timesteps_3[t].vz++;
                        timesteps_4[t].vz--;
                    } else if (three.pz > four.pz) {
                        timesteps_3[t].vz--;
                        timesteps_4[t].vz++;
                    }
                    break;
                case 4:
                    break;
                default:
                    break;
            }
        }

        // apply velocity
        timesteps_1[t].px += timesteps_1[t].vx;
        timesteps_1[t].py += timesteps_1[t].vy;
        timesteps_1[t].pz += timesteps_1[t].vz;

        timesteps_2[t].px += timesteps_2[t].vx;
        timesteps_2[t].py += timesteps_2[t].vy;
        timesteps_2[t].pz += timesteps_2[t].vz;

        timesteps_3[t].px += timesteps_3[t].vx;
        timesteps_3[t].py += timesteps_3[t].vy;
        timesteps_3[t].pz += timesteps_3[t].vz;

        timesteps_4[t].px += timesteps_4[t].vx;
        timesteps_4[t].py += timesteps_4[t].vy;
        timesteps_4[t].pz += timesteps_4[t].vz;

        //
        // check to see if a given state ever existed before
        //
        bool previous_state_detected = false;
        for (int u = 0; u < t; u++) {
            if (timesteps_1[t].pz == timesteps_1[u].pz
             && timesteps_1[t].vz == 0 
             && timesteps_2[t].pz == timesteps_2[u].pz
             && timesteps_2[t].vz == 0
             && timesteps_3[t].pz == timesteps_3[u].pz
             && timesteps_3[t].vz == 0 
             && timesteps_4[t].pz == timesteps_4[u].pz
             && timesteps_4[t].vz == 0) {

                previous_state_detected = true;
                break;
            }
        }

        if (previous_state_detected) {
            break;
        }

        repeat_z++;
    }

    repeat_x++;
    repeat_y++;
    repeat_z++;

    cout << "The number of steps until a previous state reoccurs is: ";
    cout << 2*(lcm((int64_t)repeat_z, lcm((int64_t)repeat_x, (int64_t)repeat_y))) << endl;

    return 0;
}
