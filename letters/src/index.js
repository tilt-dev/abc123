'use strict';
const http = require('http');
const immutable = require('immutable');

const letters = immutable.List([
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
    'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
    's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
])

http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
    let item = letters.get(Math.floor(Math.random()*letters.size))
    res.end(item + '\n');
}).listen(5000, "0.0.0.0");

console.log('Server running at http://127.0.0.1:5000/');
