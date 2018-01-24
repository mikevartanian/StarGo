My code works in the following way:


The code accepts an incoming TCP connection on port 4000 and passes the connection off to the handler. The handler then receives the header data from the connection and decodes the stream of bytes into a Header struct. From there it examines the incoming request code and calls the appropriate function. In the event that no known code was passed in, it calls the error function. 

One assumption I made was that the ping function would not fail since we are controlling the server. This function simply returns the correct code and a message that says that the server is up and running. 

The GetStats functions takes the TotalBytesReceived and TotalBytesSent global variables and creates a Stats struct, which then is returned along with the correct code. 

The ResetStats function resets the global variables back to 0.

The compress function takes in a payload and performs the following algorithm:

       First I check if the string contains any invalid characters and return out of the function if it does so. If the number of characters is less than 3 , I just return the original string. 
       Otherwise we iterate through the string character by character and increment a count variable until the i + 1 character does not equal the i'th character. Once this happens we drop into the else statement which appends the character being compress and the count to the output string. Finally I perform this check one more time at the very end of the function to take the last character into account. 


 After the respective function has completed , I replace the request code of the header with the status code of the function. After that has completed I convert the header into a stream of bytes and send it back over to the client, incrementing the totalBytesSent globalVariable after doing so. Since a few of these request also involve sending other data back to the client I encode the payload as well and send that back over. 
After this has completed I close the connection and exit out of the goroutine.    