---
title: 'k6/ws'
description: 'k6 WebSocket API'
weight: 11
---

# k6/ws


{{< admonition type="note" >}}

A module with a better and standard API exists.
<br>
<br>
The new [k6/experimental/websockets API](/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/websockets/) partially implements the [WebSockets API living standard](https://websockets.spec.whatwg.org/).
<br>
<br>
When possible, we recommend using the new API. It uses a global event loop for consistency with other k6 APIs and better performance.

{{< /admonition >}}



The [`k6/ws` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws) provides a [WebSocket](https://en.wikipedia.org/wiki/WebSocket) client implementing the [WebSocket protocol](http://www.rfc-editor.org/rfc/rfc6455.txt).

| Function                                                                                                  | Description                                                                                                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [connect( url, params, callback )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/connect) | Create a WebSocket connection, and provides a [Socket](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket) client to interact with the service. The method blocks the test finalization until the connection is closed. |

| Class/Method                                                                                                                      | Description                                                                                                                                                                    |
| --------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [Params](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/params)                                                    | Used for setting various WebSocket connection parameters such as headers, cookie jar, compression, etc.                                                                        |
| [Socket](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket)                                                    | WebSocket client used to interact with a WS connection.                                                                                                                        |
| [Socket.close()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-close)                               | Close the WebSocket connection.                                                                                                                                                |
| [Socket.on(event, callback)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-on)                      | Set up an event listener on the connection for any of the following events:<br />- open<br />- binaryMessage<br />- message<br />- ping<br />- pong<br />- close<br />- error. |
| [Socket.ping()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-ping)                                 | Send a ping.                                                                                                                                                                   |
| [Socket.send(data)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-send)                             | Send string data.                                                                                                                                                              |
| [Socket.sendBinary(data)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-sendbinary)                 | Send binary data.                                                                                                                                                              |
| [Socket.setInterval(callback, interval)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-setinterval) | Call a function repeatedly at certain intervals, while the connection is open.                                                                                                 |
| [Socket.setTimeout(callback, period)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-settimeout)     | Call a function with a delay, if the connection is open.                                                                                                                       |

