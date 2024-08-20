package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	//"github.com/ivantsers/alan/util"
	"github.com/evolbioinf/esa"
	"github.com/evolbioinf/sus"
	"github.com/ivantsers/ancs"
	"github.com/ivantsers/fasta"
	"os"
)

func main() {
	var optR = flag.String("r", "", "reference sequence")
	var optA = flag.Int("a", 0, "minimum anchor length")
	var optP = flag.Float64("p", 0.05,
		"p-value for a non-random shustring")
	var optVerb = flag.Bool("verbose", false, "toggle verbose mode")
	var optN = flag.Bool("n", false, "print segregation sites (Ns)"+
		"in the output sequences")
	var optS = flag.Bool("s", false, "print segregation site ranges"+
		"to stderr")
	var optF = flag.Float64("f", 1.0, "intersection sensitivity threshold")
	var optV = flag.Bool("v", false, "print version and "+
		"program information")
	u := "alan [option]..."
	p := "Find approximate local alignment"
	e := "alan -s subject.fasta -q query.fasta"
	clio.Usage(u, p, e)
	flag.Parse()
	if *optV {
		fmt.Println("alan v0.2")
	}
	fileNames := flag.Args()
	if len(fileNames) < 2 {
		fmt.Fprintf(os.Stderr, "please specify at least two input files\n")
		os.Exit(1)
	}
	queryNames := []string{}
	if *optR == "" {
		fmt.Fprintf(os.Stderr, "please specify the reference sequence\n")
		os.Exit(1)
	} else {
		for _, fileName := range fileNames {
			if fileName != *optR {
				queryNames = append(queryNames, fileName)
			}
		}
	}
	f, _ := os.Open(*optR)
	subjectContigs := fasta.ReadAll(f)
	f.Close()
	subject := fasta.Concatenate(subjectContigs, 0)
	subject.Clean()
	rev := fasta.NewSequence(subject.Header(), subject.Data())
	rev.ReverseComplement()
	subject.SetData(append(subject.Data(), rev.Data()...))
	se := esa.MakeEsa(subject.Data())
	sl := subject.Length()
	sgc := subject.GC()
	var sma int
	if *optA > 0 {
		sma = *optA
	} else {
		sma = sus.Quantile(sl/2, sgc, 1.0-*optP)
	}
	if *optVerb {
		fmt.Fprintf(os.Stderr,
			"# minimum anchor length: %d\n", sma)
	}
	numQueries := len(queryNames)
	allH := []ancs.Seg{}
	allNs := make(map[int]bool)
	for _, q := range queryNames {
		f, _ = os.Open(q)
		queryContigs := fasta.ReadAll(f)
		f.Close()
		query := fasta.Concatenate(queryContigs, 0)
		query.Clean()
		h, ns := ancs.FindHomologies(query, subject, se, sma)
		ancs.SortByStart(h)
		allH = append(allH, ancs.ReduceOverlaps(h)...)
		for pos, _ := range ns {
			allNs[pos] = true
		}
	}
	sensitivity := *optF
	intersection := ancs.Intersect(allH, numQueries, sensitivity, sl/2)
	if *optVerb {
		fmt.Fprintf(os.Stderr,
			"# %d homologous segments(s)\n", len(intersection))
		fmt.Fprintf(os.Stderr,
			"# total length: %d\n", ancs.TotalSegLen(intersection))
		fmt.Fprintf(os.Stderr,
			"# %d segregation site(s)\n", len(allNs))
	}
	if *optS {
		ancs.PrintSegsiteRanges(allNs, intersection, os.Stderr)
	}
	result := ancs.SegToFasta(intersection, se, allNs, *optN)
	for _, seq := range result {
		fmt.Fprintf(os.Stdout, "%s\n", seq)
	}
}
