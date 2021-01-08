package speaker

import "github.com/didil/goblero/pkg/blero"

var _ Queuer = (*Queue)(nil)

// Queuer contains the methods available
// to the Queue type
type Queuer interface {
	Add() error
	Remove(number int) error
	List() []Track
	Find(name string) error
}

// Queue
type Queue struct {
	Backend *blero.Blero
	Tracks  []Track
}

func (q Queue) Add() error {
	panic("implement me")
}

func (q Queue) Remove(number int) error {
	panic("implement me")
}

func (q Queue) List() []Track {
	panic("implement me")
}

func (q Queue) Find(name string) error {
	panic("implement me")
}

// Track is a single track to be added
// to the queue for playing
type Track struct {
	ID       string
	Name     string
	URL      string
	Position int
}
