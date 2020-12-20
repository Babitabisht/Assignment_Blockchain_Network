**Prerequisites**
1. Docker
1. hyperledger-fabric  docker images (v1.4)
1. fabric binaries  (v1.4)
1. golang  (go1.12.7)
1. nodejs  (v8.14.1)


**How to start**
1. ./runApp.sh
1. ./testAPIs.sh



**Where to find smart contract**
Inside artifacts -> src-> github.com

**Requirements**

1. A simple smart contract with two functionality. One should accept a hash and meta data of a file and store in the blockchain. The second functionality is to be able to take a hash of a file and give back the meta data and time at which the hash was registered.
2. Include a feature so that at any time we should be able to retrieve all the registered files and their information.
3. Create an API Server which exposes APIs for these functionality.


**Technology stack :**

Hyperledger Fabric ( Single Organization with 3 Peers )
Node.js ( For Rest API server )
Chaincode


