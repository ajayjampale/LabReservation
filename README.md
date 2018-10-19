# LAB Reservation Application (Work In Progress)

The Application helps collaborate Software Dev/Test Teams to effectively use LAB resources in a time shared manner.

## Setup Development(Window) Environment 

1. Installing Go

   Download go(latest) package from https://golang.org/dl/
   
2. Install Msys

   http://www.msys2.org/
   
3. Install git using msys pacman

   pacman -S git
   
4. Pull the code from the repository
 
   git pull https://github.com/ajayjampale/LabReservation.git
   
5. Setup PATH and GOPATH env variable.

   - PATH env variable should point to GO installation.
      
        export PATH=$PATH:/c/Go/bin
   
   - GOPATH env variable should point to the LabReservation Project.
   
        export GOPATH="<Path_to_project>/LabReservation"

6. Setting up mongoDB.

   - Install mongodb from the below link:
   
        https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/#install-mdb-edition
     
   - Start mongodb from the cmd line :
   
        c:\mongodb\bin>mongod.exe --dbpath="c:\mongodb\data\db" --logpath c:\mongodb\log\mongo.log
   
   - To start mongo shell (for debugging and understanding)
   
        c:\mongodb\bin>mongo.exe
        
   - Basic CRUD operations on mongoDB collection is demonstrated in the file mongodb_helloworld.txt
