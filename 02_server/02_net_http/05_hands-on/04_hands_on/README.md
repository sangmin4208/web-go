# Create a basic serbver using TCP

The sever should use net.Listen to listen on port 8080.  
Remeber to close the listener using defer.  
Remeber that from the "net" package you first need to Listen, then you need to ACCEPT an incoming connection.  
Now write a response back on the connection.  
Use io.WritingString to write the response: I see you connected.  
Remember to close the connection.  
Once you have all of that workingm run your TCP server and test it from telnet (telnet localhost 8080).
