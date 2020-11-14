
package knapsack

import "fmt"

const DEBUG = true

func showDP(DP [][]int) {
  N := len(DP)
  for i := 0; i < N; i++ {
    fmt.Printf("%v\n", DP[i])
  }
}

func Knapsack(K int, V []int, W []int) int {
  N := len(V)

  DP := make([][]int, N + 1)
  for i := 0; i < N + 1; i++ {
    DP[i] = make([]int, K + 1)
  }
  return knapsack(K, V, W, DP)
}

/**
 *
 * Suppose DP[i][j] represents the maximum value that can be obtained
 * from a combination of items 1 through i selected so that the total
 * weight does not exceed j 
 *
 * if j < W[i]
 *   DP[i+1][j] = DP[i][j]
 * else
 *   DP[i+1][j] = MAX DP[i][j]
 *                    DP[i][j-W[i]] + V[i]
 *
 */
func knapsack(K int, V []int, W []int, DP [][]int) int {
  N := len(V)

  if DEBUG {
    fmt.Println("--------------------------------")
    fmt.Printf("N: %d, K: %d\n", N, K)
    fmt.Printf("V: %v\n", V)
    fmt.Printf("W: %v\n", W)
    fmt.Println("--------------------------------")
  }
  for i := 0; i <= N - 1; i++ {
    for j := 1; j <= K; j++ {
      DP[i+1][j] = DP[i][j]
      if W[i] <= j {
        if DP[i+1][j] < DP[i][j - W[i]] + V[i] {
          DP[i+1][j] = DP[i][j - W[i]] + V[i]
        }
      }
    }
  }
  if DEBUG { showDP(DP) }
  return DP[N][K]
}
