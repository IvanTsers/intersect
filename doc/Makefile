NAME = intersect
all: $(NAME)Doc.pdf
	latex $(NAME)Doc
	#bibtex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	dvipdf -dALLOWPSTRANSPARENCY $(NAME)Doc
$(NAME)Doc.pdf: $(NAME)Doc.tex $(NAME).tex 
$(NAME).tex: ../$(NAME).org
	bash ../scripts/org2nw ../$(NAME).org       | awk -f ../scripts/preWeave.awk | noweave -n -x > $(NAME).tex
clean:
	rm -f $(NAME).tex $(NAME)_test.tex *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc 
