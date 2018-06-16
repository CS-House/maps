var pubnub = require('pubnub')

var pn = new pubnub({
    subscribeKey: "sub-c-1858fa40-7158-11e8-9683-aecdde7ceb31",
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    ssl: false
})


pn.publish({
    message: {
        "color": "blue"
    },
    channel: 'stream'
})
