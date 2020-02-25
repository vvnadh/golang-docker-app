# golang-docker-app
A docker based Go application that makes a http call to MAC IO API and pulls the MAC address vendor details.
Look here for more info. https://macaddress.io/


It is simple Golang based application that takes few command line arguments such as apikey, macaddress etc,  
submits http request to macaddress.io, parsed the retrieved JSON based response and dumps the inforamtion in stdout.

Actual API call to macaddress.io: https://api.macaddress.io/v1?output=json&search=44:38:39:ff:ef:57.  
The program passes APIKEY via HeaderBased Authentication, i.e X-Authentication-Token.

Here are the instructions to use this program.

Prequisites:

  *) You need to sign into macaddress.io --> API to get the API key for your login.  
     This key you need to pass via cli using -apikey

Setup:

  *) Once you clone, you can just run below command to build the docker image.
     sh docker_build.sh
     
  *) Run build script that dumps the docker run command. 
     sh docker_run.sh

  *) The list of argument you can pass to docker which inturn passed to go excutable in docker container.
  
Various commands
 
      #docker run -it golang-mac-api -h
      
      Output:

            INFO: Running MAC API...
            INFO: ARGS passed -h
            Usage of ./main:
              -apikey string
                  a string : Mandatory
              -macaddress string
                  a string (default "44:38:39:ff:ef:57")

      #docker run -it golang-mac-api -apikey='YOUR_OWN_APIKEY'
      
      Output:

            DEBUG: Making HTTP Get for  https://api.macaddress.io/v1?output=json&search=44:38:39:ff:ef:57
            DEBUG: Response body={"vendorDetails":{"oui":"443839","isPrivate":false,"companyName":"Cumulus Networks, Inc","companyAddress":"650 Castro Street, suite 120-245 Mountain View CA 94041 US","countryCode":"US"},"blockDetails":{"blockFound":true,"borderLeft":"443839000000","borderRight":"443839FFFFFF","blockSize":16777216,"assignmentBlockSize":"MA-L","dateCreated":"2012-04-08","dateUpdated":"2015-09-27"},"macAddressDetails":{"searchTerm":"44:38:39:ff:ef:57","isValid":true,"virtualMachine":"Not detected","applications":["Multi-Chassis Link Aggregation (Cumulus Linux)"],"transmissionType":"unicast","administrationType":"UAA","wiresharkNotes":"No details","comment":""}}

            vendorDetails.oui=443839
            vendorDetails.isPrivate=false
            vendorDetails.companyName=Cumulus Networks, Inc
            vendorDetails.companyAddress=650 Castro Street, suite 120-245 Mountain View CA 94041 US
            vendorDetails.countryCode=US
            blockDetails.borderRight=443839FFFFFF
            blockDetails.blockSize=1.6777216e+07
            blockDetails.assignmentBlockSize=MA-L
            blockDetails.dateCreated=2012-04-08
            blockDetails.dateUpdated=2015-09-27
            blockDetails.blockFound=true
            blockDetails.borderLeft=443839000000
            macAddressDetails.wiresharkNotes=No details
            macAddressDetails.comment=
            macAddressDetails.searchTerm=44:38:39:ff:ef:57
            macAddressDetails.isValid=true
            macAddressDetails.virtualMachine=Not detected
            macAddressDetails.applications.0=Multi-Chassis Link Aggregation (Cumulus Linux)
            macAddressDetails.transmissionType=unicast
            macAddressDetails.administrationType=UAA

       #docker run -it golang-mac-api -apikey='YOUR_OWN_APIKEY' -macaddress 'YOUR_OWN_MAC'


