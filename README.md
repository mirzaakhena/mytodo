# MyTodo

This Application created with gogen cli tools with several command : 

```text

gogen domain core

gogen entity Todo

gogen error MessageMustNotEmpty
gogen error TodoAlreadyChecked

gogen usecase RunTodoCreate
gogen usecase RunTodoCheck
gogen usecase GetAllTodo

gogen repository SaveTodo Todo RunTodoCreate
gogen repository FindOneTodo Todo RunTodoCheck
gogen repository FindAllTodo Todo GetAllTodo

gogen gateway withmongodb
gogen gateway withsqlitedb

gogen controller restapi

gogen application mytodo
```

## How to run it

1. copy file `config.sample.json`
2. paste it into new file with name `config.json`
3. open first terminal run `go mod tidy` to download the go dependency 
4. still in first terminal, run the backend apps with command `go run main.go mytodo`
5. open second terminal then change directory to web folder `cd web/` 
6. in second terminal run `npm install` to download the vue dependency
7. still in second terminal run the frontend apps with command `npm run dev`
8. open `http://127.0.0.1:5173/web/` in browser

This application by default run with database sqlite. 
if you want to switch into mongodb you need to switch in `application/app_mytodo.go`

replace line
```go
datasource := withsqlitedb.NewGateway(log, appData, cfg)
```
with

```go
datasource := withmongodb.NewGateway(log, appData, cfg)
```