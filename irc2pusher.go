package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
  "net/textproto"
  "os"
  "os/signal"
  "strings"
  "encoding/json"
  "github.com/timonv/pusher"
)

type IrcClient struct {
  Server   string
  Port     string
  Nick     string
  Channels string

  Socket net.Conn
  Pusher *pusher.Client
}

type IrcMessage struct {
  Nick    string `json:"nick"`
  Channel string `json:"channel"`
  Message string `json:"message"`
}

func parseMessage(line string) *IrcMessage {
  parts   := strings.Split(line, "PRIVMSG")
  nick    := strings.Replace(strings.Split(parts[0], "!")[0], ":", "", 1)
  channel := strings.TrimSpace(strings.Split(parts[1], ":")[0])
  message := strings.TrimSpace(strings.Replace(parts[1], channel + " :", "", 1))

  return &IrcMessage{ nick, channel, message }
}

func handleSignals(irc *IrcClient) {
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  
  for sig := range c {
    if sig == os.Interrupt {
      log.Printf("Received os.Interrupt, exiting normally.\n\n")
      irc.Send("QUIT :\n")
      irc.Socket.Close()
      os.Exit(0)
    }
  }
}

func (irc *IrcClient) Connect() {
  target := fmt.Sprintf("%s:%s", irc.Server, irc.Port)

  socket, err := net.Dial("tcp", target)
  if err != nil {
    log.Fatalf("Unable to connect to %s\nError: %v\n\n", target, err)
  }

  log.Printf("Successfully connected to %s\n", target)

  irc.Socket = socket
  irc.Send("USER " + irc.Nick + " 8 * :" + irc.Nick + "\n")
  irc.Send("NICK " + irc.Nick + "\n")

  channels := strings.Split(irc.Channels, " ")
  for _, name := range(channels) {
    irc.Send("JOIN " + name + "\n")
  }
}

func (irc *IrcClient) Send(str string) {
  data := []byte(str)

  log.Printf(str)

  _, err := irc.Socket.Write(data)
  if err != nil {
    log.Printf("Error: %v", err)
  }
}

func (irc *IrcClient) respondToPing(str string) {
  chunks := strings.Split(str, " ")
  pong := fmt.Sprintf("PONG %s\n", chunks[1])

  irc.Send(pong)
}

func (irc *IrcClient) sendToPusher(message *IrcMessage) {
  data, err := json.Marshal(message)

  if err != nil {
    fmt.Println("JSON encode error:", err)
    return
  }

  irc.Pusher.Publish(string(data), "message", "irc")
}

func (irc *IrcClient) handleLine(line string) {
  log.Printf(line)

  if strings.Contains(line, "PING") {
    irc.respondToPing(line)
    return
  }

  if strings.Contains(line, "PRIVMSG") {
    msg := parseMessage(line)
    irc.sendToPusher(msg)
    return
  }
}

func (irc *IrcClient) Run() {
  go handleSignals(irc)

  reader := bufio.NewReader(irc.Socket)
  tp := textproto.NewReader(reader)

  for {
    line, err := tp.ReadLine()
    
    if err != nil {
      log.Printf("Error reading line: %s\n", line)
      log.Fatalf("Error: %v\n", err)
    } else {
      irc.handleLine(line)
    }
  }
}

func (irc *IrcClient) InitClient() {
  irc.Server   = "irc.freenode.org"
  irc.Port     = "6667"
  irc.Nick     = "irc2pusher"
  irc.Channels = "#irc2pusher"
}

func (irc *IrcClient) InitPusher() {
  id     := os.Getenv("PUSHER_APP_ID")
  key    := os.Getenv("PUSHER_APP_KEY")
  secret := os.Getenv("PUSHER_APP_SECRET")

  if len(id) == 0 {
    fmt.Println("PUSHER_APP_ID env variable is not set")
    os.Exit(1)
  }

  if len(key) == 0 {
    fmt.Println("PUSHER_APP_KEY env variable is not set")
    os.Exit(1)
  }

  if len(secret) == 0 {
    fmt.Println("PUSHER_APP_SECRET env variable is not set")
    os.Exit(1)
  }

  irc.Pusher = pusher.NewClient(id, key, secret)
}

func main() {
  irc := new(IrcClient)

  irc.InitClient()
  irc.InitPusher()
  irc.Connect()
  irc.Run()
}