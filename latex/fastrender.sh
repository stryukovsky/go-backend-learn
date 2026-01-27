#!/bin/bash

xelatex -interaction=nonstopmode diploma.tex 
biber diploma
xelatex -interaction=nonstopmode diploma.tex 
xelatex -interaction=nonstopmode diploma.tex 

# latexmk -xelatex -silent diploma.tex

