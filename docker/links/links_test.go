package links

import (
	"github.com/nicle-lin/dockerV1.2.0/docker/nat"
	"strings"
	"testing"
)

func TestLinkNaming(t *testing.T) {
	ports := make(nat.PortSet)
	ports[nat.Port("6379/tcp")] = struct{}{}

	link, err := NewLink("172.0.17.3", "172.0.17.2", "/db/docker-1", nil, ports, nil)
	if err != nil {
		t.Fatal(err)
	}

	rawEnv := link.ToEnv()
	env := make(map[string]string, len(rawEnv))
	for _, e := range rawEnv {
		parts := strings.Split(e, "=")
		if len(parts) != 2 {
			t.FailNow()
		}
		env[parts[0]] = parts[1]
	}

	value, ok := env["DOCKER_1_PORT"]

	if !ok {
		t.Fatalf("DOCKER_1_PORT not found in env")
	}

	if value != "tcp://172.0.17.2:6379" {
		t.Fatalf("Expected 172.0.17.2:6379, got %s", env["DOCKER_1_PORT"])
	}
}

func TestLinkNew(t *testing.T) {
	ports := make(nat.PortSet)
	ports[nat.Port("6379/tcp")] = struct{}{}

	link, err := NewLink("172.0.17.3", "172.0.17.2", "/db/docker", nil, ports, nil)
	if err != nil {
		t.Fatal(err)
	}

	if link == nil {
		t.FailNow()
	}
	if link.Name != "/db/docker" {
		t.Fail()
	}
	if link.Alias() != "docker" {
		t.Fail()
	}
	if link.ParentIP != "172.0.17.3" {
		t.Fail()
	}
	if link.ChildIP != "172.0.17.2" {
		t.Fail()
	}
	for _, p := range link.Ports {
		if p != nat.Port("6379/tcp") {
			t.Fail()
		}
	}
}

func TestLinkEnv(t *testing.T) {
	ports := make(nat.PortSet)
	ports[nat.Port("6379/tcp")] = struct{}{}

	link, err := NewLink("172.0.17.3", "172.0.17.2", "/db/docker", []string{"PASSWORD=gordon"}, ports, nil)
	if err != nil {
		t.Fatal(err)
	}

	rawEnv := link.ToEnv()
	env := make(map[string]string, len(rawEnv))
	for _, e := range rawEnv {
		parts := strings.Split(e, "=")
		if len(parts) != 2 {
			t.FailNow()
		}
		env[parts[0]] = parts[1]
	}
	if env["DOCKER_PORT"] != "tcp://172.0.17.2:6379" {
		t.Fatalf("Expected 172.0.17.2:6379, got %s", env["DOCKER_PORT"])
	}
	if env["DOCKER_PORT_6379_TCP"] != "tcp://172.0.17.2:6379" {
		t.Fatalf("Expected tcp://172.0.17.2:6379, got %s", env["DOCKER_PORT_6379_TCP"])
	}
	if env["DOCKER_PORT_6379_TCP_PROTO"] != "tcp" {
		t.Fatalf("Expected tcp, got %s", env["DOCKER_PORT_6379_TCP_PROTO"])
	}
	if env["DOCKER_PORT_6379_TCP_ADDR"] != "172.0.17.2" {
		t.Fatalf("Expected 172.0.17.2, got %s", env["DOCKER_PORT_6379_TCP_ADDR"])
	}
	if env["DOCKER_PORT_6379_TCP_PORT"] != "6379" {
		t.Fatalf("Expected 6379, got %s", env["DOCKER_PORT_6379_TCP_PORT"])
	}
	if env["DOCKER_NAME"] != "/db/docker" {
		t.Fatalf("Expected /db/docker, got %s", env["DOCKER_NAME"])
	}
	if env["DOCKER_ENV_PASSWORD"] != "gordon" {
		t.Fatalf("Expected gordon, got %s", env["DOCKER_ENV_PASSWORD"])
	}
}
