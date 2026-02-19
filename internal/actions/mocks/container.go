package mocks

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/image"
	"github.com/moby/moby/api/types/network"

	pkgcontainer "github.com/naiba-forks/watchtower/pkg/container"
	wt "github.com/naiba-forks/watchtower/pkg/types"
)

// CreateMockContainer creates a container substitute valid for testing
func CreateMockContainer(id string, name string, img string, created time.Time) wt.Container {
	content := container.InspectResponse{
		ID:    id,
		Image: img,
		Name:  name,
		HostConfig: &container.HostConfig{
			PortBindings: network.PortMap{},
		},
		Config: &container.Config{
			Image:        img,
			Labels:       make(map[string]string),
			ExposedPorts: network.PortSet{},
		},
	}
	content.Created = created.String()
	return pkgcontainer.NewContainer(
		&content,
		CreateMockImageInfo(img),
	)
}

// CreateMockImageInfo returns a mock image info struct based on the passed image
func CreateMockImageInfo(img string) *image.InspectResponse {
	return &image.InspectResponse{
		ID: img,
		RepoDigests: []string{
			img,
		},
	}
}

// CreateMockContainerWithImageInfo should only be used for testing
func CreateMockContainerWithImageInfo(id string, name string, img string, created time.Time, imageInfo image.InspectResponse) wt.Container {
	return CreateMockContainerWithImageInfoP(id, name, img, created, &imageInfo)
}

// CreateMockContainerWithImageInfoP should only be used for testing
func CreateMockContainerWithImageInfoP(id string, name string, img string, created time.Time, imageInfo *image.InspectResponse) wt.Container {
	content := container.InspectResponse{
		ID:    id,
		Image: img,
		Name:  name,
		Config: &container.Config{
			Image:  img,
			Labels: make(map[string]string),
		},
	}
	content.Created = created.String()
	return pkgcontainer.NewContainer(
		&content,
		imageInfo,
	)
}

// CreateMockContainerWithDigest should only be used for testing
func CreateMockContainerWithDigest(id string, name string, img string, created time.Time, digest string) wt.Container {
	c := CreateMockContainer(id, name, img, created)
	c.ImageInfo().RepoDigests = []string{digest}
	return c
}

// CreateMockContainerWithConfig creates a container substitute valid for testing
func CreateMockContainerWithConfig(id string, name string, img string, running bool, restarting bool, created time.Time, config *container.Config) wt.Container {
	content := container.InspectResponse{
		ID:    id,
		Image: img,
		Name:  name,
		State: &container.State{
			Running:    running,
			Restarting: restarting,
		},
		HostConfig: &container.HostConfig{
			PortBindings: network.PortMap{},
		},
		Config: config,
	}
	content.Created = created.String()
	return pkgcontainer.NewContainer(
		&content,
		CreateMockImageInfo(img),
	)
}

// CreateContainerForProgress creates a container substitute for tracking session/update progress
func CreateContainerForProgress(index int, idPrefix int, nameFormat string) (wt.Container, wt.ImageID) {
	indexStr := strconv.Itoa(idPrefix + index)
	mockID := indexStr + strings.Repeat("0", 61-len(indexStr))
	contID := "c79" + mockID
	contName := fmt.Sprintf(nameFormat, index+1)
	oldImgID := "01d" + mockID
	newImgID := "d0a" + mockID
	imageName := fmt.Sprintf("mock/%s:latest", contName)
	config := &container.Config{
		Image: imageName,
	}
	c := CreateMockContainerWithConfig(contID, contName, oldImgID, true, false, time.Now(), config)
	return c, wt.ImageID(newImgID)
}

// CreateMockContainerWithLinks should only be used for testing
func CreateMockContainerWithLinks(id string, name string, img string, created time.Time, links []string, imageInfo *image.InspectResponse) wt.Container {
	content := container.InspectResponse{
		ID:    id,
		Image: img,
		Name:  name,
		HostConfig: &container.HostConfig{
			Links: links,
		},
		Config: &container.Config{
			Image:  img,
			Labels: make(map[string]string),
		},
	}
	content.Created = created.String()
	return pkgcontainer.NewContainer(
		&content,
		imageInfo,
	)
}
