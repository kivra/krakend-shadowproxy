# Krakend Shadowproxy Middleware

This package contains custom middleware to add a shadow backend to a KrakenD
proxy pipe. It can be used to mirror proxy requests to a secondary backend,
similar to KrakenD's inbuilt [shadow functionality](https://www.krakend.io/docs/backends/shadow-backends/)
(which does not support proxy requests).

## Installation

To install `shadowproxy` from GitHub:

    go get -u github.com/kivra/krakend-shadowproxy@<commit hash>

## Quick Start

Using the following configuration, an incoming request to KrakenD's `/v1/user/123`
endpoint is proxied to `https://my.backend.com/v2/user/123`, together with a shadow
request to `https://other.backend.com/v3/user/123`. The return value of the shadow
request is ignored.

```json
"endpoints": [
  {
    "endpoint": "/v1/user/{ukey}",
    "method": "POST",
    "output_encoding": "no-op",
    "extra_config": {
      "kivra/shadowproxy": {
        "host": [ "https://other.backend.com" ],
        "url_pattern": "/v3/user/{ukey}",
        "method": "POST"
      }
    },
    "backend": [
      {
        "host": [ "https://my.backend.com" ],
        "url_pattern": "/v2/user/{ukey}",
        "method": "POST",
        "encoding": "no-op"
      }
    ]
  }
]
```

## Configuration Options

The middleware configuration supports the following parameters.

---

### `host`

The host URL(s) of the backend(s) that will receive the shadow request.

---

### `url_pattern`

The URL pattern used for requests to the shadow backend.

---

### `method` (optional)

The request method used for requests to the shadow backend (default `"GET"`).

---

### `disable_host_sanitize` (optional)

Disable host sanitization for prodived `host` list (default `false`).

---

### `timeout` (optional)

The timeout for the request to the shadow backend as a `Go` duration, for
example `10s`. Defaults to the endpoint timeout.
