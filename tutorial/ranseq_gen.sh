#!/bin/bash
# This script drives the ranseq program
# (https://github.com/evolbioinf/biobox)
# to generate ten 3kb-long target sequences
# containing one to four markers:
# m10 (0..100)     in t1 to t10;
# m9  (900..1300)  in t1 to t9;
# m6  (1400..1500) in t1 to t6;
# m2  (2950..3000) in t1 to t2.

mkdir markers
mkdir t

# Generate headers
for i in $(seq 10)
do
    printf ">t$i\n" > t/t$i.fasta
done

# Generate m10
printf ">m10\n" > markers/m10.fasta
m=$(ranseq -l 100 | grep -v ">" | tr -d '\n')
printf "%s" $m >> markers/m10.fasta

# Append m10 to ten targets,
# Then append 799 random nucleotides
for i in $(seq 10)
do
    printf "%s" $m >> t/t$i.fasta
    ranseq -l 799 | grep -v ">" | tr -d '\n' >> t/t$i.fasta
done

# Generate m9
printf ">m9\n" > markers/m9.fasta
m=$(ranseq -l 401 | grep -v ">" | tr -d '\n')
printf "%s" $m >> markers/m9.fasta

# Append m9 to nine targets,
# Then append 99 random nucleotides
# t10 gets 500 random nucleotides
for i in $(seq 9)
do
    printf "%s" $m >> t/t$i.fasta
    ranseq -l 99 | grep -v ">" | tr -d '\n' >> t/t$i.fasta
done
ranseq -l 500 | grep -v ">" | tr -d '\n' >> t/t10.fasta

# Generate m6
printf ">m6\n" > markers/m6.fasta
m=$(ranseq -l 101 | grep -v ">" | tr -d '\n')
printf "%s" $m >> markers/m6.fasta

# Append m6 to six targets,
# Then append 1449 random nucleotides
# t7 to t10 get 1550 random nucleotides
for i in $(seq 6)
do
    printf "%s" $m >> t/t$i.fasta
    ranseq -l 1449 | grep -v ">" | tr -d '\n' >> t/t$i.fasta
done

for i in $(seq 7 10)
do
    ranseq -l 1550 | grep -v ">" | tr -d '\n' >> t/t$i.fasta
done

# Generate m2
printf ">m2\n" > markers/m2.fasta
m=$(ranseq -l 51 | grep -v ">" | tr -d '\n')
printf "%s" $m >> markers/m2.fasta

# Append m2 to two targets,
# t3 to t10 get 51 random nucleotides
for i in $(seq 2)
do
    printf "%s" $m >> t/t$i.fasta
done

for i in $(seq 3 10)
do
    ranseq -l 51 | grep -v ">" | tr -d '\n' >> t/t$i.fasta
done

mv t/t1.fasta .
