package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X, Y int
}

type Sudoku struct {
	board map[Pos]int
	given []Pos
}

func (self *Sudoku) print_board() {
	fmt.Println("")
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			v := self.board[Pos{x, y}]
			if v > 0 {
				fmt.Print(self.board[Pos{x, y}])
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func (self *Sudoku) read_puzzle(puzzle string) {
	fmt.Println("Reading puzzle: ", puzzle)
	file, err := os.Open(puzzle)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	self.board = make(map[Pos]int)
	self.given = make([]Pos, 0)
	for scanner.Scan() {
		row_str := strings.Split(scanner.Text(), " ")
		v := 0
		for i := 0; i < len(row_str); i++ {
			v, _ = strconv.Atoi(row_str[i])
			p := Pos{i, counter}
			self.board[p] = v
			if v > 0 {
				self.given = append(self.given, p)
			}
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func in_list(p Pos, list []Pos) bool {
	for _, b := range list {
		if p == b {
			return true
		}
	}
	return false
}

func int_in_list(p int, list []int) bool {
	for _, b := range list {
		if p == b {
			return true
		}
	}
	return false
}

func filter(c []int, f func(int) bool) []int {
	r := make([]int, 0)
	for _, v := range c {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func gt_zero(a int) bool {
	if a > 0 {
		return true
	} else {
		return false
	}
}

func (self *Sudoku) get_row(n int) []int {
	r := make([]int, 0)
	for x := 0; x < 9; x++ {
		v := self.board[Pos{x, n}]
		if v > 0 {
			r = append(r, v)
		}
	}
	return r
}

func (self *Sudoku) get_column(n int) []int {
	r := make([]int, 0)
	for y := 0; y < 9; y++ {
		v := self.board[Pos{n, y}]
		if v > 0 {
			r = append(r, v)
		}
	}
	return r
}

func (self *Sudoku) get_nine(n Pos) []int {
	nine_grids := [][]int{
		{0, 1, 2}, {0, 1, 2}, {0, 1, 2},
		{3, 4, 5}, {3, 4, 5}, {3, 4, 5},
		{6, 7, 8}, {6, 7, 8}, {6, 7, 8},
	}
	r := make([]int, 0)
	for _, y := range nine_grids[n.Y] {
		for _, x := range nine_grids[n.X] {
			v := self.board[Pos{x, y}]
			if v > 0 {
				r = append(r, v)
			}
		}
	}
	return r
}

func (self *Sudoku) try_step(step int) bool {
	if step == 81 {
		self.print_board()
		return true
	}
	p := Pos{step % 9, step / 9}
	if in_list(p, self.given) {
		return self.try_step(step + 1)
	}

	used_numbers := append(self.get_row(p.Y), self.get_column(p.X)...)
	used_numbers = append(used_numbers, self.get_nine(p)...)
	for g := 1; g <= 9; g++ {
		if int_in_list(g, used_numbers) {
			continue
		}
		self.board[p] = g
		if self.try_step(step + 1) {
			return true
		}
	}
	self.board[p] = 0
	return false
}

func main() {
	puzzle := "puzzle2"
	if len(os.Args) > 1 {
		puzzle = os.Args[1]
	}
	bd := Sudoku{}
	bd.read_puzzle(puzzle)
	bd.print_board()
	bd.try_step(0)
}
