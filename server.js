var net = require('net')
var JsonSocket = require('json-socket');

var port = 8000
var server = net.createServer()

server.listen(port);

server.on('connection', function(socket) {
    socket = new JsonSocket(socket);

    socket.on('message', function(Packet) {
        var jsonObject = {"info": Packet.info, "lat": Packet.lat, "long": Packet.long}
        var array = []
        array.push(jsonObject)

        // console.log(jsonObject.info)
        // console.log(array[0])
        console.log(jsonObject.lat)
    })
})

