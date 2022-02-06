# Buidling upon the code from the previous problem:

Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".  
Pass the connection of type net.Conn as an argument to the function.

Add "go" infront of the call to "serve" to enable concurrency and multiple connections.
