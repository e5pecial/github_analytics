package internal

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

func GetRepositoriesByWatchEvents(n int) {
	var wg sync.WaitGroup

	eventsCh := make(chan []string)
	repos := make(map[int64]*Repository)
	reader := NewReader("/assets/events.csv", eventsCh)
	wg.Add(1)
	go getRepos(repos, &wg)
	go reader.ReadCsvToChannel()
	wg.Wait()
	watchEvents, err := getWatchEventsByRepoId(eventsCh)
	if err != nil {
		fmt.Println(err)
		return
	}
	sortedArray := sortByValue(watchEvents)
	sortedArray = sortedArray[:n]
	prettifyRepo(sortedArray, repos)
}

func GetAutorsByCommits(n int) {
	var wg sync.WaitGroup

	eventsCh := make(chan []string)
	authors := make(map[int64]*Actor)
	commitsByEvent := make(map[int64]int)
	reader := NewReader("/assets/events.csv", eventsCh)

	wg.Add(2)
	go getAuthors(authors, &wg)
	go getCommits(commitsByEvent, &wg)
	go reader.ReadCsvToChannel()
	wg.Wait()

	events, err := getEventsCountByActorID(eventsCh, commitsByEvent)
	if err != nil {
		fmt.Println(err)
		return
	}

	sortedArray := sortByValue(events)
	sortedArray = sortedArray[:n]

	prettifyActor(sortedArray, authors)
}

func GetRepositoriesByCommits(n int) {
	var wg sync.WaitGroup

	eventsCh := make(chan []string)
	commitsByEvent := make(map[int64]int)
	repos := make(map[int64]*Repository)
	reader := NewReader("/assets/events.csv", eventsCh)
	wg.Add(2)
	go getRepos(repos, &wg)
	go getCommits(commitsByEvent, &wg)
	go reader.ReadCsvToChannel()
	wg.Wait()
	events, err := getEventsCountByRepoID(eventsCh, commitsByEvent)
	if err != nil {
		fmt.Println(err)
		return
	}

	sortedArray := sortByValue(events)
	sortedArray = sortedArray[:n]

	prettifyRepo(sortedArray, repos)

}

func getAuthors(actors map[int64]*Actor, wg *sync.WaitGroup) error {
	records := make(chan []string)
	defer wg.Done()

	reader := NewReader("/assets/actors.csv", records)

	go reader.ReadCsvToChannel()
	for record := range records {
		actor := Actor{}
		err := actor.Unmarshal(record)
		if err != nil {
			return errors.New("Unmarshalling problem: " + err.Error())
		}
		actors[actor.ID] = &actor
	}

	return nil
}

func getRepos(repos map[int64]*Repository, wg *sync.WaitGroup) error {
	records := make(chan []string)
	defer wg.Done()
	reader := NewReader("/assets/repos.csv", records)
	go reader.ReadCsvToChannel()
	for record := range records {
		repo := Repository{}
		err := repo.Unmarshal(record)
		if err != nil {
			return errors.New("Unmarshalling problem: " + err.Error())
		}
		repos[repo.ID] = &repo
	}
	return nil
}

func getCommits(commits map[int64]int, wg *sync.WaitGroup) error {
	commitsCh := make(chan []string)
	reader := NewReader("/assets/commits.csv", commitsCh)
	defer wg.Done()
	go reader.ReadCsvToChannel()

	for record := range commitsCh {
		commit := Commit{}
		err := commit.Unmarshal(record)
		if err != nil {
			return errors.New("Unmarshalling problem: " + err.Error())
		}
		if err != nil {
			return err
		}
		commits[commit.EventId] += 1
	}
	return nil
}

func getEventsCountByRepoID(records <-chan []string, commitsMapping map[int64]int) (map[int64]int, error) {
	events := make(map[int64]int)

	for record := range records {
		event := Event{}
		err := event.Unmarshal(record)
		if err != nil {
			return nil, errors.New("Unmarshalling problem: " + err.Error())
		}
		if event.Type == "PushEvent" {
			events[event.RepoId] += commitsMapping[event.ID]
		}
	}
	return events, nil
}

func getEventsCountByActorID(records <-chan []string, commitsMapping map[int64]int) (map[int64]int, error) {
	events := make(map[int64]int)

	for record := range records {
		event := Event{}
		err := event.Unmarshal(record)
		if err != nil {
			return nil, errors.New("Unmarshalling problem: " + err.Error())
		}
		switch event.Type {
		case "PushEvent":
			events[event.ActorId] += commitsMapping[event.ID]
		case "PullRequestEvent":
			events[event.ActorId] += 1
		}
	}
	return events, nil
}

func getWatchEventsByRepoId(records <-chan []string) (map[int64]int, error) {
	events := make(map[int64]int)

	for record := range records {
		event := Event{}
		err := event.Unmarshal(record)
		if err != nil {
			return nil, errors.New("Unmarshalling problem: " + err.Error())
		}
		if event.Type == "WatchEvent" {
			events[event.RepoId] += 1
		}
	}
	return events, nil
}

func sortByValue(mapping map[int64]int) []Counter {

	var counterArray []Counter
	for k, v := range mapping {
		counterArray = append(counterArray, Counter{k, v})
	}

	sort.Slice(counterArray, func(i, j int) bool {
		return counterArray[i].Count > counterArray[j].Count
	})

	return counterArray
}

func prettifyRepo(counterArray []Counter, repos map[int64]*Repository) {
	for _, element := range counterArray {
		if val, ok := repos[element.ID]; ok {
			fmt.Printf("| ID: %d | Repository Name: %s | Count: %d |\n", val.ID, val.Name, element.Count)
		}
	}
}

func prettifyActor(counterArray []Counter, authors map[int64]*Actor) {
	for _, element := range counterArray {
		if val, ok := authors[element.ID]; ok {
			fmt.Printf("| ID: %d | Username: %s | Count: %d |\n", val.ID, val.Username, element.Count)
		}
	}
}
