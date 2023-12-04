import re
from pprint import pprint

def part1(fname: str):
    file = open(fname)
    lines = file.readlines()
    counter = 0
    for line in lines:
        num1 = re.findall(r'^[^\d]*(\d)', line)
        num2 = re.findall(r'(\d)[^\d]*$', line)
        num = int(f"{num1[0]}{num2[0]}")
        counter += num
    file.close()
    pprint(counter)

def part2(fname: str):
    file = open(fname)
    lines = file.readlines()
    counter = 0
    tr = {
            'one' : 1,
            'two' : 2,
            'three' : 3,
            'four' : 4,
            'five' : 5,
            'six' : 6,
            'seven' : 7,
            'eight' : 8,
            'nine' : 9,
    }
    words = '|'.join(tr.keys())
    for line in lines:
        num1 = re.findall(f"^.*?(\d|{words})", line)
        num2 = re.findall(f".*(\d|{words}).*?$", line)
        if num1[0] in tr:
            n1 = tr[num1[0]]
        else:
            n1 = num1[0]
        if num2[0] in tr:
            n2 = tr[num2[0]]
        else:
            n2 = num2[0]
        num = int(f"{n1}{n2}")
        counter += num
    file.close()
    pprint(counter)

if __name__ == '__main__':
    part1('test1.txt')
    part1('input.txt')
    part2('test2.txt')
    part2('input.txt')

