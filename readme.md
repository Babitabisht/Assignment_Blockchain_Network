The objective is to create a decentralized application to `proof existence` of a document at a particular time. Here is what the dapp should do:

1. A simple smart contract with two functionality. One should accept a hash and meta data of a file and store in the blockchain. The second functionality is to be able to take a hash of a file and give back the meta data and time at which the hash was registered.
2. Include a feature so that at any time we should be able to retrieve all the registered files and their information.
3. Create an API Server which exposes APIs for these functionality.


Technology stack :

Hyperledger Fabric ( Single Organization with 3 Peers )
Node.js ( For Rest API server )
Chaincode