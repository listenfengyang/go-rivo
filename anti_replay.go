package rivo

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func normalizeAntiReplay(version string, timestamp *int64, nonce *string) error {
	ver := strings.TrimSpace(version)
	if ver == "" {
		ver = Version10
	}

	hasTS := timestamp != nil && *timestamp > 0
	hasNonce := nonce != nil && strings.TrimSpace(*nonce) != ""

	switch ver {
	case Version11:
		if !hasTS && timestamp != nil {
			*timestamp = time.Now().UnixMilli()
			hasTS = true
		}
		if !hasNonce && nonce != nil {
			*nonce = uuid.NewString()
			hasNonce = true
		}
		if !hasTS || !hasNonce {
			return fmt.Errorf("version=1.1 requires both timestamp and nonce")
		}
	default:
		// Rivo replay rule: timestamp and nonce must be both present or both absent.
		if hasTS != hasNonce {
			return fmt.Errorf("timestamp and nonce must be both provided or both omitted")
		}
	}

	return nil
}
