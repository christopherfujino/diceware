#!/usr/bin/env ruby

writeFd = File.new('diceware/wordlist.go', 'w')
readFd = File.open 'wordlist.txt'
writeFd.write "package diceware\n\nvar wordList = [...]string {\n"
readFd.readlines.each do |line|
  tuple = line.strip.split "\t"
  word = tuple[1]
  writeFd.write "\t\"#{word}\",\n"
end
writeFd.write "}\n"
writeFd.close
