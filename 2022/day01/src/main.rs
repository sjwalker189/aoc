use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Failed to read puzzle input");

    println!("--- Day 01: Calorie Counting ---");
    println!("Part one: {}", part_one(contents.clone()));
    println!("Part two: {}", part_two(contents.clone()));
}

fn part_one(input: String) -> usize {
    let sum = input.split("\n\n").map(sum_calories).max().unwrap_or(0);
    sum
}

fn part_two(_input: String) -> usize {
    0
}

fn sum_calories(lines: &str) -> usize {
    lines
        .split("\n")
        .map(|line| line.parse::<usize>().unwrap_or(0))
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;
    const SAMPLE: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";

    #[test]
    fn test_part_one() {
        assert_eq!(part_one(SAMPLE.to_string()), 24000);
    }

    #[test]
    fn test_part_two() {
        assert_eq!(part_two(SAMPLE.to_string()), 1);
    }
}
