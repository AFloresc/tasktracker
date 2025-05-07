package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	//Id          uuid.UUID
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created,omitempty"`
	UpdatedAt   time.Time `json:"updated,omitempty"`
}

type Tracker struct {
	Tasks    []Task `json:"tasks"`
	Counter  int    `json:"counter"`
	Elements int    `json:"elements"`
}

func (tracker *Tracker) AddTAsk(t Task) {
	if t.Status == "" {
		t.Status = "todo" // Default value
	}
	t.Id = tracker.Counter
	t.CreatedAt = time.Now()
	tracker.Tasks = append(tracker.Tasks, t)
	tracker.Counter++
	tracker.Elements++
}

func (tracker *Tracker) RemoveTask(taskID int) error {
	if len(tracker.Tasks) == 0 || tracker.Counter == 0 {
		return errors.New("no elements in JSON")
	}

	for index, task := range tracker.Tasks {
		if task.Id == taskID {
			tracker.Tasks = remove(tracker.Tasks, index)
			tracker.Elements--
			break
		}
	}
	return nil
}

func (tracker *Tracker) ChangeStatusTask(status string, taskID int) error {
	if status == "in-progress" || status == "done" {
		for index, task := range tracker.Tasks {
			if task.Id == taskID {
				tracker.Tasks[index].Status = status
				tracker.Tasks[index].UpdatedAt = time.Now()
				break
			}
		}
		return nil
	}
	return errors.New("not valid status")
}

func (tracker *Tracker) ListAllByStatus(status string) (result []Task, err error) {
	if status == "in-progress" || status == "done" || status == "todo" {
		for _, task := range tracker.Tasks {
			if task.Status == status {
				result = append(result, task)
			}
		}
		return result, nil
	}
	return nil, errors.New("incorrect status")
}

func (tracker *Tracker) ListAll() (result []Task, err error) {
	if tracker.Elements == 0 {
		return nil, errors.New("there are no elements")
	}
	return tracker.Tasks, nil
}

func remove(slice []Task, index int) []Task {
	return append(slice[:index], slice[index+1:]...)
}

func seekTask(tasks []Task, taskID int) (index int, err error) {
	if len(tasks) == 0 {
		return 0, errors.New("no tasks")
	}
	for index, task := range tasks {
		if task.Id == taskID {
			return index, nil
		}
	}
	return 0, errors.New("task not found")
}

// Función para guardar el contenido del Tracker en el archivo (sobrescribiendo si existe)
func SaveTrackerToFile(fileName string, tracker Tracker) error {
	// Convertir el Tracker a formato JSON
	data, err := json.MarshalIndent(tracker, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar el Tracker a JSON: %v", err)
	}

	// Escribir o sobrescribir el archivo
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	fmt.Println("Contenido guardado exitosamente en el archivo.")
	return nil
}

func HandleTrackerFile(fileName string, trackerJSON string) (Tracker, error) {
	var tracker Tracker

	// Verificar si el archivo existe
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Si no existe, crear el archivo con los datos del JSON
		err := ioutil.WriteFile(fileName, []byte(trackerJSON), 0644)
		if err != nil {
			return tracker, fmt.Errorf("error al crear el archivo: %v", err)
		}
		fmt.Println("Archivo creado exitosamente con el contenido proporcionado.")
	} else {
		// Si el archivo existe, cargar su contenido en un Tracker
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return tracker, fmt.Errorf("error al leer el archivo: %v", err)
		}
		err = json.Unmarshal(data, &tracker)
		if err != nil {
			return tracker, fmt.Errorf("error al deserializar el contenido del archivo: %v", err)
		}
		fmt.Println("Contenido del archivo cargado exitosamente.")
	}

	// Retornar el Tracker cargado o vacío si se creó el archivo
	return tracker, nil
}
