""" This problem was asked by Two Sigma. Daily Coding Problem: Problem #769 [Hard]
    Alice wants to join her school's Probability Student Club. Membership dues are computed via one of two
    simple probabilistic games. The first game: roll a die repeatedly. Stop rolling once you get a five
    followed by a six. Your number of rolls is the amount you pay, in dollars.
    The second game: same, except that the stopping condition is a five followed by a five.
    Which of the two games should Alice elect to play? Does it even matter? Write a program to simulate
    the two games and calculate their expected value.
"""

def rollADie(count_match, num1, num2):
    import random
    count = 0
    while True:
        toss = random.randint(1, 6)
        while True:
            prev_toss = toss
            toss = random.randint(1, 6)
            count += 1
            if prev_toss == num1 and toss == num2:
                count_match += 1
                break
            if count == 1000000:
                break
        if count == 1000000:
            break
    print(count_match/count * 100)
    print(count)


rollADie(0, 5, 5)
rollADie(0, 5, 6)
