---
aliases:
  - ./k6-experimental/browser # docs/k6/<K6_VERSION>/javascript-api/k6-experimental/browser
title: 'k6/browser'
description: 'An overview of the browser-level APIs from browser module.'
weight: 02
---

# browser

The browser module APIs are inspired by Playwright and other frontend testing frameworks.

You can find examples of using [the browser module API](#browser-module-api) in the [getting started guide](https://grafana.com/docs/k6/<K6_VERSION>/using-k6-browser).

{{< admonition type="note" >}}

To work with the browser module, make sure you are using the latest [k6 version](https://github.com/grafana/k6/releases).

{{< /admonition >}}

## Properties

The table below lists the properties you can import from the browser module (`'k6/browser'`).

| Property | Description                                                                                                                                                                          |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| browser  | The browser module API is the entry point for all your tests. See the [example](#example) and the [API](#browser-module-api) below.                                                  |
| devices  | Returns predefined emulation settings for many end-user devices that can be used to simulate browser behavior on a mobile device. See the [devices example](#devices-example) below. |

## Browser Module API

The browser module is the entry point for all your tests, and it is what interacts with the actual web browser via [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) (CDP). It manages:

- [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext) which is where you can set a variety of attributes to control the behavior of pages;
- and [Page](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page) which is where your rendered site is displayed.


| Method                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [browser.closeContext()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/closecontext)                                   | Closes the current [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                          |
| [browser.context()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/context)                                             | Returns the current [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                         |
| [browser.isConnected](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/isconnected) {{< docs/bwipt id="453" >}}           | Indicates whether the [CDP](https://chromedevtools.github.io/devtools-protocol/) connection to the browser process is active or not.                                                                                                                                                                                                                                                                                                                             |
| [browser.newContext([options])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/newcontext/) {{< docs/bwipt id="455" >}} | Creates and returns a new [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext).                                                                                                                                                                                                                                                                                                                                   |
| [browser.newPage([options])](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/newpage) {{< docs/bwipt id="455" >}}        | Creates a new [Page](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page) in a new [BrowserContext](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/browsercontext) and returns the page. Pages that have been opened ought to be closed using [`Page.close`](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/page/close). Pages left open could potentially distort the results of Web Vital metrics. |
| [browser.version()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/version)                                             | Returns the browser application's version.                                                                                                                                                                                                                                                                                                                                                                                                                       |


### Example

{{< code >}}

```javascript
import { browser } from 'k6/browser';

export const options = {
  scenarios: {
    browser: {
      executor: 'shared-iterations',
      options: {
        browser: {
          type: 'chromium',
        },
      },
    },
  },
  thresholds: {
    checks: ['rate==1.0'],
  },
};

export default async function () {
  const page = await browser.newPage();

  try {
    await page.goto('https://test.k6.io/');
  } finally {
    await page.close();
  }
}
```

{{< /code >}}

Then, you can run the test with this command. Also, see the [browser module options](https://grafana.com/docs/k6/<K6_VERSION>/using-k6-browser/options) for customizing the browser module's behavior using environment variables.

{{< code >}}

```bash
k6 run script.js
```

```docker
# WARNING!
# The grafana/k6:master-with-browser image launches a Chrome browser by setting the
# 'no-sandbox' argument. Only use it with trustworthy websites.
#
# As an alternative, you can use a Docker SECCOMP profile instead, and overwrite the
# Chrome arguments to not use 'no-sandbox' such as:
# docker container run --rm -i -e K6_BROWSER_ARGS='' --security-opt seccomp=$(pwd)/chrome.json grafana/k6:master-with-browser run - <script.js
#
# You can find an example of a hardened SECCOMP profile in:
# https://raw.githubusercontent.com/jfrazelle/dotfiles/master/etc/docker/seccomp/chrome.json.
docker run --rm -i grafana/k6:master-with-browser run - <script.js
```

```windows
k6 run script.js
```

```windows-powershell
k6 run script.js
```

{{< /code >}}

### Devices example

To emulate the browser behaviour on a mobile device and approximately measure the browser performance, you can import `devices` from `k6/browser`.

{{< code >}}

```javascript
import { browser, devices } from 'k6/browser';

export const options = {
  scenarios: {
    browser: {
      executor: 'shared-iterations',
      options: {
        browser: {
          type: 'chromium',
        },
      },
    },
  },
  thresholds: {
    checks: ['rate==1.0'],
  },
};

export default async function () {
  const iphoneX = devices['iPhone X'];
  const context = await browser.newContext(iphoneX);
  const page = await context.newPage();

  try {
    await page.goto('https://test.k6.io/');
  } finally {
    page.close();
  }
}
```

{{< /code >}}

## Browser-level APIs


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


## Browser module options

You can customize the behavior of the browser module by providing browser options as environment variables.


| Environment Variable           | Description                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| K6_BROWSER_ARGS                | Extra command line arguments to include when launching browser process. See [this link](https://peter.sh/experiments/chromium-command-line-switches/) for a list of Chromium arguments. Note that arguments should not start with `--` (see the command example below).                                                                                                  |
| K6_BROWSER_DEBUG               | All CDP messages and internal fine grained logs will be logged if set to `true`.                                                                                                                                                                                                                                                                                         |
| K6_BROWSER_EXECUTABLE_PATH     | Override search for browser executable in favor of specified absolute path.                                                                                                                                                                                                                                                                                              |
| K6_BROWSER_HEADLESS            | Show browser GUI or not. `true` by default.                                                                                                                                                                                                                                                                                                                              |
| K6_BROWSER_IGNORE_DEFAULT_ARGS | Ignore any of the [default arguments](https://grafana.com/docs/k6/<K6_VERSION>/using-k6-browser/options/#default-arguments) included when launching a browser process.                                                                                                                                                                                                   |
| K6_BROWSER_TIMEOUT             | Default timeout for initializing the connection to the browser instance. `'30s'` if not set.                                                                                                                                                                                                                                                                             |
| K6_BROWSER_TRACES_METADATA     | Sets additional _key-value_ metadata that is included as attributes in every span generated from browser module traces. Example: `K6_BROWSER_TRACES_METADATA=attr1=val1,attr2=val2`. This only applies if traces generation is enabled, refer to [Traces output](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/k6-options/reference#traces-output) for more details. |

The following command passes the browser options as environment variables to launch a headful browser with custom arguments.

{{< code >}}

```bash
K6_BROWSER_HEADLESS=false K6_BROWSER_ARGS='show-property-changed-rects' k6 run script.js
```

```docker
# WARNING!
# The grafana/k6:master-with-browser image launches a Chrome browser by setting the
# 'no-sandbox' argument. Only use it with trustworthy websites.
#
# As an alternative, you can use a Docker SECCOMP profile instead, and overwrite the
# Chrome arguments to not use 'no-sandbox' such as:
# docker container run --rm -i -e K6_BROWSER_ARGS='' --security-opt seccomp=$(pwd)/chrome.json grafana/k6:master-with-browser run - <script.js
#
# You can find an example of a hardened SECCOMP profile in:
# https://raw.githubusercontent.com/jfrazelle/dotfiles/master/etc/docker/seccomp/chrome.json.
docker run --rm -i -e K6_BROWSER_HEADLESS=false -e K6_BROWSER_ARGS='show-property-changed-rects' grafana/k6:master-with-browser run - <script.js
```

```windows
set "K6_BROWSER_HEADLESS=false" && set "K6_BROWSER_ARGS='show-property-changed-rects' " && k6 run script.js
```

```windows-powershell
$env:K6_BROWSER_HEADLESS="false" ; $env:K6_BROWSER_ARGS='show-property-changed-rects' ; k6 run script.js
```

{{< /code >}}

