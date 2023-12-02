import re

{
    1:  [
        {'red':4, 'blue':3},
    ]
}

def process_line(line:str) -> (bool, int):
    gid_1, bags_1 = line.split(':',2)

    # First get the game ID
    gid_2 = [int(s) for s in re.findall(r'\d+', gid_1)]
    gid = int(gid_2[0])

    # Now parse out the bag pulling
    bags_2 = bags_1.split(';')
    retval = []
    for bag in bags_2:
        cubes = bag.split(',')
        these = {}
        for cube in cubes:
            a = re.findall(r'(\d+) (red|green|blue)', cube)[0]
            cube_num = int(a[0])
            cube_col = a[1]
            these[cube_col] = cube_num
        retval.append(these)
    return gid, retval


def parse(fname: str):
    retval = {}
    with open(fname) as file:
        count = 0
        for line in file:
            line = line.strip()
            gid, bag = process_line(line)
            retval[gid] = bag
    return retval

def part1(games):
    maxcubes = {
        'red':12,
        'green':13,
        'blue':14
    }
    count = 0
    for gid, game in games.items():
        valid = True
        for bag in game:
            for col, num in bag.items():
                if num > maxcubes[col]:
                    valid = False
        if valid:
            count += gid
    print(count)

def part2(games):
    count = 0
    for gid, game in games.items():
        maxcols = {
            'red':0,
            'green':0,
            'blue':0
        }
        for bag in game:
            for col, num in bag.items():
                if num > maxcols[col]:
                    maxcols[col] = num
        power = maxcols['red'] * maxcols['green'] * maxcols['blue']
        count += power
    print(count)

if __name__ == '__main__':
    #games = parse('test.txt')
    games = parse('input.txt')
    part1(games)
    part2(games)



###
# 8 lynel guts
# 15 fire and electric lizalfos tails
# 7 bladed rhino beetles
# blue white frox parts

