package blog

import (
	"io/fs"
	"sort"
)

type OverlayFS struct {
	Upper fs.FS
	Lower fs.FS
}

func NewOverlayFS(upper, lower fs.FS) *OverlayFS {
	return &OverlayFS{
		Upper: upper,
		Lower: lower,
	}
}

func (o *OverlayFS) Open(name string) (fs.File, error) {
	f, err := o.Upper.Open(name)
	if err == nil {
		return f, nil
	}
	return o.Lower.Open(name)
}

func (o *OverlayFS) ReadDir(name string) ([]fs.DirEntry, error) {
	upperDir, upperErr := fs.ReadDir(o.Upper, name)
	lowerDir, lowerErr := fs.ReadDir(o.Lower, name)

	if upperErr != nil && lowerErr != nil {
		if upperErr != nil {
			return nil, upperErr
		}
		return nil, lowerErr
	}

	merged := make(map[string]fs.DirEntry)
	for _, e := range lowerDir {
		merged[e.Name()] = e
	}
	for _, e := range upperDir {
		merged[e.Name()] = e // upper layer overrides
	}

	entries := make([]fs.DirEntry, 0, len(merged))
	for _, e := range merged {
		entries = append(entries, e)
	}

	return entries, nil
}

func (o *OverlayFS) Glob(pattern string) ([]string, error) {
	upperMatches, _ := fs.Glob(o.Upper, pattern)
	lowerMatches, _ := fs.Glob(o.Lower, pattern)

	matchMap := make(map[string]struct{})
	for _, m := range lowerMatches {
		matchMap[m] = struct{}{}
	}
	for _, m := range upperMatches {
		matchMap[m] = struct{}{}
	}

	results := make([]string, 0, len(matchMap))
	for m := range matchMap {
		results = append(results, m)
	}

	sort.Strings(results)
	return results, nil
}
