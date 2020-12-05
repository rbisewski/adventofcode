use std::fs;

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let lines: Vec<&str> = contents.trim().split("\n").collect();

    let max_x = 31;
    // max_y is 323, as the map is 323 lines long

    let mut trees = 0;
    let mut x = 0;
    let mut y = 0;

    println!("Part 1");

    loop {
        x = (x + 3) % max_x;
        y+=1;

        if y > 322 {
            break;
        }

        let a: char = lines[y].as_bytes()[x as usize] as char;

        if a == '#' {
            trees+=1;
        }
    }

    println!("Number of trees encountered (Right 3, down 1) is: {}", trees);

    println!("Part 2");

    trees = 0;
    x = 0;
    y = 0;

    loop {
        x = (x + 1) % max_x;
        y+=1;

        if y > 322 {
            break;
        }

        let a: char = lines[y].as_bytes()[x as usize] as char;

        if a == '#' {
            trees+=1;
        }
    }

    println!("Number of trees encountered (Right 1, down 1) is: {}", trees);

    trees = 0;
    x = 0;
    y = 0;

    loop {
        x = (x + 5) % max_x;
        y+=1;

        if y > 322 {
            break;
        }

        let a: char = lines[y].as_bytes()[x as usize] as char;

        if a == '#' {
            trees+=1;
        }
    }

    println!("Number of trees encountered (Right 5, down 1) is: {}", trees);

    trees = 0;
    x = 0;
    y = 0;

    loop {
        x = (x + 7) % max_x;
        y+=1;

        if y > 322 {
            break;
        }

        let a: char = lines[y].as_bytes()[x as usize] as char;

        if a == '#' {
            trees+=1;
        }
    }

    println!("Number of trees encountered (Right 7, down 1) is: {}", trees);

    trees = 0;
    x = 0;
    y = 0;

    loop {
        x = (x + 1) % max_x;
        y+=2;

        if y > 322 {
            break;
        }

        let a: char = lines[y].as_bytes()[x as usize] as char;

        if a == '#' {
            trees+=1;
        }
    }

    println!("Number of trees encountered (Right 1, down 2) is: {}", trees);
}
