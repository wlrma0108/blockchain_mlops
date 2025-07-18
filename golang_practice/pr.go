package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    lastIdx := make(map[rune]int) 
    maxLen, start := 0, 0

    for i, ch := range s {
        if prev, found := lastIdx[ch]; found && prev >= start {
            start = prev + 1
        }
        lastIdx[ch] = i
        if curLen := i - start + 1; curLen > maxLen {
            maxLen = curLen
        }
    }
    return maxLen
}

func main() {
    tests := []struct {
        input string
        want  int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3},
        {"", 0},
        {"au", 2},
    }

    for _, tc := range tests {
        got := lengthOfLongestSubstring(tc.input)
        fmt.Printf("Input: %-10q  → Output: %d  (Expected: %d)\n",
            tc.input, got, tc.want)
    }
}
