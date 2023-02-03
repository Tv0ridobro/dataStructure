# data-structure
![GitHub Workflow Status](https://github.com/Tv0ridobro/data-structure/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Tv0ridobro/data-structure)](https://goreportcard.com/report/github.com/Tv0ridobro/data-structure)
[![GoDoc](https://godoc.org/github.com/Tv0ridobro/data-structure?status.svg)](https://godoc.org/github.com/Tv0ridobro/data-structure)


Some data structures in go, using generics
# Installation
```
go get github.com/Tv0ridobro/data-structure
```
# Usage
## List
```go
l := list.New[int]()
l.PushBack(1)
l.PushBack(2)
l.PushFront(0)
fmt.Println(l.All()) // [0 1 2]
l.Reverse()
fmt.Println(l.All()) // [2 1 0]
```
## Treap
```go
t := treap.New[int]()
t.Insert(5)
t.Insert(8)
t.Insert(10)
t.Insert(19)
fmt.Println(t.Contains(5)) // true
t.Remove(5)
fmt.Println(t.Contains(5)) // false
fmt.Println(t.All()) // [8 10 19]
```
## Segment-tree
```go
// segment tree for sum
tree := segmenttree.New([]int{17, 2, 3, 4, 6, 0, 23, 30}, func(a, b int) int { return a + b }, 0)
fmt.Println(tree.Query(0, 7)) // 85 
fmt.Println(tree.Query(1, 6)) // 38
tree.Modify(2, 40)
fmt.Println(tree.Query(0, 7)) // 122
fmt.Println(tree.Query(2, 2)) // 40
```
## Sparse-table
```go
// sparse table for greatest common divisor
table := sparsetable.New(math.GCD, []int{2, 3, 5, 4, 6, 8, 16})
fmt.Println(table.Query(0, 2)) // 1
fmt.Println(table.Query(3, 5)) // 2
fmt.Println(table.Query(2, 3)) // 1
fmt.Println(table.Query(5, 6)) // 8
```
