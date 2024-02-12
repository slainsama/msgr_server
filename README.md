# msgr_server
A message integration server.

## about scripts
the script will be added as:

```go
type Script struct {
	Id            string
	Name          string
	Command       string //such as "python3 {scriptName} {secretKey} {taskId} {arg1} {arg2}"
	Status        string
	ParamRequired []string
	DataReturn    []string
}
```

### script params

there are two ways to give args to scripts.

#### Command line parameters
you're expected to give a start command with your script uploaded,such as`python3 {scriptName} {arg1} {arg2}`.

there are two keys to get information from the server.

`{sciptName}` to refer the script file name.

`{argn}` to refer the args that sever given.

`{secretKey}` to get the server secretKey.

`{taskId}` to get the task id which is the unique id to match the proc.

`{server}` to get server address , such as `127.0.0.1:8001`.

in the most times,`{taskId}` and `{secretKey}` is necessary to callback your data to server

#### Http parameters

use `http://{server}/api/script/{secret_key}/{task_id}/params` to get a param per time.

the param is added by user's `/addParams {taskId} {arg1} {arg2} ...`

### return data

use `http://{server}/api/script/{secret_key}/{task_id}/sendData` to return data to user.

a callback request is defined as:

```go
type Callback struct {
	Action string         `json:"action"` //"sendText" or "sendPhoto"
	Msg    string         `json:"msg"`
	File   multipart.File `form:"file"`
}
```

`msg` is expected as message text sent to user.
`file` maybe a photo or video or anything.