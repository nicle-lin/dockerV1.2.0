package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestInspectImage(t *testing.T) {
	imageTest := "scratch"
	imageTestId := "511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158"
	imagesCmd := exec.Command(dockerBinary, "inspect", "--format='{{.Id}}'", imageTest)

	out, exitCode, err := runCommandWithOutput(imagesCmd)
	if exitCode != 0 || err != nil {
		t.Fatalf("failed to inspect image")
	}
	if id := strings.TrimSuffix(out, "\n"); id != imageTestId {
		t.Fatalf("Expected id: %s for image: %s but received id: %s", imageTestId, imageTest, id)
	}
	logDone("inspect - inspect an image")
}
