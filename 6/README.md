Concurrency
======

The first thing I did was create an inline version, where it runs each, and waits for the next one. This is in [inline.go](./inline.go).

Then I made a concurrent version, that runs each in a go routine, and then waits for them all to ruturn. This is in [concurrancy.go](./concurrancy.go).

Then to figure out if it was spending more or less time (less time was expected) I created a bash script to run them both and added timing.  
  
    . ./run_concurrancy.sh

Example output on my machine, with this URLS:

```
inline
starting
https://mikelynchgames.com is text/html; charset=UTF-8
https://golang.org is text/html; charset=utf-8
https://api.github.com is application/json; charset=utf-8
https://httpbin.org/xml is application/xml
go run inline.go  0.38s user 0.12s system 14% cpu 3.334 total

concurrent
starting
https://golang.org is text/html; charset=utf-8
https://api.github.com is application/json; charset=utf-8
https://httpbin.org/xml is application/xml
https://mikelynchgames.com is text/html; charset=UTF-8
go run concurrancy.go  0.42s user 0.13s system 33% cpu 1.653 total
```

These results are interesting, as it takes longer with the go routines! Depending on the sites, and how long they take to load, I can have the inline or concurrent win.

I think the reason for the result is that the sites take different amonuts of time to load, and with one long running and many almost instantanous loads, you don't see much difference, or even slower concurrent as it has more overhead.


Channels
===

At first I noticed that channels are like observables, or callbacks. Once you make the channel, 
it waits for input, and you can send values to it continuously.

