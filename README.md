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



//TODO:

There is a VerifyProof function at : server/ethercrypto/smt/proofs.go

```cgo
func VerifyProof(proof [][]byte, root []byte, key []byte, value []byte, hasher hash.Hash) bool {
    hasher.Write(key)
    path := hasher.Sum(nil)
    hasher.Reset()

    hasher.Write(value)
    currentHash := hasher.Sum(nil)
    hasher.Reset()

    if len(proof) != hasher.Size() * 8 {
        return false
    }

    for i := hasher.Size() * 8 - 1; i >= 0; i-- {
        node := make([]byte, hasher.Size())
        copy(node, proof[i])
        if len(node) != hasher.Size() {
            return false
        }
        if hasBit(path, i) == right {
            hasher.Write(append(node, currentHash...))
            currentHash = hasher.Sum(nil)
            hasher.Reset()
        } else {
            hasher.Write(append(currentHash, node...))
            currentHash = hasher.Sum(nil)
            hasher.Reset()
        }
    }

    return bytes.Compare(currentHash, root) == 0
}
```

I need to rewrite it on JS 

with keccak256() hasher

Here there is some example how it must looks like

server/assets/index.js

```js
const left = 0;
const right = 1;

function verifyProof(data, key, hash, root) {
    keyHash = getHash(key);
    dataHash = getHash(hash);

    for (let i = 255; i >= 0; i--) {
        var node = [255];
        node[i] = (data[i]);
        if (hasBit(path, i) == right) {
            dataHash = getHash(node.concat(dataHash))

        } else {
            dataHash = getHash(dataHash.concat(node))
        }
    }
    return equal(dataHash, root)
}

function hasBit(data, position) {
    if ((data[position / 8]) & (1 << (uint(position) % 8)) > 0) {
        return 1
    }
    return 0
}

/**
 * Allows to hashing data via keccak256
 * @param data Any set of data
 * @return {*} Hash of data
 */
function getHash(data) {
    const web3 = new Web3();
    return web3.utils.keccak256(data);
}
```


How it works:

1. Make a GET request http://localhost:80/proof/assetId/txNumber/hash
Where hash is keccak256() of file which was uploaded

2. Receive a JSON
JSON consists of two things: Info and Data

Info : 
key - []byte value is a key of map[string][]byte
Hash - []byte  value is a value if map map[string][]byte
Root - []byte  value is Merkle Root of Tree 

Data :
Array of []byte  

3. Using JSON with verifyProof return answer that everything is OK

Notes: 

1. verifyProof everything is in []byte values

2. Later we need to get Root from Solidity contract from client

3. I can send all values in []byte type



