# Superdense Coding

Transmit two bits of classical information in one qubit! (sort of, maybe)

### Steps
1. Eve generates an entangled qubit pair, one goes to alice, one goes to bob
2. alice encodes her 2-bit message on her single qubit
3. alice sends her qubit to bob
4. Bob is able to decode the message

### How it works

Depending on the message Alice wants to send, she performs some X,Z gates on her qubit. This will modify both |0> to |1> (and vice versa), and the phase |0> to -|0>. 

Using this Bob is able to decode the message.

### Running

```
go run superdense.go
```