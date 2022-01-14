# graph: [][]int, v: int
def dfs(graph, v, seen)
  seen[v] = true
  graph[v].each do |next_v|
    next if seen[next_v]

    dfs(graph, next_v, seen)
  end
end

seen = []
graph = []

graph[0] = [1, 2]
graph[1] = [2, 3]
graph[2] = [3, 4]
graph[3] = [4, 5]
graph[4] = [5, 6]

n = 5
n.times do
  seen.push false
  graph.push []
end

dfs(graph, 1, seen)
if seen[0]
  p true
else
  p false
end
