n = gets.to_i
totaling = {}
readlines.each do |l|
  vote = l.chomp
  totaling[vote] = 0 unless totaling[vote]
  totaling[vote] += 1
end

puts (totaling.max { |x, y| x[1] <=> y[1] })[0]
