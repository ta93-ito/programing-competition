l, r = gets.split(' ').map(&:to_i)
s = gets.strip
s[l - 1..r - 1] = s[l - 1..r - 1].reverse
puts s
