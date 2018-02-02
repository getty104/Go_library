package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "sort"
  "strconv"
)

//============================================================

var sc = bufio.NewScanner(os.Stdin)

const INF = math.MaxInt8

func max(a int, b int) int { return map[bool]int{true: a, false: b}[a > b] }
func min(a int, b int) int { return map[bool]int{true: a, false: b}[a < b] }
func maxch(a *int, b *int) { *a = max(*a, *b) }
func minch(a *int, b *int) { *a = min(*a, *b) }
func cout(out string)      { fmt.Println(out) }
func setScSplit()          { sc.Split(bufio.ScanWords) }
func stoi(str string) int {
  num, _ := strconv.Atoi(str)
  return num
}
func itos(num int) string {
  str := strconv.Itoa(num)
  return str
}
func gsf() string {
  sc.Scan()
  return sc.Text()
}
func gif() int {
  i := stoi(gsf())
  return i
}
func gff() float64 {
  f, _ := strconv.ParseFloat(gsf(), 64)
  return f
}
func uniq(list []int) []int {
  m := make(map[int]struct{})
  for _, ele := range list {
    m[ele] = struct{}{}
  }
  uni := []int{}
  for i := range m {
    uni = append(uni, i)
  }
  return uni
}
func usort(list []int) []int {
  sort.Sort(sort.IntSlice(list))
  return list
}
func dsort(list []int) []int {
  sort.Sort(sort.Reverse(sort.IntSlice(list)))
  return list
}
func push(t *[]int, x int) {
  tmp := *t
  *t = append(tmp, x)
}
func pop(t *[]int) int {
  tmp := *t
  x := tmp[len(tmp)-1]
  *t = tmp[:len(tmp)-1]
  return x
}
func deq(t *[]int) int {
  tmp := *t
  x := tmp[0]
  *t = tmp[1:]
  return x
}
func exist(s []int, e int) bool {
  for _, v := range s {
    if e == v {
      return true
    }
  }
  return false
}

//=============================================================

func main() {
  setScSplit()
}
