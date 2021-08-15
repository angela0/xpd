# xpd 

A simple xpath tool using [golang xpath library](https://github.com/antchfx/xpath).

# Install

``` shell
go get -u https://github.com/angela0/xpd
```

# Usage

``` shell
xpd <selector> [filename]
```

For example:

``` shell
xpd '//title' index.html
```

Arg `filename` is optional, and `xpd` will read data from stdin if not provided.

For example:

``` shell
less index.html | xpd '//title'
# or
xpd '//title' < index.html
```

# Licence

[MIT](LICENSE)

Some codes are from [htmlquery](github.com/antchfx/htmlquery) protected by MIT.
