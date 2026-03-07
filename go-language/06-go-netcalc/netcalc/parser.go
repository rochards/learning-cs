package netcalc

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatIPv4 converts a uint32 IPv4 value into dotted-decimal notation (a.b.c.d).
func FormatIPv4(value uint32) string {
	// Walkthrough example:
	//   IP: 192.168.1.10
	//   uint32: 3232235786
	//   binary: 11000000.10101000.00000001.00001010
	//
	// Each octet is extracted with two operations:
	// 1) Right shift (>>) to move the target octet into the lowest 8 bits.
	// 2) Mask with 0xff (11111111) to keep only those 8 bits.
	//
	// First octet example:
	//   value >> 24  => 00000000.00000000.00000000.11000000
	//   & 0xff       => 11000000 (192)
	//
	// The same logic is applied to the remaining octets with shifts 16, 8, and 0.

	octet1 := (value >> 24) & 0xff
	octet2 := (value >> 16) & 0xff
	octet3 := (value >> 8) & 0xff
	octet4 := value & 0xff

	return fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)
}

// parseOctet converts a decimal string into an IPv4 octet value.
// It returns an error if the value is not a valid integer in the range 0-255.
func parseOctet(value string) (int, error) {
	octet, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid octet %q: %v", octet, err)
	}
	if octet < 0 || octet > 255 {
		return 0, fmt.Errorf("octet out of range: %d", octet)
	}

	return octet, nil
}

// parseIPv4 parses an IPv4 string in dotted-decimal format (a.b.c.d)
// and returns its 32-bit numeric representation.
// It returns an error if the input does not have exactly 4 octets
// or if any octet is not a valid number in the range 0-255.
func parseIPv4(ip string) (uint32, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid IPv4: %s", ip)
	}

	octet1, err := parseOctet(parts[0])
	if err != nil {
		return 0, err
	}

	octet2, err := parseOctet(parts[1])
	if err != nil {
		return 0, err
	}

	octet3, err := parseOctet(parts[2])
	if err != nil {
		return 0, err
	}

	octet4, err := parseOctet(parts[3])
	if err != nil {
		return 0, err
	}

	// Build one uint32 from 4 octets (network order: a.b.c.d).
	//
	// Example for 192.168.1.10:
	//   octet1=192, octet2=168, octet3=1, octet4=10
	//
	// Shift each octet to its 8-bit slot in the 32-bit value:
	//   octet1 << 24 -> 11000000.00000000.00000000.00000000
	//   octet2 << 16 -> 00000000.10101000.00000000.00000000
	//   octet3 << 8  -> 00000000.00000000.00000001.00000000
	//   octet4       -> 00000000.00000000.00000000.00001010
	//
	// Use bitwise OR (|) to merge non-overlapping parts into the final IPv4 number:
	//   11000000.10101000.00000001.00001010
	//   = 3232235786 (uint32)
	return uint32(octet1<<24 | octet2<<16 | octet3<<8 | octet4), nil
}

// parsePrefixLength parses a CIDR prefix length string (for example, "24").
// It returns an error if the value is not an integer in the valid IPv4
// prefix range of 0 to 32.
func parsePrefixLength(prefix string) (int, error) {
	prefixLen, err := strconv.Atoi(prefix)
	if err != nil {
		return 0, fmt.Errorf("invalid prefix %q: %v", prefix, err)
	}
	if prefixLen < 0 || prefixLen > 32 {
		return 0, fmt.Errorf("prefix out of range: %d", prefixLen)
	}

	return prefixLen, nil
}

// ParseCIDR parses a CIDR string in the form "<ipv4>/<prefix>" (for example,
// "192.168.1.10/24") and returns a Subnet.
// It returns an error if the CIDR format is invalid, the IPv4 address is invalid,
// or the prefix length is outside the range 0-32.
func ParseCIDR(cidr string) (Network, error) {
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return Network{}, fmt.Errorf("invalid CIDR address: %s", cidr)
	}
	ipPart := parts[0]
	prefixPart := parts[1]

	ip, err := parseIPv4(ipPart)
	if err != nil {
		return Network{}, err
	}

	prefix, err := parsePrefixLength(prefixPart)
	if err != nil {
		return Network{}, err
	}

	return Network{ip, prefix}, nil
}
