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
    2. Install packages
        $ go get github.com/ethereum/go-ethereum/ethclient
        $ go get github.com/joho/godotenv
        $ go get github.com/stretchr/testify
        $ go get go.mongodb.org/mongo-driver
    3. Install tdm-gcc
        https://jmeubank.github.io/tdm-gcc/


## Instructions: How to run the system
    1. Initialize file ".env" eg. using format from ".env_example"
    2. Modify file "addresses.json" with list of addresses you want to monitor.
    3. Run the monitoring program with the following command
        $ go run main.go


## Useful links
- [node-service-ankr](https://www.ankr.com/rpc/eth/)
- [json-rpc-methods](https://ethereum.org/en/developers/docs/apis/json-rpc/#json-rpc-methods)

### Design Documentation
    1. Makes packages independent from each other
    2. Make packages as shallow as possible.
    3. Always use unit-tests to make systems reliable.
    4. Uses 'config' package to hide confidential information, to make the system secure.
    5. Use 'adapter' and 'mockdata' packages to mock input list of addresses, to make the system scalable.
    6. Connection with a host outside should only connect once.
    7. Use 'core' package as scenario runner.
    8. Add 'chunker' package, to handle large number of addresses by distribute the list of addresses to each worker equally.

### Note Challenges:

    1. package "ethclient" 
    Does not let you get information about sender("from") easily
    you need to use info about ChainID, to get latest signer, 
    and then use signer to get Sender.
        Ex:
            types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)

    2. Addresses in HEX
    Address cannot be compared normally, each digit have different capitalization.

        expected address from the assignment:
            0x28c6c06298d514db089934071355e5743bf21d60
      
        actual address data:
            0x28C6c06298d514Db089934071355E5743bf21d60
    
    I need to use strings.EqualFold() for this to work.

    3. Insert data to MongoDB
    Data in JSON format cannot insert into MongoDB directly,
    It needs to be converted into BSON format.

    4. Subscribe method cannot be used on 'https'
    Following function cannot be used when connected to https server
        func (*ethclient.Client).SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
    have to find another service node that supports wss endpoint.
