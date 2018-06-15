var MongoClient = require('mongodb').MongoClient;
var sleep = require('sleep')
var url = "mongodb://localhost:27017/";

const DBName = "latlong"
const collectionName = "data"

// MongoClient.connect(url, function(err, db) {
//   if (err) throw err;
//   var dbObject = db.db(DBNAME);
//   dbObject.collection(collectionName).find().toArray(function(err, result) {
//     if (err) throw err;
//     console.log(result[0].latitude, result[0].longitude);
//     db.close();
//   });
// });

MongoClient.connect(url, function(err, db) {
    if (err) throw err;
    var dbObject = db.db(DBName);

    var insertJSONType = {latitude: getRandomArbitrary(30,50), longitude: getRandomArbitrary(30,50)}

        dbObject.collection(collectionName).insert(insertJSONType, (err, res) => {
            if (err) throw err;
            console.log("Document inserted {nInserted: 1}")
            db.close()
        })

});

function getRandomArbitrary(min, max) {
    return (Math.random() * (max - min) + min).toFixed(3);
}