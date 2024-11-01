package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)



type Status string

const (
	New Status = "New"
	InProgress Status = "InProgress"
	Done Status = "Done"
)

type Todo struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Status Status `json:"status"`
	Deadline time.Time `json:"deadline"`
	Bucket string `json:"bucket"`
}

var pendingTasks []Todo

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "testttttt\n")
}


func check(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _,h := range headers{
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func addTodo(w http.ResponseWriter, req *http.Request) {

	body,err := io.ReadAll(req.Body)

	if (err != nil){
		http.Error(w,"Unable to read the request", http.StatusBadRequest)
		return
	}

	var task Todo

	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w,"Invalid Request Body format", http.StatusBadRequest)
		return
	}

	task.Id = uuid.New().String()


	if task.Status == "" {
		task.Status = New
	}

	if task.Bucket == ""{
		task.Bucket="default"
	}

	pendingTasks =  append(pendingTasks,task)

	fmt.Fprintf(w, "Task Added sucessfully\n")

	fmt.Fprintf(w, "%v", pendingTasks)

}

func updateTodp(w http.ResponseWriter, req *http.Request){

	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Unable to Read Request Body", http.StatusBadRequest)
		return
	}

	var updateTask Todo

	if err := json.Unmarshal(body, &updateTask); err != nil{
		http.Error(w, "Invalid Request Format", http.StatusBadRequest)
		return
	}

	for i,task  :=  range pendingTasks {
		if task.Id == updateTask.Id{
			pendingTasks[i] = updateTask
			fmt.Fprintf(w, "task updated successfully, %v",pendingTasks)
			return
		}
	}

	fmt.Fprintf(w,"unable to find the task requested %v",pendingTasks)



}


func deleteTask(w http.ResponseWriter, req *http.Request){

	body,err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Unable to read the body request", http.StatusBadRequest)
		return
	}

	var deleteRquest struct {
		Id string `json:"id"`
	}

	if err := json.Unmarshal(body, &deleteRquest); err != nil{
		http.Error(w, "Inavlid request format in the body", http.StatusBadRequest)
		return
	}

	id :=  deleteRquest.Id

	for ind,task := range pendingTasks {
		if task.Id == id {
			pendingTasks = append(pendingTasks[:ind], pendingTasks[ind+1:]...)
			fmt.Fprintf(w,"task deleted successfully,  %v", pendingTasks)
			return
		}
	}
	fmt.Fprintf(w,"unable to find the task requested %v",pendingTasks)


}

func main() {
	http.HandleFunc("/test",test)
	http.HandleFunc("/check",check)

	http.HandleFunc("/addTodo",addTodo)

	http.HandleFunc("/updateTodp",updateTodp)
	http.HandleFunc("/deleteTask",deleteTask)

	port := 1234
	fmt.Printf("Server running On port %d",port)
	http.ListenAndServe(fmt.Sprintf(":%d",port),nil)
}