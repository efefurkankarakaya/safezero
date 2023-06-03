package main

import (
	"fmt"
	"math/rand"
	"os"
	"safezero/cmd/safewipe"
	"strings"
	"time"
)

func main() {
	// TODO: Should I consider cryptographic seed instead of time seed?
	rand.Seed(time.Now().UnixMicro())

	root := ""

	criticalPaths := []string{
		"",
		"/", "/bin", "/boot", "/cdrom", "/dev", "/etc", "/home", "/lib", "/lost+found", "/media", "/mnt", "/opt", "/proc", "/root", "/run", "/sbin", "/selinux", "/srv", "/sys", "/tmp", "/usr", "/var", "/dev",
		"/Applications", "/Library", "System", "/Users",
		"A:", "B:", "C:", "E:", "D:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:",
		// TODO: Does Windows still allow to use another partition rather than C as a system partition?
		// TODO: Add more Windows-specific critical paths.
	}

	for _, criticalPath := range criticalPaths {
		if strings.EqualFold(criticalPath, root) {
			fmt.Println(criticalPath)
			os.Exit(1) // TODO: Use instead ReflectError but first rewrite this block as a function.
		}
	}

	// DeleteFast(root)
	// safedelete.DeleteSecure(root)
	safewipe.WipeSafe(root)

	os.Exit(0)
}
