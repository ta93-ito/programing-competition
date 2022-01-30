n, a, b = gets.chomp.split.map(&:to_i)
graph = Array.new(n) { Array.new(n, '.') }
pp, q, r, s = gets.chomp.split(' ').map(&:to_i)
([1 - a, 1 - b].max..[n - a, n - b].min).each do |k|
  graph[a + k - 1][b + k - 1] = '#' if graph[a + k - 1]
end
([1 - a, b - n].max..[n - a, b - 1].min).each do |k|
  graph[a + k - 1][b - k - 1] = '#' if graph[a + k - 1]
end
(pp..q).each do |i|
  tmp = []
  (r..s).each do |j|
    tmp.push(graph[i - 1][j - 1]) if graph[i - 1][j - 1]
  end
  puts tmp.join('')
end
