require 'colorize'

class Sudoku
  def initialize( puzzle ) 
    @board = Hash.new
    @known_keys = []
    @sub_range = [[0,1,2], [3,4,5], [6,7,8]]
    @all = (1..9).to_a.reverse
    File.open(puzzle) do |f|
      row = 0
      while line = f.gets
        line.split.each_with_index do |number, col|
          value = number.to_i
          @board[[col, row]] = value
          @known_keys << [col, row] if value > 0
        end
        row += 1
      end
    end
    show_result
    trap("USR1") do
      puts
      show_result
    end
  end

  def solve
    puts
    puts "Just a moment..."
    puts
    try
  end

  private

  def get_row( y )
    row = (0..8).to_a.product([y]).collect{|key| @board[key]}
  end

  def get_column( x )
    col = [x].product((0..8).to_a).collect{|key| @board[key]}
  end

  def get_nine( x, y )
    x_r = @sub_range[ x / 3 ]
    y_r = @sub_range[ y / 3 ]
    nine = x_r.product(y_r).collect{|key| @board[key] }
  end

  def try(step=0) 
    if step == 81
      show_result
      return true
    end

    y = step / 9
    x = step % 9
    key = [x, y]
    return try( step + 1 ) if @known_keys.include? key

    guesses = @all - get_row(y) - get_column(x)- get_nine(x, y)

    guesses.each do |g|
      @board[key] = g
      return true if try(step + 1)
    end      
    @board[key] = 0
    return false    
  end
  
  def show_result()
    for y in (0..8)
      for x in (0..8)
        if @known_keys.include? [x,y]
          print @board[[x,y]].to_s.red
        else
          print @board[[x,y]] > 0 ? @board[[x,y]] : '.'
        end
        print ' '
      end
      print "\n"
    end
  end
end

s = Sudoku.new ARGV[0]
s.solve
