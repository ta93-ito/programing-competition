n, q = gets.split(' ').map(&:to_i)
arr = gets.split(' ').map(&:to_i)
queries = readlines.map do |l|
  l.split(' ').map(&:to_i)
end

# [x][k]に値を詰める
tb = {}
arr.each do |e|
  tb[e] = []
end
arr.map.with_index do |e, i|
  tb[e].push(i)
end
queries.each do |q|
  if tb[q[0]]&.[] (q[1] - 1)
    p tb[q[0]][q[1] - 1] + 1
  else
    p(-1)
  end
end
