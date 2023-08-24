package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"robinhood/pkg/models/client/robinhood"
)

type Service struct {
	ref    *firestore.DocumentRef
	batch  *firestore.BulkWriter
	colRef *firestore.CollectionRef
}

func NewFireStore(ctx context.Context, name string, project string) (*Service, error) {
	fireClient, err := firestore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}

	col := fireClient.Collection("main")
	bw := fireClient.BulkWriter(ctx)

	ref := col.Doc(name)

	return &Service{
		ref:    ref,
		batch:  bw,
		colRef: col,
	}, nil

}

func (s *Service) InsertInstruments(data *robinhood.Instruments) error {
	return s.batchWrite(data)
}

func (s *Service) batchWrite(data *robinhood.Instruments) error {
	for _, v := range data.Results {
		ref := s.colRef.Doc(v.Symbol)
		_, err := s.batch.Set(ref, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Commit() {
	s.batch.Flush()
}
