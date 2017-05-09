
SRC=$(shell ls *.uml)
OBJ=$(SRC:%.uml=%.svg)

.SUFFIXES: .uml .svg

all: $(OBJ)

%.svg: %.uml
	plantuml -tsvg $<

clean:
	rm -f *.svg


