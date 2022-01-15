def fibo(n, memo)
  return 1 if n === 1
  return 0 if n === 0
  return memo[n] unless memo[n] === -1

  memo[n] = fibo(n - 1, memo) + fibo(n - 2, memo)
end

last = gets.to_i
memo = []
(last + 1).times do
  memo.push(-1)
end

p fibo(last, memo)
