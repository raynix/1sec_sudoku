require 'colorize'

class Sudoku
  def initialize 
    @board = Hash.new
    @known_keys = []
    @sub_range = [[0,1,2], [3,4,5], [6,7,8]]
    File.open('sample') do |f|
      row = 0
      while line = f.gets
        line.split.each_with_index do |number, col|
          value = number.to_i
          @board[[col, row]] = value
          @known_keys << [col, row] if value > 0
          #print "#{value} "
        end
        #puts
        row += 1
      end
    end
    show_result
    puts
    puts "Just a moment..."
    puts
    #puts @known_keys.to_s
  end

  def check_row(x, y)
    row = (0..8).to_a.product([y]).collect{|key| @board[key]}
    #print row
    row.delete 0
    return row == row.uniq
  end

  def check_column(x, y)
    col = [x].product((0..8).to_a).collect{|key| @board[key]}
    col.delete 0
    return col == col.uniq
  end

  def check_sub(x, y)
    x_r = @sub_range[ x / 3 ]
    y_r = @sub_range[ y / 3 ]
    nine = x_r.product(y_r).collect{|key| @board[key] }
    nine.delete 0
    return nine == nine.uniq
  end

  def try(step=0) 
    #print "\n\n"
    #puts step

    if step == 81
      show_result
      return true
    end

    y = step / 9
    x = step % 9
    key = [x, y]
    return try( step + 1 ) if @known_keys.include? key

    while true
      @board[key] += 1
      if @board[key] == 10
        @board[key] = 0
        return false
      end
      return true if check_row(x,y) && check_column(x,y) && check_sub(x,y) && try(step + 1)
      
    end    
  end
  
  def show_result()
    #(0..8).to_a.product((0..8).to_a).collect{|key| @board[key]}.each
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

s = Sudoku.new
s.check_row 1,3
s.try
