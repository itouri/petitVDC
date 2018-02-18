package instance

import "golang.org/x/net/context"

func (i *InstanceAPI) Create(ctx context.Context, req *CreateRequest) (*CreateReply, error) {
	return &CreateReply{InstanceId: "1"}, nil
}
