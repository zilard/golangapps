# Build and run in Docker container

     docker compose build
     docker compose up


# Build

     make all


# Run API server

     ./bin/ethquery --apitoken YOUR-INFURA-API-KEY --port 8080

 or

     ./bin/ethquery -t YOUR-INFURA-API-KEY -p 8080    


# Send a REST API query

    curl 127.0.0.1:8080/api/getethblocknumber

    latest block number is: 0x107056d
