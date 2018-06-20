var PubNub = require('pubnub')

var pn = new PubNub({
    subscribeKey: "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe",
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    ssl: false
});

pn.addListener({
    message: function (m) {
        var message = "'" + m.message + "'"
        console.log(message)
        var str = JSON.parse(m.message)
        console.log(str["DeviceID"]);  
        console.log(str["TimeStamp"], str["Values"][0]["Latitude"]) 
    }
});

pn.subscribe({
    channels: ['exp-channel'],
});
