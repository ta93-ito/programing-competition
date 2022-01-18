n = gets.to_i
plats = gets.split(" ").map(&:to_i)

before = plats[0]
plats[1..].each do |pl|
  if pl > before
    before = pl
  else
    break
  end
end

p before
