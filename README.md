# PortScanner
A multi-threaded port scanner made in Go.
## Requirements
* Go (latest)

## Installation
```
go build ./src/scanner
```

## Usage
```
./scanner.exe -threads=1 -protocol=tcp -host=127.0.0.1 -ports=22,80,443
```
