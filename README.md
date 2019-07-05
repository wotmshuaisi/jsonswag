# JsonSwag

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

simply generate swag doc json by specific plain text, easy to build a swagger document

## Installation

Download binary file from [Release Page](https://github.com/wotmshuaisi/jsonswag/releases)

### For linux

```bash
> curl https://github.com/wotmshuaisi/jsonswag/releases/download/latest/jsonswag-linux-amd64 -o jsonswag
> chmod +x jsonswag
> mv jsonswag ~/bin/
> jsonswag -h
  -f string
        *[required] file path
  -o string
        output file path (default "./swagger.json")
  -p    pretty print
```

### For windows

serve yourself, i don't know how

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