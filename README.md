# Debouncer

No callbacks, only channels!

```go

// noise refers to the input channel we want to debounce to
// and debounced is the output channel that tiggers a message
// when debouncing has finished
go Debounce(noise, debounced, time.Second)
```

That's it.
