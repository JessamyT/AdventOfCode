# Advent of Code 2023 Day 1 Part 1
import re

def main():
  file = open("./inputs23/day1input.txt", "r")
  total = 0
  for i in file:
    first = re.search("[0-9]", i).group()
    last = re.search("([0-9])[a-z]*$", i).group(1)
    total+=int(first + last)
  file.close()
  print("Total:", total)

main()