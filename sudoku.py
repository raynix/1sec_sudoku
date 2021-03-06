import sys

def non_zero(x):
  return x > 0 and x

class Sudoku:
  def __init__(self, puzzle):
    self.board = {}
    self.known_keys = []
    self.sub_range = [[0,1,2], [3,4,5], [6,7,8]]
    self.all = range(1, 10)
    self.all_index = range(0, 9)
    with open(puzzle) as f:
      for row, line in enumerate( f.readlines() ):
        for col, number in enumerate( line.split() ):
          value = int(number)
          self.board[ (col, row) ] = value
          if value > 0:
            self.known_keys.append( (col, row) )
    self.show_results()


  def show_results(self):
    for row in range(0, 9):
      for col in range(0, 9):
        print( self.board[ (col, row) ] > 0 and self.board[ (col, row) ] or '.', end=' ' )
      print()

  def solve(self):
    print()
    print("Just a moment...")
    print()
    self.try_next()



  def get_row(self, r):
    row = [ self.board[ (x, y) ] for x in self.all_index for y in [r] ]
    return filter( non_zero, row )

  def get_col(self, c):
    col = [ self.board[ (x, y) ] for x in [c] for y in self.all_index ]
    return filter( non_zero, col)

  def get_nine(self, c, r):
    nine = [ self.board[ (x, y) ] for x in self.sub_range[ int(c / 3) ]  for y in self.sub_range[ int(r / 3) ] ]
    return filter( non_zero, nine)

  def try_next(self, step=0):
    if step == 81:
      self.show_results()
      return True
    x = int(step / 9)
    y = step % 9
    key = (x, y)

    if key in self.known_keys:
      return self.try_next( step + 1)

    guesses = set(self.all) - set( self.get_row(y) ) - set( self.get_col(x) ) - set( self.get_nine(x,y) )
    for g in guesses:
      self.board[ key ] = g
      if self.try_next( step + 1):
        return True
    self.board[ key ] = 0
    return False


if __name__ == "__main__":
  s = Sudoku(sys.argv[1])
  s.solve()
