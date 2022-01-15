# n=4,a=(3,2,6,5),X=14
# xになるaの組み合わせ
def partial_sum(i, x, a)
  if i === 0
    return true if x === 0

    return false
  end
  # a[i-1]を選ばない場合
  return true if partial_sum(i - 1, x, a)
  # a[i-1]を選ぶ場合
  return true if partial_sum(i - 1, x - a[i - 1], a)

  false
end

n = gets.to_i
a = gets.chomp.split(' ').map(&:to_i)
x = gets.to_i
p partial_sum(n, x, a)
