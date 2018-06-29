var pn = new PubNub({
    subscribeKey: "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe",
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    ssl: false
});

pn.subscribe({
    channels: ['exp-channel'],
});

var map;

function initMap() {
    var markers = []

    map = new google.maps.Map(document.getElementById('map'), {
        center: {
            lat: -34.397,
            lng: 150.644
        },
        zoom: 8
    });

    class Point {
        constructor(id, lat, lon) {
            this.id = id;
            this.lat = lat;
            this.lon = lon;
        }
    }

    pubnub = new PubNub({
        publishKey: '',
        subscribeKey: 'sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe'
    })

    pubnub.addListener({
        message: pubnubListener,
        presence: function (presenceEvent) {
            // handle presence
        }
    });

    pubnub.subscribe({
        channels: ['exp-channel']
    });

    function pubnubListener(m) {

        var pointsFromSql = JSON.parse(m.message);

        for (var point of pointsFromSql) {

            var oldMarker = markers.find(m => m.title == point["DeviceID"]);
            var markerIndex = markers.findIndex(m => m.title == point["DeviceID"]);

            var gpsCurrentLocation = {
                lat: parseFloat(point["Lat"]),
                lng: parseFloat(point["Long"])
            }

            if (oldMarker) {
                //console.log("Found old point: " + point["DeviceID"]);
                markers[markerIndex].setMap(null);
                //markers[markerIndex].setPosition(gpsCurrentLocation);
                //markers.remove(oldMarker);
            }

            var marker = new google.maps.Marker({
                position: gpsCurrentLocation,
                map: map,
                title: point["DeviceID"]
            });
            
            markers.push(marker);
        }
    }
}