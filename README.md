# reap

_reap_ allows to once define a task an than execute it over and over again from the command line.

## Motivation

This tool allows to define _plans_. Each _plan_ consists out of one or more _tasks_. Each _task_ is a small step to achieve a bigger task.
I use it to automate the deployment process of a specific program:

* Stop the service
* Unzip the package containing the release
* Delete the old release
* Copy the new release to the correct folder
* Replace some strings in the configuration
* Execute a command to update a database
* Check the database logs for errors
* Start the service

## Usage

Instead of writing a (Power-) shell script the steps are defined in a _json_ file.

Place your _plans.json_ in the same folder as the _reap_ executable.

It is possible to define global variables that will be used to replace otherwise repeatitive values. (See parameter `Variables` in _plans.json_ example.)

A simple _plans.json_ could look like the following one setting the password in a config file.

```json
"Plans": [
        {
            "Name": "Replace the string 'PLACEHOLDER_PASSWORD' in the file 'config.ini' with 'MyPa55w0rd' in ",
            "Variables":{
              "CONFIG_PATH": "/opt/mytool/config.ini"
            },
            "Tasks": [
                {
                    "Type": "ReplaceInFileTask",
                    "Preferences": {
                        "FilePath": "${CONFIG_PATH}",
                        "Replace": "PLACEHOLDER_PASSWORD",
                        "With": "MyPa55w0rd"
                    }
                }
            ]
        }
    ]
}
```

Run the program and enter one of the following commands:

* `execute`: Execute a plan (i.e. `execute -p 1` to execute the first plan)
  * `-p [PLAN_ID]` (_PLAN_ID_ can be found using the `list` command)
* `list`: Show a list of all available plans and their task as defined in the _plans.json_ file
* `help`: Show a list of avilable commands
* `clear`: Clear all output from the console
* `exit`: Quit the program

### Available tasks

The general structure of a task looks as follows. It consists out of a _type_, and (optional) _desciption_ and _preferences_. The _preferences_ are individuell to every task.

```json
{
  "Type": "ConfirmationTask",
  "Description": "Ensure the database update was executed.",
  "Preferences": {
  }
},
```

#### ConfirmationTask

Ask the user to confirm something. This can be used to give instructions to the user to execute a task that can't be executed by this tool.

* `Question`: The question/task the user needs to answer/perform.

```json
{
  "Type": "ConfirmationTask",
  "Description": "Ensure the database update was executed.",
  "Preferences": {
    "Question": "Has the database been updated to the latest version?"
  }
},
```

#### DeleteTask

Delete a file/folder

* `Path` Path of the file/folder to be deleted.

```json
{
  "Type": "DeleteTask",
  "Preferences": {
    "Path": "/home/Temp/folderToDelete"
  }
}
```

#### ExecCommandTask

Execute a custom command as if it would be executed on the command line.

* `Command`: The command to execute.
* `Args`: Arguments to provide to the command. (One or more)

```json
{
  "Type": "ExecCommandTask",
  "Description": "Print 'Hello World'",
  "Preferences": {
    "Command": "echo",
    "Args": ["Hello", "World"]
  }
}
```

#### FindInFileTask

Search and count the occurences of a string in a file.

* `FilePath`: Path to the file to search
* `Query`: String to search for

```json
{
  "Type": "FindInFileTask",
  "Preferences": {
    "FilePath": "/home/Temp/test.txt",
    "Query": "Hello World"
  }
}
```

#### RenameFileTask

Rename a file.

* `FilePath`: Path to the folder of the file
* `OldName`: Current name of the file
* `NewName`: Name the file is renamed to

```json
{
  "Type": "RenameFileTask",
  "Description": "Rename file a to b",
  "Preferences": {
    "FilePath": "/home/temp/test.txt",
    "OldName": "testOldName",
    "NewName": "testNewName"
  }
}
```

#### ReplaceInFileTask

Replace all occurences of a given string in a file.

* `FilePath`: Path to the file to process
* `Replace`: The string to replace
* `With`: The string `Replace` will be replaced with

```json
{
  "Type": "ReplaceInFileTask",
  "Preferences": {
    "FilePath": "/home/Temp/test.txt",
    "Replace": "Hello World",
    "With": "dlorW olleH"
  }
}
```

#### ServiceTask

Start or stop a service. Only supports Windows.

* `Command`: The actiont to perform _start_/_stop_
* `Name`: The name of the service

```json
{
  "Type": "ServiceTask",
  "Description": "Start service Spooler",
  "Preferences": {
    "Command": "start",
    "Name": "Spooler"
  }
}
```

#### UnpackArchiveTask

Extract the content of a zip archive.

* `FilePath`: Path to the archive
* `DestinationPath`: Path to extract the archive to

```json
{
  "Type": "UnpackArchiveTask",
  "Description": "Extract the release package",
  "Preferences": {
    "FilePath": "/home/Downloads/release.zip",
    "DestinationPath": "/home/Downloads"
  }
}
```

#### TemplateTask

Write a template file to a given path. The mechanism of variables can be used to fill values in the template.

* `FilePath`: Path to where the template will be saved to
* `Template`: The content of the template file. It is an array. Each item in the array wil be a line in the file

```json
{
  "Type": "TemplateTask",
  "Description": "Dockerfile to creat an image writing 'Hello World' to stdout.",
  "Preferences": {
    "FilePath": "/home/user/Dockerfile",
    "Template": [
      "# Dockerfile test",
      "# Maintainers: ${MAIL}",
      "FROM ubuntu",
      "",
      "# Print stdout.",
      "CMD echo ${OUTPUT}"
    ]
  }
}
```

## Build

To build the _exe_ in a a way that it requests administrator privileges (embed a _manifest_) use the following commands:

```shell
go generate
go build
```
