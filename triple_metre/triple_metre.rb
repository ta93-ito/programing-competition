t = 'oxx' * 100
s = gets.chomp
if t =~ /#{s}/
  puts 'Yes'
else
  puts 'No'
end
