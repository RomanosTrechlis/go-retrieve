package cli

import "os"

// NonZeroExit is used for testing purposes, since
// os.Exit cannot be tested without exiting the test
var NonZeroExit func(code int) = os.Exit
