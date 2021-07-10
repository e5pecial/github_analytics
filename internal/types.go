package internal

import (
	"errors"
	"reflect"
	"strconv"
)

type Actor struct {
	ID       int64
	Username string
}

type Commit struct {
	Sha     string
	Message string
	EventId int64
}

type Event struct {
	ID      int64
	Type    string
	ActorId int64
	RepoId  int64
}

type Repository struct {
	ID   int64
	Name string
}

type Counter struct {
	ID    int64
	Count int
}

func (a *Actor) Unmarshal(row []string) error {
	if len(row) != reflect.TypeOf(*a).NumField() {
		return errors.New("Incorrect elements count")
	}
	actorID, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return errors.New("Incorrect actorID type")
	}
	username := row[1]
	a.ID = actorID
	a.Username = username
	return nil
}

func (r *Repository) Unmarshal(row []string) error {
	if len(row) != reflect.TypeOf(*r).NumField() {
		return errors.New("Incorrect elements count")
	}
	repoId, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return errors.New("Incorrect repoID type")
	}
	name := row[1]

	r.ID = repoId
	r.Name = name
	return nil
}

func (c *Commit) Unmarshal(row []string) error {
	if len(row) != reflect.TypeOf(*c).NumField() {
		return errors.New("Incorrect elements count")
	}

	sha := row[0]
	message := row[1]

	eventId, err := strconv.ParseInt(row[2], 10, 64)
	if err != nil {
		return errors.New("Incorrect eventID type")
	}
	c.Sha = sha
	c.Message = message
	c.EventId = eventId
	return nil
}

func (e *Event) Unmarshal(row []string) error {
	if len(row) != reflect.TypeOf(*e).NumField() {
		return errors.New("Incorrect elements count")
	}
	eventId, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return errors.New("Incorrect eventID type")
	}

	typeEvent := row[1]

	actorId, err := strconv.ParseInt(row[2], 10, 64)
	if err != nil {
		return errors.New("Incorrect actorID type")
	}

	repoId, err := strconv.ParseInt(row[3], 10, 64)
	if err != nil {
		return errors.New("Incorrect repoID type")
	}
	e.ID = eventId
	e.Type = typeEvent
	e.ActorId = actorId
	e.RepoId = repoId
	return nil
}
