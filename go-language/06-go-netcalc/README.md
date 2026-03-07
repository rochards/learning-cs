# go-netcalc

A small Go CLI to calculate IPv4 network information from a CIDR input.

Given an address like `192.168.1.10/24`, it prints:
- Network mask
- Network address
- Broadcast address
- Host address range
- Number of usable host IPs

## Requirements

- Go installed

## Run

```bash
go run . -cidr 192.168.1.10/24
```

### Input Format

- Flag: `-cidr`
- Format: `<ipv4>/<prefix>`
- Example: `10.0.0.15/16`

## Example Output

```text
Input:                192.168.1.10/24
Network Mask:         255.255.255.0
Network Address:      192.168.1.0
Broadcast Address:    192.168.1.255
Host Address Range:   192.168.1.1 - 192.168.1.254
Number of Host IPs:  254
```

## Project Structure

- `main.go`: CLI entry point and output
- `netcalc/parser.go`: CIDR and IPv4 parsing utilities
- `netcalc/network.go`: Network calculations (`Mask`, `Address`, `BroadcastAddress`, etc.)

## Note

For production-ready CIDR parsing/validation, prefer Go's standard library parser `net.ParseCIDR(*cidr)` from the `net` package.
