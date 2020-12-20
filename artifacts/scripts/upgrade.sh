#!/bin/bash

version=$(date +%s)
CC=$1
FOLDER="example_cc"
if [ -z "$CC" ]; then
CC="mycc"
FOLDER="example_cc"
fi
if [[ $CC == "mycc" ]]; then
    FOLDER="example_cc"
fi

if [[ $CC == "mycc2" ]]; then
echo "inids"
    FOLDER="module1"
fi

if [[ $CC == "mycc3" ]]; then
echo "qw"
    FOLDER="module8"
fi

if [[ $CC == "batteryInventoryAndOrderManagement" ]]; then
echo "batteryInventoryAndOrderManagement"
    FOLDER="batteryInventoryAndOrderManagement"
fi

echo "*********************************************************"
echo "     installing chaincode : $CC version $version  folder $FOLDER    "
echo "*********************************************************"     


set -x
docker exec -t cli-1 bash scripts/upgrade1.sh $version $CC $FOLDER
set +x

set -x
docker exec -t cli-1 bash scripts/upgrade2.sh $version $CC $FOLDER
set +x


