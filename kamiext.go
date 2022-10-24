package kamiext

import (
	"encoding/json"
	"os"

	"github.com/FrancescoMerenda/kamitypes"
)

type command struct {
	Executable string
	Command    string
	Args       []string
}

func Command() (*command, int8) {
	new_cmd := &command{}
	new_cmd.Executable = os.Args[0]
	if len(os.Args) < 2 {
		return new_cmd, 1
	}
	command_name := os.Args[1]
	new_cmd.Command = command_name[2:]
	if len(os.Args) > 2 {
		new_cmd.Args = os.Args[2:]
	}
	return new_cmd, 0
}

func Result(data any) kamitypes.Response {
	return kamitypes.Response{Data: data}
}

func Error(err error) kamitypes.Response {
	return kamitypes.Response{Error: err.Error()}
}

func ErrorStr(err string) kamitypes.Response {
	return kamitypes.Response{Error: err}
}

func Write(data kamitypes.Response) int8 {
	bytes, _ := json.Marshal(data)
	_, err := os.Stdout.Write(bytes)
	if err != nil {
		return 1
	}
	return 0
}

func (cmd *command) ArgMethod(f func(string) kamitypes.Response) kamitypes.Response {
	if len(cmd.Args) < 1 {
		return ErrorStr("Not enough arguments")
	}
	return f(cmd.Args[0])
}

func (cmd *command) Method(f func() kamitypes.Response) kamitypes.Response {
	return f()
}

// var input_data input type
// kamiext.JsonMethod(MethodName, cmd.args, input_data)
func (cmd *command) JsonMethod(f func(any) kamitypes.Response, input_data any) kamitypes.Response {
	if len(cmd.Args) < 1 {
		return ErrorStr("Not enough arguments")
	}
	json.Unmarshal([]byte(cmd.Args[0]), &input_data)
	return f(input_data)
}
