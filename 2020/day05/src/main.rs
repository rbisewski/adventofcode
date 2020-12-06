use std::fs;

#[derive(Debug)]
struct BoardingPass {
    row: u16,
    col: u16,
    id:  u16,
    raw: String,
}

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let lines: Vec<&str> = contents.trim().split("\n").collect();

    let mut boarding_passes: Vec<BoardingPass> = vec![];

    let mut highest_value = 0;

    for line in lines.iter() {
        let row_chars = &line[0..7];
        let col_chars = &line[7..10];

        let row_binary = String::from(&row_chars.replace("B","1").replace("F","0"));
        let col_binary = String::from(&col_chars.replace("R","1").replace("L","0"));

        let mut new_boarding_pass = BoardingPass {
            row: isize::from_str_radix(&row_binary, 2).unwrap() as u16,
            col: isize::from_str_radix(&col_binary, 2).unwrap() as u16,
            id:  0,
            raw: line.to_string(),
        };
        new_boarding_pass.id = (new_boarding_pass.row * 8) + new_boarding_pass.col;

        if new_boarding_pass.id > highest_value {
            highest_value = new_boarding_pass.id;
        }

        boarding_passes.push(new_boarding_pass);
    }

    println!("Part 1 - highest seat ID: {}", highest_value);

    for i in 68..965 {
        let mut is_found: bool = false;
        for boarding_pass in boarding_passes.iter() {
            if boarding_pass.id == i {
                is_found = true;
            }
        }

        if !is_found {
            println!("Part 2 - your seat ID: {}", i);
        }
    }
}
