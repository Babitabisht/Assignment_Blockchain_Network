docker rm -f $(docker ps -aq)
# docker volume prune
sudo rm -rf artifacts/backup_*  && 

sudo rm -rf artifacts/couchdb0 artifacts/couchdb1 artifacts/couchdb2 artifacts/couchdb3 &&

sudo rm -rf fabric-client-kv-org*
