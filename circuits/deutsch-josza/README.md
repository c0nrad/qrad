# Deutsch Josza Algorithm

A nifty little quantum computing trick for determining if a set of quantum gates is "balanced" or "constant".

I don't think it has any practical uses. 

The gates are constant if they always return 0, or always return 1. They are balanced if they return exactly half 0 and half 1 for all inputs.

### Commentary

What blows my mind is that the function is applied onto the last qubit, but nothing is even done with the last qubit. I think it has to do with phase kickback?. But I didn't think the X gate applied kickback. 

(So to be honest I'm not really sure how it works.) I can do the math, and the math "makes sense", but I don't have any intuition as to why.

One of the books mentioned that the line of Hadamards is akin to a simple quantum fourier transform. So there must be some phase stuff in there.

### Running the program

```
go run deutsch-josza.go
```