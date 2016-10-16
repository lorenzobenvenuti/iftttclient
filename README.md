# iftttclient
A client for the IFTTT Maker APIs. Just a simple use case for the [ifttt package](https://github.com/lorenzobenvenuti/ifttt) (and a bit more handy than using `curl` in your scripts...)

```bash
iftttclient -key myApiKey -event myEvent -value value1 -value value2 -value value3
```

Or 

```bash
export IFTTT_API_KEY=myApiKey 
iftttclient -event myEvent -value value1 -value value2 -value value3
```
