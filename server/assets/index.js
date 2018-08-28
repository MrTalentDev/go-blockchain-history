function getData(assetId, txNumber, hash) {
    var xmlHttp = new XMLHttpRequest();
    // xmlHttp.open("GET", "/proof"+"/"+ assetId +"/"+ txNumber +"/"+ hash, false); // false for synchronous request
    a = "a";
    b = "0";
    c = "96e75810b7fe519dd92f6a3f72170b00c0a8a9553f9c765a3cc681eaf7eeab38";
    xmlHttp.open("GET", "/proof/" + a + "/" + b + "/" + c, false);
    xmlHttp.send(null);
    return xmlHttp.responseText;
}

function sort() {
    var data = getData(1, 1, 1);
    var s = JSON.parse(data);

    // var group = []
    //
    // for (var i = 0; i < s.length; i++) {
    //     group.push(s[i].Number)
    // }
    //
    // var d = groupn(group)
    // while (d.length > 1) {
    //     d = groupn(d)
    //     // console.log("t")
    // }

    return s
}

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

function tryN() {
    var s = sort()
    key = "d298be85487e0f453736b6b4a5634d09cc08790493e03941a699d0444d1ee1e4";
    hash = "96e75810b7fe519dd92f6a3f72170b00c0a8a9553f9c765a3cc681eaf7eeab38";
    root = "4101c0157600045fdf8dcb8e8a78c5891606bd596a7c497b5e9091ca08f9dc6d";
    verifyProof(s, key, hash, root)
}

class PoA {
    /**
     * Allows to get file
     */
    getFile() {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            const file = document.querySelector('input[type=file]').files[0];
            if (!file)
                throw new Error('You didn\'t add a file');
            reader.readAsDataURL(file);
            reader.onloadend = () => {
                resolve(reader.result);
            };
        });
    }

    /**
     * Allows to hashing data via keccak256
     * @param data Any set of data
     * @return {*} Hash of data
     */
    getHash(data) {
        const web3 = new Web3();
        return web3.utils.keccak256(data);
    }

    /**
     * Allows to send any data to any url
     * @param data Any data
     * @param url Any server url
     * @return {Promise<*>} response
     */
    async sendData(data, url) {
        var settings = {
            "async": true,
            "crossDomain": true,
            "url": url,
            "method": "POST",
            "headers": {
                "Content-Type": "application/json",
                "Cache-Control": "no-cache"
            },
            "data": JSON.stringify(data),
            "processData": false,
        };

        const result = await $.ajax(settings);
        return result;
    }
}

async function main() {
    const p = new PoA();
    try {
        const data = await p.getFile();
        document.getElementById('file-error').innerText = '';
        const hash = p.getHash(data);
        document.getElementById('data').innerHTML = `<p>File data</p><textarea rows="4" cols="50">${data}</textarea><p>Hash</p><p>${hash}</p>`
        // const response = await p.sendData(hash, 'url');
    } catch (e) {
        document.getElementById('file-error').innerText = e.message;
    }
}

function getCell(assetId, txNumber) {
    var a = getHash(txNumber);
    var b = getHash(assetId);

    return getHash(a.substring(2) + b.substring(2))
}

function getHash(data) {
    const web3 = new Web3();
    return web3.utils.keccak256(data);
}
