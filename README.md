# Risponse

Risponse is a simple, pseudo-static HTTP server. It's design goal is to be a
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

```json
{
  "defaults": { // optional
    "cors": { // optional
      "allowOrigin": ["http://localhost:8080"],
      "allowCredentials": true,
      "exposeHeaders": ["Link"]
    },
  }
  "resources": [{
    "path": "/unauthorized",
    "status": 401,
    "headers": { // optional
      "www-authenticate": "ClientBasic",
      "link": "</login>; rel=\"authenticate\"; display=\"authenticate\""
    }
  }]
}
```
