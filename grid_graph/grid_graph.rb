@h, @w = gets.split(' ').map(&:to_i)

@graph = Array.new(@h, Array.new(@w))
@seen = Array.new(@h)
@h.times do |t|
  @seen[t] = Array.new(@w, false)
end

readlines.map.with_index do |l, i|
  @graph[i] = l.chomp.split('')
end

@dx = [1, 0, -1, 0]
@dy = [0, 1, 0, -1]

def dfs(h, w)
  @seen[h][w] = true
  4.times do |t|
    dh = h + @dy[t]
    dw = w + @dx[t]

    next if dh < 0 || dh > @h - 1 || dw < 0 || dw > @w - 1
    next if @graph[dh][dw] == '#'
    next if @seen[dh][dw]

    dfs(dh, dw)
  end
end

sh = 0
sw = 0
gh = 0
gw = 0

@graph.each_with_index do |h, hi|
  h.each_with_index do |w, wi|
    if w == 's'
      sh = hi
      sw = wi
    end
    if w == 'g'
      gh = hi
      gw = wi
    end
  end
end

dfs(sh, sw)

if @seen[gh][gw]
  puts 'Yes'
else
  puts 'No'
end
