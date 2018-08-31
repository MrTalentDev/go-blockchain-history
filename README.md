## About
This is on open source project which aims to make easy solution with timestamping data on blockchain

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





