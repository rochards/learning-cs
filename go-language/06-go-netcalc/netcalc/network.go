package netcalc

type Network struct {
	IPv4   uint32
	Prefix int
}

func (subnet Network) Mask() uint32 {
	// Build the network mask from the CIDR prefix length.
	//
	// Start with all 32 bits set to 1:
	//   0xffffffff -> 11111111.11111111.11111111.11111111
	//
	// Then shift left by (32 - Prefix) to keep Prefix bits as 1 (network bits)
	// and move the remaining bits to 0 (host bits).
	//
	// Example for Prefix = 24:
	//   shift amount: 32 - 24 = 8
	//   11111111.11111111.11111111.11111111 << 8
	//   = 11111111.11111111.11111111.00000000
	//   = 255.255.255.0
	//   = 4294967040 (uint32)
	return uint32(0xffffffff << (32 - subnet.Prefix))
}

func (subnet Network) Address() uint32 {
	// Compute the network address using bitwise AND between:
	// 1) The IP address
	// 2) The subnet mask
	//
	// Because mask network bits are 1, those IP bits are preserved.
	// Because mask host bits are 0, host bits are forced to 0.
	//
	// Example:
	//   IP   192.168.1.10  = 11000000.10101000.00000001.00001010
	//   Mask /24           = 11111111.11111111.11111111.00000000
	//   AND result         = 11000000.10101000.00000001.00000000
	//   Network address    = 192.168.1.0
	//                      = 3232235776 (uint32)
	return uint32(subnet.IPv4 & subnet.Mask())
}

func (subnet Network) BroadcastAddress() uint32 {
	// Compute the broadcast address as:
	//   network address OR inverted mask
	//
	// Step 1: subnet.Mask() has 1s in network bits and 0s in host bits.
	// Step 2: ^subnet.Mask() flips that:
	//   - network bits become 0
	//   - host bits become 1
	// Step 3: OR with network address keeps network bits unchanged
	//         and forces all host bits to 1.
	//
	// Example for 192.168.1.10/24:
	//   Network Address    = 11000000.10101000.00000001.00000000
	//   Mask/24            = 11111111.11111111.11111111.00000000
	//   ^Mask              = 00000000.00000000.00000000.11111111
	//   OR result          = 11000000.10101000.00000001.11111111
	//   Broadcast address  = 192.168.1.255
	//                      = 3232236031 (uint32)
	return uint32(subnet.Address() | ^subnet.Mask())
}

func (subnet Network) FirstHostAddress() uint32 {
	return subnet.Address() + 1
}

func (subnet Network) LastHostAddress() uint32 {
	return subnet.BroadcastAddress() - 1
}
