# stdin -> n: number of steps, steps: step height separated by line breaks
def chmin(a, b)
  if a > b
    b
  else
    a
  end
end

MAX = 10_000_000_000
@n = gets.to_i
@dp = Array.new(@n, MAX)
@dp[0] = 0
@steps = gets.chomp.split(' ').map(&:to_i)

def rec_leap(i)
  return 0 if i === 0
  return @dp[i] if @dp[i] < MAX

  @dp[i] = chmin(@dp[i], rec_leap(i - 1) + (@steps[i] - @steps[i - 1]).abs)
  @dp[i] = chmin(@dp[i], rec_leap(i - 2) + (@steps[i] - @steps[i - 2]).abs) if i > 1
  @dp[i]
end

p rec_leap(@n - 1)
