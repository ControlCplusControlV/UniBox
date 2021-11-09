# 🦄 UniBox 📦

Built out for the Unicode 2021 Hackathon!

[UniCode Showcase](https://showcase.ethglobal.com/unicode/unibox)

[Installation Help](https://vimeo.com/643752929)

[Usage Help](https://vimeo.com/643752888)
## It's Uniswap in a Box!

Uniswap allows users to interact with the Uniswap Protocol directly through a self-hosted docker container. This improves security for those who want to keep their swapping as far from broswers and some of the potential risks they pose. UniBox allows users to also take advantage of the Uniswap Protocol in new ways, as it provides an analytics interface that utilizes OmniAnalytics' uniswappeR package to fetch data. UniBox also provides users with an interface to build out and execute automated trading strategies.

### Analytics

UniBox uses uniswappeR to collect data analytics on uniswappeR to pull data directly from the Ethereum network so that it can provide high accuracy graphs to inform trading decisions from within the command line interface. 

### Trading Strategies

UniBox also exposes the data feeds it collects from the dataAggregator package. The swap package is also exposed to the strategy file, allowing users to build out complex trading strategies based on the provided analytics. Since UniBox is self-hosted users can have piece of mind knowing that their trading bot is secure,

### Command Line Interface

UniBox also has a powerful command line interface that gives user's access to a swap interface where they can swap tokens similarly to from a web interface, in addition to an analytics window where users can track stats like the performance of LP pools, and overall protocol statistics like TVL and Volume.

#### Protocol Support

It's important to note UniBox only supports UniSwap V2 but we'd be happy to add V3 support in the future, however as uniswappeR only supports Uniswap V2 we wanted to focus on supporting the protocol we can provide users with data from. There are also limitations with web3-go that we were unable to patch in time to provide support for web3-go.

## Installation

Clone this repo, cd into it and run
```
docker build -t unibox .
```
this may take around 10 minutes to build, this is to be expected as its a large branch. 
Once it is finished attach and login to the container, then run

```
cd src
go run .
```

and you will be taken into the UniBox interface!

To enable swapping you have to enter a private key, a public key, and a node url in to the swap.go file.
