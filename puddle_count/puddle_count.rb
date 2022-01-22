# 水溜りの塊をカウントする。
# 4 6
# #..#..
# ...#..
# #..#..
# #..#..

def dfs(h, w)
  @seen[h][w] = true

  (-1..1).each do |i|
    (-1..1).each do |j|
      next if @seen[h + i]&.[](w + j)
      next if h + i < 0 || h + i > @h - 1 || w + i < 0 || w + i > @w - 1
      next if @graph[h + i][w + j] != '#'

      dfs(h + i, w + j)
    end
  end
end

@h, @w = gets.split(' ').map(&:to_i)
@graph = @h.times.map { @w.times.map { [] } }
readlines.each_with_index do |l, i|
  @graph[i] = l.chomp.split('')
end
@seen = @h.times.map { @w.times.map { false } }
@ans = 0

@graph.each_with_index do |h, i|
  h.each_with_index do |cell, j|
    next if @seen[i][j]

    if cell == '#'
      dfs(i, j)
      @ans += 1
    end
  end
end

puts @ans
