//go:build darwin

package touchbar

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/internal/darwin"
)

// New allows to create a new Touch Bar for this application
// Note: only one Touch Bar can be active at a given time.
var New = darwin.NewTouchBar
