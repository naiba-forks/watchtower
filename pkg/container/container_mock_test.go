package container

import (
	dockerspec "github.com/moby/docker-image-spec/specs-go/v1"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/image"
	"github.com/moby/moby/api/types/network"
)

type MockContainerUpdate func(*container.InspectResponse, *image.InspectResponse)

func MockContainer(updates ...MockContainerUpdate) *Container {
	containerInfo := container.InspectResponse{
		ID:         "container_id",
		Image:      "image",
		Name:       "test-containrrr",
		HostConfig: &container.HostConfig{},
		Config: &container.Config{
			Labels: map[string]string{},
		},
	}
	img := image.InspectResponse{
		ID:     "image_id",
		Config: &dockerspec.DockerOCIImageConfig{},
	}

	for _, update := range updates {
		update(&containerInfo, &img)
	}
	return NewContainer(&containerInfo, &img)
}

func WithPortBindings(portBindingSources ...string) MockContainerUpdate {
	return func(c *container.InspectResponse, i *image.InspectResponse) {
		portBindings := network.PortMap{}
		for _, pbs := range portBindingSources {
			portBindings[network.MustParsePort(pbs)] = []network.PortBinding{}
		}
		c.HostConfig.PortBindings = portBindings
	}
}

func WithImageName(name string) MockContainerUpdate {
	return func(c *container.InspectResponse, i *image.InspectResponse) {
		c.Config.Image = name
		i.RepoTags = append(i.RepoTags, name)
	}
}

func WithLinks(links []string) MockContainerUpdate {
	return func(c *container.InspectResponse, i *image.InspectResponse) {
		c.HostConfig.Links = links
	}
}

func WithLabels(labels map[string]string) MockContainerUpdate {
	return func(c *container.InspectResponse, i *image.InspectResponse) {
		c.Config.Labels = labels
	}
}

func WithContainerState(state container.State) MockContainerUpdate {
	return func(cnt *container.InspectResponse, img *image.InspectResponse) {
		cnt.State = &state
	}
}

func WithHealthcheck(healthConfig container.HealthConfig) MockContainerUpdate {
	return func(cnt *container.InspectResponse, img *image.InspectResponse) {
		cnt.Config.Healthcheck = &healthConfig
	}
}

func WithImageHealthcheck(healthConfig container.HealthConfig) MockContainerUpdate {
	return func(cnt *container.InspectResponse, img *image.InspectResponse) {
		img.Config.Healthcheck = &healthConfig
	}
}
