class Heap
  def initialize(heap: [])
    @heap = heap
  end

  attr_reader :heap

  def push(vertex)
    heap.push(vertex)
    i = heap.size - 1

    # 親 > 子の条件を満たすまで入れ替えを行う。
    while i > 0
      pi = (i - 1) / 2
      break if heap[pi] >= vertex

      heap[i] = heap[pi]
      i = pi
    end
    heap[i] = vertex
  end
end

heap = Heap.new
heap.push(1)
heap.push(2)
heap.push(3)
heap.push(5)
heap.push(2)
heap.push(10)
p heap.heap
