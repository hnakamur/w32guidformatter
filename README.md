w32guidformatter
================

A command line tool to format a GUID for using it with https://github.com/mattn/go-ole

## Build

```
go build
```

## Usage example

```
$ ./w32guidformatter 30cbe57d-d9d0-452a-ab13-7ac5ac4825ee
&ole.GUID{0x30cbe57d, 0xd9d0, 0x452a, [8]byte{0xab, 0x13, 0x7a, 0xc5, 0xac, 0x48, 0x25, 0xee}}
```

## License

MIT
