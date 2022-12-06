package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []rune
}

func (s *Stack) Push(x []rune) {
	for _, ea := range x {
		s.Items = prepend(s.Items, ea)
	}
}

func (s *Stack) Pop(n int) []rune {
	s.underflow(n)

	popped := []rune{}
	for n > 0 {
		popped = prepend(popped, s.Items[0])
		s.Items = s.Items[1:]
		n--
	}
	return popped
}

func PopPush(n int, s1, s2 *Stack) (*Stack, *Stack) {
	s1.underflow(n)
	fmt.Printf("Moving %s from %s to %s - ", string(s1.Items[:n]), string(s1.Items), string(s2.Items))

	for n > 0 {
		s2.Items = prepend(s2.Items, s1.Items[0])
		s1.Items = s1.Items[1:]
		n--
	}

	fmt.Printf("%s, %s\n", string(s1.Items), string(s2.Items))
	return s1, s2
}

func Move(n int, s1, s2 *Stack) (*Stack, *Stack) {
	s1.underflow(n)
	fmt.Printf("Moving %s from %s to %s - ", string(s1.Items[:n]), string(s1.Items), string(s2.Items))

	s2.Items, s1.Items = append(s1.Items[:n:n], s2.Items...), s1.Items[n:]

	fmt.Printf("%s, %s\n", string(s1.Items), string(s2.Items))
	return s1, s2
}

func (s *Stack) String() string {
	return string(s.Items)
}

func (s *Stack) underflow(n int) {
	if n > len(s.Items) {
		fmt.Printf("Failed to take %d items from %s\n", n, string(s.Items))
		panic(errors.New("Stack underflow"))
	}
}

func prepend(s []rune, v rune) []rune {
	return append([]rune{v}, s...)
}
