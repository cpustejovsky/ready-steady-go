package filedescriptorreflect

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
)

func MessageDescriptorFromFileDescriptor(name string, fd protoreflect.FileDescriptor, resolver protodesc.Resolver) (protoreflect.MessageDescriptor, error) {
	fdp := protodesc.ToFileDescriptorProto(fd)
	f, err := protodesc.NewFile(fdp, resolver)
	if err != nil {
		return nil, err
	}

	md := f.Messages().ByName(protoreflect.Name(name))

	if md == nil {
		return nil, errors.New("message descriptor was nil")
	}

	return md, nil
}

func MarshalFromFileDescriptor(name string, fd protoreflect.FileDescriptor, data []byte) (proto.Message, error) {
	md, err := MessageDescriptorFromFileDescriptor(name, fd, nil)
	if err != nil {
		return nil, err
	}
	msg := dynamicpb.NewMessage(md)

	log.Printf("Message JSON:\n%#v\n", msg)
	err = proto.Unmarshal(data, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
