package config

import "testing"

func TestSetTelegraf(t *testing.T) {
	SetTelegraf("test-2")
}

func TestRestartTelegraf(t *testing.T) {
	RestartTelegraf()
}