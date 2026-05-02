//go:build !windows

package main

func maybeRelaunchAsSystem() (bool, error) {
	return false, nil
}
