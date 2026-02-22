#!/usr/bin/env ruby

writeFd = File.new('wordlist.cpp', 'w')
readFd = File.open 'wordlist.txt'
writeFd.write "#include <string>\n#include <vector>\n\nstd::vector<std::string> words = {\n"
readFd.readlines.each do |line|
  tuple = line.strip.split "\t"
  word = tuple[1]
  writeFd.write "  \"#{word}\",\n"
end
writeFd.write "};\n"
writeFd.close
