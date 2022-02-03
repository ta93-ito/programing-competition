def main(lines)
  # 0が黒、1が白
  record = [0, 1]
  turns = lines[0].chomp.split('')
  next_color = 0
  turns.each do |t|
    unless record.find(next_color)
      if t == 'L'
        record.unshift(next_color)
      else
        record.push(next_color)
      end
      next_color = (next_color - 1).abs
      next
    end
    if t == 'L'
      record.each_with_index do |r, i|
        break if r === next_color

        record[i] = next_color
      end
      record.unshift(next_color)
    else
      record.reverse_each.with_index do |r, i|
        break if r === next_color

        record[i] = next_color
      end
      record.push(next_color)
    end
    next_color = (next_color - 1).abs
  end
  black_num = record.count(0)
  white_num = record.size - black_num
  puts "#{black_num} #{white_num}"
end

main(readlines)
