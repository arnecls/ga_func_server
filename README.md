# ga_func_server
A very simplisitc demo RPC-Server project as a Games Academy task endpoint.  
Start with `go build && ./ga_func_server`.  
Once started it listens on port 5880 (tcp).

## Protocol
The protocol is binary (little endian) and is defined in protocol.go

```
request     := [header][data]  
header      := [function id][data count]  
function id := uint32  
data count  := uint32  
data        := []float32  

response    := [result][error code]  
result      := float32  
error code  := uint32
```

Please note that <function id> currently has to be in the range of [0..2] where

```
0 = Median  
1 = Average  
2 = Sum
3 = Number of Cores reported (request data is ignored)
4 = Memory used by service (request data is ignored)
```

The response may have one of the following error codes

```
0 = All OK  
1 = Error reading request  
2 = Error processing data   
```
