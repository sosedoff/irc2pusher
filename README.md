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

### Pusher environment variables

- `PUSHER_ID` - Application ID (required)
- `PUSHER_KEY` - Application access key (required)
- `PUSHER_SECRET` - Application secret key (required)
- `PUSHER_CHANNEL` - Channel to send events to (optional, default: "irc")
- `PUSHER_EVENT` - Message event name (optional, default: "message")

### Options

```
Usage:
  irc2pusher [OPTIONS]

Application Options:
  -s, --server=   IRC server hostname or IP
  -p, --port=     IRC server port
  -n, --nick=     Nickname
  -c, --channels= Channels to join

Help Options:
  -h, --help      Show this help message
```

## Compile

Clone repository and install dependencies:

```bash
git clone https://github.com/sosedoff/irc2pusher.git
cd irc2pusher
go get
```

Make sure you have Go >= 1.1 installed. Build:

```bash
go build
```

## License

The MIT License (MIT)

Copyright (c) 2013 Dan Sosedoff, <dan.sosedoff@gmail.com>