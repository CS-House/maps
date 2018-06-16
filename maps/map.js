function initMap() {

    var map = new google.maps.Map(document.getElementById('map'), {
            center: {lat: 25.344, lng: 81.036},
            zoom: 12,
            scaleControl: true,
            styles: [
              {elementType: 'geometry', stylers: [{color: '#242f3e'}]},
              {elementType: 'labels.text.stroke', stylers: [{color: '#242f3e'}]},
              {elementType: 'labels.text.fill', stylers: [{color: '#746855'}]},
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
                  featureType: 'administrative.neighborhood',
                  elementType: 'geometry',
                  stylers: [{colors: '#F41FE1'}]
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
          });
  
          
          var broadway = {
            info: 'A',
            lat: 41.976,
            long: -87.659
          };
        
          var belmont = {
            info: 'B',
            lat: 41.939,
            long: -87.655
          };
        
          var sheridan = {
            info: 'C',
            lat: 42.002,
            long: -87.661
          };
        
          var locations = [
              [broadway.info, broadway.lat, broadway.long, 0],
              [belmont.info, belmont.lat, belmont.long, 1],
              [sheridan.info, sheridan.lat, sheridan.long, 2],
          ];
        
          var infowindow = new google.maps.InfoWindow({});
        
          var marker, i;
        
          for (i = 0; i < locations.length; i++) {
            marker = new google.maps.Marker({
              position: new google.maps.LatLng(locations[i][1], locations[i][2]),
              map: map
            });
          }
}
