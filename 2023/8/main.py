import re
from pprint import pprint

def parse_instructions(t: str):
    return list(t.strip())


def next_hop(network, place, instruction):
    return network[place][instruction]


def part1(fname: str):
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
                bits = re.findall('^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$', line.strip())
                network[bits[0][0]] = {'L':bits[0][1], 'R':bits[0][2]}
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


if __name__ == '__main__':
    part1('test1.txt')
    part1('test2.txt')
    part1('input1.txt')
