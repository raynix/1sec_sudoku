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
	board  map[Pos]int
	ranks  []Pos
	given  []Pos
}

type Delta struct {
	pos Pos
	v int
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

func (self *Sudoku) get_row(n int, ds []Delta) []int {
	r := make([]int, 0)
	for x := 0; x < 9; x++ {
		v := self.board[Pos{x, n}]
		if v > 0 {
			r = append(r, v)
		}
	}
	for _, d := range ds {
		if d.pos.Y == n {
			r = append(r, d.v)
		}
	}
	return r
}

func (self *Sudoku) get_column(n int, ds []Delta) []int {
	r := make([]int, 0)
	for y := 0; y < 9; y++ {
		v := self.board[Pos{n, y}]
		if v > 0 {
			r = append(r, v)
		}
	}
	for _, d := range ds {
		if d.pos.X == n {
			r = append(r, d.v)
		}
	}
	return r
}

func (self *Sudoku) get_nine(n Pos, ds []Delta) []int {
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
			for _, d := range ds {
				if d.pos.X == x && d.pos.Y == y {
					r = append(r, d.v)
				}
			}
		}
	}
	return r
}

func unique_int(list []int) []int {
	r := make([]int, 0)
	for _, i := range list {
		if int_in_list(i, r) {
			continue
		}
		r = append(r, i)
	}
	return r
}

func (self *Sudoku) assess_order() {
	self.ranks = make([]Pos, 0)
	banned := make(map[Pos]int)
	ds := make([]Delta, 0)
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			p := Pos{x, y}
			if in_list(p, self.given) {
				continue
			}
			c := make([]int, 0)
			c = append(self.get_row(y, ds), self.get_column(x, ds)...)
			c = append(c, self.get_nine(p, ds)...)
			banned[p] = len(unique_int(c))
		}
	}
	for o := 8; o >= 0; o-- {
		for k, v := range banned {
			if o == v {
				self.ranks = append(self.ranks, k)
			}
		}
	}
}

func (self *Sudoku) try_step(ds []Delta) bool {
	step := len(ds)
	if step == len(self.ranks) {
		self.print_board()
		return true
	}
	p := self.ranks[step]

	used_numbers := append(self.get_row(p.Y, ds), self.get_column(p.X, ds)...)
	used_numbers = append(used_numbers, self.get_nine(p, ds)...)
	for g := 1; g <= 9; g++ {
		if int_in_list(g, used_numbers) {
			continue
		}
		new_ds := append(ds, Delta{p, g})
		self.board[p] = g
		if self.try_step(new_ds) {
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
	bd.assess_order()
	bd.try_step([]Delta{})
}
