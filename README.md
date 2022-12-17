# ETIASSIGNMENT1
<h2>My Assignment for Emerging Trends In IT, this assignment includes a console application simulating a Ride-Sharing Platform, which is created along with microservices.</h2>

<h3>Design of the ride sharing platform:</h3>

<h4>Microservices Involved<h4>
1. Passengers Microservice
2. Driver Microservice
3. Trip Microservice

The Business Logic of this application is mainly providing a convenient platform for both passengers and drivers to utilise. Like Shopee's business logic which focusses on providing a convenient platform for sellers and buyers

<h4>Explanation of what features each microservice helps to accomplish</h4>
Passengers Microservice:
The Passengers Microservice enables users to sign up for an account as a passenger, they can do so by providing their credentials such as their first name,last name, email address, mobile number, username and password. After which the passengers are allowed to log in into their account using their username and password.

The Passenger Microservice also allows users to update their credentials if they wish to, they can key in their updated information accordingly and the microservice will help them to update it automatically. Passengers are not allowed to delete their account.

Drivers Microservice:
Similarly to the Passenger Microservice. The Drivers Microservice enables users to sign up for an account as a driver, they can do so by providing their credentials such as their first name,last name, email address, mobile number, username,password,Identification Number and License number. After which the drivers are allowed to log in into their account using their username and password.

The Driver Microservice also allows users to update their credentials if they wish to, they can key in their updated information accordingly and the microservice will help them to update it automatically. Like Passengers, drivers are also not allowed to delete their account.

Trip Microservice:
The trip microservice can be accessed from both the passenger and driver, microservices. It is slightly different from the passenger and driver microservice because it does not have a console like the passenger and driver microservice. The trip microservice helps passengers and drivers accomplish several features. For the passenger side, the passengers can utilise the trip microservice to book a trip, see their previous history of their trips taken, as well as leave a rating on the trip they have taken. As for the drivers, they can utilise the trip microservice to view the available trip requests to them which are booked by different passengers, they can then choose to display all the trip requests and choose which trip they want to accept. Drivers can also display their trip history. Each Driver comes with a availability status, which the default sets it to 'available' when they are not driving a passenger, when drivers accept a trip request from a passenger their trip has officially started and the driver microservice will then change their availability status to 'Hired' which means that the driver is not hired by a passenger as the driver have accepted their trip request. Upon completing the trip, drivers can also choose to end their trip, after which the driver microservice will change the driver's availability status back to available.

<h4>REST APIS Used</h4>
1. HTTP GET: To Retrieve and display information from the database, E.g displaying user's credentials when they log in, displaying trip history to passengers and drivers, displaying trip requests to the driver.
  
2. HTTP POST: To Post new Information entered by the user to the database. This is mainly used when the user creates a new account and enters in their credentials. E.g when a customer or driver creates a new account.
  
3. HTTP PUT: To Update any information in the database. E.g. When Users choose to update their profile credentials like their first name, last name or email addresses. To Update the driver's availability, update driver's names in a new trip request when they have accepted the trip.





