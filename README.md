
### REST API with the following functionalities:

- **View all tasks**: Retrieve a list of all tasks.
- **Add a new task**: Create a new task and add it to the list.
- **Update an existing task**: Modify the details of an existing task.
- **Delete a task**: Remove a task from the list.

Initialize a Go module in the project directory 

     go mod init todo 

and install Gorilla Mux (a powerful routing package in Go) by

    go get -u github.com/gorilla/mux


### API Endpoints to Use:

- GET /tasks → List all tasks
- POST /tasks → Add a new task
- PUT /tasks → Update an existing task
- DELETE /tasks/{id} → Delete a task