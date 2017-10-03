# iftttclient

A client for the IFTTT Maker APIs.
I wrote this command to be notified when certain long tasks are finished (or they fail):

```
$ ./perform-long-task.sh && iftttclient trigger notify "Task finished" || iftttclient trigger notify "Task failed"
```

## Usage

```
$ iftttclient trigger -k my-api-key my-event [my-value1 my-value2 my-value3]
```

API key can be specified using an environment variable:

```
$ export IFTTT_API_KEY=my-api-key
$ iftttclient trigger my-event [my-value1 my-value2 my-value3]
```

Or you can use the `iftttclient store` command to store it in your home directory:

```
$ iftttclient store my-api-key
$ iftttclient trigger my-event [my-value1 my-value2 my-value3]
```

Please notice that, for now, you API key is written in plain text (in `$HOME/.iftttclient/config.json`)

## Downloads

Binaries [here](https://github.com/lorenzobenvenuti/iftttclient/releases).

## TODO

* Support encryption for locally stored API key
