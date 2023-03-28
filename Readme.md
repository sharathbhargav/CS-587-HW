Instructions to start the server:

- Copy the files banking-client.go, banking-server.go, Makefile, myRpc.t to a folder accessible by multiple users. 
- Run `make install` in the folder where all the above files are present.
- Change to directory `client` and run `sudo -E ethosRun` . This starts the RPC server
- Open another terminal with same user
- Navigate to the `client` directory 
- Run `etAl client.ethos`. This will move the prompt to ethos VM
- Run `syncClient` which will start the client that communicates to server using RPC.
- The above step (client) can be imitated with other users.