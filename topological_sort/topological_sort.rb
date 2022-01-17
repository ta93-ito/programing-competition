def dfs(v)
  @seen[v] = true

  @graph[v].each do |nv|
    next if @seen[nv]

    dfs(nv)
  end
  @order.push(v)
end

@order = []

v, e = gets.split(' ').map(&:to_i)
@seen = v.times.map { false }
@graph = v.times.map { [] }

readlines.each do |l|
  s, e = l.split(' ').map(&:to_i)
  @graph[e].push(s)
end

v.times do |t|
  next if @seen[t]

  dfs(t)
end
@order.each {|o| p o}
