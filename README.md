## About
This is on open source project which aims to make easy solution with timestamping data on blockchain

## Swagger 
https://history.bankex.team:3001/

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

project structure: TODO: Explain project structure with refactoring 

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

**Route:** /update/:assetId/:hash

**Description:** Allow to add new asset to assetId. Returns txNumber of this hash, timesamp

**Return:** JSON 

```
{
   "assetId": Id of current asset chaid
   "timestamp": UNIX format time when server got hash of file
   "txNumber": Number of asset from assetId
}
```
---

**GET**

**Route:** /get/:assetId/:txNumber

**Description:** Return asset hash by assetId and txNumber

**Return:** JSON 

```
{
   "assets": current asset
}
```
---

**GET**

**Route:** /proof/:assetId/:txNumber/:hash/:timestamp

**Description:** Return list of merkle proofs by assetId, txNumber, hash, timestamp

**Return:** JSON 

**More about return:** Merkle proof for assetId, txNumber, hash, timestamp (Actually send a JSON File with two arrays **Data** and **Info**
                       
**Data** is a list of merkle proofs leaves from end to start (256 Hashes of type Buffer)
                       
**Info** has parameters: 
1. Key - array key
2. Hash - array value
3. Root - current merkle tree Root Hash

Response looks like:
```
{
  {
      "Data": [
          {
              "Hash": "QGTfJZ5sF0U5U0nwQDI0q+FXE7p+87DGZ1bhijbapPU="
          },
          {
              "Hash": "hBp5I5E3E57YRPCIRziHXVdlPSF3nWCNKmRRcS+nQZc="
          },
          {
              "Hash": "SZEJoTogdMeznCpdpIIqXM+ztBfXnLxnFCOUYTl4Jm4="
          }
      ], 
      "Info": {
              "Key": "VCRbbhhUHqe//lRV3RDBawTnATBTeZNsm9FQtwR9JMw=",
              "Hash": "2TmoNwyUYfmxtInasAyC9xyKM7hcZq9MokNwAoQxwek=",
              "Root": "5JX8dfEibcncG2fGp0YcG5UTY9LgrNdQoq4TWL8WpUs="
          }
}
```



---

**GET**

**Route:** /list

**Description:** Return all assets info

**Return:** JSON 

```
{
  {
      "assets": [
          {
              "_id": "5b869ee5ca2985e06552a49d",
              "data": "",
              "hash": "qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c=",
              "created_on": 1535549157514,
              "updated_on": 1535549157514,
              "assetId": "bf",
              "txNumber": 0,
              "assets": {
                  "0": "ludYELf+UZ3ZL2o/chcLAMCoqVU/nHZaPMaB6vfuqzg="
              },
              "assetsTimestamp": {
                  "0": 1535549157514
              }
          }
      ]
}
```


## DevOps

There are 3 servers here 

**Product server:** works with MongoDb and is working on port 3000

**Blockchain server:** works with MongoDb and is working on port 8080

**Client server:** works with 2 servers on port 7070

**NOTE**: 

It's better to use more than 1 domain for project.

Blockchain server is just a tool, so there is no point to change it.

But product and client servers can be different. You can only run Blokchain server and make you own product and client server.

You just need to have a Verify Function implemented in client part to verify merkle proofs. 

It's **here** https://github.com/BANKEX/poa-history/blob/client/assets/download/index.js


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





