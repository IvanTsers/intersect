# [`Intersect`](https://owncloud.gwdg.de/index.php/s/G3ur0inxw8Dlymd)
This program is a demo CLI for `chr.Intersect` package. The program
finds homologous regions common for a set of closely related input
nucleotide sequences.

## Compile
Given that you have `git`, `golang`, and `make`:

	git clone https://github.com/IvanTsers/intersect
	cd intersect
	make

## Tutorial
To run `intersect`, specify a reference (a `.fasta` file)
and a query directory (targets to intersect the reference with), and
set a sensitivity threshold `f`, that is, a minimal fraction of
queries to be included in the intersection:

	intersect -r r.fasta -d t -f 0.95

Note that this directory should not contain the reference. Think of
the input as components of a `fur` database:

```bash
# stan.db is a fur database built for 4 target genomes, 
# t1.fasta became a representative target (r.fasta)
stan.db
├── dblist.txt
<...>
├── r.fasta
└── t
    ├── t2.fasta
    ├── t3.fasta
    └── t4.fasta
```

To reproduce results of a `phylonium -p` run, use these parameters:

	intersect -r r.fasta -d t -s -clean-reference -clean-queries -one-based-output
	
 - the default sensitivity threshold `-f` is 1.0 by default (complete intersection)
 - `-s` prints segregation site total number an positions into headers
 - `-clean-reference` removes non-ATGC nucleotides from the reference
 - `-clean-queries` removes non-ATGC nucleotides from the queries
 - `-one-based-output` converts headers' coordinates to 1-based end-exclusive

Use `-h` to see all available options. Toggle the `-verbose` mode to
print some statistics to `stderr`.