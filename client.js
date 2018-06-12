var net = require('net'),
    JsonSocket = require('json-socket');

var port = 8000; 
var host = '127.0.0.1';
var socket = new JsonSocket(new net.Socket()); 
socket.connect(port, host);
socket.on('connect', function() { 
    socket.sendMessage(Packet);
    console.log(Packet);
});

function generateRandomLatLng()
{
    var num = Math.random()*180;
    var posorneg = Math.floor(Math.random());
    if (posorneg == 0)
    {
        num = num * -1;
    }
    return num.toFixed(3);
}

function makeid() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  
    for (var i = 0; i < 5; i++)
      text += possible.charAt(Math.floor(Math.random() * possible.length));
  
    return text;
}

var Packet = {
    info: makeid(),
    lat : generateRandomLatLng(),
    long: generateRandomLatLng(),
}
