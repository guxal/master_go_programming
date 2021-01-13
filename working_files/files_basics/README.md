# files basics


### OPENING a FILE WITH MORE OPTIONS

```go
file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
```

We can Use opening attributes individually or combined using an OR between them e.g.
>  `os.O_CREATE|os.O_APPEND`
or `os.O_CREATE|os.O_TRUNC|os.O_WRONLY`
-----------
### Options

|attributes|values        |
|---------|---------------|
|`os.O_RDONLY`| Read only |
|`os.O_WRONLY`| Write only |
|`os.O_RDWR` | Read and write |
|`os.O_APPEND`| Append to end of file |
|`os.O_CREATE`| Create is none exist |
|`os.O_TRUNC` | Truncate file when opening |
