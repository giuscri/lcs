package main

import "fmt"

func main() {
	n := lcs("abc", "abc")
	n = lcs("abzzc", "abczz")
	fmt.Println(n)
}

func max(numbers ...int) int {
	M := numbers[0]
	for _, n := range numbers {
		if n > M {
			M = n
		}
	}

	return M
}

func lcs(s string, t string) int {
	if s == "" || t == "" {
		return 0
	}

	n := len(s)
	m := len(t)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		if s[0] == t[j] {
			dp[0][j] = 1
		} else if j >= 1 {
			dp[0][j] = dp[0][j-1]
		} else {
			dp[0][j] = 0
		}
	}

	for i := 0; i < n; i++ {
		if s[i] == t[0] {
			dp[i][0] = 1
		} else if i >= 1 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if s[i] == t[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n-1][m-1]
}
