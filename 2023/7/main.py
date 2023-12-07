from functools import cmp_to_key

class Hand:
    def __init__(self):
        self.cards = []
        self.bid = 0

    @classmethod
    def from_line(cls, line: str):
        cards, bid = line.split(" ")
        a = Hand()
        a.cards = list(cards)
        a.bid = int(bid)
        a.rank = a.score()
        return a
    
    def n_of_a_kind(self, n: int) -> bool:
        counts = {}
        for char in self.cards:
            if char in counts:
                counts[char] += 1
            else:
                counts[char] = 1
        return any(counts[x] == n for x in counts)

    def two_pair(self) -> bool:
        counts = {}
        for char in self.cards:
            if char in counts:
                counts[char] += 1
            else:
                counts[char] = 1
        num_twos = 0
        for i in counts:
            if counts[i] == 2:
                num_twos += 1
        return num_twos == 2

    def score(self) -> int:

        # Five of a kind
        if self.n_of_a_kind(5):
            return 7

        # Four of a kind
        if self.n_of_a_kind(4):
            return 6

        # Full house
        if self.n_of_a_kind(3) and self.n_of_a_kind(2):
            return 5

        # Three of a kind
        if self.n_of_a_kind(3):
            return 4

        # Two pair
        if self.two_pair():
            return 3

        # One pair
        if self.n_of_a_kind(2):
            return 2

        # High card
        return 1


def highest_card(a, b) -> int:
    rank = ['A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2']
    if a == b:
        return 0
    elif rank.index(a) < rank.index(b):
        return 1
    elif rank.index(a) > rank.index(b):
        return -1
    else:
        return 0


def highest_hand(a, b) -> int:
    if a.rank > b.rank:
        return 1
    elif a.rank < b.rank:
        return -1
    else:
        for i in range(len(a.cards)):
            c = highest_card(a.cards[i], b.cards[i])
            if c != 0:
                return c
    return 0


def part1(fname: str):
    hands = []
    with open(fname) as file:
        for line in file:
            hands.append(Hand.from_line(line.strip()))
    hands.sort(key = cmp_to_key(highest_hand))
    count = 0
    for i in range(len(hands)):
        count += (i+1) * hands[i].bid
    print(count)


if __name__ == '__main__':
    #part1("test.txt")
    part1("input.txt")

