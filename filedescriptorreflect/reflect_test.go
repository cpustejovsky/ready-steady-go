package filedescriptorreflect_test

import (
	"github.com/cpustejovsky/ready-steady-go/filedescriptorreflect"
	"github.com/cpustejovsky/ready-steady-go/helloworld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"testing"
)

var protoname = "HelloRequest"

func TestMessageDescriptorFromFileDescriptor(t *testing.T) {
	fd := helloworld.File_helloworld_helloworld_proto
	md, err := filedescriptorreflect.MessageDescriptorFromFileDescriptor(protoname, fd, nil)
	require.Nil(t, err)
	assert.Equal(t, protoname, string(md.FullName().Name()))
}

func TestMarshalProtoFromFileDescriptor(t *testing.T) {
	msg := helloworld.HelloRequest{
		Name: "Charles",
	}
	bin, err := proto.Marshal(&msg)
	require.Nil(t, err)
	fd := helloworld.File_helloworld_helloworld_proto
	tmp := &helloworld.HelloRequest{}
	err = proto.Unmarshal(bin, tmp)
	require.Nil(t, err)
	assert.Equal(t, msg.GetName(), tmp.GetName())
	newMsg, err := filedescriptorreflect.MarshalFromFileDescriptor(protoname, fd, bin)
	assert.Nil(t, err)

	//log.Printf("New Message:\t%v\n\n", newMsg)
	//log.Printf("New Message:\t%+v\n\n", newMsg)
	//log.Printf("New Message:\t%#v\n\n", newMsg)
	//log.Printf("New Message Descriptor:\t%#v\n\n", newMsg.ProtoReflect().Descriptor())
	//log.Printf("Old Message Descriptor:\t%#v\n\n", msg.ProtoReflect().Descriptor())
	t.Log(protojson.Format(newMsg))
}
