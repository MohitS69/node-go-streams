consuming the stream from the other end `pipe` will not throw any error.
`pipeline` will throw error `ERR_STREAM_PREMATURE_CLOSE` so we have to wrap it
inside a trycatch  block cause it supports await syntax and it is meant when we want to
consume the whole stream.