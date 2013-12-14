# irc2pusher

Send IRC messages to a Pusher channel

## Usage

```bash
export PUSHER_ID=12345
export PUSHER_KEY=key
export PUSHER_SECRET=secret

irc2pusher -h irc.freenode.org -p 6667 -n bot -c #linux
```

irc2pusher will send JSON messages in the following format:

```json
{
  "nick": "bot",
  "channel": "#linux",
  "message": "hello there"
}
```