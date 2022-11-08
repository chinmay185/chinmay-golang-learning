package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

func pointers() {
	i, j := 42, 2701

	p := &i           // point to i
	pp := &p          // pointer to a pointer
	fmt.Println(p)    // read i through the pointer
	fmt.Println(**pp) // read i through the pointer of a pointer
	*p = 21           // set i through the pointer
	fmt.Println(i)    // see the new value of i

	**pp = 24      // set i through the pointer of a pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

// override how string representation of Vertex works
func (v Vertex) String() string {
	return fmt.Sprintf("Vertext x: %d, y: %d", v.X, v.Y)
}

func pointer_to_a_struct() {
	v := Vertex{10, 20}
	p := &v
	p.X = 1e3
	fmt.Println(v)
}

func struct_literals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}

func array_basics() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := []string{"foo", "bar", "baz"}
	fmt.Println(b)
}

func slices_basics() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // first two elements
	b := names[1:3] // second and third element
	fmt.Println(a, b)

	b[0] = "XXX" // set the element in underlying array
	fmt.Println(a, b)
	fmt.Println(names)
}

func nil_slices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func making_slices() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	bb := append(b, 1, 2, 3, 4, 5)
	printSlice("b", b)
	printSlice("bb", bb)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func appending_slices() {
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// append works on nil slices.
	s = append(s, 0)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// The slice grows as needed.
	s = append(s, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5, 6, 7, 9)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice_literals() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func slice_defaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[1:4])
	fmt.Println(s[:2])
	fmt.Println(s[1:])
	fmt.Println(s[:])
}

func range_slices() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func mutating_maps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	split := strings.Split(s, " ")
	for _, v := range split {
		if _, ok := wc[v]; !ok {
			wc[v] = 0
		}
		wc[v]++
		fmt.Println(v)
	}
	return wc
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func function_as_values() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func function_as_closures() {
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() {
	fibonacci_closure := func() func() int {
		prev := 0
		current := 1
		return func() int {
			next := prev + current
			prev = current
			current = next
			return next
		}
	}

	f := fibonacci_closure()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func more_types() {
	pointers()

	// struct printing
	fmt.Println(Vertex{1, 2})

	// simple struct equality
	fmt.Println(Vertex{1, 2} == Vertex{1, 2})

	pointer_to_a_struct()
	struct_literals()
	array_basics()
	slices_basics()
	slice_literals()
	slice_defaults()
	nil_slices()
	making_slices()
	appending_slices()
	range_slices()
	maps()
	mutating_maps()
	wc.Test(WordCount)
	function_as_values()
	function_as_closures()
	fibonacci()
}
