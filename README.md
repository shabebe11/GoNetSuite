# GoNetSuite

A command-line network toolkit written in Go using **almost only the standard library**. It's a learning project for understanding networking at a low level (sockets, concurrency, and eventually raw packet capture) by building the tools from scratch instead of relying on third-party libraries.

The one dependency is [`golang.org/x/net`](https://pkg.go.dev/golang.org/x/net), the Go team's extended networking package, used by `scan` to build, send and parse ICMP packets.

> **AI disclosure:** All of the Go code in this repository is written by me. The documentation (this README) and the code comments are AI-generated.

## Requirements

- Go 1.27+
- macOS (the main development platform; the planned packet sniffer uses macOS BPF devices)

## Running

```sh
go run .
```

This starts an interactive shell. Type a command and press Enter:

```
> Enter a command or type 'exit' to quit:
help
```

While a command is running, press Ctrl+C to stop it and return to the prompt. If a command fails (for example, a missing or invalid argument), the error is printed right away and you're back at the prompt.

## Commands

| Command         | Status      | Description                                           |
|-----------------|-------------|-------------------------------------------------------|
| `listen <port>` | Working     | Listen for TCP connections and print what clients send |
| `scan`          | Working     | Find devices on your local network with a ping sweep  |
| `sniff`         | Placeholder | Capture and decode packets on a network interface     |
| `help`          | Working     | List the available commands                           |
| `exit`          | Working     | Quit the shell (Ctrl-D also works)                    |

### `listen`

In one terminal:

```sh
go run .
listen 9000
```

In a second terminal, connect and type a message. It appears in the first terminal:

```sh
nc localhost 9000
```

You can check that the port is open with `lsof -iTCP:9000 -sTCP:LISTEN -n -P`. Opening `http://localhost:9000` in a browser shows the raw HTTP request the browser sends.

### `scan`

```sh
go run .
scan
```

```
> Scanning 192.168.1.23/24 (254 hosts)...
  192.168.1.1      2.1ms
  192.168.1.4      38.7ms
  192.168.1.10     5.3ms
> 3 hosts up
```

How it works:

1. **Find the local subnet.** Walks the machine's network interfaces and takes the first IPv4 address on one that is up and not loopback. The address and its subnet mask (e.g. `192.168.1.23/24`) give the range of addresses on the network.
2. **List every host address.** Treats the network address as a 32-bit number and counts up through the range, skipping the network address, the broadcast address and this machine's own IP.
3. **Ping sweep.** Opens a single ICMP socket, starts one goroutine that reads replies, then sends an ICMP Echo Request to every address. Each Echo Reply is recorded with its round-trip time. The sweep ends after 3 seconds, or straight away on Ctrl+C (any replies so far are still printed).

It uses an unprivileged ICMP socket, so it doesn't need `sudo` on macOS.

Limitations:

- Only devices that answer ping show up. Phones, Windows machines and anything with a strict firewall often ignore ICMP echo requests, so they'll be missing. A planned ARP sweep will catch these.
- If the machine is on several networks (Wi-Fi plus Ethernet, a VPN, a VM bridge), it scans whichever one it finds first.
- Subnets larger than a `/22` (1,022 hosts) are refused rather than swept.

## Project layout

```
main.go              entry point; starts the shell
internal/cli/        the interactive shell (reads lines, dispatches commands)
internal/command/    one handler per command; validates arguments
internal/netcat/     TCP listen (and later connect) logic
internal/scan/       local subnet discovery and the ICMP ping sweep
```

`internal/command` handles user input, and the packages it calls (`netcat`, `scan`, and `sniffer`) do the networking without knowing about the shell.

## Roadmap

1. **CLI shell:** interactive prompt with Ctrl+C returning to the prompt instead of exiting
2. **Scanner:** host discovery with an ICMP ping sweep (done), then TCP port scanning of a host with a concurrent worker pool, port ranges, banner grabbing and JSON output
3. **Netcat:** TCP/UDP listen and connect, raw hex payloads, a raw-syscall version of `listen`
4. **Packet sniffer:** macOS BPF capture, hand-written Ethernet/IPv4/TCP/UDP/ARP/ICMP parsers, `.pcap` output
5. **arpwatch:** detect ARP spoofing by tracking IP-to-MAC changes