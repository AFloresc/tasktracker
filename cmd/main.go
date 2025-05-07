package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"tasktracker/model"
)

var statuses = [...]string{
	"todo",
	"in-progress",
	"done",
}

const fileName string = "tracker.json"

func main() {
	argWithProg := os.Args

	if len(argWithProg) < 2 {
		fmt.Println("Por favor, proporciona una acción como argumento (add, update, delete, list).")
		return
	}
	//argsWithoutProg := os.Args[1:]

	//arg := os.Args[3]
	tracker, err := loadTrackerFromFile(fileName)
	if err != nil {
		fmt.Printf("Error al cargar el archivo JSON: %v\n", err)
		return
	}
	switch argWithProg[1] {

	case "add", "Add", "ADD":
		Add(argWithProg)
	case "list", "List", "LIST":
		//List(argWithProg)
		fmt.Println("Opción seleccionada: list")
		var statusFilter string
		if len(argWithProg) > 2 {
			statusFilter = argWithProg[2]
		}
		listTasks(tracker, statusFilter)
	case "delete":
		Delete(argWithProg)
	case "update":
		Update(argWithProg)
	case "":
		fmt.Println("no arguments")
	}
}

func List(argsWithoutProg []string) {
	var statusFilter string
	tracker, err := model.HandleTrackerFile(fileName, "")
	if err != nil {
		fmt.Println(err)
	}
	if len(argsWithoutProg) > 2 {
		statusFilter = argsWithoutProg[1]
		tasks, _ := tracker.ListAll()
		b, err := json.Marshal(&tasks)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}
	tasks, _ := tracker.ListAllByStatus(statusFilter)
	b, err := json.Marshal(&tasks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

func Update(argsWithoutProg []string) {
	panic("unimplemented")
}

func Delete(argsWithoutProg []string) {
	panic("unimplemented")
}

func Add(argsWithoutProg []string) {
	if len(argsWithoutProg) < 2 {
		panic("no argument")
	}
	task := model.Task{
		Id:          1,
		Description: argsWithoutProg[1],
		Status:      "todo",
	}
	tracker, err := model.HandleTrackerFile(fileName, "")
	if err != nil {
		fmt.Println(err)
	}
	tracker.AddTAsk(task)
	err = model.SaveTrackerToFile(fileName, tracker)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTrackerFromFile(fileName string) (model.Tracker, error) {
	var tracker model.Tracker

	// Verificar si el archivo existe
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return tracker, fmt.Errorf("el archivo %s no existe", fileName)
	}

	// Leer el archivo JSON
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return tracker, fmt.Errorf("error al leer el archivo JSON: %v", err)
	}

	// Deserializar el contenido del archivo en la estructura Tracker
	err = json.Unmarshal(data, &tracker)
	if err != nil {
		return tracker, fmt.Errorf("error al deserializar el archivo JSON: %v", err)
	}

	return tracker, nil
}

func listTasks(tracker model.Tracker, statusFilter string) {
	if len(tracker.Tasks) == 0 {
		fmt.Println("No hay tareas para listar.")
		return
	}

	if statusFilter == "" {
		// Listar todas las tareas
		fmt.Println("Listado de todas las tareas:")
		for _, task := range tracker.Tasks {
			fmt.Printf("Id: %d, Description: %s, Status: %s, Created: %s, Updated: %s\n",
				task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	} else {
		// Listar tareas filtradas por estado
		fmt.Printf("Listado de tareas con estado '%s':\n", statusFilter)
		found := false
		for _, task := range tracker.Tasks {
			if task.Status == statusFilter {
				fmt.Printf("Id: %d, Description: %s, Status: %s, Created: %s, Updated: %s\n",
					task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
				found = true
			}
		}
		if !found {
			fmt.Printf("No se encontraron tareas con el estado '%s'.\n", statusFilter)
		}
	}
}
