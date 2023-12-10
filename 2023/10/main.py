from pprint import pprint


def parse(fname: str):
    tiles = []
    with open(fname) as file:
        for line in file:
            row = list(line.strip())
            # Add a left and right border
            tiles.append(['.'] + row + ['.'])

    # Add a top and bottom border
    tiles.insert(0, ['.'] * len(tiles[0]))
    tiles.append(['.'] * len(tiles[0]))

    # Find the start
    start = (0, 0)
    for y in range(len(tiles)):
        for x in range(len(tiles[y])):
            if tiles[y][x] == 'S':
                start = (y, x)
                break

    pprint(tiles)
    return (start, tiles)


def get_next(tiles, curr, avoid):
    nexts = []

    # Up
    if tiles[curr[0]][curr[1]] in ['S', '|', 'J', 'L'] and tiles[curr[0]-1][curr[1]] in ['|', 'F', '7']:
        nexts.append((curr[0]-1,curr[1]))
    # Right
    if tiles[curr[0]][curr[1]] in ['S', '-', 'L', 'F'] and tiles[curr[0]][curr[1]+1] in ['-', 'J', '7']:
        nexts.append((curr[0],curr[1]+1))
    # Down
    if tiles[curr[0]][curr[1]] in ['S', '|', '7', 'F'] and tiles[curr[0]+1][curr[1]] in ['|', 'J', 'L']:
        nexts.append((curr[0]+1,curr[1]))
    # Left
    if tiles[curr[0]][curr[1]] in ['S', '-', '7', 'J'] and tiles[curr[0]][curr[1]-1] in ['-', 'L', 'F']:
        nexts.append((curr[0],curr[1]-1))

    if len(nexts) == 1:
        # We must be next to the start
        if nexts[0] == avoid:
            return []
        else:
            return [nexts[0]]
    if nexts[0] == avoid:
        return [nexts[1]]
    elif nexts[1] == avoid:
        return [nexts[0]]
    else:
        return nexts


def part1(fname: str):
    start, tiles = parse(fname)
    route1, route2 = get_next(tiles, start, (0,0))
    print(f"Start at {start}")

    curr_pos = route1
    prev_hop = start
    counter = 0
    while True:
        #print(f"Current pos {curr_pos}")
        n = get_next(tiles, curr_pos, prev_hop)
        #print(f"Next options are {n}")
        if len(n) == 0:
            break
        prev_hop = curr_pos
        curr_pos = n[0]
        counter += 1
    print((counter / 2) + 1)


if __name__ == '__main__':
    #part1('test4.txt')
    part1('input.txt')
