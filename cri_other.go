//go:build !darwin

package cri

import "errors"

const (
	CriSocketPath      = "/run/containerd/containerd.sock"
	PodVolRoot         = "/run/vk-cri/volumes/"
	PodSecretVolDir    = "/secrets"
	PodConfigMapVolDir = "/configmaps"
)

func getOSImage() (string, error) {
	return "", errors.New("not implemented")
}
