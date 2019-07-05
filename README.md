# JsonSwag

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

simply generate swag doc json by specific plain text, easy to build a swagger document

## How to use

[![asciicast](https://asciinema.org/a/lrb9iXEtYy4UfgHqVzIAhfHHR.svg)](https://asciinema.org/a/lrb9iXEtYy4UfgHqVzIAhfHHR)

```bash
> ./jsonswag 
  -f string
        *[required] file path
  -o string
        output file path (default "./swagger.json")
  -p    pretty print
```

## Plain Text Standard

- `*` required
- nested json supported
- path/query parameters follow `name:type` format, separate by `,` 

> [Data Types in swagger](https://swagger.io/specification/#dataTypes)

```md
# [*title] - [*description] - [*version]
## [*method] | [*uri] | [*description]
### [path parameters] | [query parameters] | [json]
### [json response]
```

check example on [plain_sample](plain_sample)

## Know the Limitation

- only supported json request & response
- enum / require / etc... limitation are not supported in request yet

## LICENSE

[![](http://www.wtfpl.net/wp-content/uploads/2012/12/wtfpl-badge-4.png)](http://www.wtfpl.net/)