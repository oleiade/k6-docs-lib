package k6docslib

import (
	"errors"
	"time"
)

const (
	// MaxFileSize is the maximum allowed size for a single file during extraction.
	// This prevents decompression bombs (gosec G110).
	MaxFileSize = 50 << 20 // 50 MB

	// MaxBundleSize caps the compressed bundle download.
	// Doc bundles are typically <2 MB; this prevents memory exhaustion
	// from a malicious or corrupted asset before extraction starts.
	MaxBundleSize = 100 << 20 // 100 MB

	// StalenessCheckInterval is how often we re-check the remote ETag
	// to see if a newer doc bundle is available.
	StalenessCheckInterval = 24 * time.Hour

	// DefaultRefreshTimeout caps how long a staleness HEAD or refresh GET can take.
	// Keeps the offline-first experience fast when the network is slow.
	// Override via K6_DOCS_REFRESH_TIMEOUT env var (parsed as time.Duration).
	DefaultRefreshTimeout = 10 * time.Second
)

// homeDirFromEnv returns the user's home directory from environment variables.
// It checks HOME first, then USERPROFILE as a fallback (for Windows).
func homeDirFromEnv(env map[string]string) (string, error) {
	if home := env["HOME"]; home != "" {
		return home, nil
	}
	if home := env["USERPROFILE"]; home != "" {
		return home, nil
	}
	return "", errors.New("neither HOME nor USERPROFILE is set")
}

func resolveRefreshTimeout(env map[string]string) time.Duration {
	if s := env["K6_DOCS_REFRESH_TIMEOUT"]; s != "" {
		if d, err := time.ParseDuration(s); err == nil && d > 0 {
			return d
		}
	}
	return DefaultRefreshTimeout
}
