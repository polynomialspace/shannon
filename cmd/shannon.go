package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/polynomialspace/shannon"
)

type result struct {
	filename        string
	err             error
	entropy         float64
	compressability float64
}

func (res result) Print(w *tabwriter.Writer) {
	if res.err != nil {
		fmt.Fprintf(w, "%v\tn/a\tn/a\n", res.err)
		return
	}
	fmt.Fprintf(w, "%s\t%.2f\t%.2f\n", res.filename, res.entropy, res.compressability)
}

func main() {
	flag.Parse()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 8, 2, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "filename\tbits/Byte\t~%compressible")
	defer w.Flush()

	if flag.NArg() < 1 {
		var res result
		defer res.Print(w)

		entropy, err := shannon.Shannon(os.Stdin)
		if err != nil {
			res = result{err: err}
			return
		}
		pct := 100.0 - ((entropy / 8.0) * 100.0)

		res = result{
			filename:        "stdin",
			err:             nil,
			entropy:         entropy,
			compressability: pct,
		}
		return
	}

	results := make(chan result, 1)

	wg := new(sync.WaitGroup)
	wg.Add(flag.NArg())

	for _, arg := range flag.Args() {
		go func(arg string, wg *sync.WaitGroup, c chan result) {
			defer wg.Done()
			f, err := os.Open(arg)
			if err != nil {
				c <- result{err: err}
				return
			}
			defer f.Close()

			entropy, err := shannon.Shannon(f)
			if err != nil {
				c <- result{err: err}
				return
			}

			pct := 100.0 - ((entropy / 8.0) * 100.0)

			c <- result{
				filename:        f.Name(),
				err:             nil,
				entropy:         entropy,
				compressability: pct,
			}
		}(arg, wg, results)
	}

	go func(c chan result, wg *sync.WaitGroup) {
		wg.Wait()
		close(c)
	}(results, wg)

	for res := range results {
		res.Print(w)
	}
}
