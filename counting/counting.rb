n, q = gets.split(' ').map(&:to_i)
ss = gets.split(' ').map(&:to_i).sort
readlines.each do |l|
  i = ss.bsearch_index { |v| v >= l.to_i }
  unless i
    (puts 0
     next)
  end
  puts ss.length - i
end
