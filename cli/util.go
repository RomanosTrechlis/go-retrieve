package cli

import "os"

// nonZeroExit is used for testing purposes, since
// os.Exit cannot be tested without exiting the test
var nonZeroExit func(code int) = os.Exit
