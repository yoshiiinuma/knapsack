
package knapsack

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  var K int
  var W []int
  var V []int

  K = 5
  V = []int{3, 2, 4, 2}
  W = []int{2, 1, 3, 2}
  Equal(t, 7, Knapsack(K, V, W))

  K = 50
  V = []int{40, 80, 100}
  W = []int{10, 20, 30}
  Equal(t, 180, Knapsack(K, V, W))

  K = 5
  V = []int{1, 2, 3, 4}
  W = []int{1, 2, 3, 4}
  Equal(t, 5, Knapsack(K, V, W))

  K = 5
  V = []int{4, 3, 2, 1}
  W = []int{1, 2, 3, 4}
  Equal(t, 7, Knapsack(K, V, W))

  K = 3
  V = []int{1, 2, 3, 4}
  W = []int{3, 3, 3, 3}
  Equal(t, 4, Knapsack(K, V, W))

  K = 12
  V = []int{1, 2, 3, 4}
  W = []int{3, 3, 3, 3}
  Equal(t, 10, Knapsack(K, V, W))

  K = 11
  V = []int{1, 2, 3, 4}
  W = []int{3, 3, 3, 3}
  Equal(t, 9, Knapsack(K, V, W))
}
