FROM golang:latest

RUN apt-get update -y && apt-get upgrade -y

RUN apt-get install r-base -y
RUN apt-get install build-essential libcurl4-gnutls-dev libxml2-dev libssl-dev -y

# Now that R is setup its time to setup all the needed golang packages
COPY setupUniswappe.R .
COPY RServe.R .
COPY src/* src/
RUN Rscript setupUniswappe.R

RUN go get github.com/senseyeio/roger
RUN go get github.com/0xKitsune/go-web3@unibox

CMD [ "/usr/bin/nohup", "/usr/bin/Rscript", "/go/RServe.R", "&" ]
