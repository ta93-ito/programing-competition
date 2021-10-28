def count_score(a, b)
  c = (a - b).abs
  if c <= 5
    0
  elsif c <= 10
    1
  elsif c <= 20
    2
  elsif c <= 30
    3
  else
    5
  end
end

n, l = gets.split(' ').map(&:to_i)
hzl = []
l.times do
  hzl.push(gets.to_i)
end

sl = readlines.map(&:to_i)
results = []

n.times do
  tmp = []
  (0..l - 1).each do |i|
    sc = count_score(sl.shift, hzl[i])
    tmp.push(sc)
  end
  results.push(tmp.sum)
end

ans = 100 - results.min
if ans.negative?
  puts 0
else
  puts ans
end
