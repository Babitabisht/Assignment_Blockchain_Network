
version=$1
chaincode=$2
folder=$3
set -x
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
set +x

set -x
peer chaincode install -n "$chaincode" -v "$version" -p github.com/"$folder"/go
set +x


