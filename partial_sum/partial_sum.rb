# n=4,a=(3,2,6,5),X=14
# xになるaの組み合わせ
def partial_sum(i, x, a, dp)
  if i === 0
    return true if x === 0

    return false
  end
  return false if x < 0
  # a[i-1]を選ばない場合
  return dp[i][x] = true if partial_sum(i - 1, x, a, dp)
  # a[i-1]を選ぶ場合
  return dp[i][x] = true if partial_sum(i - 1, x - a[i - 1], a, dp)

  dp[i][x] = false
end

n = gets.to_i
a = gets.chomp.split(' ').map(&:to_i)
x = gets.to_i

dp = Array.new(n)
(n + 1).times do |i|
  dp[i] = Array.new(x + 1, false)
end

p partial_sum(n, x, a, dp)
