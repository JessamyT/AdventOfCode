# Advent of Code 2023 Day 1 Part 1
import re

def swap(dictionary, line):
  for x in dictionary.keys():
    if re.search(x, line):
      line = re.sub(x, dictionary[x], line)
  return line

# Returns a string number for first found written word
def find_first(line, mylist, dictionary):
  result = ""
  thenum = ""
  strindex = 700
  if re.search("[0-9]", line):
    strindex = re.search("[0-9]", line).start()
  for num in mylist:
    newstrindex = line.find(num)
    if -1 < newstrindex < strindex:
      strindex = newstrindex
      result = mylist.index(num)
      thenum = num
  line = re.sub(thenum, str(result), line)
  line = swap(dictionary, line)
  first = re.search("[0-9]", line).group()
  return first

def find_last(line, mylist, dictionary):
  result = ""
  thenum = ""
  strindex = 0
  for num in mylist:
    newstrindex = line.find(num)
    if newstrindex > strindex:
      strindex = newstrindex
      result = mylist.index(num)
      thenum = num
  line = re.sub(thenum, str(result), line)
  line = swap(dictionary, line)
  last = re.search("([0-9])[a-z]*$", line).group(1)
  return last

def main():
  file = open("./inputs23/day1input.txt", "r")
  total = 0
  dictionary = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
  }
  mylist = ["zero", "one", "two", "three", "four",
            "five", "six", "seven", "eight", "nine"]
  # i is a string line in the file
  for i in file:
    print(i)
    first = find_first(i, mylist, dictionary)
    last = find_last(i, mylist, dictionary)
    #print(i)
    print(first, last)
    total+=int(first + last)
  file.close()
  print("Total:", total)

main()