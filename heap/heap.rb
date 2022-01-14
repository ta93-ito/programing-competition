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

  def top
    heap.first
  end

  # 最大値を削除する
  def pop
    return if heap.empty?

    # 先頭の要素を削除し末尾要素を先頭に移動させる。
    heap.shift
    last_child = heap.pop
    heap.unshift(last_child)

    # 値の大きい子とlast_childを比較し、last_childが子よりも小さくなるまで移動していく。
    # 交換の際は子を上にスライドしていき、交換しないタイミングでlast_childをその位置に挿入する。
    i = 0
    while 2 * i + 1 < (heap.size - 1)
      larger_index = larger_child_index(i)
      break if heap[larger_index] <= last_child

      heap[i] = heap[larger_index]
      i = larger_index
    end
    heap[i] = last_child
  end

  def larger_child_index(i)
    if heap[2 * i + 1] > heap[2 * i + 2]
      2 * i + 1
    else
      2 * i + 2
    end
  end
end

heap = Heap.new
heap.push(1)
heap.push(2)
heap.push(3)
heap.push(5)
heap.push(2)
heap.push(10)

p "before: #{heap.heap}"
heap.pop
p "after: #{heap.heap}"
