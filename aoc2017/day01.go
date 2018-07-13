package aoc2017;

import (
  "io/ioutil"
)

type increment func(string) int

func Solve1(input string) int {
  return solve(input, one)
}

func Solve2(input string) int {
  return solve(input, half)
}

func solve(input string, fxn increment) int {
  ch := []rune(input)
  sum, len, incr := 0, len(ch), fxn(input)

  for i := 0; i < len; i++ {
    ix := (i + incr)

    if ix >= len {
      ix -= len
    }

    if ch[i] == ch[ix] {
      sum += (int(ch[i]) - '0')
    }
  }

  return sum
}

func one(s string) int {
  return 1
}

func half(s string) int {
  return len(s) / 2
}

func Input() string {
  data, err := ioutil.ReadFile("input01.txt")
  if err != nil {
      panic(err)
  }
  return string(data)
}
