# stdin -> n: number of steps, steps: step height separated by line breaks
def chmin(a, b)
  if a > b
    b
  else
    a
  end
end

MAX = 10_000_000_000

n = gets.to_i
# dpテーブル。n番目の段差に辿り着くまでの最小コストを登録する。
dp = Array.new(n, MAX)
dp[0] = 0
steps = gets.chomp.split(' ').map(&:to_i)

(1...n).each do |i|
  dp[i] = chmin(dp[i], dp[i - 1] + (steps[i] - steps[i - 1]).abs)
  dp[i] = chmin(dp[i], dp[i - 2] + (steps[i] - steps[i - 2]).abs) if i > 1
end

p dp.last
