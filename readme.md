# Pyorha

Pyorha (별하) is static serving tooling

# Usage
* Next.js
```sh
$npx next build
$npx next export
$pyorha build ./out out
$pyorha serve out
```

# Build
```sh
$go build ./cmd/pyorha
```

# Requirement
* go1.19
* gcc

# Performance

```sh
$gobench -u http://localhost:3000 -k=true -c 500 -t 10
```

* next.js + Pyorha static serving (Hello World website)
```
Requests:                          2329810 hits
Successful requests:               2329810 hits
Network failed:                          0 hits
Bad requests failed (!2xx):              0 hits
Successful requests rate:             3222 hits/sec
Read throughput:                  14133681 bytes/sec
Write throughput:                   277191 bytes/sec
Test time:                             723 sec
```

* next.js + echo static serving (Hello World website)
```
Requests:                          1826186 hits
Successful requests:               1826186 hits
Network failed:                          0 hits
Bad requests failed (!2xx):              0 hits
Successful requests rate:             2536 hits/sec
Read throughput:                   3748754 bytes/sec
Write throughput:                   218187 bytes/sec
Test time:                             720 sec
```
### File Size

* next.js + Pyorha static serving (Hello World website)
> 128KB
* next.js + echo static serving (Hello World website)
> 375KB