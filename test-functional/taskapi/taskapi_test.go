package taskapi_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testfunctional/taskapi"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const baseUrl = "http://localhost:8080/tasks"

func Test_GetTasks(t *testing.T) {
	// GIVEN
	taskResponse := createNewTask(t, readFile("../files/Test_GetTasks/post_tasks_request_body.json"))

	// WHEN
	response, err := http.Get(baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 200, response.StatusCode)
	tasksResponse := parseBody[taskapi.TasksResponse](t, response.Body)
	assert.Contains(t, tasksResponse, taskResponse)
}

func Test_PostTasks(t *testing.T) {
	// GIVEN
	requestBody := readFile("../files/Test_PostTasks/post_tasks_request_body.json")

	// WHEN
	response, err := http.Post(baseUrl, "application/json", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 201, response.StatusCode)
	taskResponse := parseBody[taskapi.TaskResponse](t, response.Body)
	assert.Equal(t, "the newest task ever", taskResponse.Title)
	assert.Equal(t, "with the newest description ever", *taskResponse.Description)
}

func Test_PutTask(t *testing.T) {
	// GIVEN
	client := &http.Client{}
	taskResponse := createNewTask(t, readFile("../files/Test_PutTask/post_tasks_request_body.json"))

	// WHEN
	request, err := http.NewRequest(
		http.MethodPut,
		baseUrl+fmt.Sprintf("/%v", taskResponse.TaskId),
		readFile("../files/Test_PutTask/put_task_request_body.json"),
	)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 204, response.StatusCode)
}

func Test_PutTask_TaskNotFound(t *testing.T) {
	// GIVEN
	client := &http.Client{}
	taskId, err := uuid.Parse("44bd6239-7e3d-4d4a-90a0-7d4676a00f5c")
	if err != nil {
		t.Fatal(err)
	}
	requestBody := readFile("../files/Test_PutTask_TaskNotFound/put_task_request_body.json")

	// WHEN
	request, err := http.NewRequest(http.MethodPut, baseUrl+fmt.Sprintf("/%v", taskId), requestBody)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 404, response.StatusCode)
}

func Test_DeleteTask(t *testing.T) {
	// GIVEN
	client := &http.Client{}
	taskResponse := createNewTask(t, readFile("../files/Test_DeleteTask/post_tasks_request_body.json"))

	// WHEN
	request, err := http.NewRequest(http.MethodDelete, baseUrl+fmt.Sprintf("/%v", taskResponse.TaskId), nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 204, response.StatusCode)
}

func Test_DeleteTask_TaskNotFound(t *testing.T) {
	// GIVEN
	client := &http.Client{}
	taskId, err := uuid.Parse("808843a8-1736-4597-a044-e8b491e61307")
	if err != nil {
		t.Fatal(err)
	}

	// WHEN
	request, err := http.NewRequest(http.MethodDelete, baseUrl+fmt.Sprintf("/%v", taskId), nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 404, response.StatusCode)
}

func Test_GetTask(t *testing.T) {
	// GIVEN
	postTaskResponse := createNewTask(t, readFile("../files/Test_GetTask/post_tasks_request_body.json"))

	// WHEN
	response, err := http.Get(baseUrl + fmt.Sprintf("/%v", postTaskResponse.TaskId))
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 200, response.StatusCode)
	taskResponse := parseBody[taskapi.TaskResponse](t, response.Body)
	assert.Equal(t, postTaskResponse.TaskId, taskResponse.TaskId)
	assert.Equal(t, "the shiniest task ever", taskResponse.Title)
	assert.Equal(t, "with the shiniest description ever", *taskResponse.Description)
}

func Test_GetTask_TaskNotFound(t *testing.T) {
	// GIVEN
	taskId, err := uuid.Parse("9d1ab41f-582f-4aaa-97f0-da975624a2fa")
	if err != nil {
		t.Fatal(err)
	}

	// WHEN
	response, err := http.Get(baseUrl + fmt.Sprintf("/%v", taskId))
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// THEN
	assert.Equal(t, 404, response.StatusCode)
}

func parseBody[T any](t *testing.T, reader io.Reader) T {
	var structType T
	err := json.NewDecoder(reader).Decode(&structType)
	if err != nil {
		t.Fatal(err)
	}
	return structType
}

func createNewTask(t *testing.T, body io.Reader) taskapi.TaskResponse {
	response, err := http.Post(baseUrl, "application/json", body)
	if err != nil {
		t.Fatal(err)
	}
	return parseBody[taskapi.TaskResponse](t, response.Body)
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return file
}
