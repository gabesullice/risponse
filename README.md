# Risponse

Risponse is a simple, pseudo-static HTTP server. Its design goal is to be a
quick and easy tool to build RESTful API backends for testing and debugging.

## Usage

`risponse -p [PORT] -d [BASE DIRECTORY]`

`risponse` will look in the current directory for a `config.json` file. This
file is required; an example is provided below.

You may provide response payloads by adding files within the base directory.
`risponse` expects these files to be located and named using the pattern
`{BASE DIRECTORY}/{path}/{method}.json` where:

* `{BASE DIRECTORY}` is either the current directory in which `risponse` was
  executed or a directory provided via the `-d` command line flag.
* `{path}` is the path provided in `config.json` for the requested resource.
* `{method}` is the HTTP method of the request being processed.

## Configuration

Here is an example `config.json` file. Values in the `defaults` object are
applied to all resources unless they are specificied in the resource object
explicitly. At this time, only `cors` and `headers` may be specificied as
defaults.

```json
{
  "defaults": {
    "cors": {
      "allowOrigin": ["http://localhost:8080"],
      "allowCredentials": true
    },
    "headers": {
      "content-type": "application/vnd.api+json"
    }
  },
  "resources": [{
    "path": "/unauthorized",
    "status": 401,
    "headers": {
      "www-authenticate": "Basic",
      "link": "</login>; rel=\"authenticate\""
    },
    "cors": {
      "exposeHeaders": ["Link"]
    }
  }]
}
```

Provided there exists the following file at `/unauthorized/get.json`:

```json
{
  "errors": {
    "status": "401 Unauthorized"
  }
}
```

An HTTP request to `/unauthorized` will return this response:

```http
HTTP/1.1 401 Unauthorized
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: http://localhost:8080
Access-Control-Expose-Headers: Link
Content-Type: application/vnd.api+json
Link: </login>; rel="authenticate"
WWW-Authenticate: Basic

{
  "errors": {
    "status": "401 Unauthorized"
  }
}
```
