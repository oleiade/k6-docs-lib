---
title: 'Socket.send(data)'
description: 'Send a data string through the connection.'
---

# Socket.send(data)


{{< admonition type="note" >}}

A module with a better and standard API exists.
<br>
<br>
The new [k6/experimental/websockets API](/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/websockets/) partially implements the [WebSockets API living standard](https://websockets.spec.whatwg.org/).
<br>
<br>
When possible, we recommend using the new API. It uses a global event loop for consistency with other k6 APIs and better performance.

{{< /admonition >}}


Send a data string through the connection.
You can use `JSON.stringify` to convert a JSON or JavaScript values to a JSON string.

| Parameter | Type   | Description       |
| --------- | ------ | ----------------- |
| data      | string | The data to send. |

### Example

{{< code >}}

```javascript
import ws from 'k6/ws';

export default function () {
  const url = 'wss://echo.websocket.org';
  const response = ws.connect(url, null, function (socket) {
    socket.on('open', function () {
      socket.send('my-message');
      socket.send(JSON.stringify({ data: 'hola' }));
    });
  });
}
```

{{< /code >}}

- See also [Socket.sendBinary(data)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-ws/socket/socket-sendbinary)
