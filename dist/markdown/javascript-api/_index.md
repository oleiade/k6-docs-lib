---
title: JavaScript API
weight: 600
---

# JavaScript API

The list of k6 modules natively supported in your k6 scripts.

## Init context


Before the k6 starts the test logic, code in the _init context_ prepares the script.
A few functions are available only in init context.
For details about the runtime, refer to the [Test lifecycle](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/test-lifecycle).

| Function                                                                                              | Description                                          |
| ----------------------------------------------------------------------------------------------------- | ---------------------------------------------------- |
| [open( filePath, [mode] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/init-context/open) | Opens a file and reads all the contents into memory. |


## import.meta


`import.meta` is only available in ECMAScript modules, but not CommonJS ones.

| Function                                                                                           | Description                                               |
| -------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| [import.meta.resolve](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/import.meta/resolve) | resolve path to URL the same way that an ESM import would |


## k6


The [`k6` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6) contains k6-specific functionality.

| Function                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| [check(val, sets, [tags])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6/check) | Runs one or more checks on a value and generates a pass/fail result but does not throw errors or otherwise interrupt execution upon failure. |
| [fail([err])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6/fail)               | Throws an error, failing and aborting the current VU script iteration immediately.                                                           |
| [group(name, fn)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6/group)          | Runs code inside a group. Used to organize results in a test.                                                                                |
| [randomSeed(int)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6/random-seed)    | Set seed to get a reproducible pseudo-random number using `Math.random`.                                                                     |
| [sleep(t)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6/sleep)                 | Suspends VU execution for the specified duration.                                                                                            |


## k6/browser

The [`k6/browser` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental) provides browser-level APIs to interact with browsers and collect frontend performance metrics as part of your k6 tests.


| Method                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [browser.closeContext()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/closecontext)                                   | Closes the current [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                          |
| [browser.context()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/context)                                             | Returns the current [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                         |
| [browser.isConnected](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/isconnected) {{< docs/bwipt id="453" >}}           | Indicates whether the [CDP](https://chromedevtools.github.io/devtools-protocol/) connection to the browser process is active or not.                                                                                                                                                                                                                                                                                                                             |
| [browser.newContext([options])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/newcontext/) {{< docs/bwipt id="455" >}} | Creates and returns a new [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                   |
| [browser.newPage([options])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/newpage) {{< docs/bwipt id="455" >}}        | Creates a new [Page](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page) in a new [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext) and returns the page. Pages that have been opened ought to be closed using [`Page.close`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page/close). Pages left open could potentially distort the results of Web Vital metrics. |
| [browser.version()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/version)                                             | Returns the browser application's version.                                                                                                                                                                                                                                                                                                                                                                                                                       |



| k6 Class                                                                                                               | Description                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext) {{< docs/bwipt >}} | Enables independent browser sessions with separate [Page](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page)s, cache, and cookies. |
| [ElementHandle](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/elementhandle) {{< docs/bwipt >}}   | Represents an in-page DOM element.                                                                                                                       |
| [Frame](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/frame) {{< docs/bwipt >}}                   | Access and interact with the [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page).'s `Frame`s.                              |
| [JSHandle](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/jshandle)                                | Represents an in-page JavaScript object.                                                                                                                 |
| [Keyboard](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/keyboard)                                | Used to simulate the keyboard interactions with the associated [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page).        |
| [Locator](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator)                                  | The Locator API makes it easier to work with dynamically changing elements.                                                                              |
| [Mouse](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/mouse)                                      | Used to simulate the mouse interactions with the associated [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page).           |
| [Page](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page) {{< docs/bwipt >}}                     | Provides methods to interact with a single tab in a browser.                                                                                             |
| [Request](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/request) {{< docs/bwipt >}}               | Used to keep track of the request the [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page) makes.                           |
| [Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/response) {{< docs/bwipt >}}             | Represents the response received by the [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page).                               |
| [Touchscreen](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/touchscreen)                          | Used to simulate touch interactions with the associated [`Page`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page).               |
| [Worker](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/worker)                                    | Represents a [WebWorker](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API).                                                              |


## k6/crypto


The [`k6/crypto` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto) provides common hashing functionality available in the GoLang [crypto](https://golang.org/pkg/crypto/) package.

| Function                                                                                                                | Description                                                                                                                  |
| ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| [createHash(algorithm)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/createhash)                   | Create a Hasher object, allowing the user to add data to hash multiple times, and extract hash digests along the way.        |
| [createHMAC(algorithm, secret)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/createhmac)           | Create an HMAC hashing object, allowing the user to add data to hash multiple times, and extract hash digests along the way. |
| [hmac(algorithm, secret, data, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/hmac) | Use HMAC to sign an input string.                                                                                            |
| [md4(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/md4)                     | Use MD4 to hash an input string.                                                                                             |
| [md5(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/md5)                     | Use MD5 to hash an input string.                                                                                             |
| [randomBytes(int)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/randombytes)                       | Return an array with a number of cryptographically random bytes.                                                             |
| [ripemd160(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/ripemd160)         | Use RIPEMD-160 to hash an input string.                                                                                      |
| [sha1(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha1)                   | Use SHA-1 to hash an input string.                                                                                           |
| [sha256(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha256)               | Use SHA-256 to hash an input string.                                                                                         |
| [sha384(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha384)               | Use SHA-384 to hash an input string.                                                                                         |
| [sha512(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha512)               | Use SHA-512 to hash an input string.                                                                                         |
| [sha512_224(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha512_224)       | Use SHA-512/224 to hash an input string.                                                                                     |
| [sha512_256(input, outputEncoding)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/sha512_256)       | Use SHA-512/256 to hash an input string.                                                                                     |

| Class                                                                              | Description                                                                                                                                                                                           |
| ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Hasher](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/hasher) | Object returned by [crypto.createHash()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-crypto/createhash). It allows adding more data to be hashed and to extract digests along the way. |


## k6/data


The [`k6/data` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-data) provides helpers to work with data.

| Class/Method                                                                               | Description                                                   |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------- |
| [SharedArray](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-data/sharedarray) | read-only array like structure that shares memory between VUs |


## k6/encoding


The [`k6/encoding` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-encoding) provides [base64](https://en.wikipedia.org/wiki/Base64)
encoding/decoding as defined by [RFC4648](https://tools.ietf.org/html/rfc4648).

| Function                                                                                                                 | Description             |
| ------------------------------------------------------------------------------------------------------------------------ | ----------------------- |
| [b64decode(input, [encoding], [format])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-encoding/b64decode/) | Base64 decode a string. |
| [b64encode(input, [encoding])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-encoding/b64encode/)           | Base64 encode a string. |


## k6/execution


The [`k6/execution` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-execution) provides the capability to get information about the current test execution state inside the test script. You can read in your script the execution state during the test execution and change your script logic based on the current state.

`k6/execution` provides the test execution information with the following properties:

- [instance](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-execution#instance)
- [scenario](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-execution#scenario)
- [test](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-execution#test)
- [vu](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-execution#vu)


## k6/experimental

[`k6/experimental` modules](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental) are stable modules that may introduce breaking changes. Once they become fully stable, they may graduate to become k6 core modules.


| Modules                                                                                          | Description                                                                                                                |
| ------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------- |
| [csv](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/csv)               | Provides support for efficient and convenient parsing of CSV files.                                                        |
| [fs](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/fs)                 | Provides a memory-efficient way to handle file interactions within your test scripts.                                      |
| [streams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/streams)       | Provides an implementation of the Streams API specification, offering support for defining and consuming readable streams. |
| [websockets](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-experimental/websockets) | Implements the browser's [WebSocket API](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket).                      |


## k6/html


The [`k6/html` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html) contains functionality for HTML parsing.

| Function                                                                                    | Description                                                                                                                        |
| ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| [parseHTML(src)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html/parsehtml) | Parse an HTML string and populate a [Selection](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html/selection) object. |

| Class                                                                                  | Description                                                                                                                        |
| -------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| [Element](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html/element)     | An HTML DOM element as returned by the [Selection](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html/selection) API. |
| [Selection](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-html/selection) | A jQuery-like API for accessing HTML DOM elements.                                                                                 |


## k6/http


The [`k6/http` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http) contains functionality for performing HTTP transactions.

| Function                                                                                                                       | Description                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------- |
| [batch( requests )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/batch)                                     | Issue multiple HTTP requests in parallel (like e.g. browsers tend to do).                                                                 |
| [cookieJar()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/cookiejar-method)                                | Get active HTTP Cookie jar.                                                                                                               |
| [del( url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/del)                            | Issue an HTTP DELETE request.                                                                                                             |
| [file( data, [filename], [contentType] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/file)                | Create a file object that is used for building multi-part requests.                                                                       |
| [get( url, [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/get)                                    | Issue an HTTP GET request.                                                                                                                |
| [head( url, [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/head)                                  | Issue an HTTP HEAD request.                                                                                                               |
| [options( url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/options)                    | Issue an HTTP OPTIONS request.                                                                                                            |
| [patch( url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/patch)                        | Issue an HTTP PATCH request.                                                                                                              |
| [post( url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/post)                          | Issue an HTTP POST request.                                                                                                               |
| [put( url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/put)                            | Issue an HTTP PUT request.                                                                                                                |
| [request( method, url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/request)            | Issue any type of HTTP request.                                                                                                           |
| [asyncRequest( method, url, [body], [params] )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/asyncrequest)  | Issue any type of HTTP request asynchronously.                                                                                            |
| [setResponseCallback(expectedStatuses)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/set-response-callback) | Sets a response callback to mark responses as expected.                                                                                   |
| [url\`url\`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/url)                                              | Creates a URL with a name tag. Read more on [URL Grouping](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/http-requests#url-grouping). |
| [expectedStatuses( statusCodes )](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/expected-statuses)           | Create a callback for setResponseCallback that checks status codes.                                                                       |

| Class                                                                                  | Description                                                                              |
| -------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| [CookieJar](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/cookiejar) | Used for storing cookies, set by the server and/or added by the client.                  |
| [FileData](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/filedata)   | Used for wrapping data representing a file when doing multipart requests (file uploads). |
| [Params](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/params)       | Used for setting various HTTP request-specific parameters such as headers, cookies, etc. |
| [Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-http/response)   | Returned by the http.\* methods that generate HTTP requests.                             |


## k6/metrics


The [`k6/metrics` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-metrics) provides functionality to [create custom metrics](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/metrics/create-custom-metrics) of various types.

| Metric type                                                                           | Description                                                                                   |
| ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| [Counter](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-metrics/counter) | A metric that cumulatively sums added values.                                                 |
| [Gauge](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-metrics/gauge)     | A metric that stores the min, max and last values added to it.                                |
| [Rate](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-metrics/rate)       | A metric that tracks the percentage of added values that are non-zero.                        |
| [Trend](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-metrics/trend)     | A metric that calculates statistics on the added values (min, max, average, and percentiles). |


## k6/net/grpc


The [`k6/net/grpc` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc) provides a [gRPC](https://grpc.io/) client for Remote Procedure Calls (RPC) over HTTP/2.

| Class/Method                                                                                                                                 | Description                                                                                                                                                                         |
| -------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Client](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client)                                                         | gRPC client used for making RPC calls to a gRPC Server.                                                                                                                             |
| [Client.load(importPaths, ...protoFiles)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client/client-load)            | Loads and parses the given protocol buffer definitions to be made available for RPC requests.                                                                                       |
| [Client.connect(address [,params])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client/client-connect)               | Connects to a given gRPC service.                                                                                                                                                   |
| [Client.invoke(url, request [,params])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client/client-invoke)            | Makes an unary RPC for the given service/method and returns a [Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/response).                             |
| [Client.asyncInvoke(url, request [,params])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client/client-async-invoke) | Asynchronously makes an unary RPC for the given service/method and returns a Promise with [Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/response). |
| [Client.close()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/client/client-close)                                    | Close the connection to the gRPC service.                                                                                                                                           |
| [Params](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/params)                                                         | RPC Request specific options.                                                                                                                                                       |
| [Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/response)                                                     | Returned by RPC requests.                                                                                                                                                           |
| [Constants](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/constants)                                                   | Define constants to distinguish between [gRPC Response](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/response) statuses.                                     |
| [Stream(client, url, [,params])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream)                                 | Creates a new gRPC stream.                                                                                                                                                          |
| [Stream.on(event, handler)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream/stream-on)                            | Adds a new listener to one of the possible stream events.                                                                                                                           |
| [Stream.write(message)](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream/stream-write)                             | Writes a message to the stream.                                                                                                                                                     |
| [Stream.end()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream/stream-end)                                        | Signals to the server that the client has finished sending.                                                                                                                         |
| [EventHandler](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream/event-handler)                                     | The function to call for various events on the gRPC stream.                                                                                                                         |
| [Metadata](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-net-grpc/stream/message-metadata)                                      | The metadata of a gRPC stream’s message.                                                                                                                                            |


## k6/secrets


The [`k6/secrets` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-secrets) gives access to secrets provided by configured [secret sources](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/secret-source).

| Property                                                                                      | Description                                                                                         |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| [get([String])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-secrets#get)       | asynchrounsly get a secret from the default secret source                                           |
| [source([String])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-secrets#source) | returns a source for the provided name that can than be used to get a secret from a concrete source |


## k6/timers


The [`k6/timers` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-timers) implements timers to work with k6's event loop. They mimic the functionality found in browsers and other JavaScript runtimes.

| Function                                                                      | Description                                          |
| :---------------------------------------------------------------------------- | :--------------------------------------------------- |
| [setTimeout](https://developer.mozilla.org/en-US/docs/Web/API/setTimeout)     | Sets a function to be run after a given timeout.     |
| [clearTimeout](https://developer.mozilla.org/en-US/docs/Web/API/clearTimeout) | Clears a previously set timeout with `setTimeout`.   |
| [setInterval](https://developer.mozilla.org/en-US/docs/Web/API/setInterval)   | Sets a function to be run on a given interval.       |
| [clearInterval](https://developer.mozilla.org/en-US/docs/Web/API/setInterval) | Clears a previously set interval with `setInterval`. |

{{< admonition type="note" >}}

The timer methods are available globally, so you can use them in your script without including an import statement.

{{< /admonition >}}


## k6/ws


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


## crypto


The [`crypto` module](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto) provides a WebCrypto API implementation.

| Class/Method                                                                                      | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [getRandomValues](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/getrandomvalues) | Fills the passed `TypedArray` with cryptographically sound random values.                                                                                                                                          |
| [randomUUID](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/randomuuid)           | Returns a randomly generated, 36 character long v4 UUID.                                                                                                                                                           |
| [subtle](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/subtlecrypto)             | The [SubtleCrypto](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/subtlecrypto) interface provides access to common cryptographic primitives, such as hashing, signing, encryption, or decryption. |

{{< admonition type="note" >}}

The `crypto` object is available globally, so you can use it in your script without including an import statement.

{{< /admonition >}}


## Error codes

The following specific error codes are currently defined:


- 1000: A generic error that isn't any of the ones listed below.
- 1010: A non-TCP network error - this is a place holder there is no error currently known to trigger it.
- 1020: An invalid URL was specified.
- 1050: The HTTP request has timed out.
- 1100: A generic DNS error that isn't any of the ones listed below.
- 1101: No IP for the provided host was found.
- 1110: Blacklisted IP was resolved or a connection to such was tried to be established.
- 1111: Blacklisted hostname using The [Block Hostnames](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/k6-options/reference#block-hostnames) option.
- 1200: A generic TCP error that isn't any of the ones listed below.
- 1201: A "broken pipe" on write - the other side has likely closed the connection.
- 1202: An unknown TCP error - We got an error that we don't recognize but it is from the operating system and has `errno` set on it. The message in `error` includes the operation(write,read) and the errno, the OS, and the original message of the error.
- 1210: General TCP dial error.
- 1211: Dial timeout error - the timeout for the dial was reached.
- 1212: Dial connection refused - the connection was refused by the other party on dial.
- 1213: Dial unknown error.
- 1220: Reset by peer - the connection was reset by the other party, most likely a server.
- 1300: General TLS error
- 1310: Unknown authority - the certificate issuer is unknown.
- 1311: The certificate doesn't match the hostname.
- 1400 to 1499: error codes that correspond to the [HTTP 4xx status codes for client errors](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#4xx_Client_errors)
- 1500 to 1599: error codes that correspond to the [HTTP 5xx status codes for server errors](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#5xx_Server_errors)
- 1600: A generic HTTP/2 error that isn't any of the ones listed below.
- 1610: A general HTTP/2 GoAway error.
- 1611 to 1629: HTTP/2 GoAway errors with the value of the specific [HTTP/2 error code](https://tools.ietf.org/html/rfc7540#section-7) added to 1611.
- 1630: A general HTTP/2 stream error.
- 1631 to 1649: HTTP/2 stream errors with the value of the specific [HTTP/2 error code](https://tools.ietf.org/html/rfc7540#section-7) added to 1631.
- 1650: A general HTTP/2 connection error.
- 1651 to 1669: HTTP/2 connection errors with the value of the specific [HTTP/2 error code](https://tools.ietf.org/html/rfc7540#section-7) added to 1651.
- 1701: Decompression error.


Read more about [Error codes](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/error-codes).

## jslib


[jslib](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib) is a collection of JavaScript libraries maintained by the k6 team that can be used in k6 scripts.

| Library                                                                                                                        | Description                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------- |
| [aws](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/aws)                                                       | Library allowing to interact with Amazon AWS services                                                                  |
| [httpx](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/httpx)                                                   | Wrapper around [k6/http](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/#k6http) to simplify session handling |
| [k6chaijs](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/k6chaijs)                                             | BDD assertion style                                                                                                    |
| [http-instrumentation-pyroscope](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/http-instrumentation-pyroscope) | Library to instrument k6/http to send baggage headers for pyroscope to read back                                       |
| [http-instrumentation-tempo](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/http-instrumentation-tempo)         | Library to instrument k6/http to send tracing data                                                                     |
| [testing](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/testing)                                               | Advanced assertion library with Playwright-inspired API for protocol and browser testing                             |
| [totp](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/totp)                                                     | TOTP (Time-based One-Time Password) generation and verification                                                        |
| [utils](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/jslib/utils)                                                   | Small utility functions useful in every day load testing                                                               |

