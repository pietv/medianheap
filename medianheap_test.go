package medianheap_test

import (
	. "math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	. "github.com/pietv/medianheap"
)

var Tests = []struct {
	name string
	in   []int
	want []int
}{
	{"1", []int{0}, []int{0}},
	{"2", []int{MaxInt32}, []int{MaxInt32}},
	{"3", []int{MinInt32}, []int{MinInt32}},

	{"4", []int{0, 1}, []int{0, 0}},
	{"5", []int{-1, 2}, []int{-1, -1}},
	{"6", []int{2, -1}, []int{2, -1}},
	{"7", []int{2, 1}, []int{2, 1}},
	{"8", []int{2, 2}, []int{2, 2}},
	{"9", []int{MaxInt32, MinInt32}, []int{MaxInt32, MinInt32}},
	{"10", []int{MinInt32, MaxInt32}, []int{MinInt32, MinInt32}},
	{"11", []int{MinInt32, 0}, []int{MinInt32, MinInt32}},
	{"12", []int{0, MinInt32}, []int{0, MinInt32}},
	{"13", []int{0, MaxInt32}, []int{0, 0}},
	{"14", []int{MinInt32, 0}, []int{MinInt32, MinInt32}},
	{"15", []int{MaxInt32, 0}, []int{MaxInt32, 0}},

	{"16", []int{1, 2, 3, 4, 5}, []int{1, 1, 2, 2, 3}},
	{"17", []int{5, 4, 3, 2, 1}, []int{5, 4, 4, 3, 3}},
	{"18", []int{2, 4, 5, 3, 1}, []int{2, 2, 4, 3, 3}},
	{"19", []int{20, 40, 50, 30, 10}, []int{20, 20, 40, 30, 30}},
	{"20", []int{0, 0, 0, 0, 1}, []int{0, 0, 0, 0, 0}},
	{"21", []int{0, 0, 0, 1, 1}, []int{0, 0, 0, 0, 0}},
	{"22", []int{0, 0, 1, 1, 1}, []int{0, 0, 0, 0, 1}},
	{"23", []int{0, 1, 1, 1, 1}, []int{0, 0, 1, 1, 1}},
	{"24", []int{1, 0, 0, 0, 0}, []int{1, 0, 0, 0, 0}},
	{"25", []int{1, 1, 0, 0, 0}, []int{1, 1, 1, 0, 0}},
	{"26", []int{1, 1, 1, 0, 0}, []int{1, 1, 1, 1, 1}},

	{"27", []int{0, 0, MaxInt32, 0, 0},
		[]int{0, 0, 0, 0, 0}},
	{"28", []int{MaxInt32, MinInt32, 0, MaxInt32, 0},
		[]int{MaxInt32, MinInt32, 0, 0, 0}},
	{"29", []int{MinInt32, 0, MinInt32, MinInt32, 0},
		[]int{MinInt32, MinInt32, MinInt32, MinInt32, MinInt32}},
	{"30", []int{0, 0, 0, MaxInt32, MaxInt32},
		[]int{0, 0, 0, 0, 0}},
	{"31", []int{0, 0, MaxInt32, MinInt32, MaxInt32},
		[]int{0, 0, 0, 0, 0}},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func checkAdd(a []int) []int {
	h := New()
	out := make([]int, 0)
	for _, elem := range a {
		h.Add(elem)
		out = append(out, h.Median())
	}
	return out
}

func checkUpdate(a []int) []int {
	h := New()
	out := make([]int, 0)
	for _, elem := range a {
		out = append(out, h.Update(elem))
	}
	return out
}

func checkReference(a []int) []int {
	median := func(a []int) int {
		b := make([]int, len(a))
		copy(b, a)

		sort.Ints(b)
		if len(b)%2 == 0 {
			return b[(len(b)-1)/2]
		} else {
			return b[len(b)/2]
		}
	}
	out := make([]int, 0)

	for i := 1; i <= len(a); i++ {
		out = append(out, median(a[:i]))
	}
	return out
}

func genRandomIntSlice(size int) []int {
	seq := make([]int, 0)

	N := rand.Intn(size)
	for i := 0; i < N; i++ {
		// Generate both positive and negative integers.
		seq = append(seq, rand.Int()*(-1*rand.Intn(2)))
	}
	return seq
}

func TestMedianWithoutAdd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("median didn't panic when expected to")
		}
	}()

	h := New()
	_ = h.Median()
}

func TestMultipleCalls(t *testing.T) {
	h := New()
	h.Add(-1)

	if h.Median() != h.Median() {
		t.Errorf("consecutive calls don't yield the same result")
	}
}

func TestAdd(t *testing.T) {
	for _, test := range Tests {
		if actual := checkAdd(test.in); reflect.DeepEqual(actual, test.want) != true {
			t.Errorf("%q; got %v, want %v", test.name, actual, test.want)
		}
	}
}

func TestUpdate(t *testing.T) {
	for _, test := range Tests {
		if actual := checkUpdate(test.in); reflect.DeepEqual(actual, test.want) != true {
			t.Errorf("%q: got %v, want %v", test.name, actual, test.want)
		}
	}
}

func TestRandomAdd10(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := genRandomIntSlice(256)

		ref, actual := checkReference(a), checkAdd(a)
		if reflect.DeepEqual(ref, actual) != true {
			t.Errorf("got %v, want %v", actual, ref)
		}
	}
}

func TestRandomUpdate10(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := genRandomIntSlice(256)

		ref, actual := checkReference(a), checkUpdate(a)
		if reflect.DeepEqual(ref, actual) != true {
			t.Errorf("got %v, want %v", actual, ref)
		}
	}
}

func BenchmarkManyAddsOneMedian(b *testing.B) {
	b.StopTimer()
	h := New()

	for i := 0; i < b.N; i++ {
		n := rand.Int()
		b.StartTimer()
		h.Add(n)
		b.StopTimer()
	}
	b.StartTimer()
	_ = h.Median()
	b.StopTimer()
}

func BenchmarkManyAddsManyMedians(b *testing.B) {
	b.StopTimer()
	h := New()

	for i := 0; i < b.N; i++ {
		n := rand.Int()
		b.StartTimer()
		h.Add(n)
		_ = h.Median()
		b.StopTimer()
	}
}
