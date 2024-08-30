package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/ivantsers/chr"
	"github.com/ivantsers/fasta"
	"os"
)

func main() {
	var optR = flag.String("r", "", "reference sequence")
	var optD = flag.String("d", "", "directory of target sequences")
	optP := flag.Float64("p", 0.05,
		"p-value for a non-random matching shustring length")
	optVerb := flag.Bool("verbose", false, "toggle verbose mode")
	optN := flag.Bool("n", false, "print segregation sites (Ns)"+
		"in the output sequences")
	optS := flag.Bool("s", false, "print segregation site ranges"+
		"in the headers")
	optF := flag.Float64("f", 1.0, "intersection sensitivity threshold")
	optCleanR := flag.Bool("clean-reference", true,
		"remove non-ATGC nucleotides from the reference")
	optCleanQ := flag.Bool("clean-queries", true,
		"remove non-ATGC nucleotides from the queries")
	optOneBased := flag.Bool("one-based-output", true,
		"print one-based end-exclusive coordinates in the "+
			"output headers. The default coordinates "+
			"are zero-based end-inclusive.")
	u := "intersect [option]..."
	p := "Find common homologous regions in a set of genomes"
	e := "intersect -r subject.fasta -d query_dir"
	clio.Usage(u, p, e)
	flag.Parse()
	numFiles := 0
	dirEntries, err := os.ReadDir(*optD)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"error reading %v: %v\n", *optD, err)
		os.Exit(1)
	}
	numFiles = len(dirEntries)
	if numFiles < 1 {
		fmt.Fprintf(os.Stderr, "the target dir contains no files\n")
		os.Exit(1)
	}
	if *optR == "" {
		fmt.Fprintf(os.Stderr, "please specify the reference sequence\n")
		os.Exit(1)
	}
	f, _ := os.Open(*optR)
	referenceContigs := fasta.ReadAll(f)
	f.Close()
	if *optP > 1 || *optP < 0 {
		fmt.Fprintf(os.Stderr,
			"can't use %v as a sensitivity threshold\n", *optF)
		os.Exit(1)
	}
	pval := 1.0 - *optP
	if *optF > 1.0 || *optF <= 0.0 {
		fmt.Fprintf(os.Stderr,
			"can't use %v as a sensitivity threshold\n", *optF)
		os.Exit(1)
	}
	parameters := chr.Parameters{
		Reference:       referenceContigs,
		TargetDir:       *optD,
		Threshold:       *optF,
		ShustrPval:      pval,
		CleanSubject:    *optCleanR,
		CleanQuery:      *optCleanQ,
		PrintSegSitePos: *optS,
		PrintN:          *optN,
		PrintOneBased:   *optOneBased,
	}
	isc := chr.Intersect(parameters)
	if *optVerb {
		totalLen := 0
		for _, seq := range isc {
			totalLen += seq.Length()
		}
		fmt.Fprintf(os.Stderr, "# Intersected sequences from %d"+
			" n files\n", numFiles+1)
		fmt.Fprintf(os.Stderr,
			"#  common homologous region(s): %d\n", len(isc))
		fmt.Fprintf(os.Stderr,
			"#  intersection's total length: %d\n", totalLen)
	}
	for _, seq := range isc {
		fmt.Fprintf(os.Stdout, "%s\n", seq)
	}
}
