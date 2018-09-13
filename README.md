## About
This is on open source project which aims to make easy solution with timestamping data on blockchain

## Architecture 

How it works: 

Blockchain server - server with SMT (Sparse Merkle tree) and Ethereum connetion to Contract

Product server - server with DB which push merkle hash of assets to Blockchain server and stores this file uncompressed

Client can send file to Product server and download it. Product server can send file to client and send a Merklee Proof to client. Client can verify data with provided merkle proof.



## About ENV 

```
PVT_KEY: Private key at Rinkeby network

CONTRACT_ADDRESS: target contract address

LOGIN_DB: Login of Mongo

PASSWORD_DB: Password of Mongo

IP: IP/URL of Mongo

LOGIN: Login for basic auth from users

PASSWORD: Password for basic auth from users

```
## Deploying

```
mkdir docker

cd docker 

sudo nano docker-compose.yml

docker swarm init

docker stack deploy -c docker-compose.yml poa_hist
```
## More
There are 3 servers:

Blockchain validator server

Product storage server

Frontend example





