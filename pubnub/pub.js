var PubNub = require('pubnub')

const channel = "mychannel"

var pubnub = new PubNub ({
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    subscribeKey: "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe",
    secretKey: "sec-c-MGMxMDBmYTMtOGZmYy00ZjMyLTkwMmUtMjE4YWE5MzJiYjg5",
    ssl: false
})

pubnub.publish(
    {
        message: {
            lat: generateRandomLatLng(),
            long: generateRandomLatLng()
        },
        channel: "map2-channel",
        sendByPost: false,
        storeInHistory: false,
    },
    function(status, response) {
        if (status.error) {
            console.log(error)
        } else {
            console.log("Message published w/ timetoken", response.timetoken)
        }
    }
);

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

