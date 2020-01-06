package main

import (
	"fmt"
	"io"
)

type Exchanger interface {
	Exchange()
}

type StringPair struct{ first, second string }

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

type Point [2]int

func (point *Point) Exchange() { point[0], point[1] = point[1], point[0] }

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.first, pair.second)
}

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
}

func main() {
	jekyll := StringPair{"Henry", "Jekyll"}
	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}
	fmt.Println("Before: ", jekyll, hyde, point)
	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()
	fmt.Println("After #1: ", jekyll, hyde, point)
	exchangeThese(&jekyll, &hyde, &point)
	fmt.Println("After #2: ", jekyll, hyde, point)

	const size = 16
	robert := &StringPair{"Robert L.", "Stevenson"}
	david := StringPair{"David", "Balfour"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", raw)
	}
}
