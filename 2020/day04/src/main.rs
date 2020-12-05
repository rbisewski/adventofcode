use std::fs;
use regex::Regex;

#[derive(Debug)]
struct Passport {
    byr: String,
    iyr: String,
    eyr: String,
    hgt: String,
    hcl: String,
    ecl: String,
    pid: String,
    cid: String,
    raw: String,
}

fn main() {
    let contents = fs::read_to_string("input1.txt")
        .expect("Something went wrong reading the file");

    let lines: Vec<&str> = contents.trim().split("\n").collect();

    let hair_colour_regex = Regex::new(r"^#[0-9a-f]{6}$").unwrap();

    let height_inches_regex = Regex::new(r"^5[7-9]in|6[0-9]in|7[0-6]in$").unwrap();
    let height_metres_regex = Regex::new(r"^15[0-9]cm|16[0-9]cm|17[0-9]cm|18[0-9]cm|19[0-3]cm$").unwrap();

    let mut valid_part1 = 0;
    let mut valid_part2 = 0;
    let mut passports: Vec<Passport> = vec![];

    // sort raw data
    let mut raw: String = String::from("");
    for line in lines.iter() {
        if line == &"" {
            let new_passport = Passport {
                byr: "".to_string(),
                iyr: "".to_string(),
                eyr: "".to_string(),
                hgt: "".to_string(),
                hcl: "".to_string(),
                ecl: "".to_string(),
                pid: "".to_string(),
                cid: "".to_string(),
                raw: raw.to_string(),
            };
            passports.push(new_passport);
            raw = "".to_string();
        } else if raw == "" {
            raw = line.to_string();
        } else {
            raw = [raw, " ".to_string(), line.to_string()].concat();
        }
    }
    
    let mut byr: String;
    let mut iyr: String;
    let mut eyr: String;
    let mut hgt: String;
    let mut hcl: String;
    let mut ecl: String;
    let mut pid: String;
    let mut cid: String;

    // convert raw data into Passport struct
    for passport in passports.iter_mut() {
        byr = String::from("");
        iyr = String::from("");
        eyr = String::from("");
        hgt = String::from("");
        hcl = String::from("");
        ecl = String::from("");
        pid = String::from("");
        cid = String::from("");

        let elements: Vec<&str> = passport.raw.trim().split(" ").collect();
        for element in elements.iter() {
            let key_and_value: Vec<&str> = element.trim().split(":").collect();
            match key_and_value[0] {
                "byr" => byr = key_and_value[1].to_string(),
                "iyr" => iyr = key_and_value[1].to_string(),
                "eyr" => eyr = key_and_value[1].to_string(),
                "hgt" => hgt = key_and_value[1].to_string(),
                "hcl" => hcl = key_and_value[1].to_string(),
                "ecl" => ecl = key_and_value[1].to_string(),
                "pid" => pid = key_and_value[1].to_string(),
                "cid" => cid = key_and_value[1].to_string(),
                _ => (),
            }
        }

        *passport = Passport {
            byr: byr,
            iyr: iyr,
            eyr: eyr,
            hgt: hgt,
            hcl: hcl,
            ecl: ecl,
            pid: pid,
            cid: cid,
            raw: passport.raw.clone(),
        };

        if passport.byr == "" || passport.iyr == "" || passport.eyr == "" || passport.hgt == "" || passport.hcl == "" || passport.ecl == "" || passport.pid == "" {
            continue;
        } 

        valid_part1+=1;

        // validation rules for part 2

        let byr_int: u16 = passport.byr.parse().unwrap();
        if passport.byr.len() != 4 || byr_int < 1920 || byr_int > 2002 {
            continue;
        }

        let iyr_int: u16 = passport.iyr.parse().unwrap();
        if passport.iyr.len() != 4 || iyr_int < 2010 || iyr_int > 2020 {
            continue;
        }

        let eyr_int: u16 = passport.eyr.parse().unwrap();
        if passport.eyr.len() != 4 || eyr_int < 2020 || eyr_int > 2030 {
            continue;
        }

        if height_inches_regex.is_match(&passport.hgt) {
            // do nothing as the height is valid
        } else if height_metres_regex.is_match(&passport.hgt) {
            // do nothing as the height is valid
        } else {
            continue;
        }

        if hair_colour_regex.is_match(&passport.hcl) {
            // do nothing as the hair colour is valid
        } else {
            continue;
        }

        if passport.ecl == "amb" || passport.ecl == "blu" || passport.ecl == "brn" || passport.ecl == "gry" || passport.ecl == "grn" || passport.ecl == "hzl" || passport.ecl == "oth" {
            // do nothing as the eye colour is valid
        } else {
            continue;
        }

        if passport.pid.len() != 9 {
            continue;
        }
        let _: u32 = match passport.pid.parse() {
            Ok(n) => n,
            Err(_) => continue,
        };

        valid_part2+=1;
    }

    println!("Part 1 - Valid Passports: {}", valid_part1);
    println!("Part 2 - Valid Passports: {}", valid_part2);
}
