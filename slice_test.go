package radix

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSortSlice(t *testing.T) {
	data := [...]string{"", "Hello", "foo", "fo", "xb", "xa", "bar", "foo", "f00", "%*&^*&^&", "***"}
	sorted := data[0:]
	sort.Strings(sorted)

	a := data[0:]
	SortSlice(a, func(i int) string { return a[i] })
	if !reflect.DeepEqual(a, sorted) {
		t.Errorf(" got %v", a)
		t.Errorf("want %v", sorted)
	}

	SortSlice(nil, func(i int) string { return a[i] })
	a = []string{}
	SortSlice(a, func(i int) string { return a[i] })
	if !reflect.DeepEqual(a, []string{}) {
		t.Errorf(" got %v", a)
		t.Errorf("want %v", []string{})
	}
	a = []string{""}
	SortSlice(a, func(i int) string { return a[i] })
	if !reflect.DeepEqual(a, []string{""}) {
		t.Errorf(" got %v", a)
		t.Errorf("want %v", []string{""})
	}
}

func TestSortSlice1k(t *testing.T) {
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	sorted := make([]string, len(data))
	copy(sorted, data)
	sort.Strings(sorted)

	SortSlice(data, func(i int) string { return data[i] })
	if !reflect.DeepEqual(data, sorted) {
		t.Errorf(" got %v", data)
		t.Errorf("want %v", sorted)
	}
}

func TestSortSliceBible(t *testing.T) {
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	sorted := make([]string, len(data))
	copy(sorted, data)
	sort.Strings(sorted)

	SortSlice(data, func(i int) string { return data[i] })
	if !reflect.DeepEqual(data, sorted) {
		for i, s := range data {
			if s != sorted[i] {
				t.Errorf("%v  got: %v", i, s)
				t.Errorf("%v want: %v", i, sorted[i])
			}
		}
	}
}

func BenchmarkRadixSortSliceBible(b *testing.B) {
	b.StopTimer()
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		SortSlice(a, func(i int) string { return a[i] })
		b.StopTimer()
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkSortSliceBible(b *testing.B) {
	b.StopTimer()
	var data []string
	f, err := os.Open("res/bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	for sc := bufio.NewScanner(f); sc.Scan(); {
		data = append(data, sc.Text())
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		b.StopTimer()
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkRadixSortSlice1k(b *testing.B) {
	b.StopTimer()
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		SortSlice(a, func(i int) string { return a[i] })
		b.StopTimer()
	}
}

func BenchmarkSortSlice1k(b *testing.B) {
	b.StopTimer()
	data := make([]string, 1<<10)
	for i := range data {
		data[i] = strconv.Itoa(i ^ 0x2cc)
	}

	a := make([]string, len(data))
	for i := 0; i < b.N; i++ {
		copy(a, data)
		b.StartTimer()
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		b.StopTimer()
	}
}
