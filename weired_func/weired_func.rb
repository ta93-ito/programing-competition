t = gets.to_i
def f(t)
  t * t + 2 * t + 3
end

p f(f(f(t) + t) + f(f(t)))
