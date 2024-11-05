# [`intersect`](https://owncloud.gwdg.de/index.php/s/G3ur0inxw8Dlymd)
This program is a CLI for `chr.Intersect` package. The program finds
homologous regions common for a set of closely related input
nucleotide sequences.

## Compile
Given that you have `git`, `golang`, and `make`:

	git clone https://github.com/IvanTsers/intersect
	cd intersect
	make

## Tutorial
Change to `tutorial`:

	cd tutorial

There are 10 sequences containing from 1 to 4 marker regions. We run
`intersect` using `t1.fasta` as a reference and sequences in `t` as
queries:

	intersect -r t1.fasta -d t

This will return a perfect intersection, that is, a region shared by
all sequences. It contains only one of the markers. We lower the
sensitivity threshold to `0.9`:

	intersect -r t1.fasta -d t -f 0.9

This returns two markers shared by at leat 9 out of 10 genomes. We
continue to lower the sensitivity of intersection and add switches to
print mutation positions:

	intersect -r t1.fasta -d t -f 0.6 -n

This time the intersection includes three markers, two of them are
separated with mismatches shown as `N`. These regions are found in at
least 6 out of 10 genomes. We lower the sensitivity threshold again
and add the `-verbose` flag to print some statistics:

	intersect -r t1.fasta -d t -f 0.2 -n -verbose

The intersection now contains all four markers.

## To reproduce results of a `phylonium -p` run
use these parameters:

    intersect -r <reference> -d <targets> -s -clean-reference -clean-queries
	
 - the default sensitivity threshold `-f` is 1.0 (perfect intersection)
 - `-s` prints the positions of segregating sites into headers
 - `-clean-reference` removes non-ATGC nucleotides from the reference
 - `-clean-queries` removes non-ATGC nucleotides from the queries

Use `-h` to see all available options.