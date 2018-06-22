var PubNub = require('pubnub')

var pn = new PubNub({
    subscribeKey: "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe",
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    ssl: false
});

pn.addListener({
    message: function (m) {
        var message = "'" + m.message + "'"
        var str = JSON.parse(m.message)
        //console.log(str)
        var lat = str["DeviceID"]
        // var long = str["Values"][0]["Longitude"]
        console.log(lat)
    }
});

pn.subscribe({
    channels: ['exp-channel'],
});
