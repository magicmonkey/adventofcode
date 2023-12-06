import re


def distance(time, speed):
    return time * speed


def best_distance(time, record):
    dists = []
    for charge_time in range(0, time):
        remaining = time - charge_time
        d = distance(remaining, charge_time)
        if d > record:
            dists.append(d)
    return dists


def part1(fname: str):
    with open(fname) as file:
        lines = file.readlines()
    times = [int(t) for t in re.split(' +', lines[0].split(":")[1].strip())]
    dists = [int(d) for d in re.split(' +', lines[1].split(":")[1].strip())]

    output = 1
    for i in range(len(times)):
        print(f"--- looking at race {i} ---")
        d = best_distance(times[i], dists[i])
        output *= len(d)
    print(output)
    

if __name__ == '__main__':
    #part1("test.txt")
    part1("input.txt")

