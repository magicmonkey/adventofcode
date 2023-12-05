import re
from pprint import pprint
import sys


class Mapping:
    
    def __init__(self, dest):
        self.dest = dest
        self.ranges = []

    def add_range(self, src_start, dst_start, size):
        self.ranges.append({'src':src_start, 'dst':dst_start, 'size':size})

    def find_next(self, num) -> int:
        for r in self.ranges:
            if num >= r['src'] and num <= r['src']+r['size']:
                offset = num - r['src']
                return r['dst'] + offset
        return num

    def __str__(self):
        return f'Mapping with {len(self.ranges)} ranges'


def parse(fname: str):
    re_seeds = re.compile('^seeds: ([\d ]+)$')
    re_category = re.compile('^([a-z]+)-to-([a-z]+) map:$')
    re_range    = re.compile('^(\d+) (\d+) (\d+)$')
    seeds = []
    mappings  = {}
    curr_mapping = None
    with open(fname) as file:
        for line in file:
            row = line.strip()
            if row == "":
                continue
            elif re_seeds.match(row):
                bits = re_seeds.findall(row)
                for seed in bits[0].split(" "):
                    seeds.append(int(seed))
            elif re_category.match(row):
                bits = re.findall(re_category, row)
                curr_mapping = Mapping(bits[0][1])
                mappings[bits[0][0]] = curr_mapping
            elif re_range.match(row):
                bits = re.findall(re_range, row)
                curr_mapping.add_range(int(bits[0][1]), int(bits[0][0]), int(bits[0][2]))
            else:
                raise "Unknown line format"
    return (seeds, mappings)


def part1(fname: str):
    seeds, mappings = parse(fname)
    lowest = sys.maxsize
    for seed in seeds:
        next_type = 'seed'
        this_num = seed
        while next_type != "location":
            n = mappings[next_type].find_next(this_num)
            next_type = mappings[next_type].dest
            this_num = n
        if n < lowest:
            lowest = n
        print(n)
    print(f"Lowest is {lowest}")


if __name__ == '__main__':
    #part1("test.txt")
    part1("input.txt")
