#!
from turtle import pos


def arrayManipulation(n, queries):
    arr = [0] * (n+2)

    for start, end, value in queries:
        arr[start] += value
        if end+1 <= n:
            arr[end+1] -= value

    max = temp = 0
    for v in arr:
        temp += v
        if max < temp:
            max = temp
    return max


if __name__ == '__main__':
    fptr = open('./test-data', 'r')
    lines = fptr.readlines()
    fptr.close()

    header = lines[0].rstrip().split()
    n = int(header[0])
    m = int(header[1])
    queries = []
    for i in range(1, m+1):
        queries.append(list(map(int, lines[i].rstrip().split())))

    result = arrayManipulation(n, queries)
    print(result)
