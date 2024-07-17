package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	//"github.com/ivantsers/alan/util"
	"github.com/evolbioinf/esa"
	"github.com/ivantsers/ancs"
	"github.com/ivantsers/fasta"
	"os"
)

func main() {
	var optS = flag.String("s", "", "subject sequence")
	var optQ = flag.String("q", "", "query sequence")
	var optA = flag.Int("a", 0, "minimum anchor length")
	var optP = flag.Float64("p", 0.025,
		"p-value for a non-random shustring")
	var optVerb = flag.Bool("verbose", false, "toggle verbose mode")
	var optN = flag.Bool("n", false, "print segregation sites (Ns)"+
		"in the output sequences")
	var optR = flag.Bool("r", false, "print segregation site ranges"+
		"to stderr")
	var optV = flag.Bool("v", false, "print version and "+
		"program information")
	u := "alan [option]..."
	p := "Find approximate local alignment of two sequences"
	e := "alan -s subject.fasta -q query.fasta"
	clio.Usage(u, p, e)
	flag.Parse()
	if *optV {
		fmt.Println("alan v0.1")
	}
	if !(*optQ != "" && *optS != "") {
		fmt.Fprintf(os.Stderr, "please give names "+
			"of subject and query files.\n")
		os.Exit(1)
	}
	f, _ := os.Open(*optS)
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
		sma = ancs.MinAncLen(sl/2, sgc, 1.0-*optP)
	}
	if *optVerb {
		fmt.Fprintf(os.Stderr,
			"# minimum anchor length: %d\n", sma)
	}
	f, _ = os.Open(*optQ)
	queryContigs := fasta.ReadAll(f)
	f.Close()
	query := fasta.Concatenate(queryContigs, 0)
	query.Clean()
	homologies, segsites := ancs.FindHomologies(query, se, sl, sma)
	ancs.SortByStart(homologies)
	homologies = ancs.ReduceOverlaps(homologies)
	if *optVerb {
		fmt.Fprintf(os.Stderr,
			"# %d homologous segments(s)\n", len(homologies))
		fmt.Fprintf(os.Stderr,
			"# total length: %d\n", ancs.TotalSegLen(homologies))
		fmt.Fprintf(os.Stderr,
			"# %d segregation site(s)\n", len(segsites))
	}
	if *optR {
		ancs.PrintSegsiteRanges(segsites, os.Stderr)
	}
	result := ancs.SegToFasta(homologies, se, segsites, *optN)
	for _, seq := range result {
		fmt.Fprintf(os.Stdout, "%s\n", seq)
	}
}
