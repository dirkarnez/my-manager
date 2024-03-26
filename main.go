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

	// Parse the YAML file
	var todos TodoList
	err = yaml.Unmarshal(yamlFile, &todos)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	ntpTime, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ntpTime)

	timer := time.NewTicker(1 * time.Second)
	// Access the parsed todos
	for {
		// Wait for the timer to tick
		<-timer.C

		for _, todo := range todos.Todos {
			if ntpTime.After(todo.StartTime) && ntpTime.Before(todo.EndTime) {
				fmt.Println("You should be doing:", todo.Title)
			}
			// // fmt.Println("Title:", todo.Title)
			// fmt.Println("Start Time:", todo.StartTime)
			// // fmt.Println("End Time:", todo.EndTime)
			// fmt.Println()
		}
	}

}
