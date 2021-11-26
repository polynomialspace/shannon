package shannon

import (
	"bufio"
	"io"
	"math"
)

// Shannon calculates the shannon entropy of a source 'r' and
// returns a bits/byte 'entropy' value between 0.0 and 8.0,
// 0 being no, and 8 being maximal entropy.
//
// On error, 0 is returned for the 'entropy' value.
func Shannon(r io.Reader) (entropy float64, err error) {
	br := bufio.NewReader(r)
	var frequency [256]int64
	var len int64
loop:
	for ; ; len++ {
		switch b, err := br.ReadByte(); err {
		case io.EOF:
			break loop
		case nil:
			frequency[b]++
		default:
			return 0, err
		}
	}

	for _, freq := range frequency {
		if freq > 0 {
			f := float64(freq) / float64(len)
			entropy += f * math.Log2(f)
		}
	}

	return -entropy, nil
}
