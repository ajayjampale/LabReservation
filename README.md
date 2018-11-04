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

   - Once familiar with above mongoDB queries, you can use mongo Compass GUI tool to View/query/delete collections
   
        https://www.mongodb.com/download-center?initial=true#compass
        
7. Setting up swagger and generate REST client code, server code and API Documentation.

        https://github.com/swagger-api/swagger-codegen

   - Installation
 
         wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar -O ~/bin/swagger-codegen-cli.jar

         Set your $PATH to point to java binary. C:\Program Files\Docker\Docker\resources\bin:C:\Program Files (x86)\Java\jre7\bin
         
         java -jar ~/bin/swagger-codegen-cli.jar help
         
         Set an alias in bashrc to have this tool handy.
         alias swagger-codegen="java -jar ~/bin/swagger-codegen-cli.jar"  

   - Use the online yaml editor to view and modify the openapi.yaml
        
        https://editor.swagger.io/ 

   - Generating HTML Documentation

    swagger-codegen generate -i ./src/app-server/infra/rest/openapi.yaml -l html -o ./src/app-server/infra/rest/generated/html

   - Generating Golang Client code

        swagger-codegen generate -i ./src/app-server/infra/rest/openapi.yaml -l go -o ./src/app-server/infra/rest/generated/client

   - Generating Golang Server code
   
    swagger-codegen generate -i ./src/app-server/infra/rest/openapi.yaml -l go-server -o ./src/app-server/infra/rest/generated/server
