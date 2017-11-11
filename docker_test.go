package main

import (
	"os"
	"testing"

	"github.com/spiffe/spire/proto/agent/workloadattestor"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type DockerPluginSuite struct{}

var _ = Suite(&DockerPluginSuite{})

var expectedImageID = "sha256:86f1955c7430d56b89b41182909329143e6494e1e411bc04d0c290cc3172d41b"

func (dps *DockerPluginSuite) Test_CreateSelector_WithTypeAndImgID(c *C) {
	dp := &DockerPlugin{&mockDockerClient{}}

	resp, _ := dp.Attest(&workloadattestor.AttestRequest{Pid: int32(os.Getpid())})

	c.Assert(resp.Selectors[0].Type, Equals, "docker")
	c.Assert(resp.Selectors[0].Value, Equals, "imageid:"+expectedImageID)
}

type mockDockerClient struct{}

func (cl *mockDockerClient) imageIdentify(pid int32) string {
	return expectedImageID
}

// TODO: Error tests
// TODO:	No docker image id
// TODO:	Invalid docker image id
// TODO: Finish workloadattestor API
// TODO:	Configure()
// TODO:	GetPluginInfo()
