package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/ivantsers/chr"
	"github.com/ivantsers/fastautils"
	"os"
)

func main() {
	optR := flag.String("r", "", "reference sequence")

	optP := flag.Float64("p", 0.975,
		"threshold p-value for a shustring length")

	optN := flag.Bool("n", false, "print segregating sites (Ns) "+
		"in the output sequences")

	optF := flag.Float64("f", 1.0,
		"intersection sensitivity threshold")

	optT := flag.Int("t", 1, "number of threads")
	u := "intersect [option]..."
	p := "Find common homologous regions in a set of genomes"
	e := "intersect -r subject.fna query1.fna query2.fna ..."
	clio.Usage(u, p, e)
	flag.Parse()
	if *optR == "" {
		fmt.Fprintf(os.Stderr,
			"please specify the reference sequence\n")
		os.Exit(1)
	}
	f, _ := os.Open(*optR)
	referenceContigs := fastautils.ReadAll(f)
	f.Close()
	if *optP > 1 || *optP < 0 {
		fmt.Fprintf(os.Stderr,
			"can't use %v as a sensitivity threshold\n",
			*optF)
		os.Exit(1)
	}
	if *optF > 1.0 || *optF <= 0.0 {
		fmt.Fprintf(os.Stderr,
			"can't use %v as a sensitivity threshold, "+
				"please use a value in the interval (0,1]",
			*optF)
		os.Exit(1)
	}
	parameters := chr.Parameters{
		Reference:  referenceContigs,
		QueryPaths: flag.Args(),
		Threshold:  *optF,
		ShustrPval: *optP,
		PrintN:     *optN,
		NumThreads: *optT,
	}
	isc := chr.Intersect(parameters)
	for _, seq := range isc {
		fmt.Fprintf(os.Stdout, "%s\n", seq)
	}
}
