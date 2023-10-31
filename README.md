# Background

I am planning to integrate clash to my own Tailscale distro(a.k.a xEdge), especially on Android/iOS devices,
since the platform only allow single VPN instance (both Android/iOS is based on Unix tuntap mechanism).

Since clash OSS code does not provide the tunnel support, this is actually not the issue with Tailscale,
since it will create the tunnel, only the glue code needed to be written.

However, to understand how clash works more, and also trying to do some POC, also enhance my learning of
the GIOUI project which is the GUI toolkit used by xEdge both in Android/iOS version (for Tailscale, do not
the detail of iOS GUI toolkit, since it make every proprietary app close-source (I can understand that since
my xEdge iOS code is close-source too.)

This project will be Open Source, the LICENSE file will be added later.

# Plan

I like break the actually development into smally pieces as I do in most of my project.

So there should be a plan for this project, though I can't guarantee when it will be finished.

- [x] init the projet
- [x] create the tun interface using tailscale libraries.
- [ ] add tunnel support using some go tun library which is cross-platform.
- [ ] add config generator for clash basically add v2ray config support.
- [ ] test, test, test, make it at least can be used on MacOS (using UNIX api).
  1. use a oversea dns server, make the server IP address routed to the tun, and use
     gVisor to extract the UDP traffic and use outbound to do that and handle reply too.
- [ ] use this project experience to port clash to my app.

# Others

~~Nothing to be noted now.~~

I do not know if there will be code that use CGO... :(

After add CGO_ENABLED=0 to Makefile it make me realized gio actually use CGO in some gl libraries...


## Need sudo

obviously if we want to create tun, we need sudo.
