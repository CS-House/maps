var app = require('express')()
var http = require('http').Server(app)
var io = require("socket.io")(http)

const DBName = "latlong"
const collectionName = "data"

app.get('/', function(req, res){
    res.sendFile('/home/gowtham/random/index.html');
});

http.listen(8000, function(){
    console.log('listening on *:3000');
});

var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://localhost:27017/";

// MongoClient.connect(url, function(err, db) {
//   if (err) throw err;
//   var dbo = db.db("latlong");
//   var query = { address: "Park Lane 38" };
//   dbo.collection("data").find().toArray(function(err, result) {
//     if (err) throw err;
//     console.log(result[0].latitude, result[0].longitude);
//     db.close();
//   });
// });

io.on('connection', (socket) => {
    MongoClient.connect(url, function(err, db) {
        if (err) throw err;
        var dbo = db.db(DBName);
        dbo.collection(collectionName).find().toArray(function(err, result) {
          if (err) throw err;
        //  console.log(result[0].latitude, result[0].longitude);
          socket.emit('showrows', result)
          db.close();
        });
      }); 
})