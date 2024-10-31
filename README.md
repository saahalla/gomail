# Gomail
[![Build Status](https://travis-ci.org/go-mail/mail.svg?branch=master)](https://travis-ci.org/go-mail/mail) [![Code Coverage](http://gocover.io/_badge/github.com/go-mail/mail)](http://gocover.io/github.com/go-mail/mail) [![Documentation](https://godoc.org/github.com/go-mail/mail?status.svg)](https://godoc.org/github.com/go-mail/mail)

This is an actively maintained fork of [Gomail][1] and includes fixes and
improvements for a number of outstanding issues. The current progress is
as follows:

 - [x] Add Authentication for NTLM.

[1]: https://github.com/go-mail/mail

## Documentation

https://godoc.org/github.com/go-mail/mail


## Examples

See the [examples in the documentation](https://godoc.org/github.com/go-mail/mail#example-package).

### Transitioning Existing Codebases

If you're already using the original Gomail or mail.v2, you can add this on go.mod

```
replace gopkg.in/mail.v2 => github.com/saahalla/gomail v0.0.0-20241031070749-1f30f01cc509
```

If smtp server use NTLM Authentication, you can Call function UseAuthNTLM()
```
package main

import (
	"gopkg.in/mail.v2"
)

func main() {
	d := mail.NewDialer("smtp.example.com", 587, "user", "123456")
	
	// if smtp server use NTML Authentication
	d.UseAuthNTLM()
	
	// Send emails using d.
}
```

## Contribute

Contributions are more than welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for
more info.


## Change log

See [CHANGELOG.md](CHANGELOG.md).


## License

[MIT](LICENSE)


## Support & Contact

You can ask questions on the [Gomail
thread](https://groups.google.com/d/topic/golang-nuts/jMxZHzvvEVg/discussion)
in the Go mailing-list.
