# STANDUP-MANAGEMENT-TOOL
The standup management api is used for recording and tracking daily updates from employees in a company.

### Setup Locally
- Use the .env_sample to set up your .env file
- Start up docker on your system
- Run `make docker-up` on your terminal to build app image and its dependencies
- Run `make run-server` to start up app.

## API ENDPOINTS DOCUMENTATION

### Company Admin Sign Up (POST)
The admin for the company registers on the management tool

Endpoint
```
localhost:{SERVER_PORT}/signup
```
Payload
```
{
    "first_name": "Luka",
    "last_name": "Modric",
    "email": "lukamodric@email.com",
    "password": "password",
    "confirm_password" :  "password",
    "company": "Real Madrid Football Club Limited"
}
```

### Login (POST)
Registered user or users created by the admin can log in and receive a token for subsequent request into the system.

Endpoint
```
localhost:{SERVER_PORT}/login
```
Payload
```
{
    "email": "lukamodric@email.com",
    "password": "password"
}
```

### Create Employee (POST)
An admin can create an employee/ add an employee to the system.

Endpoint
```
localhost:{SERVER_PORT}/employees
```
Payload
```
{
    "first_name": "Guler",
    "last_name": "Arda",
    "email": "ardaguler@email.com"
}
```

### Create A Sprint (POST)
A logged in user can create a sprint for a project with a unique project name.

Endpoint
```
localhost:{SERVER_PORT}/sprints
```
Payload
```
{
    "name":  "Sprint 2",
    "project_name": "Project 2",
    "start_date_time": "2024-05-30T15:00:14.225584+01:00",
    "duration": 2,// in weeks
    "daily_update_start_time": "10:00AM"
}
```

### Get All Sprints (GET)
A logged in user can get all sprints.

Endpoint
```
localhost:{SERVER_PORT}/sprints
```

### Create Standup Update (POST)
A logged in user can create a sprint for a project with a unique project name.

Endpoint
```
localhost:{SERVER_PORT}/sprints/:id/updates
```
Url Param
```
id: {unique uuid}
```
Payload
```
{
	"task_id": "Get All Partners to come for a meeting",
	"next_update_todo": "I will get call all partners",
	"previous_update_done": "I was able to get all executive to call for a meeting",
	"blocked_by_id": "fd3ecf4a-8703-4f19-a743-95a9f6df075d",
	"break_away": "break away"
}
```

### Get All Standup Updates (GET)
A logged in user can get all standup updates.

Endpoint
```
localhost:{SERVER_PORT}/sprints?page={page}&sprint={sprint}&owner={owner}&week_start={week_start}&week_end={week_end}
```
Url Query Param
```
sprint: aee106db-531d-48b0-b1ad-2c4091bcbb27
day: 2024-05-22
owner: 06207852-dfce-45fb-831b-efb44067931a
page: 1
week_start: 2024-05-07
week_end: 2024-05-22
```