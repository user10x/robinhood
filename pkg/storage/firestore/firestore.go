package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Service struct {
	ref *firestore.DocumentRef
}

func NewFireStore(name string, project string) *Service {
	ctx := context.Background()
	fireClient, err := firestore.NewClient(ctx, project)
	if err != nil {
		//
	}
	col := fireClient.Collection(name)

	ref := col.Doc(name)

	return &Service{
		ref: ref,
	}

}

func (s *Service) Insert(ctx context.Context, data interface{}) {

	wr, err := s.ref.Create(ctx, data)
	if err != nil {
		// TODO: Handle error.
	}

	fmt.Println(wr)
}
