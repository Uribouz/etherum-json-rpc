# etherum-json-rpc

### Name. Vatcharit Opas

## Assignment:
You are tasked with building a Golang system that monitors incoming and outgoing transactions of specified Ethereum addresses. The system should have the following features:

    1. The ability to read incoming and outgoing transactions of specified Ethereum addresses.
    2. The ability to store transaction data in a database.
        a. Don’t need to extract any data. Just store the whole transaction data.
    3. The ability to scale to handle a large number of Ethereum addresses.

## Requirements:
    1. Use the Golang programming language to implement the system.
    2. Use the Ethereum JSON-RPC API to interact with the Ethereum network.
    3. Use a database of your choice to store transaction data.
    4. Implement a scalable architecture that can handle the monitoring of a large number of Ethereum addresses.
    5. Ensure the system is secure, reliable, and scalable.

## Submission Guidelines:
    1. Include clear instructions on how to run the system.
    2. Include a brief description of your approach and any design decisions you made.
    3. Submit the link to your repository (or zip file) and any necessary credentials to the interviewer.

This assignment will test the interviewee's knowledge and skills in building a Golang system that
monitors incoming and outgoing transactions of specified Ethereum addresses using the JSONRPC API. Additionally, the submission should demonstrate their ability to design and implement
a scalable architecture, store data in a database, send notifications, and write clean and
maintainable code, as well as their understanding of security, scalability and reliability concerns
in building such a system.


## Installation guidelines: (For windows)
    1. Install Go
    2. Install go-ethereum package 
        $ go get -u github.com/ethereum/go-ethereum/ethclient
    3. Install tdm-gcc
        https://jmeubank.github.io/tdm-gcc/
        
## Useful links
- [node-service-ankr](https://www.ankr.com/rpc/eth/)
- [json-rpc-methods](https://ethereum.org/en/developers/docs/apis/json-rpc/#json-rpc-methods)