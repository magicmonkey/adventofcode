import re

def is_symbol(c) -> bool:
    if c == '.':
        return False
    elif c.isdigit():
        return False
    else:
        return True

def char_adj_to_symbol(chars, x, y) -> bool:
    if is_symbol(chars[y-1][x-1]):
        return True
    elif is_symbol(chars[y-1][x]):
        return True
    elif is_symbol(chars[y-1][x+1]):
        return True
    elif is_symbol(chars[y][x-1]):
        return True
    elif is_symbol(chars[y][x]):
        return True
    elif is_symbol(chars[y][x+1]):
        return True
    elif is_symbol(chars[y+1][x-1]):
        return True
    elif is_symbol(chars[y+1][x]):
        return True
    elif is_symbol(chars[y+1][x+1]):
        return True
    else:
        return False

def part1(fname: str):
    chars = []
    with open(fname) as file:
        count = 0
        for line in file:
            row = list(line.strip())
            row.insert(0, '.')
            row.append('.')
            chars.append(row)
    blankline = ['.' for i in range(len(chars[0]))]
    chars.insert(0, blankline)
    chars.append(blankline)

    valid_coords = []
    counter = 0
    for y in range(1, len(chars)-1):
        found = False
        for x in range(1, len(chars[y])-1):
            if not chars[y][x].isdigit():
                continue
            # Is this the first digit in a number?  If so, parse it out
            if not chars[y][x-1].isdigit():
                n = re.findall('^(\d+)', "".join(chars[y][x:]))
                thisnum = int(n[0])
                found = False
            if not found and char_adj_to_symbol(chars, x, y):
                print(f"{thisnum} -> ({x}, {y})")
                counter += thisnum
                found = True
    print(counter)


def get_full_num_at_coord(chars, x, y) -> int:
    num = ''
    s = "".join(chars[y])
    # Iterate from the given location to the end
    for i in range(x, len(s)):
        if s[i].isdigit():
            num += s[i]
        else:
            break
    # Iterate from the given location to the start
    for i in range(x-1, -1, -1):
        if s[i].isdigit():
            num = s[i] + num
        else:
            break
    return(int(num))


def is_gear(chars, x, y) -> int:
    nums = []
    if chars[y-1][x-1].isdigit():
        nums.append(get_full_num_at_coord(chars, x-1, y-1))
    if chars[y-1][x].isdigit() and not chars[y-1][x-1].isdigit():
        nums.append(get_full_num_at_coord(chars, x, y-1))
    if chars[y-1][x+1].isdigit() and not chars[y-1][x].isdigit():
        nums.append(get_full_num_at_coord(chars, x+1, y-1))
    if chars[y][x-1].isdigit():
        nums.append(get_full_num_at_coord(chars, x-1, y))
    if chars[y][x+1].isdigit():
        nums.append(get_full_num_at_coord(chars, x+1, y))
    if chars[y+1][x-1].isdigit():
        nums.append(get_full_num_at_coord(chars, x-1, y+1))
    if chars[y+1][x].isdigit() and not chars[y+1][x-1].isdigit():
        nums.append(get_full_num_at_coord(chars, x, y+1))
    if chars[y+1][x+1].isdigit() and not chars[y+1][x].isdigit():
        nums.append(get_full_num_at_coord(chars, x+1, y+1))
    if len(nums) == 2:
        return nums[0] * nums[1]
    else:
        return -1


def part2(fname: str):
    chars = []
    with open(fname) as file:
        count = 0
        for line in file:
            row = list(line.strip())
            row.insert(0, '.')
            row.append('.')
            chars.append(row)
    blankline = ['.' for i in range(len(chars[0]))]
    chars.insert(0, blankline)
    chars.append(blankline)

    valid_coords = []
    counter = 0
    for y in range(1, len(chars)-1):
        found = False
        for x in range(1, len(chars[y])-1):
            if chars[y][x] != '*':
                continue
            ratio = is_gear(chars, x, y)
            if ratio >= 0:
                counter += ratio
                #print(f"({x}, {y}) -> {ratio}")
    print(counter)



if __name__ == '__main__':
    #part1('test.txt')
    #part1('input.txt')
    #part2('test.txt')
    part2('input.txt')


