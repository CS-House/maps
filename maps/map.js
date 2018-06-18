var pn = new PubNub({
    subscribeKey: "sub-c-18580a92-f8cc-11e5-9086-02ee2ddab7fe",
    publishKey: "pub-c-f3cae627-a107-45d2-a3cc-256467b09e6a",
    ssl: false
});

pn.subscribe({
    channels: ['exp-channel'],
});

function initMap() {

    var markerMap = new Map();

    function getRandomArbitrary(min, max) {
        return Math.random() * (max - min) + min;
    }

    class Location {
        constructor(latitude, longitude) {
            this.latitude = latitude;
            this.longitude = longitude;
        }
    }

    var mapOptions = {
        center: new google.maps.LatLng(25.36,13.044),
        zoom: 4,
        styles: [
            {elementType: 'geometry', stylers: [{color: '#242f3e'}]},
            {elementType: 'labels.text.stroke', stylers: [{color: '#242f3e'}]},
            {elementType: 'labels.text.fill', stylers: [{color: '#746855'}]},
            {
                featureType: 'administrative.locality',
                elementType: 'labels.text.fill',
                stylers: [{color: '#d59563'}]
            },
            {
                featureType: 'poi',
                elementType: 'labels.text.fill',
                stylers: [{color: '#d59563'}]
            },
            {
                featureType: 'poi.park',
                elementType: 'geometry',
                stylers: [{color: '#263c3f'}]
            },
            {
                featureType: 'poi.park',
                elementType: 'labels.text.fill',
                stylers: [{color: '#6b9a76'}]
            },
            {
                featureType: 'road',
                elementType: 'geometry',
                stylers: [{color: '#38414e'}]
            },
            {
                featureType: 'road',
                elementType: 'geometry.stroke',
                stylers: [{color: '#212a37'}]
            },
            {
                featureType: 'road',
                elementType: 'labels.text.fill',
                stylers: [{color: '#9ca5b3'}]
            },
            {
                featureType: 'road.highway',
                elementType: 'geometry',
                stylers: [{color: '#746855'}]
            },
            {
                featureType: 'road.highway',
                elementType: 'geometry.stroke',
                stylers: [{color: '#1f2835'}]
            },
            {
                featureType: 'road.highway',
                elementType: 'labels.text.fill',
                stylers: [{color: '#f3d19c'}]
            },
            {
                featureType: 'transit',
                elementType: 'geometry',
                stylers: [{color: '#2f3948'}]
            },
            {
                featureType: 'transit.station',
                elementType: 'labels.text.fill',
                stylers: [{color: '#d59563'}]
            },
            {
                featureType: 'water',
                elementType: 'geometry',
                stylers: [{color: '#17263c'}]
            },
            {
                featureType: 'water',
                elementType: 'labels.text.fill',
                stylers: [{color: '#515c6d'}]
            },
            {
                featureType: 'water',
                elementType: 'labels.text.stroke',
                stylers: [{color: '#17263c'}]
            }
        ]
    }

    var map = new google.maps.Map(document.getElementById('map'), mapOptions)

    function addToMap(deviceId, location) {
        var marker = new google.maps.Marker({
            position: new google.maps.LatLng(location.latitude, location.longitude),
        })
        marker.setMap(map);
        markerMap.set(deviceId, marker);
    }

        pn.addListener({
            message: function (m) {
                // handle response from publisher message
                addToMap(1, new Location(JSON.parse(m.message)["Latitude"], JSON.parse(m.message)["Longitude"]));

                console.log(JSON.parse(m.message)["Latitude"], JSON.parse(m.message)["Longitude"]);   
            }
        });

}
