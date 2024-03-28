package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
	"gopkg.in/yaml.v3"
)

type Todo struct {
	Title     string    `yaml:"title"`
	StartTime time.Time `yaml:"start_time"`
	EndTime   time.Time `yaml:"end_time"`
}

type TodoList struct {
	Todos []Todo `yaml:"todos"`
}

// // Custom marshaller for time.Time
// func (t *Todo) MarshalYAML() (interface{}, error) {
// 	type Alias Todo
// 	return &struct {
// 		*Alias
// 		StartTime string `yaml:"start_time"`
// 		EndTime   string `yaml:"end_time"`
// 	}{
// 		Alias:     (*Alias)(t),
// 		StartTime: t.StartTime.Format("2006-01-02 15:04:05"),
// 		EndTime:   t.EndTime.Format("2006-01-02 15:04:05"),
// 	}, nil
// }

// // Custom unmarshaller for time.Time
// func (t *Todo) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	type Alias Todo
// 	aux := &struct {
// 		*Alias
// 		StartTime string `yaml:"start_time"`
// 		EndTime   string `yaml:"end_time"`
// 	}{Alias: (*Alias)(t)}

// 	err := unmarshal(aux)
// 	if err != nil {
// 		return err
// 	}

// 	t.StartTime, err = time.Parse("03:04 PM", aux.StartTime)
// 	if err != nil {
// 		return err
// 	}

// 	t.EndTime, err = time.Parse("03:04 PM", aux.EndTime)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	// Read the YAML file
	yamlFile, err := os.ReadFile("todos.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	ntpTime, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println(err)
	}
	timeDifference := time.Since(ntpTime).Abs().Seconds()
	fmt.Println("Google time", ntpTime, "system time", time.Now(), "time difference in seconds", timeDifference)
	if timeDifference > 60 {
		log.Fatalln("inaccurate time")
	}

	// Parse the YAML file
	var todos TodoList
	err = yaml.Unmarshal(yamlFile, &todos)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {

		var s = false
		for _, todo := range todos.Todos {
			fmt.Println(time.Now(), "-", todo.StartTime, "-", todo.EndTime)

			if time.Now().After(todo.StartTime) && time.Now().Before(todo.EndTime) {
				fmt.Println("You should be doing:", todo.Title)
				s = true
			}
			// // fmt.Println("Title:", todo.Title)
			// fmt.Println("Start Time:", todo.StartTime)
			// // fmt.Println("End Time:", todo.EndTime)
			// fmt.Println()
		}

		if !s {
			fmt.Println("Nothing")
		}
	}
}
