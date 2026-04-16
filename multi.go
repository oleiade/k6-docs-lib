package k6docslib

import (
	"fmt"
	"strings"
)

// MultiIndex holds sections across multiple versions with fast lookups.
// It is not safe for concurrent use. Callers must complete all Add/SetLatest
// calls before using Lookup/GetAll/etc. from multiple goroutines.
type MultiIndex struct {
	versions []string
	latest   string
	indices  map[string]*Index
	bySlug   map[string]map[string]*Section // version -> slug -> *Section
}

// NewMultiIndex creates an empty MultiIndex.
func NewMultiIndex() *MultiIndex {
	return &MultiIndex{
		indices: make(map[string]*Index),
		bySlug:  make(map[string]map[string]*Section),
	}
}

// Add registers a version's Index. The first version added becomes Latest
// unless SetLatest is called explicitly.
func (mi *MultiIndex) Add(version string, idx *Index) {
	if _, exists := mi.indices[version]; !exists {
		mi.versions = append(mi.versions, version)
	}
	mi.indices[version] = idx

	slugMap := make(map[string]*Section, len(idx.Sections))
	for i := range idx.Sections {
		slugMap[strings.ToLower(idx.Sections[i].Slug)] = &idx.Sections[i]
	}
	mi.bySlug[version] = slugMap

	if mi.latest == "" {
		mi.latest = version
	}
}

// SetLatest explicitly sets the latest version.
func (mi *MultiIndex) SetLatest(version string) {
	mi.latest = version
}

// Lookup returns a section by slug for a specific version (case-insensitive).
// Empty version uses latest.
func (mi *MultiIndex) Lookup(slug, version string) (*Section, bool) {
	version = mi.resolveVersion(version)
	versionMap, ok := mi.bySlug[version]
	if !ok {
		return nil, false
	}
	sec, ok := versionMap[strings.ToLower(slug)]
	return sec, ok
}

// GetAll returns all sections for a version. Empty version uses latest.
func (mi *MultiIndex) GetAll(version string) []Section {
	version = mi.resolveVersion(version)
	idx, ok := mi.indices[version]
	if !ok {
		return nil
	}
	return idx.Sections
}

// GetByCategory returns sections in a category for a version.
// Empty version uses latest.
func (mi *MultiIndex) GetByCategory(category, version string) []Section {
	version = mi.resolveVersion(version)
	idx, ok := mi.indices[version]
	if !ok {
		return nil
	}

	var results []Section
	for _, sec := range idx.Sections {
		if sec.Category == category {
			results = append(results, sec)
		}
	}
	return results
}

// GetCategories returns distinct categories for a version.
// Empty version uses latest.
func (mi *MultiIndex) GetCategories(version string) []string {
	version = mi.resolveVersion(version)
	idx, ok := mi.indices[version]
	if !ok {
		return nil
	}

	seen := make(map[string]bool)
	var categories []string
	for _, sec := range idx.Sections {
		if sec.Category != "" && !seen[sec.Category] {
			seen[sec.Category] = true
			categories = append(categories, sec.Category)
		}
	}
	return categories
}

// GetVersions returns all registered versions.
func (mi *MultiIndex) GetVersions() []string {
	return mi.versions
}

// GetLatestVersion returns the latest version string.
func (mi *MultiIndex) GetLatestVersion() string {
	return mi.latest
}

// HasVersion reports whether a version is registered.
func (mi *MultiIndex) HasVersion(version string) bool {
	_, ok := mi.indices[version]
	return ok
}

// MatchVersion maps a user version (e.g., "v1.4.0") to the best available
// docs version (e.g., "v1.4.x"). Returns (matched, nil) on success or
// (latest, error) if no exact match is found.
func (mi *MultiIndex) MatchVersion(userVersion string) (string, error) {
	if userVersion == "" {
		return mi.latest, nil
	}

	// Direct match (e.g., "v1.4.x").
	if mi.HasVersion(userVersion) {
		return userVersion, nil
	}

	// Try to extract major.minor and match to ".x" version.
	parts := strings.Split(userVersion, ".")
	if len(parts) >= 2 {
		majorMinor := parts[0] + "." + parts[1] + ".x"
		if mi.HasVersion(majorMinor) {
			return majorMinor, nil
		}
	}

	return mi.latest, fmt.Errorf("no exact match for version %s, using latest", userVersion)
}

func (mi *MultiIndex) resolveVersion(version string) string {
	if version == "" {
		return mi.latest
	}
	return version
}
