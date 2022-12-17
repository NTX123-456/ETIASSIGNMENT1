package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gorilla/mux"
)

type Passenger struct {
	//ID           int    `json:ID`
	FirstName    string `json:First Name`
	LastName     string `json:" Last Name`
	MobileNumber string `json: Mobile Number`
	EmailAddress string `json: Email Address`
	Username     string `json: Username`
	Password     string `json: Password`
}

type Driver struct {
	//ID                   int    `json:ID`
	FirstName            string `json:First Name`
	LastName             string `json:" Last Name`
	MobileNumber         string `json: Mobile Number`
	EmailAddress         string `json: Email Address`
	IdentificationNumber int    `json: Identification Number`
	LicenseNumber        string `json: License Number`
	Username             string `json: Username`
	Password             string `json: Password`
	Availability         string `json: Availability`
	//StartTrip            bool   `json: StartTrip`
	//EndTrip              bool   `json: EndTrip`
}
type Trip struct {
	PickUpPostal   string `json:Pick Up Postal Code`
	DropOffPostal  string `json:Drop Off Postal Code`
	driverUsername string `json:Driver`
}

type Drivers struct {
	Driver map[string]Driver `json:"Drivers"`
}

type Passengers struct {
	Passenger map[string]Passenger `json:"Passengers"`
}

var passengers map[string]Passenger = map[string]Passenger{
	"JoeLee":  Passenger{"Joe", "Lee", "+6590986789", "JoeL@Gmail.com", "JoeL645", "JE@3425"},
	"JackDoe": Passenger{"Jack", "Doe", "+6591903243", "JackD@Gmail.com", "JackD909", "JD@0987"},
	"TomLin":  Passenger{"Tom", "Lin", "+6590986789", "TomL@Gmail.com", "TomL877", "TL@4343"},
}

var drivers map[string]Driver = map[string]Driver{
	"JohnTan":  Driver{"John", "Tan", "+6590986789", "JohnT@Gmail.com", 1342, "S45673E", "JohnT234", "JT@1234", "Available"},
	"JamesLim": Driver{"James", "Lim", "+6591903243", "JamesL@Gmail.com", 1455, "S32424J", "JamL342", "JL@2134", "Available"},
	"JimmyLoh": Driver{"Jimmy", "Loh", "+6590986789", "JimmyL@Gmail.com", 1342, "S23455T", "Jims224", "J@23424", "Available"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/CreatePassengers/{passengerUsername}", Createpassenger).Methods("POST")
	router.HandleFunc("/api/v1/GetPassenger/{passengerUsername}", Getpassenger).Methods("GET")
	router.HandleFunc("/api/v1/UpdatePassenger/{passengerUsername}", Updatepassenger).Methods("PUT", "PATCH")
	router.HandleFunc("/api/v1/CreateDrivers/{driverUsername}", Createdriver).Methods("POST")
	router.HandleFunc("/api/v1/GetDriver/{driverUsername}", Getdriver).Methods("GET")
	router.HandleFunc("/api/v1/UpdateDriver/{driverUsername}", Updatedriver).Methods("PUT", "PATCH")
	router.HandleFunc("/api/v1/CreateTrip/{PassengerFirstName}", CreateTrip).Methods("POST")
	router.HandleFunc("/api/v1/GetTrip/{PassengerPassword}", GetTrip).Methods("GET")
	router.HandleFunc("/api/v1/GetTrip", GetTripD).Methods("GET")
	router.HandleFunc("/api/v1/UpdateTrip/{driverPassword}", UpdateTrip).Methods("PUT")
	//router.HandleFunc("/api/v1/GetDriver/{driverUsername}", Getdriver).Methods("GET")
	//router.HandleFunc("/api/v1/UpdateDriver/{driversUsername}", Updatedriver).Methods("PUT", "PATCH")
	//router.HandleFunc("/api/v1/courses", allcourses)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func CreateTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["PassengerFirstName"]]; !ok {
					fmt.Println(data)
					passengers[params["PassengerFirstName"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Error")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func UpdateTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["driverPassword"]]; !ok {
					fmt.Println(data)
					passengers[params["driverPassword"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Error")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["PassengerPassword"]]; !ok {
					fmt.Println(data)
					passengers[params["PassengerPassword"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Error")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func GetTripD(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["tripid"]]; !ok {
					fmt.Println(data)
					passengers[params["tripid"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Error")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func Updatepassenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["passengerUsername"]]; ok {
					fmt.Println(data)

					passengers[params["passengerUsername"]] = data
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Passenger Username does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := passengers[params["passengerUsername"]]; ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "First Name":
							orig.FirstName = v.(string)
						case "Last Name":
							orig.LastName = v.(string)
						case "Mobile Number":
							orig.MobileNumber = v.(string)
						case "Email Address":
							orig.EmailAddress = v.(string)
						case "Username":
							orig.Username = v.(string)
						case "Password":
							orig.Password = v.(string)

						}
					}
					passengers[params["passengerUsername"]] = orig
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Passenger Username does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}

func Createpassenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["passengerUsername"]]; !ok {
					fmt.Println(data)
					passengers[params["passengerUsername"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Passenger Username exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func Getpassenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := passengers[params["passengerUsername"]]; !ok {
					fmt.Println(data)
					passengers[params["passengerUsername"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Passenger Username exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func Getdriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := drivers[params["driverUsername"]]; !ok {
					fmt.Println(data)
					drivers[params["driverUsername"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "No Driver Found")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Course ID")
	}
}

func Updatedriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := drivers[params["driverUsername"]]; ok {
					fmt.Println(data)

					drivers[params["driverUsername"]] = data
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver Username does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := drivers[params["driverUsername"]]; ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "First Name":
							orig.FirstName = v.(string)
						case "Last Name":
							orig.LastName = v.(string)
						case "Mobile Number":
							orig.MobileNumber = v.(string)
						case "Email Address":
							orig.EmailAddress = v.(string)
						case "Identification Number":
							orig.Username = v.(string)
						case "License Number":
							orig.Password = v.(string)
						case "IdentificationNumber":
							orig.IdentificationNumber = int(v.(float64))
						case "LicenseNumber":
							orig.LicenseNumber = v.(string)
						case "Availability":
							orig.Availability = v.(string)
						}
					}
					drivers[params["driverUsername"]] = orig
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver Username does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}

func Createdriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := drivers[params["driverUsername"]]; !ok {
					fmt.Println(data)
					drivers[params["driverUsername"]] = data

					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver Username exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if val, ok := drivers[params["driverUsername"]]; ok {
		if r.Method == "DELETE" {
			fmt.Fprintf(w, params["driverUsername"]+" Deleted")
			delete(drivers, params["driverUsername"])
		} else {
			json.NewEncoder(w).Encode(val)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Driver Username")
	}
}

func alldrivers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if v, ok := drivers[params["driverUsername"]]; ok { //IF COURSE ID EXIST
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		} else if r.Method == "POST" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Course ID exists")

		} else if r.Method == "PUT" {
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				var data Driver

				if err := json.Unmarshal(body, &data); err == nil {
					//fmt.Printf("%v", data)
					w.WriteHeader(http.StatusOK)
					drivers[params["driverUsername"]] = data
				}
			}

		} else {
			delete(drivers, params["driverUsername"])
			fmt.Fprintf(w, params["driverUsername"]+" Deleted")
		}
	} else if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				//fmt.Printf("%v", data)
				w.WriteHeader(http.StatusOK)
				drivers[params["driverUsername"]] = data
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Driver Username dose not exist")
		} else {
			fmt.Fprintf(w, "Invalid Driver Username")
		}

	}

}
