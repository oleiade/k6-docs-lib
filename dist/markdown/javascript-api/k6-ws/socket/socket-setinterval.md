---
title: 'Socket.setInterval(callback, interval)'
description: 'Call a function repeatedly, while the WebSocket connection is open.'
---

# Socket.setInterval(callback, interval)


{{< admonition type="note" >}}

A module with a better and standard API exists.
<br>
<br>
The new [k6/experimental/websockets API](/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/websockets/) partially implements the [WebSockets API living standard](https://websockets.spec.whatwg.org/).
<br>
<br>
When possible, we recommend using the new API. It uses a global event loop for consistency with other k6 APIs and better performance.

{{< /admonition >}}


Call a function repeatedly, while the WebSocket connection is open.

| Parameter | Type     | Description                                                 |
| --------- | -------- | ----------------------------------------------------------- |
| callback  | function | The function to call every `interval` milliseconds.         |
| interval  | number   | The number of milliseconds between two calls to `callback`. |

### Example

{{< code >}}

```javascript
import ws from 'k6/ws';
import { check } from 'k6';

export default function () {
  const url = 'wss://echo.websocket.org';
  const params = { tags: { my_tag: 'hello' } };

  const res = ws.connect(url, params, function (socket) {
    socket.on('open', function open() {
      console.log('connected');

      socket.setInterval(function timeout() {
        socket.ping();
        console.log('Pinging every 1sec (setInterval test)');
      }, 1000);
    });

    socket.on('pong', function () {
      console.log('PONG!');
    });
  });

  check(res, { 'status is 101': (r) => r && r.status === 101 });
}
```

{{< /code >}}
