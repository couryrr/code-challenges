def minimumSwaps(arr):
    swaps = 0
    for i, v in enumerate(arr):
        while i+1 != v:
            swaps += 1
            temp = arr[v-1]
            arr[v-1] = v
            arr[i] = temp
            v = temp

    return swaps


if __name__ == '__main__':
    arr = [2, 31, 1, 38, 29, 5, 44, 6, 12, 18, 39, 9, 48, 49, 13, 11, 7, 27, 14, 33, 50, 21, 46, 23, 15,
           26, 8, 47, 40, 3, 32, 22, 34, 42, 16, 41, 24, 10, 4, 28, 36, 30, 37, 35, 20, 17, 45, 43, 25, 19]

    print(minimumSwaps(arr))
