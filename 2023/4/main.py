import re

def parse(fname: str):
    cards = {}
    with open(fname) as file:
        for line in file:
            row = line.strip()
            print(row)
            bits = re.findall('^Card +(\d+): ([^\|]*) \| (.*)$', row)
            for bit in bits:
                card = {
                    'winning': [int(x) for x in re.split(f' (?! )', bit[1].strip())],
                    'have': [int(x) for x in re.split(f' (?! )', bit[2].strip())]
                }
                yield card


def part1(fname: str):
    total = 0
    for card in parse(fname):
        score = 0
        for winning_num in card['winning']:
            if winning_num in card['have']:
                print(f"Got a match for {winning_num}")
                if score == 0:
                    score = 1
                else:
                    score *= 2
        total += score
        print(f"{card['have']} - {card['winning']} - {score} - {total}")
        print("")
    print(total)


if __name__ == '__main__':
    #part1("test.txt")
    part1("input.txt")

