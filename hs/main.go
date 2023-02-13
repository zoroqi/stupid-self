package main

func main() {
  for i := 0; i < 3; i++ {
    defer func() {print(i)}()
  }
  for i := range [3]int{} {
    defer func() {print(i)}()
  }
}
