#!/bin/bash

xelatex -interaction=nonstopmode diploma.tex 
bibtex diploma.aux 
xelatex -interaction=nonstopmode diploma.tex 
xelatex -interaction=nonstopmode diploma.tex 
