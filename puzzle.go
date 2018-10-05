package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jonbodner/sidecar"
)

func main() {
	allVals := [][]float64{
		[]float64{3, 3, 8, 8},
		[]float64{8, 8, 3, 3},
		[]float64{8, 3, 3, 8},
		[]float64{3, 8, 8, 3},
		[]float64{8, 3, 8, 3},
		[]float64{3, 8, 3, 8},
	}
	for _, vals := range allVals {
		results := calc(vals)
		for _, v := range results {
			total := sidecar.Calc(v)
			if math.Abs(total-24) < 0.000001 {
				fmt.Println(v, "=", total)
			}
		}
	}
}

func calc(vals []float64) []string {
	if len(vals) == 1 {
		return []string{fmt.Sprintf("%d", int(vals[0]))}
	}
	head := vals[0]
	rest := vals[1:]
	var out []string
	results := calc(rest)
	for _, v := range results {
		for j := 0; j <= 1; j++ {
			openParens := strings.Repeat("( ", j)
			closeParens := strings.Repeat(" )", j)
			out = append(out, fmt.Sprintf("%s%d + %s%s", openParens, int(head), v, closeParens))
			out = append(out, fmt.Sprintf("%s%d - %s%s", openParens, int(head), v, closeParens))
			out = append(out, fmt.Sprintf("%s%d * %s%s", openParens, int(head), v, closeParens))
			out = append(out, fmt.Sprintf("%s%d / %s%s", openParens, int(head), v, closeParens))
		}
	}
	return out
}
