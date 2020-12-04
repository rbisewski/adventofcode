use std::fs;
use std::collections::HashMap;

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let list_of_numbers: Vec<u16> = contents.trim()
                                            .split("\n")
                                            .map(|x| x.parse::<u16>().unwrap())
                                            .collect();

    let mut hashmap = HashMap::new();

    for i in 0..list_of_numbers.len()-1 {
        hashmap.insert(list_of_numbers[i], i);
    }

    // prints two elements that sum to 2020
    println!("Two elements that sum to 2020:");
    for i in 0..list_of_numbers.len()-1 {
        let key = 2020-list_of_numbers[i];
        if hashmap.contains_key(&key) {
            if hashmap[&key] != i {
                println!("Index: {}, Element: {}", i, key);
            }
        }
    }

    // prints three elements that sum to 2020
    println!("\nThree elements that sum to 2020:");
    for i in 0..list_of_numbers.len()-2 {
        for j in i+1..list_of_numbers.len()-1 {
            for k in j+1..list_of_numbers.len() {
                if list_of_numbers[i] + list_of_numbers[j] + list_of_numbers[k] == 2020 {
                    println!("{}, {}, {}", list_of_numbers[i], list_of_numbers[j], list_of_numbers[k]);
                }
            }
        }
    }
}
