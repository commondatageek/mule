# mule

> Because the best mules don't ask any questions.

**What is it?**

- secure, synchronous file transfer
- a long-running service and a command line client
- can run on a t2.micro instance, probably t2.nano

## Next Steps

- **Refactor** - merge `mule-send` and `mule-recv` into one client, `mule`
- **Usability** - On producer end of things, have it produce a command
  line invocation for the consumer that can be copied and pasted and
  sent to the receiver that already has all of the necessary parameters.
- **Usability** - Read data from STDIN
- **Usability** - Write data to STDOUT
- **Usability** - Show some kind of progress bar for files that take
  more than five seconds to transfer
- **Configuration** - On the client side, take server params (host, port) from:
    - a local dotfile
    - environment variables
- **Security** - for both producer and consumer, print out a SHA256 hash after it
  has been transferred
- **Configuration** - Make server ports configurable
- **Security** - Add encryption
    - symmetric encryption?
        - `mule send this_file secret12345`
        - `mule recv this_file secret12345`
    - public/private key encryption?
        - `mule send this_file johnny.pub`
        - `mule recv this_file johnny`
- **Project** - consider changing name to `skymule`? Just in case there are already
  a lot of "mule" projects out there
- **Security** - add some kind of authentication on the server with Okta
- **Security** - encrypt the connections to the server
    - TLS?
    - SSH?

## Motivation

Needed a dead simple way to securely transfer a single private key
file to a coworker.

- can't use Slack, email, etc.
- can't use thumb drives, etc.
- privbin was on a different VPN that we didn't have immediate access to
- don't want to store secrets on some server anyway

They know it's coming.  They need it now.  I just need a simple way to send it now.

**Benefits**

- It never stores your secret on the server; it's just a middle man for transport,
  and never more than one buffer's worth of data are in its memory at a time.
- The encrypted data and the encryption secret are never in the same channel together.
- You and the receiver need to coordinate synchronously — it should be pretty obvious if it worked or not.
- `mule` will only let one connection publish and one connection consume.  If it got consumed, but your partner didn't consume it, then who did? Better change that password!
- Doesn't transfer ANY data until both publisher and receiver are ready to go.
- Can do cryptographic hash validation on either end.