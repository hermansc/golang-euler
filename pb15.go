package main

import (
  "fmt"
)

func main(){
  // Fill a table where we add paths until the cell in question
  const dim = 20
  grid := [dim+1][dim+1]int{}
  for i:=1;i<dim+1;i++{
    grid[0][i] = 1
    grid[i][0] = 1
  }
  // Add the paths
  for i:=1;i<dim+1;i++{
    for j:=1;j<dim+1;j++{
      grid[i][j] = grid[i][j-1] + grid[i-1][j]
    }
  }
  fmt.Println(grid[dim][dim])
}
