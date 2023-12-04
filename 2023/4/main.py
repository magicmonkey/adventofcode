import re
from pprint import pprint

def parse(fname: str):
    cards = {}
    with open(fname) as file:
        for line in file:
            row = line.strip()
            bits = re.findall('^Card +(\d+): ([^\|]*) \| (.*)$', row)
            for bit in bits:
                card = {
                    'winning': [int(x) for x in re.split(f' (?! )', bit[1].strip())],
                    'have': [int(x) for x in re.split(f' (?! )', bit[2].strip())]
                }
                yield card


def num_wins(card) -> int:
    count = 0
    for winning_num in card['winning']:
        if winning_num in card['have']:
            count += 1
    return count


def part1(fname: str):
    total = 0
    for card in parse(fname):
        score = 0
        n = num_wins(card)
        if n == 0:
            score = 0
        elif n == 1:
            score = 1
        else:
            score = 2 ** (n-1)
        total += score
    print(total)


def part2(fname: str):
    total = 0
    cards = []

    for card in parse(fname):
        n = num_wins(card)
        card['wins'] = n
        card['instances'] = 1
        cards.append(card)

    for c in range(len(cards)):
        for i in range(cards[c]['wins']):
            cards[c+i+1]['instances'] += 1 * cards[c]['instances']

    count = 0
    for card in cards:
        count += card['instances']

    pprint(count)


if __name__ == '__main__':
    #part1("test.txt")
    #part1("input.txt")
    #part2("test.txt")
    part2("input.txt")

