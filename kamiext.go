package kamiext

import (
	"encoding/json"
	"os"
)

type command struct {
	Executable string
	Command    string
	Args       []string
}

func Command() (*command) {
	new_cmd := &command{}
	new_cmd.Executable = os.Args[0]
	if len(os.Args) < 2 {
		return new_cmd
	}
	command_name := os.Args[1]
	new_cmd.Command = command_name[2:]
	if len(os.Args) > 2 {
		new_cmd.Args = os.Args[2:]
	}
	return new_cmd
}

func Result(data any) Response {
	return Response{Data: data}
}

func Error(err error) Response {
	return Response{Error: err.Error()}
}

func ErrorStr(err string) Response {
	return Response{Error: err}
}

func Write(data Response) int8 {
	bytes, _ := json.Marshal(data)
	_, err := os.Stdout.Write(bytes)
	if err != nil {
		return 1
	}
	return 0
}

func (cmd *command) ArgMethod(f func(string) Response) Response {
	if len(cmd.Args) < 1 {
		return ErrorStr("Not enough arguments")
	}
	return f(cmd.Args[0])
}

func (cmd *command) Method(f func() Response) Response {
	return f()
}

func (cmd *command) NoMethod() Response {
	return ErrorStr("No Command Given")
}

func (cmd *command) MethodNotFound() Response {
	return ErrorStr("Method Not Found")
}

// var input_data input_type
// kamiext.JsonMethod(MethodName, cmd.args, input_data)
func (cmd *command) JsonMethod(f func(any) Response, input_data any) Response {
	if len(cmd.Args) < 1 {
		return ErrorStr("Not enough arguments")
	}
	json.Unmarshal([]byte(cmd.Args[0]), &input_data)
	return f(input_data)
}
