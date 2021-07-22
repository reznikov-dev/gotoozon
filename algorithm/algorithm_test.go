package algorithm

import (
    "math/rand"
    "testing"
    "time"
)

type testCase struct {
    slice []int
    subSlice []int
    expected bool
}

func init()  {
    rand.Seed(time.Now().UnixNano())
}

func TestIsInclude(t *testing.T) {
    slice := []int{1, 2, 3, 5, 7, 9, 11}

    tCases := []testCase{
        {slice, []int{}, true},
        {slice, []int{3, 5, 7}, true},
        {slice, []int{4, 5, 7}, false},
    }

    for _, tCase := range tCases {
        t.Run("test include", testInclude(&tCase, IsInclude))
    }
}

func TestIsIncludeWithRepeatingNumbers(t *testing.T) {
    slice := []int{1, 1, 1, 2, 3, 5, 5, 7, 9, 10}

    tCases := []testCase{
        {slice, []int{}, true},
        {slice, []int{3, 5, 5, 7}, true},
        {slice, []int{4, 5, 7}, false},
        {slice, []int{1, 1, 2, 3}, true},
        {slice, []int{1, 2, 3}, true},
        {slice, []int{5, 5, 5, 7}, false},
        {slice, []int{1, 1, 1}, true},
        {slice, []int{1, 1, 1, 1}, false},
    }

    for _, tCase := range tCases {
        t.Run("test include with repeating numbers", testInclude(&tCase, IsIncludeWithRepeatingNumbers))
    }
}

func TestBinarySearchConsequenceEndNotFound(t *testing.T) {
    tCases := [][]int{
        0 : {1,1,2,3},
        1 : {0,0,3,4},
        2 : {1,1,1,5,6,7},
    }

    for n, s := range tCases {
        if idx := binarySearchConsequenceEnd(s, n); idx != notFound {
            t.Error("Expected :", notFound, "actual:", idx, "slice:", s, "needle:", n)
        }
    }
}

func TestBinarySearchConsequenceEndFound(t *testing.T) {
    needle := 3

    tCases := [][]int{
        0 : {3,4,4,4},
        1 : {3,3,4,5},
        2 : {1,3,3,4},
        3 : {1,3,3,3},
        4 : {1,2,2,2,3,4},
        5 : {1,1,1,3,3,3,4,4,5},
    }

    for i, s := range tCases {
        if idx := binarySearchConsequenceEnd(s, needle); idx != i {
            t.Error("Expected :", i, "actual:", idx, "slice:", s, "needle:", needle)
        }
    }
}

func TestGeneratedSequences(t *testing.T) {
    tCases := generateTestCases(20)

    for _, tCase := range tCases {
        t.Run("test include with repeating numbers on generated sequences", testInclude(tCase, IsIncludeWithRepeatingNumbers))
    }
}

func testInclude(tCase *testCase, f func (a, b []int) bool) func(*testing.T) {
    return func(t *testing.T) {
        if tCase.expected != f(tCase.slice, tCase.subSlice) {
            t.Error("expected:", tCase.expected, "actual:", !tCase.expected, "slice:", tCase.slice, "subSlice:", tCase.subSlice)
        }
    }
}

func generateTestCases(n int) []*testCase {
    var (
        num int
        idx int
        s []int
    )

    var tc []*testCase

    for i := 1; i <= n; i++ {
        num, idx, s = 0, -1, []int{}

        for j := 0; j < 10 * i; j++ {
            r := rand.Intn(10)
            if r >= 5 {
                num++
            }

            s = append(s, num)

            if idx == -1 && r == 0 {
                idx = j
            }
        }

        if idx == -1 {
            idx = len(s) - 1
        }

        subLen := min(idx + 1 + rand.Intn(5), len(s))

        cs := &testCase{
            slice: s,
            subSlice: []int{},
            expected: true,
        }

        for k := idx; k < subLen; k++ {
            cs.subSlice = append(cs.subSlice, s[k])
        }

        if subLen == idx + 5 {
            // тестовая последовательность для под массива будет не отсортирована, но мы точно будем уверены,
            // что результат проверки должен быть отрицательный
            // tricky to get guaranteed false result sequence (subSlice will not be sorted)
            cs.subSlice[3] = -1
            cs.expected = false
        }

        tc = append(tc, cs)
    }

    return tc
}
