use std::env;
use std::fs;

struct Pos {
    x: usize,
    y: usize,
}

struct Sudoku {
    board: [[u8; 9]; 9],
}

impl Sudoku {
    fn new() -> Self {
        Sudoku { board: [[0; 9]; 9] }
    }

    fn from_file(path: &str) -> Result<Self, std::io::Error> {
        let content = fs::read_to_string(path)?;
        let mut sudoku = Sudoku::new();

        for (i, line) in content.lines().enumerate() {
            if i >= 9 {
                break;
            }

            for (j, num) in line.split_whitespace().enumerate() {
                if j >= 9 {
                    break;
                }

                sudoku.board[i][j] = num.parse().unwrap_or(0);
            }
        }

        Ok(sudoku)
    }

    fn print(&self) {
        for row in &self.board {
            for &cell in row {
                print!("{} ", cell);
            }
            println!();
        }
    }

    fn is_valid_move(&self, row: usize, col: usize, num: u8) -> bool {
        // Check row
        for j in 0..9 {
            if self.board[row][j] == num {
                return false;
            }
        }

        // Check column
        for i in 0..9 {
            if self.board[i][col] == num {
                return false;
            }
        }

        // Check 3x3 box
        let box_start_row = (row / 3) * 3;
        let box_start_col = (col / 3) * 3;

        for i in 0..3 {
            for j in 0..3 {
                if self.board[box_start_row + i][box_start_col + j] == num {
                    return false;
                }
            }
        }

        true
    }

    fn solve(&mut self) -> bool {
        match self.find_empty_cell() {
            None => true, // Solved
            Some(Pos { x, y }) => {
                for num in 1..=9 {
                    if self.is_valid_move(x, y, num) {
                        self.board[x][y] = num;

                        if self.solve() {
                            return true;
                        }

                        self.board[x][y] = 0; // Backtrack
                    }
                }
                false
            }
        }
    }

    fn find_empty_cell(&self) -> Option<Pos> {
        for i in 0..9 {
            for j in 0..9 {
                if self.board[i][j] == 0 {
                    return Some(Pos { x: i, y: j });
                }
            }
        }
        None
    }
}

fn main() -> Result<(), std::io::Error> {
    let args: Vec<String> = env::args().collect();
    let puzzle_path = args
        .get(1)
        .unwrap_or(&String::from("../puzzle2"))
        .to_string();

    let mut sudoku = Sudoku::from_file(&puzzle_path)?;

    println!("Original puzzle:");
    sudoku.print();

    if sudoku.solve() {
        println!("\nSolution:");
        sudoku.print();
    } else {
        println!("\nNo solution exists!");
    }

    Ok(())
}
