#!/usr/bin/env ruby

file = $stdin
lines = file.readlines
file.close

a,b = lines[0].split.map{|x| x.to_i}

puts a + b
