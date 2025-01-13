progs = intersect ref2query

all:

	test -d bin || mkdir bin
		for prog in $(progs); do \
			make -C $$prog; \
			cp $$prog/$$prog bin; \
		done

doc:
	make -c doc

clean:

	for prog in $(progs) doc; do \
		make clean -C $$prog; \
	done
	rm -f bin/*
