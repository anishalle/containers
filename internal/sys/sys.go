package sys

import (
	"os"
	"strings"
	"syscall"

	"github.com/Code-Hex/vz/v3"
	"github.com/anishalle/containers/internal/log"
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

var l = log.GetInstance()

type Image struct {
	vmlinuz                    string
	initrd                     string
	disk                       string
	diskBlock                  string // we need a disk block for the installation disc, im making this required
	cpuLimit                   uint64
	memoryLimit                uint64 // in bytes e.g. 2 * 1024 * 1024 * 1024 for 2GB
	setRawMode                 bool   // if true, this uses setRawMode to pipe stdin raw, which disables local echo, input canonicalization etc.
	kernelCommandLineArguments []string
}

func setRawMode(f *os.File) {
	var attr unix.Termios

	// Get settings for terminal
	termios.Tcgetattr(f.Fd(), &attr)

	// Put stdin into raw mode, disabling local echo, input canonicalization,
	// and CR-NL mapping.
	attr.Iflag &^= syscall.ICRNL
	attr.Lflag &^= syscall.ICANON | syscall.ECHO

	// Set minimum characters when reading = 1 char
	attr.Cc[syscall.VMIN] = 1

	// set timeout when reading as non-canonical mode
	attr.Cc[syscall.VTIME] = 0

	// reflects the changed settings
	termios.Tcsetattr(f.Fd(), termios.TCSANOW, &attr)
}

func (i *Image) Run() error {
	return nil
}

func (i *Image) config() *vz.VirtualMachineConfiguration {

	bootLoader, err := vz.NewLinuxBootLoader(
		i.vmlinuz,
		vz.WithCommandLine(strings.Join(i.kernelCommandLineArguments, " ")),
		vz.WithInitrd(i.initrd),
	)
	if err != nil {
		l.Error("bootloader creation failed: %s", err)
	}

}
