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
	banned map[Pos]int
	ranks  []Pos
	given  []Pos
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
	self.banned = make(map[Pos]int)
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			p := Pos{x, y}
			if in_list(p, self.given) {
				continue
			}
			c := make([]int, 0)
			c = append(self.get_row(y), self.get_column(x)...)
			c = append(c, self.get_nine(p)...)
			self.banned[p] = len(unique_int(c))
		}
	}
	for o := 8; o >= 0; o-- {
		for k, v := range self.banned {
			if o == v {
				self.ranks = append(self.ranks, k)
			}
		}
	}
}

func copy_board(prev_board Sudoku) (new_board Sudoku) {
	new_board = Sudoku{}
	new_board.board = make(map[Pos]int)
	for k, v := range prev_board.board {
		new_board.board[k] = v
	}
	copy(new_board.ranks, prev_board.ranks)
	copy(new_board.given, prev_board.given)
	return
}

func (board Sudoku) try_step(step int, done chan bool) {
	if step == len(board.ranks) {
		board.print_board()
		done <- true
		return
	}
	p := board.ranks[step]

	used_numbers := append(board.get_row(p.Y), board.get_column(p.X)...)
	used_numbers = append(used_numbers, board.get_nine(p)...)
	sub_done := make(chan bool, 9)
	for g := 1; g <= 9; g++ {
		if int_in_list(g, used_numbers) {
			sub_done <- false
			continue
		}
		board.board[p] = g
		sub_board := copy_board(board)
		sub_board.print_board()
		go sub_board.try_step(step+1, sub_done)
	}
	super_done := false
	for s := 1; s <= 9; s++ {
		found := <-sub_done
		if found {
			super_done = true
		}
	}
	if super_done {
		done <- true
	} else {
		done <- false
	}
	board.board[p] = 0
	return
}

func main() {
	puzzle := "puzzle2"
	if len(os.Args) > 1 {
		puzzle = os.Args[1]
	}
	top_done := make(chan bool, 1)
	bd := Sudoku{}
	bd.read_puzzle(puzzle)
	bd.print_board()
	bd.assess_order()
	bd.try_step(0, top_done)
	<-top_done
}
