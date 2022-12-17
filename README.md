# ETIASSIGNMENT1
<h2>My Assignment for Emerging Trends In IT, this assignment includes a console application simulating a Ride-Sharing Platform, which is created along with microservices.</h2>

<h3>Design of the ride sharing platform:</h3>

<h4>Microservices Involved<h4>
1. Passengers Microservice
  
2. Driver Microservice
  
3. Trip Microservice

The Business Logic of this application is mainly providing a convenient platform for both passengers and drivers to ride and accept ride requests. Like Shopee's business logic which focusses on providing a convenient platform for sellers and buyers to buy and sell.

<h4>Explanation of what features each microservice helps to accomplish</h4>
<h4>Passengers Microservice:</h4>
The Passengers Microservice enables users to sign up for an account as a passenger, they can do so by providing their credentials such as their first name,last name, email address, mobile number, username and password. After which the passengers are allowed to log in into their account using their username and password.

The Passenger Microservice also allows users to update their credentials if they wish to, they can key in their updated information accordingly and the microservice will help them to update it automatically. Passengers are not allowed to delete their account.

<h4>Drivers Microservice:</h4>
Similarly to the Passenger Microservice. The Drivers Microservice enables users to sign up for an account as a driver, they can do so by providing their credentials such as their first name,last name, email address, mobile number, username,password,Identification Number and License number. After which the drivers are allowed to log in into their account using their username and password.

The Driver Microservice also allows users to update their credentials if they wish to, they can key in their updated information accordingly and the microservice will help them to update it automatically. Like Passengers, drivers are also not allowed to delete their account and they are not allowed to update their Identification Number.

<h4>Trip Microservice:</h4>
The trip microservice can be accessed from both the passenger and driver, microservices. It is slightly different from the passenger and driver microservice because it does not have a console like the passenger and driver microservice. The trip microservice helps passengers and drivers accomplish several features. For the passenger side, the passengers can utilise the trip microservice to book a trip, see their previous history of their trips taken, as well as leave a rating on the trip they have taken. As for the drivers, they can utilise the trip microservice to view the available trip requests to them which are booked by different passengers, they can then choose to display all the trip requests and choose which trip they want to accept. Drivers can also display their trip history. Each Driver comes with a availability status, which the default sets it to 'available' when they are not driving a passenger, when drivers accept a trip request from a passenger their trip has officially started and the driver microservice will then change their availability status to 'Hired' which means that the driver is not hired by a passenger as the driver have accepted their trip request. Upon completing the trip, drivers can also choose to end their trip, after which the driver microservice will change the driver's availability status back to available.

<h4>REST APIS Used</h4>
1. <h4>HTTP GET:</h4> To Retrieve and display information from the database, E.g displaying user's credentials when they log in, displaying trip history to passengers and drivers, displaying trip requests to the driver.
  
2. <h4>HTTP POST:</h4> To Post new Information entered by the user to the database. This is mainly used when the user creates a new account and enters in their credentials. E.g when a customer or driver creates a new account.
  
3. <h4>HTTP PUT:</h4> To Update any information in the database. E.g. When Users choose to update their profile credentials like their first name, last name or email addresses. To Update the driver's availability, update driver's names in a new trip request when they have accepted the trip.
  
<h4>Database Design</h4>
The database that I am utilising is MySQL which is an open-source relational database management system. There are a total of 3 different tables created in the         database. The three tables include:
<h4> 1. Passenger </h4>
The Passenger Table stores the passenger's details and credentials. The list of details are as follows: Refer to: ![image](https://user-images.githubusercontent.com/73065899/208236063-003316bc-ab21-46b8-a8bb-f8e4851e92b4.png). Passenger ID, which stores the ID of the passenger record. PFirstName, which records the passenger's first name. PLastName, which records the passenger's last name. MobileNo, which records the passenger's mobile number. EmailAddr, which records the passenger's email address. Username, which records the passenger's username and finally Password, which records the passenger's password.

 <h4>2. Driver </h4>
 The Driver Table stores the driver's details and credentials. The list of details are as follows: Refer to: ![image](https://user-images.githubusercontent.com/73065899/208236202-37c50f8d-1591-4122-9191-0a684c263da9.png). Driver ID, which stores the ID of the driver record. DFirstName, which records the driver's first name. DLastName, which records the driver's last name. MobileNo, which records the driver's mobile number. EmailAddr, which records the driver's email address. Username, which records the driver's username. Password, which records the driver's password. IdentificationNumber which records the driver's Identification number. LicenseNumber which records the driver's License Number and finally the Avalability which records the driver's availability status. The default availability status is set to available.

 <h4>3. Trip </h4>
The Trip table stores all the information of a trip record booked by a passenger which is then accepted by a driver. The list of details are as follows: Refer to: 
![image](https://user-images.githubusercontent.com/73065899/208236383-67e8e70c-ba33-40df-bc4d-e6e099f51e37.png). Trip ID, which stores the ID of the trip record. PFirstName, which stores the first name of the passenger. PLastName, which stores the last name of the passenger, DFirstName, which stores the first name of the drive. DLastName, which stores the last name of the driver. DateOfTrip, which stores the date and time of the trip that is booked by the passenger. PickUpPostalCode, which stores the Pick Up Postal Code that the passenger has keyed in. DropOffPostalCode, which stores the drop off postal code that the passenger has keyed in and finally, the Rating which stores the passenger's rating for the particular trip





