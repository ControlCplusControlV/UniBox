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
RUN go get github.com/D-Cous/go-web3@unibox

RUN mkdir /go/src/strategy
RUN mkdir /go/src/swap
RUN mkdir /go/src/dataAggregator
RUN mkdir /go/src/abi

RUN mv /go/src/strategy.go /go/src/strategy
RUN mv /go/src/swap.go /go/src/swap
RUN mv /go/src/dataAggregator.go /go/src/dataAggregator

RUN mv /go/src/erc20ABI.json /go/src/abi
RUN mv /go/src/uniswapRouterV3ABI.json /go/src/abi
RUN mv /go/src/uniswapRouterV2ABI.json /go/src/abi


CMD [ "/usr/bin/nohup", "/usr/bin/Rscript", "/go/RServe.R", "&" ]
