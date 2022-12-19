#!
from turtle import pos


def minimumBribes(q):
    bribes = 0
    # Start at the back
    position = len(q)-1
    while position >= 0:
        '''
        Checking to see if the current value matches
        the index +1 (value = index +1)
        In the end this is sorting with a constraint
        '''
        if q[position] != position+1:
            # When it does not see if 1 or 2 indexs away is the correct value
            if position-1 >= 0 and q[position-1] == position+1:
                # Swap the positions (inc bribes)
                v = q[position-1]
                temp = q[position:position+1]
                temp.append(v)
                q[position-1] = temp[0]
                q[position] = temp[1]
                bribes += 1
            elif position-2 >= 0 and q[position-2] == position+1:
                # Swap both of the positions (inc bribes)
                v = q[position-2]
                temp = q[position-1:position+1]
                temp.append(v)
                q[position-2] = temp[0]
                q[position-1] = temp[1]
                q[position] = temp[2]
                bribes += 2
            # More than 2 places fails constraint
            else:
                print("Too chaotic")
                bribes = -1
                break
        position -= 1
    if bribes != -1:
        print(bribes)


if __name__ == '__main__':
    q = list(map(int, "1 2 5 3 7 8 6 4".rstrip().split()))
    minimumBribes(q)
