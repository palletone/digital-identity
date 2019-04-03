# digital-identity

Digital identity management relies on the fabric ca server for the distribution of digital certificates.

This project provides CA client support for Palletone.



# To Do

- Do some preparation before using them.

- Install the server  binary commands into the $GOPATH/bin directory

````
$ go get -u github.com/hyperledger/fabric-ca/cmd/...
````

- Switch to the source directory:

````
$ cd $GOPATH/src/github.com/hyperledger/fabric-ca/
````

- Compile using the make command:

````
$ make fabric-ca-server
````

- Generate the bin directory, which contains the fabric-ca-server executable.

- Setting environment variables

````
$ export PATH=$GOPATH/src/github.com/hyperledger/fabric-ca/bin:$PATH
````

# *initialize* 

- Return to the user directory

````
$ cd ~
$ mkdir cawork
$ cd cawork
$ mkdir root immediateca
$ cd root
````

- Init    

````
$ fabric-ca-server init -b admin:pass
````

- Start   

````
$ fabric-ca-server start -b admin:pass
````

- Start   immediateca

​       $  cd immediateca

````
$ fabric-ca-server start -b admin:pass -p 7064 -u http://admin:pass@localhost:7054
````

! If port occupancy occurs, modify the configuration file.

