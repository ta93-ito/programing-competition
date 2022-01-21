max = 0

n = gets.to_i
points = readlines.map do |l|
  l.split(' ').map(&:to_i)
end

def chmax(a,b)
  if a > b
    a
  else
    b
  end
end

def hypotenuse(a, b)
  Math.sqrt(a * a + b * b)
end

points.each_with_index do |e, i|
  points[(i+1..)].each do |f|
    a = (e[0] - f[0]).abs
    b = (e[1] - f[1]).abs
    max = chmax(hypotenuse(a, b), max)
  end
end
p max
