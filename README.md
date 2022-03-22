# mule

> Because the best mules don't ask any questions.

**What is it?**

- secure, synchronous file transfer
- a long-running service and a command line client
- can run on a t2.micro instance, probably t2.nano

## Next Steps

- **Deployment** - Create some terraform to run this on an EC2 server
- **Refactor** - merge `mule-send` and `mule-recv` into one client, `mule`
- **Project** - consider changing name to `skymule`? Just in case there are already
  a lot of "mule" projects out there
- **Configuration** - On the client side, take server params (host, port) from:
    - a local dotfile
    - environment variables
- **Configuration** - Make server ports configurable
- **Deployment** - Create a Docker image with this to run
- **Security** - encrypt the connections to the server
    - TLS?
    - SSH?
- **Security** - Add symmetric encryption
    - `mule send this_file secret12345`
    - `mule recv this_file secret12345`
- **Usability** - On producer end of things, have it produce a command
  line invocation for the consumer that can be copied and pasted and
  sent to the receiver that already has all of the necessary parameters.
- **Usability** - Show some kind of progress bar for files that take
  more than five seconds to transfer
- **Security** - Add public key encryption
    - `mule send this_file johnny.pub`
    - `mule recv this_file johnny`
- **Security** - add some kind of authentication on the server with Okta
- **Usability** - add an option on the receiving end to copy the data as text to the clipboard
- **Usability** - specify a file to send
    - on the receiving end, output the data into the appropriately named file
    - this probably requires us to have a more robust data model of transactions
- **Refactor** - Use a more robust data model
    - PipeReader
    - sending IP address
    - receiving IP address
    - key name
    - date and time
    - cryptographic authentication of both sender and receiver
- **Security** - add cryptographic authentication for both sender and receiver
- **Security** - if public keys are on the server, sender has the option to use them to encrypt the data
- **Reliability** - TimeToLive -- destroy a transaction if it gets too old
- **Usability** - Make it so we can specify a user to receive it
    - all the user has to do is say `mule`, and the system will authenticate them, and retrieve anything that is waiting for them
- **Security** - Do sha256 sums everywhere and make sure they all match up
    - can send an sha256 sum as a trailer back on the PUT
    - can send an sha256 sum as a trailer back on the GET
- **Improve** - make keys case-insensitive (convert to lowercase on back end)
- **information** - print out bytes transferred when it's over

## Motivation

Needed a dead simple way to securely transfer a single private key
file to a coworker.

- can't use Slack, email, etc.
    - well, I CAN do it in good conscience, but only if I encrypt it first
    - but even then, it's still there in the Slack or email system forever
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