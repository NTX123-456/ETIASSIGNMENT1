package main

import (

	//"bytes"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"
	"strings"
	"time"

	//"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct {
	ID           int    `json:ID`
	FirstName    string `json:First Name`
	LastName     string `json:" Last Name`
	MobileNumber string `json: Mobile Number`
	EmailAddress string `json: Email Address`
	Username     string `json: Username`
	Password     string `json: Password`
}

type Driver struct {
	ID                   int    `json:ID`
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
	ID                int       `json:ID`
	PFirstName        string    `json:Passenger First Name`
	DFirstName        string    `json:" Driver First Name`
	PLastName         string    `json:Passenger Last Name`
	DLastName         string    `json:" Driver Last Name`
	DateOfTrip        time.Time `json: DateOfTrip`
	PickUpPostalCode  int       `json: PickUpPostalCode`
	DropOffPostalCode int       `json: DropOffPostalCode`
	Rating            int       `json: Rating`
	DPassword         string    `json: Driver Password`
}

type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

type Trips struct {
	Trips map[string]Trip `json:"Trips"`
}

func main() {

outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Ride Sharing Platform Login\n",
			"1. Passenger\n",
			"2. Driver\n",
			"3. Quit\n")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			PassengerHome()
		case 2:
			DriverHome()
		case 3:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func PassengerHome() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Passenger Home\n",
			"1. Login\n",
			"2. Sign Up\n",
			"3. Quit\n")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			PassengerLogIn()
			break outer
		case 2:
			PassengerSignUp()
		case 3:
			main()
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func DriverHome() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Driver Home\n",
			"1. Login\n",
			"2. Sign Up\n",
			"3. Quit\n")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			DriverLogIn()
		case 2:
			DriverSignUp()
		case 3:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func PassengerLogIn() {
	var p Passenger

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &p.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &p.Password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/GetPassenger/"+p.Username, nil); err == nil {
		if res, err := client.Do(req); err == nil {

			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			if err != nil {
				panic(err.Error())

			}
			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", &p.Username, &p.Password)
			err1 := result1.Scan(&p.ID, &p.FirstName, &p.LastName, &p.EmailAddress, &p.MobileNumber, &p.Username, &p.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				PassengerHome()
			} else {

				if res.StatusCode == 200 {
					fmt.Println("You have successfully logged in", res.StatusCode)

				} else if res.StatusCode == 404 {
					fmt.Println("Error", res.StatusCode)
				}

				results, err := db.Query("select * from Passengers where Username = ? and Password = ?", p.Username, p.Password)
				if err != nil {
					fmt.Println("That is not a valid username and password!")
					panic(err.Error())

				}
				for results.Next() {

					err = results.Scan(&p.ID, &p.FirstName, &p.LastName, &p.EmailAddress, &p.MobileNumber, &p.Username, &p.Password)
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("\nFirst Name: "+p.FirstName, "\nLast Name: "+p.LastName, "\nEmail Address: "+p.EmailAddress, "\nMobile Number: "+p.MobileNumber, "\nUsername: "+p.Username, "\nPassword: "+p.Password)
				}

				defer db.Close()

			outer:
				for {
					fmt.Println(strings.Repeat("=", 10))
					fmt.Println("Your Profile\n",
						"1. Update Details\n",
						"2. Book a Trip\n",
						"3. Log Out\n",
						"4. View Trips\n",
						"5. Give a Rating\n",
						"6. Quit\n")
					fmt.Print("Enter an option: ")

					var choice int
					fmt.Scanf("%d\n", &choice)

					switch choice {
					case 1:
						UpdatePagePassenger()
					case 2:
						BookTrip()
					case 3:
						PassengerHome()
					case 4:
						ViewTripHistory()
					case 5:
						GiveRating()
					case 6:
						break outer
					default:
						fmt.Println("### Invalid Input ###")
					}
				}
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func BookTrip() {
	var trip Trip
	fmt.Println("Enter your First Name: ")
	fmt.Scanf("%s\n", &trip.PFirstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scanf("%s\n", &trip.PLastName)

	fmt.Println("Enter your Pick Up Postal Code: ")
	fmt.Scanf("%d\n", &trip.PickUpPostalCode)

	fmt.Println("Enter your Drop Off Postal Code: ")
	fmt.Scanf("%d\n", &trip.DropOffPostalCode)

	postBody, _ := json.Marshal(trip)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/CreatePassengers/"+trip.PFirstName, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}
			loc := time.Now().UTC()
			location, err := time.LoadLocation("Singapore")
			trip.DateOfTrip = loc.In(location)
			result, err := db.Exec("insert into Trip (PFirstName,PLastName,DateOfTrip, PickUpPostalCode,DropOffPostalCode)values(?,?,?,?,?)",
				trip.PFirstName, trip.PLastName, trip.DateOfTrip, trip.PickUpPostalCode, trip.DropOffPostalCode)
			if err != nil {
				panic(err.Error())
			}
			id, err := result.LastInsertId()
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("inserted id: %d\n", id)

			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("Trip has been booked successfully", res.StatusCode)
			} else if res.StatusCode == 404 {
				fmt.Println("Error", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func GiveRating() {
	var trip Trip
	fmt.Println("Enter the Rating: ")
	fmt.Scanf("%d\n", &trip.Rating)

	fmt.Println("Enter the Trip ID: ")
	fmt.Scanf("%d\n", &trip.ID)

	fmt.Println("Enter your password: ")
	fmt.Scanf("%s\n", &trip.DPassword)

	postBody, _ := json.Marshal(trip.DPassword)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateTrip/"+trip.DPassword, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result, err := db.Exec("UPDATE trip t,passengers p SET t.Rating = ? WHERE t.ID = ? and p.Password = ? and t.PFirstName = p.PFirstName", trip.Rating, trip.ID,
				trip.DPassword)
			if err != nil {
				panic(err.Error())
			}
			id, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Rows Affected: %d", id)
			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("\nUpdated")
			} else if res.StatusCode == 404 {
				fmt.Println("\nError", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func ViewTripHistory() {
	var trip Trip
	var password string

	fmt.Println("Enter your password: ")
	fmt.Scanf("%s\n", &password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/GetTrip/"+password, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db?parseTime=true")
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Password= ?", &password)
			err1 := result1.Scan(&password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {

				if res.StatusCode == 200 {
					fmt.Println("Success", res.StatusCode)
				} else if res.StatusCode == 404 {
					fmt.Println("Error", res.StatusCode)
				}

				results, err := db.Query("Select t.ID, t.PFirstName,t.PLastName,t.DFirstName,t.DLastName,t.DateOfTrip,t.PickUpPostalCode,t.DropOffPostalCode,t.Rating from trip t INNER JOIN passengers p ON p.PFirstName = t.PFirstName WHERE p.Password = ? ORDER BY t.DateOfTrip DESC ", password)
				if err != nil {
					fmt.Println("That is not a valid username and password!")
					panic(err.Error())

				}
				for results.Next() {

					err = results.Scan(&trip.ID, &trip.PFirstName, &trip.PLastName, &trip.DFirstName, &trip.DLastName, &trip.DateOfTrip, &trip.PickUpPostalCode, &trip.DropOffPostalCode, &trip.Rating)
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("\nTrip ID: ", trip.ID, "\nPassenger First Name: "+trip.PFirstName, "\nPassenger Last Name: "+trip.PLastName, "\nDriver First Name: "+trip.DFirstName, "\nDriver Last Name: "+trip.DLastName, "\nDateOfTrip: "+trip.DateOfTrip.Format(time.RFC822), "\nPick Up Postal Code: ", trip.PickUpPostalCode, "\nDrop Off Postal Code: ", trip.DropOffPostalCode, "\nRating: ", trip.Rating)
				}

				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func ViewTripRequests() {
	var trip Trip

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/GetTrip", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db?parseTime=true")

			results, err := db.Query("select * from trip where DFirstName = 'Awaiting Driver' and DLastName = 'Awaiting Driver'")
			if err != nil {
				fmt.Println("That is not a valid username and password!")
				panic(err.Error())

			}
			for results.Next() {

				err = results.Scan(&trip.ID, &trip.PFirstName, &trip.PLastName, &trip.DFirstName, &trip.DLastName, &trip.DateOfTrip, &trip.PickUpPostalCode, &trip.DropOffPostalCode, &trip.Rating)
				if err != nil {
					panic(err.Error())
				}
				fmt.Println("\nTrip ID: ", trip.ID, "\nPassenger First Name: "+trip.PFirstName, "\nPassenger Last Name: "+trip.PLastName, "\nDriver First Name: "+trip.DFirstName, "\nDriver Last Name: "+trip.DLastName, "\nDateOfTrip: "+trip.DateOfTrip.Format(time.RFC822), "\nPick Up Postal Code: ", trip.PickUpPostalCode, "\nDrop Off Postal Code: ", trip.DropOffPostalCode, "\nRating: ", trip.Rating)
			}

			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("Sucess", res.StatusCode)

			} else if res.StatusCode == 404 {
				fmt.Println("Error", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func AcceptTripRequests() {
	var trip Trip

	fmt.Println("Enter the Trip ID: ")
	fmt.Scanf("%d\n", &trip.ID)

	fmt.Println("Enter your password: ")
	fmt.Scanf("%s\n", &trip.DPassword)

	postBody, _ := json.Marshal(trip.DPassword)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+trip.DPassword, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result, err := db.Exec("UPDATE trip t,drivers d SET t.DFirstName = d.DFirstName, t.DLastName = d.DLastName, d.Availability = 'Hired' WHERE t.ID = ? and d.Password = ?", trip.ID,
				trip.DPassword)
			if err != nil {
				panic(err.Error())
			}
			id, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Rows Affected: %d", id)
			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("\nUpdated")
			} else if res.StatusCode == 404 {
				fmt.Println("\nError", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func EndTrip() {
	var d Driver

	fmt.Println("Enter your password: ")
	fmt.Scanf("%s\n", &d.Password)

	postBody, _ := json.Marshal(d.Password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+d.Password, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result, err := db.Exec("UPDATE drivers SET Availability = 'Available' WHERE Password = ?",
				d.Password)
			if err != nil {
				panic(err.Error())
			}
			id, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Rows Affected: %d", id)
			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("\nUpdated")
			} else if res.StatusCode == 404 {
				fmt.Println("\nError", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func ViewTripHistoryD() {
	var trip Trip
	var password string

	fmt.Println("Enter your password: ")
	fmt.Scanf("%s\n", &password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/GetTrip/"+password, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db?parseTime=true")

			results, err := db.Query("Select t.ID, t.PFirstName,t.PLastName,t.DFirstName,t.DLastName,t.DateOfTrip,t.PickUpPostalCode,t.DropOffPostalCode,t.Rating from trip t INNER JOIN drivers d ON d.DFirstName = t.DFirstName WHERE d.Password = ? ORDER BY t.DateOfTrip DESC ", password)
			if err != nil {
				fmt.Println("That is not a valid username and password!")
				panic(err.Error())

			}
			for results.Next() {

				err = results.Scan(&trip.ID, &trip.PFirstName, &trip.PLastName, &trip.DFirstName, &trip.DLastName, &trip.DateOfTrip, &trip.PickUpPostalCode, &trip.DropOffPostalCode, &trip.Rating)
				if err != nil {
					panic(err.Error())
				}
				fmt.Println("\nTrip ID: ", trip.ID, "\nPassenger First Name: "+trip.PFirstName, "\nPassenger Last Name: "+trip.PLastName, "\nDriver First Name: "+trip.DFirstName, "\nDriver Last Name: "+trip.DLastName, "\nDateOfTrip: "+trip.DateOfTrip.Format(time.RFC822), "\nPick Up Postal Code: ", trip.PickUpPostalCode, "\nDrop Off Postal Code: ", trip.DropOffPostalCode, "\nRating: ", trip.Rating)
			}

			defer db.Close()

			if res.StatusCode == 200 {
				fmt.Println("Success", res.StatusCode)
			} else if res.StatusCode == 404 {
				fmt.Println("Error", res.StatusCode)
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func DriverLogIn() {
	var d Driver
	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &d.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &d.Password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/GetDriver/"+d.Username, nil); err == nil {
		if res, err := client.Do(req); err == nil {

			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			if err != nil {
				panic(err.Error())

			}
			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", &d.Username, &d.Password)
			err1 := result1.Scan(&d.ID, &d.FirstName, &d.LastName, &d.EmailAddress, &d.MobileNumber, &d.Username, &d.Password, &d.IdentificationNumber, &d.LicenseNumber, &d.Availability)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				DriverHome()
			} else {
				if res.StatusCode == 200 {

					fmt.Println("You have successfully logged in", res.StatusCode)

				} else if res.StatusCode == 404 {
					fmt.Println("Error", res.StatusCode)
				}

				results, err := db.Query("select * from Drivers where Username = ? and Password = ?", d.Username, d.Password)
				if err != nil {
					fmt.Println("That is not a valid username and password!")
					panic(err.Error())

				}
				for results.Next() {

					err = results.Scan(&d.ID, &d.FirstName, &d.LastName, &d.EmailAddress, &d.MobileNumber, &d.Username, &d.Password, &d.IdentificationNumber, &d.LicenseNumber, &d.Availability)
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("\nFirst Name: "+d.FirstName, "\nLast Name: "+d.LastName, "\nEmail Address: "+d.EmailAddress, "\nMobile Number: "+d.MobileNumber, "\nUsername: "+d.Username, "\nPassword: "+d.Password, "\nIdentification Number: ", d.IdentificationNumber, "\nLicenseNumber: "+d.LicenseNumber)
				}
				defer db.Close()

			outer:
				for {
					fmt.Println(strings.Repeat("=", 10))
					fmt.Println("Your Profile\n",
						"1. Update Details\n",
						"2. View Trip Requests\n",
						"3. Accept Trip Request\n",
						"4. View Trip History\n",
						"5. End Trip\n",
						"6. Log Out\n",
						"7. Quit\n")
					fmt.Print("Enter an option: ")

					var choice int
					fmt.Scanf("%d\n", &choice)

					switch choice {
					case 1:
						UpdatePageDriver()
					case 2:
						ViewTripRequests()
					case 3:
						AcceptTripRequests()
					case 4:
						ViewTripHistoryD()
					case 5:
						EndTrip()
					case 6:
						DriverHome()
					case 7:
						break outer
					default:
						fmt.Println("### Invalid Input ###")
					}
				}

			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func PassengerSignUp() {
	var passenger Passenger
	fmt.Println("Enter your First name: ")
	fmt.Scanf("%s\n", &passenger.FirstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scanf("%s\n", &passenger.LastName)

	fmt.Println("Enter your email address: ")
	fmt.Scanf("%s\n", &passenger.EmailAddress)

	fmt.Println("Enter your mobile number: ")
	fmt.Scanf("%s\n", &passenger.MobileNumber)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/CreatePassengers/"+passenger.Username, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where EmailAddr= ? OR MobileNo= ?", &passenger.EmailAddress, &passenger.MobileNumber)
			err1 := result1.Scan(&passenger.ID, &passenger.FirstName, &passenger.LastName, &passenger.EmailAddress, &passenger.MobileNumber, &passenger.Username, &passenger.Password)
			if err1 == sql.ErrNoRows {
				if res.StatusCode == 200 {
					fmt.Println("You have successfully created a new account", res.StatusCode)

				} else if res.StatusCode == 404 {
					fmt.Println("Error - username", passenger.Username, "exists")
				}
				result, err := db.Exec("insert into Passengers (PFirstname, PLastname, EmailAddr,MobileNo,Username,Password)values(?, ?, ?, ?, ?, ?)",
					passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.LastInsertId()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("inserted id: %d", id)
				defer db.Close()

			} else {
				fmt.Println("This Email or Mobile Number has already been used, please try again")
				PassengerHome()
			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func DriverSignUp() {
	var driver Driver
	fmt.Println("Enter your First name: ")
	fmt.Scanf("%s\n", &driver.FirstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scanf("%s\n", &driver.LastName)

	fmt.Println("Enter your email address: ")
	fmt.Scanf("%s\n", &driver.EmailAddress)

	fmt.Println("Enter your mobile number: ")
	fmt.Scanf("%s\n", &driver.MobileNumber)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	fmt.Println("Enter your Account Identification Number: ")
	fmt.Scanf("%d\n", &driver.IdentificationNumber)

	fmt.Println("Enter your Account License Number: ")
	fmt.Scanf("%s\n", &driver.LicenseNumber)

	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/CreateDrivers/"+driver.Username, resBody); err == nil {
		if res, err := client.Do(req); err == nil {

			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where EmailAddr= ? OR MobileNo= ?", &driver.EmailAddress, &driver.MobileNumber)
			err1 := result1.Scan(&driver.ID, &driver.FirstName, &driver.LastName, &driver.EmailAddress, &driver.MobileNumber, &driver.Username, &driver.Password, &driver.IdentificationNumber, &driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				if res.StatusCode == 200 {
					fmt.Println("You have successfully created a new account", res.StatusCode)

				} else if res.StatusCode == 404 {
					fmt.Println("Error - username", driver.Username, "exists")
				}

				result, err := db.Exec("insert into Drivers (DFirstname, DLastname, EmailAddr,MobileNo,Username,Password,IdentificationNumber,LicenseNumber)values(?, ?, ?, ?, ?, ?, ?, ?)",
					driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.LastInsertId()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("inserted id: %d", id)
				defer db.Close()

			} else {
				fmt.Println("This Email or Mobile Number has already been used, please try again")
				DriverHome()
			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func UpdateFirstNameD() {
	var driver Driver
	fmt.Println("Enter your First Name to be updated: ")
	fmt.Scanf("%s\n", &driver.FirstName)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%v\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.FirstName)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}
			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set DFirstName = ? where Username = ? and Password = ?",
					driver.FirstName, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateLastNameD() {
	var driver Driver
	fmt.Println("Enter your Last name to be updated: ")
	fmt.Scanf("%s\n", &driver.LastName)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.LastName)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set DLastName = ? where Username = ? and Password = ?",
					driver.LastName, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateEmailAddrD() {
	var driver Driver
	fmt.Println("Enter your Email Address to be updated: ")
	fmt.Scanf("%s\n", &driver.EmailAddress)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.EmailAddress)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}
			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {

				if res.StatusCode == 200 {

					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set EmailAddr = ? where Username = ? and Password = ?",
					driver.EmailAddress, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateMobileNoD() {
	var driver Driver
	fmt.Println("Enter your Mobile Number to be updated: ")
	fmt.Scanf("%s\n", &driver.MobileNumber)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.MobileNumber)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set MobileNo = ? where Username = ? and Password = ?",
					driver.MobileNumber, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateFirstName() {
	var passenger Passenger
	fmt.Println("Enter your First name to be updated: ")
	fmt.Scanf("%s\n", &passenger.FirstName)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.FirstName)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set PFirstName = ? where Username = ? and Password = ?",
					passenger.FirstName, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateLastName() {
	var passenger Passenger
	fmt.Println("Enter your Last name to be updated: ")
	fmt.Scanf("%s\n", &passenger.LastName)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.LastName)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set PLastName = ? where Username = ? and Password = ?",
					passenger.LastName, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateEmailAddr() {
	var passenger Passenger
	fmt.Println("Enter your Email Address to be updated: ")
	fmt.Scanf("%s\n", &passenger.EmailAddress)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.EmailAddress)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set EmailAddr = ? where Username = ? and Password = ?",
					passenger.EmailAddress, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateMobileNo() {
	var passenger Passenger
	fmt.Println("Enter your Mobile Number to be updated: ")
	fmt.Scanf("%s\n", &passenger.MobileNumber)

	fmt.Println("Enter your Account Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.MobileNumber)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set MobileNo = ? where Username = ? and Password = ?",
					passenger.MobileNumber, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateUsername() {
	var passenger Passenger
	fmt.Println("Enter your Username to be updated: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.Username)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set Username = ? where Username = ? and Password = ?",
					passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateUsernameD() {
	var driver Driver
	fmt.Println("Enter your Username to be updated: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.Username)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {

				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set Username = ? where Username = ? and Password = ?",
					driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdatePassword() {
	var passenger Passenger
	fmt.Println("Enter your Username: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password to be updated: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger.Password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from passengers where Username= ? OR Password= ?", passenger.Username, passenger.Password)
			err1 := result1.Scan(passenger.ID, passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set Password = ? where Username = ? and Password = ?",
					passenger.Password, passenger.Username)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdatePasswordD() {
	var driver Driver
	fmt.Println("Enter your Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password to be updated: ")
	fmt.Scanf("%s\n", &driver.Password)

	postBody, _ := json.Marshal(driver.Password)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set Password = ? where Username = ? and Password = ?",
					driver.Password, driver.Username)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()
			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdateLicenseNoD() {
	var driver Driver
	fmt.Println("Enter your Username: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password: ")
	fmt.Scanf("%s\n", &driver.Password)

	fmt.Println("Enter your License Number to be updated: ")
	fmt.Scanf("%s\n", &driver.LicenseNumber)

	postBody, _ := json.Marshal(driver.LicenseNumber)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}

			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", driver.Username, driver.Password)
			err1 := result1.Scan(driver.ID, driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.Username, driver.Password, driver.IdentificationNumber, driver.LicenseNumber)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set LicenseNumber = ? where Username = ? and Password = ?",
					driver.LicenseNumber, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func UpdatePagePassenger() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Update Personal Details Passenger\n",
			"1. Update First Name\n",
			"2. Update Last Name\n",
			"3. Update Email Address\n",
			"4. Update Mobile Number\n",
			"5. Update Username\n",
			"6. Update Password\n",
			"7. Update all Details\n",
			"8. Quit\n")

		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			UpdateFirstName()
		case 2:
			UpdateLastName()
		case 3:
			UpdateEmailAddr()
		case 4:
			UpdateMobileNo()
		case 5:
			UpdateUsername()
		case 6:
			UpdatePassword()
		case 7:
			PassengerUpdate()
		case 8:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func UpdatePageDriver() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Update Personal Details\n",
			"1. Update First Name\n",
			"2. Update Last Name\n",
			"3. Update Email Address\n",
			"4. Update Mobile Number\n",
			"5. Update Username\n",
			"6. Update Password\n",
			"7. Update License Number\n",
			"8. Update all Details\n",
			"9. Quit\n")

		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			UpdateFirstNameD()
		case 2:
			UpdateLastNameD()
		case 3:
			UpdateEmailAddrD()
		case 4:
			UpdateMobileNoD()
		case 5:
			UpdateUsernameD()
		case 6:
			UpdatePasswordD()
		case 7:
			UpdateLicenseNoD()
		case 8:
			DriverUpdate()
		case 9:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func PassengerUpdate() {
	var passenger Passenger
	fmt.Println("Enter your First name to be updated: ")
	fmt.Scanf("%s\n", &passenger.FirstName)

	fmt.Println("Enter your Last Name to be updated: ")
	fmt.Scanf("%s\n", &passenger.LastName)

	fmt.Println("Enter your email address to be updated: ")
	fmt.Scanf("%s\n", &passenger.EmailAddress)

	fmt.Println("Enter your mobile number to be updated: ")
	fmt.Scanf("%s\n", &passenger.MobileNumber)

	fmt.Println("Enter your Account Username to confirm update: ")
	fmt.Scanf("%s\n", &passenger.Username)

	fmt.Println("Enter your Account Password to confirm update: ")
	fmt.Scanf("%s\n", &passenger.Password)

	postBody, _ := json.Marshal(passenger)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdatePassenger/"+passenger.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}
			result1 := db.QueryRow("select * from passengers where Username = ? and Password = ?", &passenger.Username, &passenger.Password)
			err1 := result1.Scan(&passenger.Username, &passenger.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePagePassenger()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")
				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Passengers set PFirstName = ?, PLastName = ?,EmailAddr = ?, MobileNo = ? where Username = ? and Password = ?",
					passenger.FirstName, passenger.LastName, passenger.EmailAddress, passenger.MobileNumber, passenger.Username, passenger.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()

			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func DriverUpdate() {
	var driver Driver
	fmt.Println("Enter your First name to be updated: ")
	fmt.Scanf("%s\n", &driver.FirstName)

	fmt.Println("Enter your Last Name to be updated: ")
	fmt.Scanf("%s\n", &driver.LastName)

	fmt.Println("Enter your email address to be updated: ")
	fmt.Scanf("%s\n", &driver.EmailAddress)

	fmt.Println("Enter your mobile number to be updated: ")
	fmt.Scanf("%s\n", &driver.MobileNumber)

	fmt.Println("Enter your Account Username to be updated: ")
	fmt.Scanf("%s\n", &driver.Username)

	fmt.Println("Enter your Account Password to confirm update: ")
	fmt.Scanf("%s\n", &driver.Password)

	fmt.Println("Enter your License Number to confirm update: ")
	fmt.Scanf("%s\n", &driver.LicenseNumber)

	postBody, _ := json.Marshal(driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/UpdateDriver/"+driver.Username, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/assg_db")
			// handle error
			if err != nil {
				panic(err.Error())
			}
			result1 := db.QueryRow("select * from drivers where Username= ? OR Password= ?", &driver.Username, &driver.Password)
			err1 := result1.Scan(&driver.Username, &driver.Password)
			if err1 == sql.ErrNoRows {
				fmt.Println("This Account does not exist, please try again")
				UpdatePageDriver()
			} else {
				if res.StatusCode == 200 {
					fmt.Println("\nUpdated")

				} else if res.StatusCode == 404 {
					fmt.Println("\nError", res.StatusCode)
				}

				result, err := db.Exec("Update Drivers set DFirstName = ?,DLastName = ?,EmailAddr = ?,MobileNo = ?,LicenseNumber = ? where Username = ? and Password = ?",
					driver.FirstName, driver.LastName, driver.EmailAddress, driver.MobileNumber, driver.LicenseNumber, driver.Username, driver.Password)
				if err != nil {
					panic(err.Error())
				}
				id, err := result.RowsAffected()
				if err != nil {
					panic(err.Error())
				}
				fmt.Printf("Rows Affected: %d", id)
				defer db.Close()
			}

		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}
