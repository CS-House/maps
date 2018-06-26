package parsepub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Currently supports AIS140
type GPSParsed struct {
	Raw            string  // Just a copy of the raw data that was parsed
	Protocol       string  // Which protocol this message was parsed as
	PacketType     string  // Type of packet (Status? Alert? OverSpeed?)
	Uniqid         string  // Unique identifier (Used as device ID)
	TS_Millis      int64   // The timestamp from the message in unix millis
	ActualLat      float64 // Latitude in conventional format (Signed)
	ActualLng      float64 // Longitude in conventional format (Signed)
	Speed          int
	OdoMeter       int
	Direction      int
	NoOfSatellites int
	StatusBox      bool    //true = Box Open, false = Box Closed
	GSMSignal      int     // Signal Strength
	StatusBattery  bool    // true = Battery Connected, false = Battery Disconnected
	BatteryLow     bool    // true = Battery Low, false = Battery Normal
	StatusIgnition bool    // true = Ignition on, false = Ignition off
	Voltage        float64 // Analog voltage
}

// type ExportPacket struct {
// 	DeviceID  string `json:"DeviceID"`
// 	TimeStamp string `json:"TimeStamp`
// 	Values    []Values
// }

// type Values struct {
// 	Latitude  string `json:"Latitude"`
// 	Longitude string `json:"Longitude"`
// 	Speed     string `json:"Speed"`
// 	Box       string `json:"Box"`
// 	Battery   string `json:"Battery"`
// 	Ignition  string `json:"Ignition"`
// }

/*
Constants ...
*/
const (
	DDMMYYHHMMSS = "020106:150405"
	HHMMSSDDMMYY = "150405:010206"
	DEBUGGING    = false
)

func main() {

	input2 := "GTPL $9,867322035135813,A,290518,062804,18.709738,S,80.068397,W,0#"

	o1 := Parse(input2)

	jsonobj, _ := json.Marshal(o1)

	fmt.Println(string(jsonobj))
}

// Parse function takes in a raw string and puts its GPS data in the channel
// Silently fails if it cannot parse
func Parse(raw string) string {
	if raw == "" {
		if DEBUGGING {
			log.Printf("Empty message received")
		}
		return "Empty message received"
	}

	if strings.HasPrefix(raw, "GTPL") {
		ais := AIS140Parse(raw)
		return ais
	} else if strings.HasPrefix(raw, "*ZJ") {
		zt := WTDParse(raw)
		return zt
	} else {
		if DEBUGGING {
			log.Printf("Invalid or unsupported protocol")
		}
		return "Invalid or unsupported protocol"
	}
}

/*
Parsing AIS140 Device Data ...
*/
func AIS140Parse(raw string) string {
	g := &GPSParsed{}
	g.Raw = raw
	g.Protocol = "AIS140"

	// This format can have multiple messages delimited by #
	messages := strings.Split(raw, "#")

	for _, message := range messages {
		fields := strings.Split(message, ",")
		if len(fields) == 1 {
			if DEBUGGING {
				log.Printf("Not a CSV message: %s", message)
			}
			continue
		}
		switch len(fields) {
		case 10:
			break
		case 18:
			break
		default:
			if DEBUGGING {
				log.Printf("Invalid number of fields in CSV: %s", message)
			}
			continue

		}

		// 0th field has GTPL $1, GTPL $2 etc for different types of packets
		g.PacketType = strings.Split(fields[0], " ")[1]

		// For any packet type we have the following fields Uniqid, TimeDate, Lat, Long
		g.Uniqid = fields[1]

		Dd_mm_yy_hh_mm_ss := strings.Join([]string{fields[3], fields[4]}, ":")
		timestamp, e := time.Parse(DDMMYYHHMMSS, Dd_mm_yy_hh_mm_ss)
		if e != nil {
			if DEBUGGING {
				log.Printf("%s", e)
			}
		}
		// GPSParser returns in unix seconds, but thingsboard wants it in millis
		g.TS_Millis = timestamp.Unix() * 1000

		//export time stamp to the export packet

		// 5th field contains latitude as a float
		if lat, err := strconv.ParseFloat(fields[5], 64); err != nil {
			if DEBUGGING {
				log.Printf("Parsing error for latitude %s", fields[5])
			}
			continue
		} else {
			g.ActualLat = float64(lat)
		}
		// 6th field contains latitude direction information
		if fields[6] == "S" {
			g.ActualLat = -g.ActualLat
		}

		//now put it in &Values{} and later put the whole thing into exportPacket all at once.

		// 7th field contains longitude as a float
		if lng, err := strconv.ParseFloat(fields[7], 64); err != nil {
			if DEBUGGING {
				log.Printf("Parsing error for longitude %s", fields[7])
			}
			continue
		} else {
			g.ActualLng = float64(lng)
		}
		// 8th field contains lngitude direction information
		if fields[8] == "W" {
			g.ActualLng = -g.ActualLng
		}

		//export longitude into values

		var jsonBuffer bytes.Buffer
		jsonBuffer.WriteString(`{`) // Start the Json Object
		// Add whatever we've parsed so far into the JSON Object
		//jsonBuffer.WriteString(fmt.Sprintf(`"%s":[{"ts":%d,"values":{"latitude":%f,"longitude":%f`, g.Uniqid, g.TS_Millis, g.ActualLat, g.ActualLng))
		//Note that no comma has been inserted at the end

		jsonBuffer.WriteString(fmt.Sprintf(`"DeviceID":"%s","TimeStamp":"%d","Latitude":"%f","Longitude":"%f"`, g.Uniqid, g.TS_Millis, g.ActualLat, g.ActualLng))

		// Now each packetType has its own specific parameters and syntax
		switch g.PacketType {
		// Status packet ($1)
		case "$1":
			if speed, err := strconv.ParseInt(fields[9], 10, 64); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for speed %s", fields[9])
				}
				continue
			} else {
				g.Speed = int(speed)

				// Add the parsed Speed to the json
				jsonBuffer.WriteString(fmt.Sprintf(`,"Speed":%d`, g.Speed))
			}

			if boxOpen, err := strconv.ParseBool(fields[13]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for boxOpen %s", fields[13])
				}
				continue
			} else {
				g.StatusBox = boxOpen

				// Add the box Status to json
				// Adds true or false (Json booleans)
				jsonBuffer.WriteString(fmt.Sprintf(`,"Box":%v`, g.StatusBox))
			}

			if batConnected, err := strconv.ParseBool(fields[15]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for batConnected %s", fields[15])
				}
				continue
			} else {
				g.StatusBattery = batConnected

				// Add the battery Status to json
				// Adds true or false (Json booleans)
				jsonBuffer.WriteString(fmt.Sprintf(`,"Battery":%v`, g.StatusBattery))
			}

			if ignition, err := strconv.ParseBool(fields[16]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for ignition %s", fields[16])
				}
				continue
			} else {
				g.StatusIgnition = ignition

				// Add the ignition Status to json
				// Adds true or false (Json booleans)
				jsonBuffer.WriteString(fmt.Sprintf(`,"Ignition":%v`, g.StatusIgnition))
			}
		// Ignition Alert packet ($2)
		case "$2":
			if ignition, err := strconv.ParseBool(fields[9]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for ignition %s", fields[9])
				}
				continue
			} else {
				g.StatusIgnition = ignition
				if ignition {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Ignition on"))
				} else {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Ignition off"))
				}
			}
		// Main Battery Alert packet ($3)
		case "$3":
			if batConnected, err := strconv.ParseBool(fields[9]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for batConnected %s", fields[9])
				}
				continue
			} else {
				g.StatusBattery = batConnected
				if batConnected {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Battery connected"))
				} else {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Battery disconnected"))
				}
			}
		// Low Battery Alert packet ($4)
		case "$4":
			jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Battery low"))
		// Harsh Acceleration Alert packet ($5)
		case "$5":
			jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Harsh Acceleration"))
		// Harsh Braking Alert packet ($6)
		case "$6":
			jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Harsh Braking"))
		// Overspeeding Alert packet ($7)
		case "$7":
			jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "OverSpeeding Alert"))
			if speed, err := strconv.ParseInt(fields[9], 10, 64); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for speed %s", fields[9])
				}
				continue
			} else {
				g.Speed = int(speed)
				// Add the parsed Speed to the json
				jsonBuffer.WriteString(fmt.Sprintf(`,"Speed":%d`, g.Speed))
			}
		// Box Alert packet ($8)
		case "$8":
			if boxOpen, err := strconv.ParseBool(fields[9]); err != nil {
				if DEBUGGING {
					log.Printf("Parsing error for boxOpen %s", fields[9])
				}
				continue
			} else {
				g.StatusBattery = boxOpen

				if boxOpen {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Box Opened"))
				} else {
					jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "Box Closed"))
				}
			}
		// SOS Alert packet ($9)
		case "$9":
			jsonBuffer.WriteString(fmt.Sprintf(`,"Alert":"%s"`, "SOS"))
		}

		jsonBuffer.WriteString(`}`)
		jsonString := jsonBuffer.String()
		//c <- &jsonString
		// c <- g
		if DEBUGGING {
			log.Printf("Parsed dumped")
		}

		return jsonString
	}
	return ""
}

/*
Parsing a different format of device ...
*/
func WTDParse(raw string) string {
	if raw == "*ZJ#" {
		if DEBUGGING {
			log.Printf("Empty message")
		}
		return "Empty message"
	}

	g := &GPSParsed{}
	g.Raw = raw
	g.Protocol = "WTD"
	fields := strings.Split(raw, ",")
	if len(fields) == 1 {
		if DEBUGGING {
			log.Printf("Not a CSV message: %s", raw)
		}
		return "Not a CSV message"
	}

	// For any packet type we have the following fields Uniqid, TimeDate, Lat, Long
	g.Uniqid = fields[1]
	Hh_mm_ss_dd_mm_yy := strings.Join([]string{fields[3], fields[11]}, ":")
	timestamp, e := time.Parse(HHMMSSDDMMYY, Hh_mm_ss_dd_mm_yy)
	if e != nil {
		if DEBUGGING {
			log.Printf("%s", e)
		}
		return ""
	}
	// fmt.Println(timestamp)
	// GPSParser returns in unix seconds, but thingsboard wants it in millis
	g.TS_Millis = timestamp.Unix() * 1000
	// 5th field contains latitude as a float
	if lat, err := strconv.ParseFloat(fields[5], 64); err != nil {
		if DEBUGGING {
			log.Printf("Parsing error for latitude %s", fields[5])
		}
		return "Not a CSV message"
	} else {
		g.ActualLat = float64(lat) / 100
	}
	// 6th field contains latitude direction information
	if fields[6] == "S" {
		g.ActualLat = -g.ActualLat
	}

	// 7th field contains longitude as a float
	if lng, err := strconv.ParseFloat(fields[7], 64); err != nil {
		if DEBUGGING {
			log.Printf("Parsing error for longitude %s", fields[7])
		}
		return "Parsing error for longitude"
	} else {
		g.ActualLng = float64(lng) / 100
	}
	// 8th field contains lngitude direction information
	if fields[8] == "W" {
		g.ActualLng = -g.ActualLng
	}

	// // _ = timestamp
	var jsonBuffer bytes.Buffer
	jsonBuffer.WriteString("{") // Start the Json Object
	// Add whatever we've parsed so far into the JSON Object
	//jsonBuffer.WriteString(fmt.Sprintf(`"%s":[{"TimeStamp":%d,"Values":{"Latitude":%f,"Longitude":%f`, g.Uniqid, g.TS_Millis, g.ActualLat, g.ActualLng))

	jsonBuffer.WriteString(fmt.Sprintf(`"DeviceID":"%s","TimeStamp":"%d","Latitude":"%f","Longitude":"%f"`, g.Uniqid, g.TS_Millis, g.ActualLat, g.ActualLng))

	jsonBuffer.WriteString(`}`)

	jsonString := jsonBuffer.String()
	log.Printf(jsonString)
	return jsonString
}

// {
// 	"DeviceID":"867322035135813",
// 	"TimeStamp":1527575284000,
// 	"values": [
// 		{
// 			"Latitude":18.709738,
// 			"Longitude":80.068397,
// 			"Speed":0,
// 			"Box":false,
// 			"Bat":true,
// 			"Ign":false
// 		}
// 	]
// }
