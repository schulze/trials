NOTANGLE = notangle 
NOWEAVE = noweave -index -delay
# NOWEAVE = noweave -autodefs lisp -index -delay 

all: turingmacro

turingmacro: turingmacro.scm.nw
	@ ${NOTANGLE} turingmacro.scm.nw > turingmacro.scm
	@ ${NOWEAVE} turingmacro.scm.nw > turingmacro.scm.tex
	@ pdflatex turingmacro.scm.tex
	@ # makeindex turingmacro.scm.idx -o turingmacro.scm.ind
	@ # pdflatex turingmacro.scm.tex
	
.PHONY: all clean

clean:
	@ rm -f turingmacro.scm.aux turingmacro.scm.log turingmacro.scm.toc \
		turingmacro.scm.ind turingmacro.scm.idx turingmacro.scm.out \
		turingmacro.scm.tex turingmacro.scm.dvi turingmacro.scm \
		turingmacro.scm.pdf
