var app = require('express')()
var http = require('http').Server(app)
var io = require("socket.io-client")

var socket = io('http://localhost');

var insertJSONType = {latitude: getRandomArbitrary(30,50), longitude: getRandomArbitrary(30,50)}

socket.emit('stream', insertJSONType)

function getRandomArbitrary(min, max) {
    return parseInt((Math.random() * (max - min) + min).toFixed(3), 10);
}