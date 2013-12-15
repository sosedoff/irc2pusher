# irc2pusher

Send IRC messages to a Pusher channel

## Usage

Set pusher application settings:

```
export PUSHER_ID=12345
export PUSHER_KEY=key
export PUSHER_SECRET=secret
```

Run daemon:

```
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

## Compile

Clone repository and install dependencies:

```
git clone https://github.com/sosedoff/irc2pusher.git
cd irc2pusher
go get
```

Build:

```
go build
```

## License

The MIT License (MIT)

Copyright (c) 2013 Dan Sosedoff, <dan.sosedoff@gmail.com>