1sec_sudoku
===========

Solve a sudoku puzzle within 1 second.

Ruby Sample:

```
time ruby sudoku.mk3.rb puzzle2
```

Python:

```
time python2 sudoku.py puzzle2
```

Go:

```
go build src/sudoku.go
time ./sudoku puzzle2
```

The puzzle0 is very hard to backtrack. My Python solver tooks 6.5 minutes but the Go solver is much faster, around 28 seconds.
