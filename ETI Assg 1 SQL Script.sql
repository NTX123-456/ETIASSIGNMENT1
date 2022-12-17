CREATE database assg_db
use assg_db
SET SQL_SAFE_UPDATES = 0;

CREATE TABLE Passengers (ID int AUTO_INCREMENT NOT NULL PRIMARY KEY,PFirstName VARCHAR(30), PLastName VARCHAR(30), EmailAddr VARCHAR(40),MobileNo VARCHAR(30), Username VARCHAR(30), Password VARCHAR(30));

INSERT INTO Passengers (ID, PFirstName, PLastName, EmailAddr,MobileNo,Username,Password) VALUES
(1,"Joe", "Lee", "JoeL@Gmail.com", "+6590986789", "JoeL645", "JE@3425"),(2,"Jack", "Doe","JackD@Gmail.com", "+6591903243", "JackD909", "JD@0987"),(3,"Tom", "Lin", "TomL@Gmail.com", "+6590986789", "TomL877", "TL@4343");

CREATE TABLE Drivers (ID int AUTO_INCREMENT NOT NULL PRIMARY KEY, 
DFirstName VARCHAR(30), DLastName VARCHAR(30), EmailAddr VARCHAR(40),MobileNo VARCHAR(30), Username VARCHAR(30), Password VARCHAR(30),IdentificationNumber VARCHAR(10), LicenseNumber VARCHAR(10), Availability VARCHAR(10) DEFAULT "Available");

INSERT INTO drivers (ID, DFirstName, DLastName, EmailAddr,MobileNo,Username,Password,IdentificationNumber,LicenseNumber,Availability) VALUES
(1,"John", "Tan", "JohnT@Gmail.com", "+6590986789", "JohnT234", "JT@1234", 1342, "S45673E", "Available"),(2,"James", "Lim", "JamesL@Gmail.com", "+6591903243", "JamL342", "JL@2134", 1455, "S32424J", "Available"),(3,"Jimmy", "Loh", "JimmyL@Gmail.com", "+6590986789", "Jims224", "J@23424", 1342, "S23455T", "Available");

CREATE TABLE Trip (ID int AUTO_INCREMENT NOT NULL PRIMARY KEY,PFirstName VARCHAR(30) DEFAULT "Awaiting Passenger", PLastName VARCHAR(30) DEFAULT "Awaiting Passenger",DFirstName VARCHAR(30) DEFAULT "Awaiting Driver",DLastName VARCHAR(30) DEFAULT "Awaiting Driver", DateOfTrip datetime,PickUpPostalCode int, DropOffPostalCode int, Rating int Default 0);


