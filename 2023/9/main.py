def get_gaps(nums):
    retval = []
    for i in range(1, len(nums)):
        retval.append(nums[i] - nums[i-1])
    return retval


def get_next(nums):
    gaps = [nums,]
    while True:
        next_gaps = get_gaps(gaps[len(gaps)-1])
        gaps.append(next_gaps)
        if len(set(next_gaps)) == 1:
            break
    # Now unroll the right-hand-side of all of the gaps
    thing = 0
    for i in range(len(gaps), 0, -1):
        thing = gaps[i-1][len(gaps[i-1])-1]
        gaps[i-2][len(gaps[i-2])-1] += thing
    return gaps[0][len(gaps[0])-1]


def part1(fname: str):
    with open(fname) as file:
        counter = 0
        for line in file:
            nums = [int(x) for x in line.split(" ")]
            counter += get_next(nums)
    print(counter)

if __name__ == '__main__':
    part1("test.txt")
    part1("input.txt")

