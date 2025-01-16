EXE = ref2query

all: $(EXE)

$(EXE): $(EXE).go
	go build $(EXE).go
$(EXE).go: $(EXE).org
	bash ../scripts/org2nw $(EXE).org | notangle -R$(EXE).go | gofmt > $(EXE).go

#.PHONY: doc clean

#doc:
#	make -C doc

clean:
	rm -f $(EXE) *.go
#	make clean -C doc
