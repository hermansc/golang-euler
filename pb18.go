package main

import (
  "fmt"
)

// See pb67 for a better, smaller and faster(!) solution.
// This solution builds a tree of Node interfaces, which takes a
// lot memory and time. Instead it is faster to do a bottom-up summation.

type Node struct{
  num int
  children []Node
}

func build_tree(input [][]int, x int, y int) Node{
  if x == len(input) - 1{
    return Node{num: input[x][y], children: []Node{}}
  }
  root := Node{num: input[x][y]}
  left_child := build_tree(input, x+1, y)
  right_child := build_tree(input, x+1, y+1)
  root.children = []Node{left_child, right_child}
  return root
}

func high_score(root Node) int{
  if len(root.children) > 1{
    l_score := high_score(root.children[0])
    r_score := high_score(root.children[1])
    if l_score > r_score{
      return root.num + l_score
    } else {
      return root.num + r_score
    }
  } else{
    return root.num
  }
}

func main(){
  triangle := [][]int{
  {75},
  {95, 64},
  {17,47,82},
  {18,35,87,10},
  {20,04,82,47,65},
  {19,01,23,75,03,34},
  {88,02,77,73,07,63,67},
  {99,65,04,28,06,16,70,92},
  {41,41,26,56,83,40,80,70,33},
  {41,48,72,33,47,32,37,16,94,29},
  {53,71,44,65,25,43,91,52,97,51,14},
  {70,11,33,28,77,73,17,78,39,68,17,57},
  {91,71,52,38,17,14,91,43,58,50,27,29,48},
  {63,66,04,68,89,53,67,30,73,16,69,87,40,31},
  {04,62,98,27,23,9,70,98,73,93,38,53,60,04,23}}
  root_node := build_tree(triangle, 0, 0)
  m := high_score(root_node)
  fmt.Println(m)
}
