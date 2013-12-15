# irc2pusher

Send IRC messages to a Pusher channel

## Usage

Set pusher application settings:

```bash
export PUSHER_ID=12345
export PUSHER_KEY=key
export PUSHER_SECRET=secret
```

Run daemon:

```bash
irc2pusher -s irc.freenode.org -p 6667 -n bot -c linux
irc2pusher -s irc.freenode.org -c "channel1,channel2"
```

Messages will be sent in JSON format:

```json
{
  "nick": "bot",
  "channel": "#linux",
  "message": "hello there"
}
```

Extra options:

```bash
# Set a different pusher channel, defaults to "irc"
export PUSHER_CHANNEL=mychannel

# Set a different pusher event name, defaults to "message"
export PUSHER_EVENT=myevent
```

## Compile

Clone repository and install dependencies:

```bash
git clone https://github.com/sosedoff/irc2pusher.git
cd irc2pusher
go get
```

Build:

```bash
go build
```

## License

The MIT License (MIT)

Copyright (c) 2013 Dan Sosedoff, <dan.sosedoff@gmail.com>