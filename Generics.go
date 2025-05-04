package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// ====== SECTION 2: Generic Struct ======
type Pair[T, U any] struct {
	First  T
	Second U
}

func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// ====== SECTION 3: Generic Constraints ======
type Adder interface {
	~int | ~float64 // allows int and float64 (also aliases)
}

func Add[T Adder](a, b T) T {
	return a + b
}

// ====== SECTION 4: Custom Constraint with Interface ======
type Comparable interface {
	Compare(other any) int
}

type Number struct {
	Value int
}

func (n Number) Compare(other any) int {
	o := other.(Number)
	return n.Value - o.Value
}

func Max[T Comparable](a, b T) T {
	if a.Compare(b) > 0 {
		return a
	}
	return b
}

// ====== SECTION 5: Generic Map Function ======
func Map[T any, U any](input []T, f func(T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = f(v)
	}
	return output
}

// ====== SECTION 6: Generic Stack ======
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		panic("stack is empty")
	}
	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// ====== MAIN FUNCTION TO TEST ALL ======
func main() {
	fmt.Println("=== Basic Generic Function ===")
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})

	fmt.Println("\n=== Generic Struct ===")
	p := Pair[string, int]{First: "Age", Second: 30}
	fmt.Println(p)

	fmt.Println("\n=== Generic Constraints ===")
	fmt.Println(Add(3, 4))
	fmt.Println(Add(3.5, 2.5))

	fmt.Println("\n=== Custom Constraint with Interface ===")
	n1 := Number{Value: 10}
	n2 := Number{Value: 20}
	fmt.Println("Max:", Max(n1, n2))

	fmt.Println("\n=== Generic Map Function ===")
	squared := Map([]int{1, 2, 3}, func(x int) int { return x * x })
	PrintSlice(squared)

	fmt.Println("\n=== Generic Stack ===")
	stack := Stack[string]{}
	stack.Push("Go")
	stack.Push("Lang")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
