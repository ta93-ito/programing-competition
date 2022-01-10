def main
  vertex_num, edge_num = gets.split(' ').map(&:to_i)

  graph = []
  readlines.each do |l|
    point, adjacent_point = l.split(' ').map(&:to_i)
    graph[point] = [] unless graph[point]
    graph[point].push(adjacent_point)
  end
  puts 'accepted graph is below'
  graph.each_with_index do |g, i|
    puts "#{i}:#{g}"
  end
end

main
