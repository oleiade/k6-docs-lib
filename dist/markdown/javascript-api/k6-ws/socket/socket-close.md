---
title: 'Socket.close([code])'
description: 'Close the WebSocket connection.'
---

# Socket.close([code])


{{< admonition type="note" >}}

A module with a better and standard API exists.
<br>
<br>
The new [k6/experimental/websockets API](/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/websockets/) partially implements the [WebSockets API living standard](https://websockets.spec.whatwg.org/).
<br>
<br>
When possible, we recommend using the new API. It uses a global event loop for consistency with other k6 APIs and better performance.

{{< /admonition >}}


Close the WebSocket connection.

| Parameter       | Type   | Description                            |
| --------------- | ------ | -------------------------------------- |
| code (optional) | number | WebSocket status code. (default: 1001) |

### Example

{{< code >}}

```javascript
import ws from 'k6/ws';

export default function () {
  const url = 'wss://echo.websocket.org';
  const response = ws.connect(url, null, function (socket) {
    socket.on('open', function () {
      socket.close();
    });
  });
}
```

{{< /code >}}
