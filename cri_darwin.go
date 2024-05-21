package cri

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	CriSocketPath      = "/var/run/containerd/containerd.sock"
	PodVolRoot         = "/var/run/vk-cri/volumes/"
	PodSecretVolDir    = "/var/lib/vk-cri/secrets"
	PodConfigMapVolDir = "/var/lib/vk-cri/configmaps"
)

func getOSImage() (string, error) {
	productNameCmd := exec.Command("sw_vers", "-productName")
	productVersionCmd := exec.Command("sw_vers", "-productVersion")
	buildVersionCmd := exec.Command("sw_vers", "-buildVersion")

	productName, err := productNameCmd.Output()
	if err != nil {
		return "", err
	}

	productVersion, err := productVersionCmd.Output()
	if err != nil {
		return "", err
	}

	buildVersion, err := buildVersionCmd.Output()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s (%s)", strings.TrimSpace(string(productName)), strings.TrimSpace(string(productVersion)), strings.TrimSpace(string(buildVersion))), nil
}
