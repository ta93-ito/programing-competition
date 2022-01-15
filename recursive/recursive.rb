def rec_add(n, i)
  p "#{i += 1}回目の呼び出し"
  return 0 if n === 0

  n + rec_add(n - 1, i)
end

i = 0
n = gets.to_i
sum = rec_add(n, i)
p "#{n}までの総和:#{sum}"
