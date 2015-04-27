chromedriver_helper in Go
=========================

## Description

This is a clone of [chromedriver-helper](https://rubygems.org/gems/chromedriver-helper) written in Go.

chromedriver_helper provides some commands to install and manage a [chromedriver](https://sites.google.com/a/chromium.org/chromedriver/) executable, and some functions to access it.


## As Command Line Tool

### Installation

    $ go get github.com/a2ikm/chromedriver_helper

### Usage

To install latest version, run `install` command:

    $ chromedriver_helper install

This will install `chromedriver` to `~/.chromedriver-helper`.

Note that chromedriver_helper doesn't use the platform specific directory like `~/.chromedriver-helper/linux64`, and place chromedriver just below `~/.chromedriver-helper` like `~/.chromedriver-helper/chromedriver`. This is different from original [chromedriver-helper](https://rubygems.org/gems/chromedriver-helper).

If you want to tell what version of chromedriver is installed, run `installed` command:

    $ chromedriver_helper installed


## As Package

### Import

```go
import (
  github.com/a2ikm/chromedriver_helper/chromedriver_helper
)
```

### Functions

See [chromedriver_helper/chromedriver_helper.go](chromedriver_helper/chromedriver_helper.go).


## Development

[Rakefile](Rakefile) defines some tasks like test:

    $ rake test


## Contributing

1. Fork it ( https://github.com/a2ikm/chromedriver_helper/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

