FROM golang:latest

RUN apt-get install r-base -y
RUN apt-get install build-essential libcurl4-gnutls-dev libxml2-dev libssl-dev -y
RUN apt-get update && apt-get upgrade

# Now that R is setup its time to setup all the needed golang packages
COPY setupUniswappe.R .
RUN Rscript setupUniswappe.R

RUN go get github.com/senseyeio/roger
