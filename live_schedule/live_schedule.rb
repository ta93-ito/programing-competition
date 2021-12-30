a = {}
b = {}
gets.to_i.times do
  a[gets.to_i] = true
end

gets.to_i.times do
  b[gets.to_i] = true
end

prev = 'B'

(1..31).each do |i|
  result = ''
  if a[i] && b[i]
    result = (prev == 'A') ? 'B' : 'A'
    prev = result
  elsif a[i]
    result = 'A'
  elsif b[i]
    result = 'B'
  else
    result = 'x'
    puts result
    next
  end
  puts result
end
