package main

import (
	"context"
	"fmt"

	//	plugin "github.com/hashicorp/go-plugin"

	"github.com/docker/docker/api/types"
	moby "github.com/docker/docker/client"
	"github.com/spiffe/spire/proto/agent/workloadattestor"
	"github.com/spiffe/spire/proto/common"
)

type imageIdentifier interface {
	imageIdentify(int32) string
}

type DockerPlugin struct {
	docker imageIdentifier
}

func (p *DockerPlugin) Attest(req *workloadattestor.AttestRequest) (*workloadattestor.AttestResponse, error) {
	return &workloadattestor.AttestResponse{
		[]*common.Selector{&common.Selector{
			Type:  "docker",
			Value: "imageid:" + p.docker.imageIdentify(req.Pid),
		}},
	}, nil
}

type docker struct {
	cli *moby.Client
}

func (d *docker) imageIdentify(pid int32) string {
	cs, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	for _, c := range cs {
		cjson, err := d.cli.ContainerInspect(context.Background(), c.ID)
		if err != nil {
			fmt.Println("err: " + err.Error())
			return ""
		}
		if pid == int32(cjson.State.Pid) {
			return c.ImageID
		}
	}
	return ""
}

// TODO: use better logging statements
func main() {
	// TODO: incorporate docker configuration from spire/conf/agent/plugin/docker.conf
	cli, err := moby.NewEnvClient()
	if err != nil {
		fmt.Println("err creating client: " + err.Error())
	}
	d := &DockerPlugin{&docker{cli}}

	// Manual testing: change bellow pid to your docker container pid
	// docker inspect --format {{.State.Pid}} <enter container id>
	samplePid := int32(7429)

	resp, err := d.Attest(&workloadattestor.AttestRequest{samplePid})
	fmt.Println("result: " + resp.Selectors[0].Value)

	//plugin.Serve(&plugin.ServeConfig{
	//	HandshakeConfig: workloadattestor.Handshake,
	//	Plugins: map[string]plugin.Plugin{
	//		"docker": workloadattestor.WorkloadAttestorPlugin{
	//			WorkloadAttestorImpl: &DockerPlugin{
	//				&docker{cli},
	//			},
	//		},
	//	},
	//	GRPCServer: plugin.DefaultGRPCServer,
	//})
}
