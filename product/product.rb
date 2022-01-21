def dfs(pos, product)
  if pos === @n
    @ans += 1 if product === @x
    return
  end
  @graph[pos].each do |g|
    next if product * g > @x
    dfs(pos+1, product*g)
  end
end

@n, @x = gets.chomp.split(" ").map(&:to_i)
@ans = 0
@graph = readlines.map do |l|
  l.split(" ").map(&:to_i)[1..]
end

dfs(0, 1)
p @ans
