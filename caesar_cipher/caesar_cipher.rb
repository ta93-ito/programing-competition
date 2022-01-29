i = 1
str = 'a'
convert_table = {}
26.times do
  convert_table.merge!({ str => i })
  str = str.succ
  i += 1
end
base = gets.chomp.split('').map { |r| convert_table[r] }
target = gets.chomp.split('').map { |r| convert_table[r] }
result = false
(0..26).each do |i|
  dup_base = base.dup
  dup_base.each_with_index do |_, bi|
    dup_base[bi] = if dup_base[bi] + i > 26
                     (dup_base[bi] + i) % 26
                   else
                     dup_base[bi] + i
                   end
  end
  break if result = dup_base === target
end

puts result ? 'Yes' : 'No'
