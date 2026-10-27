# GoNetSuite

A command-line network toolkit written in Go using **only the standard library**. It's a learning project for understanding networking at a low level (sockets, concurrency, and eventually raw packet capture) by building the tools from scratch instead of relying on third-party libraries.

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
Enter a command or type 'exit' to quit:
help
```

## Commands

| Command         | Status      | Description                                           |
|-----------------|-------------|-------------------------------------------------------|
| `listen <port>` | Working     | Listen for TCP connections and print what clients send |
| `scan`          | Placeholder | Scan a host for open TCP ports                        |
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

## Project layout

```
main.go              entry point; starts the shell
internal/cli/        the interactive shell (reads lines, dispatches commands)
internal/command/    one handler per command; validates arguments
internal/netcat/     TCP listen (and later connect) logic
```

`internal/command` handles user input, and the packages it calls (`netcat`, and later `scanner` and `sniffer`) do the networking without knowing about the shell.

## Roadmap

1. **CLI shell:** interactive prompt with Ctrl+C returning to the prompt instead of exiting
2. **Port scanner:** concurrent worker pool, port ranges, banner grabbing, JSON output
3. **Netcat:** TCP/UDP listen and connect, raw hex payloads, a raw-syscall version of `listen`
4. **Packet sniffer:** macOS BPF capture, hand-written Ethernet/IPv4/TCP/UDP/ARP/ICMP parsers, `.pcap` output
5. **arpwatch:** detect ARP spoofing by tracking IP-to-MAC changes