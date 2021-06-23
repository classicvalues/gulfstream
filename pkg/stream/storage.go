package stream

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Storage interface {
	BlankStream() *Stream
	Persist(ctx context.Context, s *Stream) error
	Load(ctx context.Context, streamName string, streamID uuid.UUID, owner uuid.UUID) (*Stream, error)
	MarkUnpublished(ctx context.Context, s *Stream) error
	Walk(ctx context.Context, streamName string, streams []uuid.UUID, owner uuid.UUID, iter func(*Stream) error) error
}

func NewStorage(newStream func() *Stream) Storage {
	return &stateStorage{
		data:        make(map[key][]byte),
		versions:    make(map[key]int),
		blankStream: newStream,
	}
}

type stateStorage struct {
	mu          sync.RWMutex
	data        map[key][]byte
	blankStream func() *Stream
	versions    map[key]int
}

type key struct {
	streamType string
	streamID   uuid.UUID
	owner      uuid.UUID
}

func (s *stateStorage) BlankStream() *Stream {
	return s.blankStream()
}

func (s *stateStorage) Persist(_ context.Context, stream *Stream) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := key{streamType: stream.Name(), owner: stream.Owner(), streamID: stream.ID()}
	rawData, found := s.data[k]
	if found {
		prev := s.blankStream()
		if err := prev.UnmarshalBinary(rawData); err != nil {
			return err
		}
		if prev.Version() != stream.PreviousVersion() {
			return fmt.Errorf("mismatch stream version. got v%d, expected v%d",
				prev.Version(), stream.PreviousVersion())
		}
	}
	stream.version = stream.Version()
	data, err := stream.MarshalBinary()
	if err != nil {
		return err
	}
	s.data[k] = data
	return nil
}

func (s *stateStorage) Load(_ context.Context, streamName string, streamID uuid.UUID, owner uuid.UUID) (*Stream, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	k := key{streamType: streamName, owner: owner, streamID: streamID}
	rawData, found := s.data[k]
	if !found {
		return nil, fmt.Errorf("%s{StreamID:%s, Owner:%s} not found",
			streamName, streamID, owner)
	}
	blankStream := s.blankStream()
	if err := blankStream.UnmarshalBinary(rawData); err != nil {
		return nil, err
	}
	return blankStream, nil
}

func (s *stateStorage) MarkUnpublished(ctx context.Context, cur *Stream) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := key{streamType: cur.Name(), owner: cur.Owner(), streamID: cur.ID()}
	s.versions[k] = cur.Version()
	return nil
}

func (s *stateStorage) Walk(_ context.Context, streamName string, streams []uuid.UUID, owner uuid.UUID, iter func(*Stream) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var k key
	for _, streamID := range streams {
		k = key{streamType: streamName, owner: owner, streamID: streamID}
		rawData, found := s.data[k]
		if !found {
			return fmt.Errorf("%s{StreamID:%s, Owner:%s} not found",
				streamName, streamID, owner)
		}
		blankStream := s.blankStream()
		if err := blankStream.UnmarshalBinary(rawData); err != nil {
			return err
		}
		return iter(blankStream)
	}
	return nil
}
