#!/bin/bash

pdflatex -interaction=nonstopmode diploma.tex 
bibtex diploma.aux 
pdflatex -interaction=nonstopmode diploma.tex 
pdflatex -interaction=nonstopmode diploma.tex 
