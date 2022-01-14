class Edge
  attr_reader :to, :weight

  def initialize(to, weight)
    @to = to
    @weight = weight
  end
end

graph = []
edge1 = Edge.new(1, 2)
edge2 = Edge.new(2, 1)
edge3 = Edge.new(3, 4)
graph.push([edge1])
graph.push([edge2])
graph[0].push(edge3)
p graph
