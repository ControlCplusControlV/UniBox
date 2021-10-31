#!/usr/bin/Rscript

install.packages("devtools")
install.packages("Rserve")
devtools::install_github("Omni-Analytics-Group/uniswappeR")
library(reticulate)
require('Rserve')
