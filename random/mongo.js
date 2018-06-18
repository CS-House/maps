var MongoClient = require('mongodb').MongoClient;
var sleep = require('sleep');

var url = "mongodb://localhost:27017/";

const DBName = "latlong"
const collectionName = "data"

async function main() {
  var dbConn = await MongoClient.connect(url).catch(err => {
    console.log("Something went wrong while connecting");
  });
  var dbHandle = dbConn.db(DBName);
  
  for (var i = 0; i < 10; i++) {
    var insertJSONType = {
      latitude: getRandomArbitrary(30,50), 
      longitude: getRandomArbitrary(30,50),
    }
    var insertResult = await dbHandle.collection(collectionName).insert(insertJSONType).catch(err => {
      console.err("Major blowup: ", err);
    });
    sleep.sleep(2)
  }
  dbConn.close();
}

main();

function getRandomArbitrary(min, max) {
    return (Math.random() * (max - min) + min).toFixed(3);
}