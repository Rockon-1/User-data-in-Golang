This application is made using Golang.

user's data is inside  object variable as mentioned in ASSIGNMENT
we can change user's data by changing values of variable stored in "data.go" file of "dataservices " Folder

----------------------------------------------------------
2 Option to run this application

1> I have build my application so executable file works fine in windows
	To launch the executable file
	-> double click the Project1.exe application
		OR
	->  open the terminal , set the path under Project1 and enter the below commabd
	./PROJECT1.exe

2> If Go is installed in your system, then you can run this application by opening the terminal and enter the below command:-
	go run main.go 

----------------------------------------------------------


I have made 4 API's

For windows the API's are:-

localhost:8080/ping  
//[GET API]we can use this API to check application is working or not(trial mode)

localhost:8080/all
//[GET API]use to retrieve all user's data

localhost:8080/id?id=2
//[GET API] use to retrieve particular user by specific id param
//we can change the value of id param)

localhost:8080/ListOfID?id=2,1
//[GET API] use to retrieve list of user by specific id param
//we can enter many id as we can , just don't forgot to put comma in between

--------------------------------------------------------------------------

All the API's working fine in my system (even checked by unit testiung)
..........................................................................