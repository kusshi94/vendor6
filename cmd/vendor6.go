package cmd

import (
	"bufio"
	"encoding/csv"
	"io"
	"net"
	"net/netip"
	"os"

	"github.com/kusshi94/vendor6/ouidb"
	"github.com/spf13/cobra"
)

func NewVendor6Command() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:                   "vendor6",
		Short:                 "",
		Long:                  ``,
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// with files
			if len(args) > 0 {
				// open files
				var files []io.Reader
				for _, arg := range args {
					f, err := os.Open(arg)
					if err != nil {
						return err
					}
					defer f.Close()
					files = append(files, f)
				}
				// concat files
				reader := io.MultiReader(files...)
				// run gipp
				return Vendor6(reader, os.Stdout, os.Stderr)
			}

			// without files
			return Vendor6(os.Stdin, os.Stdout, os.Stderr)
		},
	}

	return rootCmd
}

func Vendor6(in io.Reader, out, eout io.Writer) error {
	db, err := ouidb.NewOUIDb("oui.txt")
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(out)

	csvwriter.Write([]string{"IPv6 Address", "Vendor Name"})

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse user input as IPv6 address.
		// If the address is not IPv6, skip it.
		nip, err := netip.ParseAddr(line)
		if err != nil {
			continue
		}

		// Chek if the address is IPv6
		if !nip.Is6() {
			continue
		}

		// Get IID from IPv6 address
		iid := getIID(nip)

		// Check if IID is EUI-64
		if !iid.isEUI64() {
			continue
		}

		// Get MAC address from IID
		mac := getMAC(iid)

		// Get OUI information from MAC address
		oui := db.Lookup(mac)

		// Return vendor name if OUI is found
		if oui != nil {
			csvwriter.Write([]string{line, oui.Company})
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Execute() {
	err := NewVendor6Command().Execute()
	if err != nil {
		os.Exit(1)
	}
}

type IID [8]byte

// getIID returns IID from IPv6 address.
func getIID(nip netip.Addr) IID {
	nip16 := nip.As16()
	return IID(nip16[8:])
}

// isEUI64 returns true if IID is EUI-64.
func (iid IID) isEUI64() bool {
	// Check if IID is EUI-64
	return iid[3] == 0xff && iid[4] == 0xfe
}

// getMAC returns MAC address from EUI-64 IID.
func getMAC(eui64 IID) net.HardwareAddr {
	mac := make(net.HardwareAddr, 6)
	mac[0] = eui64[0] ^ 0x02
	mac[1] = eui64[1]
	mac[2] = eui64[2]
	mac[3] = eui64[5]
	mac[4] = eui64[6]
	mac[5] = eui64[7]
	return mac
}
