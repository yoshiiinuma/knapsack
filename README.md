# 0-1 Knapsack Problem

Given a set of different N items, each one with an associated value and weight, determine the maximum value of items that can be packed into a knapsack of capacity K.
The values and weights of N items are given as integer arrays; V { v1, v2, ..., vn }, W { w1, w2, ..., wn }, respectively.

## Constraints

```
0 < N <= 100
0 < Vi, Wi <= 100
0 < K <= 10000
```

## Examples

```
INPUT:

  N: 4
  K: 5
  V: [3, 2, 4, 2]
  W: [2, 1, 3, 2]

OUTPUT:

  7 (= 3 + 2 + 2; the total weight is 5 = 2 + 1 + 2)
  
```

```
INPUT:

  N: 3
  K: 50
  V: [40, 80, 100]
  W: [10, 20 ,30]

OUTPUT:

  180 (= 80 + 100; the total weight is 50)
  
```

## Approach

This problem can be solved with Dynamic Programming.

Suppose DP[i][j] represents the maximum value that can be obtained from a combination of items 1 through i selected so that the total weight does not exceed j, then: 

```
i) When j < W[i]

  DP[i+1][j] = DP[i][j]

ii) When j >= W[i]

  DP[i+1][j] = MAX(DP[i][j], D[i][j - W[i]] + V[i])
```

## Complexity Analysis

* Time Complexity: O(N*K)
* Space Complexity: O(N*K)
