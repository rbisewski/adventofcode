use std::fs;
use regex::Regex;

struct Entry {
    min: u16,
    max: u16,
    rule: char,
    password: String,
}

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let lines: Vec<&str> = contents.trim().split("\n").collect();

    let re = Regex::new(r"(\d+)-(\d+) ([a-z]): ([a-z]+)").unwrap();

    let mut valid_rules = 0;
    for line in lines.iter() {
        let caps = re.captures(line).unwrap();
        let new_entry = Entry {
            min:      caps.get(1).map_or("", |m| m.as_str()).parse().unwrap(),
            max:      caps.get(2).map_or("", |m| m.as_str()).parse().unwrap(),
            rule:     caps.get(3).map_or("", |m| m.as_str()).chars().next().unwrap(),
            password: caps.get(4).map_or("", |m| m.as_str()).to_string(),
        };

        let pieces: Vec<&str> = new_entry.password.split(new_entry.rule).collect();
        let count: u16 = (pieces.len()-1) as u16;
        if count >= new_entry.min && count <= new_entry.max {
            valid_rules+=1;
        }
    }
    println!("Day 2, part 1 - valid rules: {}", valid_rules);

    valid_rules = 0;
    for line in lines.iter() {
        let caps = re.captures(line).unwrap();
        let new_entry = Entry {
            min:      caps.get(1).map_or("", |m| m.as_str()).parse().unwrap(),
            max:      caps.get(2).map_or("", |m| m.as_str()).parse().unwrap(),
            rule:     caps.get(3).map_or("", |m| m.as_str()).chars().next().unwrap(),
            password: caps.get(4).map_or("", |m| m.as_str()).to_string(),
        };

        let a: char = new_entry.password.as_bytes()[(new_entry.min-1) as usize] as char;
        let b: char = new_entry.password.as_bytes()[(new_entry.max-1) as usize] as char;

        if a == new_entry.rule && b == new_entry.rule {
            // do nothing
        } else if a != new_entry.rule && b != new_entry.rule {
            // do nothing
        } else {
            valid_rules+=1;
        }
    }
    println!("Day 2, part 2 - valid rules: {}", valid_rules);
}
