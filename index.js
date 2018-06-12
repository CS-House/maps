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

console.log(generateRandomLatLng())

function makeid() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  
    for (var i = 0; i < 5; i++)
      text += possible.charAt(Math.floor(Math.random() * possible.length));
  
    return text;
}

console.log(makeid());

var Packet = {
    info: makeid(),
    lat : generateRandomLatLng(),
    long: generateRandomLatLng(),
}

console.log(Packet)

function Changeloc(lat, long) {
    lat = lat + 0.001;
    long = long + 0.001;
}