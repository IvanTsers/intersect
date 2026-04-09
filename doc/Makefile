NAME = intersect

# ---------- Helper scripts ----------

ORG2NW   := bash ../scripts/org2nw.sh
PREWEAVE := awk -f ../scripts/preWeave.awk

# ---------- Build the version string ----------

TAG = $(shell git tag -l | sort -V | tail -n1)
VERSION = $(shell git log --pretty=format:"%cs, ${TAG}-%h" -n 1)

# ---------- Basic weaving ----------

all: $(NAME)Doc.pdf
	echo $(VERSION) > version.txt
	latex $(NAME)Doc
#	bibtex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	dvipdf -dALLOWPSTRANSPARENCY $(NAME)Doc

$(NAME)Doc.pdf: $(NAME)Doc.tex $(NAME).tex

$(NAME).tex: ../$(NAME).org
	$(ORG2NW) ../$(NAME).org | $(PREWEAVE) | noweave -n -x > $(NAME).tex

# ---------- Basic make subcommands ----------

.PHONY: clean

clean:
	rm -f $(NAME).tex *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc version.txt
