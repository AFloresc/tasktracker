package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIt(t *testing.T) {
	task := Task{
		Id:          1,
		Description: "Descripción de la tarea 1",
		Status:      "pending",
	}

	task2 := Task{
		Id:          1,
		Description: "Descripción de la tarea 2",
		Status:      "pending",
	}

	task3 := Task{
		Id:          1,
		Description: "Descripción de la tarea 3",
		Status:      "done",
	}

	task4 := Task{
		Id:          1,
		Description: "Descripción de la tarea 4",
		Status:      "done",
	}

	task5 := Task{
		Id:          1,
		Description: "Descripción de la tarea 5",
		Status:      "in-progress",
	}

	fmt.Println(task)
	fmt.Println(task2)

	tracker := Tracker{}

	tracker.AddTAsk(task)
	tracker.AddTAsk(task2)
	tracker.AddTAsk(task3)
	tracker.AddTAsk(task4)
	tracker.AddTAsk(task5)

	fmt.Println(tracker)

	tracker.RemoveTask(1)
	fmt.Println(tracker)

	fmt.Println(tracker.ListAllByStatus("done"))
	fmt.Println(tracker.ListAllByStatus("in-progress"))

	fmt.Println(tracker.ListAll())
	//JSON
	b, err := json.Marshal(&tracker)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fileName := "tracker.json"

	// Llamar a la función para manejar el archivo JSON
	trackerJSON, err := HandleTrackerFile(fileName, string(b))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Imprimir el contenido del Tracker resultante
	fmt.Printf("Tracker: %+v\n", trackerJSON)
}

func TestArgs(t *testing.T) {
	args := []string{"add", "buy groceries"}

	switch args[0] {
	case "add", "Add", "ADD":
		break
	case "list", "List", "LIST":
		break
	case "":
		fmt.Println("no arguments")

	}
}
