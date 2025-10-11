# Titly

Titly is a blazingly fast URL shortner written in go.

## Features

- Lightning-fast URL shortening
- Handles 10,000+ requests per second
- Supports high concurrency via goroutines
- Redis caching for persistent short codes

## Benchmarks

> Benchmarks are measured on localhost; production performance may vary depending on database and network latency.

<details>
  <summary>Apache Benchmark</summary>

  ```bash
    This is ApacheBench, Version 2.3 <$Revision: 1923142 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 1000 requests
    Completed 2000 requests
    Completed 3000 requests
    Completed 4000 requests
    Completed 5000 requests
    Completed 6000 requests
    Completed 7000 requests
    Completed 8000 requests
    Completed 9000 requests
    Completed 10000 requests
    Finished 10000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            4000

    Document Path:          /
    Document Length:        31 bytes

    Concurrency Level:      100
    Time taken for tests:   0.548 seconds
    Complete requests:      10000
    Failed requests:        0
    Total transferred:      1540000 bytes
    HTML transferred:       310000 bytes
    Requests per second:    18233.63 [#/sec] (mean)
    Time per request:       5.484 [ms] (mean)
    Time per request:       0.055 [ms] (mean, across all concurrent requests)
    Transfer rate:          2742.17 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    2   1.0      2       8
    Processing:     0    4   3.1      3      36
    Waiting:        0    3   3.0      2      35
    Total:          0    5   3.3      5      37

    Percentage of the requests served within a certain time (ms)
      50%      5
      66%      5
      75%      5
      80%      6
      90%      7
      95%     13
      98%     17
      99%     21
    100%     37 (longest request)
 ```
</details>

<details>
  <summary>wrk</summary>

  ```bash
    Running 30s test @ http://localhost:4000/
      4 threads and 100 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     1.21ms    0.93ms  23.04ms   69.90%
        Req/Sec    22.34k     1.29k   25.53k    77.92%
      2667393 requests in 30.01s, 391.75MB read
    Requests/sec:  88894.04
    Transfer/sec:     13.06MB
  ```
  
</details>

## Tech Stack

- Go
- Redis
- Sveltekit

## Running with Docker

To start both the server and client using Docker Compose, run:

```bash
docker compose up --build
```

**Client Port:** `4173`  
**Server Port:** `4000`
**Redis Port:** `6379`

![image](https://github.com/user-attachments/assets/197a892f-9376-4958-896e-f6e7e5b416b9)
