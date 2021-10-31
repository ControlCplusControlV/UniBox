#!/usr/bin/Rscript

install.packages("devtools")
install.packages('Rserve',,"http://rforge.net/",type="source")
devtools::install_github("Omni-Analytics-Group/uniswappeR")
require('Rserve')
library(reticulate)
library(uniswappeR)
