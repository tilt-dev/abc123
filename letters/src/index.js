'use strict';
const http = require('http');
const immutable = require('immutable');

const letters = immutable.List([
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
    'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
    's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
])

const argv = require('yargs')
  .option('port', {
    describe: 'the http port to listen on',
    default: 5000,
    type: 'number'
  }).argv

http.createServer(function (req, res) {
    console.log('Request received!');
    res.writeHead(200, {'Content-Type': 'text/plain'});
    let item = letters.get(Math.floor(Math.random()*letters.size))
    res.end(item + '\n');
}).listen(argv.port, "0.0.0.0");

console.log('~~ Setting up the "letters" service ~~')
console.log('Calibrating the subsonic subspace receiver...');
console.log('Aligning tricyclic engines with the port transmitters...');
console.log('Modifying the ventral zero point power converter...');
console.log('OK: server running at http://127.0.0.1:%d/', argv.port);
