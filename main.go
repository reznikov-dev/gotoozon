package main

import (
    "bufio"
    "fmt"
    "gotoozon/algorithm"
    "log"
    "os"
    "sort"
    "strconv"
    "strings"
)

const delimiter = ","

// Простое демо
// Simple demo
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter integer number sequence (comma separated : 1, 1, 2, 3) and press enter: \n >")

    s, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalln("Fail to read input, err:", err)
    }

    sl := getIntSliceFromString(s)
    if len(sl) > 0 {
        sort.Ints(sl)
    }

    fmt.Println("Your array sorted number sequence is: ", sl)
    fmt.Print("Enter integer number sequence for sub array to check (comma separated : 1, 1, 2, 3) and press enter: \n > ")

    s, err = reader.ReadString('\n')
    if err != nil {
        log.Fatalln("Fail to read input, err:", err)
    }

    subSlice := getIntSliceFromString(s)
    if len(subSlice) > 0 {
        sort.Ints(subSlice)
    }

    fmt.Println("Your sub array sorted number sequence is: ", subSlice)
    fmt.Println("\nCheck result is subarray is include in array is :", algorithm.IsIncludeWithRepeatingNumbers(sl, subSlice))
    fmt.Println("Good Bye!")
}

func getIntSliceFromString(s string) []int {
    var sl []int

    for _, sn := range strings.Split(s, delimiter) {
        if n, err := strconv.Atoi(strings.TrimSpace(sn)); err == nil {
            sl = append(sl, n)
        } else {
            log.Println("String to int convert err:", err, " string: ", sn)
        }
    }

    return sl
}
