package registry

import (
	"context"

	ref "github.com/distribution/reference"
	mobyClient "github.com/moby/moby/client"
	"github.com/naiba-forks/watchtower/pkg/registry/helpers"
	watchtowerTypes "github.com/naiba-forks/watchtower/pkg/types"
	log "github.com/sirupsen/logrus"
)

func GetPullOptions(imageName string) (mobyClient.ImagePullOptions, error) {
	auth, err := EncodedAuth(imageName)
	log.Debugf("Got image name: %s", imageName)
	if err != nil {
		return mobyClient.ImagePullOptions{}, err
	}

	if auth == "" {
		return mobyClient.ImagePullOptions{}, nil
	}

	// CREDENTIAL: Uncomment to log docker config auth
	// log.Tracef("Got auth value: %s", auth)

	return mobyClient.ImagePullOptions{
		RegistryAuth:  auth,
		PrivilegeFunc: DefaultAuthHandler,
	}, nil
}

// DefaultAuthHandler will be invoked if an AuthConfig is rejected
// It could be used to return a new value for the "X-Registry-Auth" authentication header,
// but there's no point trying again with the same value as used in AuthConfig
func DefaultAuthHandler(_ context.Context) (string, error) {
	log.Debug("Authentication request was rejected. Trying again without authentication")
	return "", nil
}

// WarnOnAPIConsumption will return true if the registry is known-expected
// to respond well to HTTP HEAD in checking the container digest -- or if there
// are problems parsing the container hostname.
// Will return false if behavior for container is unknown.
func WarnOnAPIConsumption(container watchtowerTypes.Container) bool {

	normalizedRef, err := ref.ParseNormalizedNamed(container.ImageName())
	if err != nil {
		return true
	}

	containerHost, err := helpers.GetRegistryAddress(normalizedRef.Name())
	if err != nil {
		return true
	}

	if containerHost == helpers.DefaultRegistryHost || containerHost == "ghcr.io" {
		return true
	}

	return false
}
