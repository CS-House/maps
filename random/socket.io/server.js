var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
  res.sendFile(__dirname + '/update.html');
});

http.listen(8000, function(){
  console.log('listening on *:8000');
});

io.on('connection', function (socket) {
    console.log('a client connected');
    socket.emit('stream', insertJSONType)
});

var insertJSONType = {latitude: getRandomArbitrary(30,50), longitude: getRandomArbitrary(30,50)}

function getRandomArbitrary(min, max) {
    return parseInt((Math.random() * (max - min) + min).toFixed(3), 10);
}
