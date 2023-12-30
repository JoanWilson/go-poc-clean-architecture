# API Endpoints

## Create a Task 

### Method
POST

### Syntax
localhost:8080/task

### Description
Creates a new task with an automatically generated UUID.

### Parameters

| Name           | Type  | Required | Description                     |
|-----------------|--------|----------|-----------------------------------|
| Title          | string | Yes     | The title of the task.          |
| Description    | string | No      | A description of the task.      |
| isCompleted    | boolean| No      | Indicates whether the task is completed (defaults to false). |



## Get All Tasks

### Method
GET

### Syntax
localhost:8080/task

### Description
Retrieves a list of all existing tasks.

### Parameters
None

### Examples

#### Request


