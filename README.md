# mule

> Because the best mules don't ask any questions.

**What is it?**

- secure, synchronous file transfer
- a long-running service and a command line client
- can run on a t2.micro instance, probably t2.nano

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