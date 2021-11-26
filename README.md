## Shannon
Package `shannon` provides one function: `Shannon`. 

`func Shannon(r io.Reader) (entropy float64, err error)`

Shannon calculates the shannon entropy of a source `r` and
returns a bits/byte `entropy` value between 0.0 and 8.0,
0 being no, and 8 being maximal entropy.

On error, 0 is returned for the 'entropy' value.

#

This package is "[unlicense](https://choosealicense.com/licenses/unlicense/)d" because I'm pretty sure I can't claim copyright on a simplistic implementation of an algorithm published in 1948.