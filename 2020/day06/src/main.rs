use std::fs;
use std::collections::HashMap;

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let lines: Vec<&str> = contents.trim().split("\n").collect();

    let mut groups: Vec<String> = vec![];

    // sort raw data
    let mut group: String = String::from("");
    for line in lines.iter() {
        if line == &"" {
            groups.push(group);
            group = "".to_string();
        } else if group == "" {
            group = line.to_string();
        } else {
            group = [group, " ".to_string(), line.to_string()].concat();
        }
    }
    groups.push(group);

    // iterate thru groups, find unique chars in each group and add them
    let mut count = 0;
    for g in groups.iter() {
        let remove_spaces = String::from(&g.replace(" ",""));
        let mut vec_of_chars: Vec<char> = remove_spaces.chars().collect();
        vec_of_chars.sort();
        vec_of_chars.dedup();
        count+=vec_of_chars.len();
    }

    println!("Part 1 - sum of question counts: {}", count);

    // iterate thru groups, find chars in every member of a group
    count = 0;
    for g in groups.iter() {
        let members: Vec<&str> = g.trim().split(" ").collect();

        // if has only one member in a group
        if members.len() == 1 {
            let mut vec_of_chars: Vec<char> = g.chars().collect();
            vec_of_chars.sort();
            vec_of_chars.dedup();
            count+=vec_of_chars.len();
            continue;
        }

        // else has multiple members

        let mut sorted_members: Vec<Vec<char>> = vec![];
        for m in members.iter() {
            let mut vec_of_chars: Vec<char> = m.chars().collect();
            vec_of_chars.sort();
            vec_of_chars.dedup();
            sorted_members.push(vec_of_chars);
        }

        let mut hashmap: HashMap<char, u16> = HashMap::new();
        for m in sorted_members.iter() {
            for character in m.iter() {
                if hashmap.contains_key(character) {
                    hashmap.insert(character.clone(), hashmap[&character]+1);
                } else {
                    hashmap.insert(character.clone(), 1);
                }
            }
        }

        // iter thru the hashmap and find the letters that match the number of members
        let length: u16 = sorted_members.len() as u16;
        let mut questions_everyone_said_yes_to = 0;
        for (_k, v) in hashmap.iter() {
            if *v == length {
                questions_everyone_said_yes_to+=1;
            }
        }

        count+=questions_everyone_said_yes_to;
    }

    println!("Part 2 - sum of everyone counts: {}", count);
}
