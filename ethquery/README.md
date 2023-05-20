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


# Run Load Test

Install k6:

https://k6.io/docs/get-started/installation/

    sudo gpg -k
    sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
    echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
    sudo apt-get update
    sudo apt-get install k6


Run k6 API Load Test:

https://k6.io/docs/get-started/running-k6/


     cd k6apiloadtest

     k6 run script.js


     golangapps/ethquery/k6apiloadtest$ k6 run script.js 

             /\      |‾‾| /‾‾/   /‾‾/   
        /\  /  \     |  |/  /   /  /    
       /  \/    \    |     (   /   ‾‾\  
      /          \   |  |\  \ |  (‾)  | 
     / __________ \  |__| \__\ \_____/ .io

     execution: local
          script: script.js
          output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 1m0s max duration (incl. graceful stop):
               * default: 10 looping VUs for 30s (gracefulStop: 30s)


          data_received..................: 41 kB 1.3 kB/s
          data_sent......................: 27 kB 882 B/s
          http_req_blocked...............: avg=10.62µs  min=1.18µs   med=6.77µs   max=135.49µs p(90)=10.34µs  p(95)=23.12µs 
          http_req_connecting............: avg=2.76µs   min=0s       med=0s       max=99.09µs  p(90)=0s       p(95)=0s      
          http_req_duration..............: avg=134.11ms min=112.88ms med=119.13ms max=458.32ms p(90)=177.1ms  p(95)=213.56ms
          { expected_response:true }...: avg=134.11ms min=112.88ms med=119.13ms max=458.32ms p(90)=177.1ms  p(95)=213.56ms
          http_req_failed................: 0.00% ✓ 0        ✗ 270 
          http_req_receiving.............: avg=84.33µs  min=7.6µs    med=81.94µs  max=290.39µs p(90)=115.35µs p(95)=141.53µs
          http_req_sending...............: avg=28.3µs   min=4.01µs   med=26.88µs  max=108.85µs p(90)=41.55µs  p(95)=47.86µs 
          http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s      
          http_req_waiting...............: avg=134ms    min=112.76ms med=119.04ms max=458.17ms p(90)=176.96ms p(95)=213.34ms
          http_reqs......................: 270   8.735383/s
          iteration_duration.............: avg=1.13s    min=1.11s    med=1.11s    max=1.45s    p(90)=1.17s    p(95)=1.21s   
          iterations.....................: 270   8.735383/s
          vus............................: 10    min=10     max=10
          vus_max........................: 10    min=10     max=10


     running (0m30.9s), 00/10 VUs, 270 complete and 0 interrupted iterations
     default ✓ [======================================] 10 VUs  30s

