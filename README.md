# Mastering WebSockets With Go

___Tutorial on how to use WebSockets to build real-time APIs in Go___

[link](https://towardsdatascience.com/mastering-websockets-with-go-c30d0ac48081)


1. frontend: `HTML`+`Javascript`
2. websocket: `gorilla\webSocket`
3. WebSocket connection in Go only allows one concurrent writer, so we can handle this by using an unbuffered channel
4. The HeartBeats — Ping & Pong