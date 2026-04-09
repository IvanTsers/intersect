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

	intersect -r t1.fasta t/*

This will return a perfect intersection, that is, a region shared by
all sequences. It contains only one of the markers. We lower the
sensitivity threshold to `0.9`:

	intersect -r t1.fasta -f 0.9 t/*

This returns two markers shared by at leat 9 out of 10 genomes. We
continue to lower the sensitivity of intersection and add switches to
print mutation positions:

	intersect -r t1.fasta -f 0.6 -n t/*

This time the second marker includes an additional region separated
with mismatches shown as `N`. These regions are found in at least 6
out of 10 genomes. We lower the sensitivity threshold again:

	intersect -r t1.fasta -f 0.2 -n t/*

The intersection now contains all four markers from the directory
`markers`.