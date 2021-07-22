package algorithm

const notFound = -1

// IsInclude
// Алгоритм для простого случая - когда мы предполагаем,
// что все числа в массиве (слайсе) могут встречаться только один раз.
// Т.к. в условиях задачи не оговорен этот момент и в примерах указаны
// именно такие последовательности
//
// Simple case algorithm - as we assume that all numbers in sorted slice
// can occurs only once. Because there is no information about other cases
// (repeating numbers sequences, for example) in task description and
// test examples.
func IsInclude(array []int, subarray []int) bool {
    l := len(subarray)
    if l == 0 {
        return true
    }

    //O (log N)
    idx := binarySearch(array, subarray[0])
    if idx == notFound {
        return false
    }

    // isSlicesEqual O (M) => O (log(N) + M)
    return isSlicesEqual(array[idx:min(idx + l, len(array))], subarray)
}

// IsIncludeWithRepeatingNumbers
// Алгоритм с обработкой ситуации, когда в массиве (слайсе)
// числа могут повторяться несколько раз []int{1, 1, 2, 2, 2, 3, 5}
//
// Complex algorithm with handling case for possible repeating numbers
// sequences in slice []int{1, 1, 2, 2, 2, 3, 5}
func IsIncludeWithRepeatingNumbers(array []int, subarray []int) bool {
    l := len(subarray)
    if l == 0 {
        return true
    }

    subSeqCount := 0
    if l > 1 {
        // O (M - 1) => O (M)
        for i := 1; i <= l - 1 && subarray[i - 1] == subarray[i]; i++ {
            subSeqCount++
        }
    }

    // O (log N)
    idx := binarySearchConsequenceEnd(array, subarray[0])
    if idx == notFound {
        return false
    }

    // Нет повторяющейся последовательности в начале вложенного массива (слайса)
    // No subsequence on sub array (slice) start
    if subSeqCount == 0 {
        return isSlicesEqual(array[idx:min(idx + l, len(array))], subarray)
    }

    st := max(0, idx - subSeqCount)
    end := min(st + l, len(array))

    // isSlicesEqual O (M) => O (log(N) + 2M)
    return isSlicesEqual(array[st:end], subarray)
}

func binarySearch(s []int, n int) int {
    lt, rt := 0, len(s) - 1

    //Мы знаем, что слайс отсортирован, поэтому можем сразу обработать пограничные случаи
    //As we know, that slice is sorted, we can handle edge cases at first
    if n < s[lt] || n > s[rt] {
        return notFound
    }

    for lt <= rt {
        m := (lt + rt) / 2

        if s[m] == n {
            return m
        }

        if s[m] < n {
            lt = m + 1
        } else {
            rt = m - 1
        }
    }

    return notFound
}

// Поиск индекса i где s[i + 1] > n, индекс последнего n в повторяющейся последовательности
//
// Search for index i where s[i + 1] > n, it would be the index of last n value in consequence
func binarySearchConsequenceEnd(s []int, n int) int {
    lt, rt, m := 0, len(s) - 1, notFound

    if n < s[lt] || n > s[rt] {
        return m
    }

    if s[rt] == n {
        return rt
    }

    for lt <= rt {
        m = (lt + rt) / 2

        if s[m] <= n {
            lt = m + 1
        } else {
            rt = m - 1
        }
    }

    if s[max(0, m)] == n {
        return m
    }

    if m == notFound || s[max(0, m - 1)] != n {
        return notFound
    }

    return m - 1
}

func isSlicesEqual(s1, s2 []int) bool {
    if len(s1) != len(s2) {
        return false
    }

    // O (M)
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            return false
        }
    }

    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}