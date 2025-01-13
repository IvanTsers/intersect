SRC1 =   intersect.tex ref2query.tex
SRC2 =   header.tex intersectDoc.tex

NAME = intersect
all: $(NAME)Doc.pdf
	latex $(NAME)Doc
	#bibtex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	dvipdf -dALLOWPSTRANSPARENCY $(NAME)Doc

$(NAME)Doc.pdf: $(NAME)Doc.tex $(SRC1) $(SRC2) 

#$(NAME).tex: ../$(NAME).org
#	bash ../scripts/org2nw ../$(NAME).org | awk -f ../scripts/preWeave.awk | noweave -n -x > #$(NAME).tex

%.tex: ../*/%.org $(NAME)Doc.tex
	bash ../scripts/org2nw $< | awk -f ../scripts/preWeave.awk | noweave -n -x | sed 's/_test/\\_test/g' > $@

clean:
	rm -f $(NAME).tex $(NAME)_test.tex $(SRC1) *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc 
