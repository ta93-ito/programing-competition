# 2 2 :w h
# 0 1 : 0 -> ocean, 1 -> land
# 1 0

DIRS = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]]
@w, @h = gets.chomp.split(' ').map(&:to_i)
@field = readlines.map do |l|
  l.split(' ').map(&:to_i)
end

def dfs(h, w)
  @field[h][w] = 0
  DIRS.each do |d|
    dh = h + d[0]
    dw = w + d[1]
    next if dh < 0 || dh > @h - 1 || dw < 0 || dw > @w - 1

    begin
      next if @field[dh][dw] == 0
    rescue StandardError
      binding.irb
    end

    dfs(dh, dw)
  end
end

def main
  count = 0
  @h.times do |i|
    @w.times do |j|
      next if @field[i][j] === 0

      dfs(i, j)
      count += 1
    end
  end
  p "number of islands is #{count}"
end

main
