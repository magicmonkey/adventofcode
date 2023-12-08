import re
from pprint import pprint
import math

def parse_instructions(t: str):
    return list(t.strip())


def next_hop(network, place, instruction):
    return network[place][instruction]


def parse(fname: str):
    line1 = True
    instr = []
    network = {}
    with open(fname) as file:
        for line in file:
            if line1:
                line1 = False
                instr = parse_instructions(line)
            elif line.strip() == "":
                continue
            else:
                bits = re.findall('^([\dA-Z]{3}) = \(([\dA-Z]{3}), ([\dA-Z]{3})\)$', line.strip())
                network[bits[0][0]] = {'L':bits[0][1], 'R':bits[0][2]}
    return (network, instr)


def part1(fname: str):
    network, instr = parse(fname)
    pos = 'AAA'
    counter = 0
    #pprint(network)
    #pprint(instr)
    while True:
        #print(f"Assessing instruction {counter} ({counter%len(instr)}) from position {pos}")
        pos = next_hop(network, pos, instr[counter%len(instr)])
        counter += 1
        #print(f"... next pos is {pos}")
        if pos == 'ZZZ':
            break
    print(counter)


def part2(fname: str, start: str):
    network, instr = parse(fname)
    # Find a starting point
    pos = start
    counter = 0
    while True:
        #print(f"Assessing instruction {counter} ({counter%len(instr)}) from position {pos}")
        pos = next_hop(network, pos, instr[counter%len(instr)])
        counter += 1
        #print(f"... next pos is {pos}")
        if pos[2] == 'Z':
            return counter


if __name__ == '__main__':
    #part1('test1.txt')
    #part1('test2.txt')
    #part1('input.txt')
    #part2('test3.txt')
    nums = [part2('input.txt', 'LCA'), part2('input.txt', 'NVA'), part2('input.txt', 'GCA'), part2('input.txt', 'SXA'), part2('input.txt', 'AAA'), part2('input.txt', 'GMA')]
    print(math.lcm(*nums))

