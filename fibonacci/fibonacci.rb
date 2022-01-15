def fibo(n)
  return 1 if n === 1
  return 0 if n === 0

  fibo(n - 1) + fibo(n - 2)
end

last = gets.to_i
(0..last).each do |n|
  p "#{n}項目目の値:#{fibo(n)}"
end
