## About
This is on open source project which aims to make easy solution with timestamping data on blockchain

## FAQ

How to proof that file was uploaded with hash N and timestamp T
Upload:
1) Make a Sparse Merkle tree, where value is (N,T) (Before all, it's important to save N, T, Id of file)
2) Put merkle root to ethereum server 
Proove:
1) Download file and get N and T
2) Check that N = N saved before uploading 
3) Do the same for T 
4) Ask for Merkle proof from server
5) Get Merkle root from ethereum contract
6) check that merkle proof is correct ( it's a function with inputs: Hash file, timestamp file, assetId, txNumber - all these parameters client at the beggining)
7) if merkle proof is correct - than all is OK 

![image](https://raw.githubusercontent.com/BANKEX/poa-history/master/docs/info.svg?sanitize=true)


## Architecture 

How it works: 

Blockchain server - server with SMT (Sparse Merkle tree) and Ethereum connetion to Contract

Product server - server with DB which push merkle hash of assets to Blockchain server and stores this file uncompressed

Client can send file to Product server and download it. Product server can send file to client and send a Merklee Proof to client. Client can verify data with provided merkle proof.

## Backend handlers 

**POST:**

**Route:** a/new/:assetId/:hash 

**Description:** Allow to create new AssetID with Hash. 

**Return:** JSON 

```
{
   "assetId": Id of current asset chaid
   "hash": hash of file what we've got from product server
   "merkleRoot": root of merkle tree at Ethereum
   "timestamp": UNIX format time when server got hash of file
   "txNumber": Number of asset from assetId
}
```

**Return:** JSON if Error 

```
{
    "Answer": "This assetId is already created"
}
```

**Example:** http://localhost:8080/a/new/testAsset/0293a80682dc2a192c683baf434dd67343cedd70

---

**POST:**
/update/:assetId/:hash
Allow to add new asset to assetId. Returns txNumber of this hash, timesamp

**Description:** Allow to add new asset by assetId

**Return:** JSON 
```

```
---

**GET**
/get/:assetId/:txNumber
Return asset hash by assetId and txNumber

---

**GET**
/proof/:assetId/:txNumber/:hash/:timestamp
Return list of merkle proofs

---

**GET**
/list
Return all assets info



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





