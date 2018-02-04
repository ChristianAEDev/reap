# reap

*reap* allows to once define a task an than execute it over and over again from the command line.

## Motivation

This tool allows to define *plans*. Each *plan* consists out of one or more *tasks*. Each *task* is a small step to achieve a bigger task.
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

Instead of writing a (Power-) shell script the steps are defined in a *json* file.

Place your *plans.json* in the same folder as the *reap* executable.

A simple *plans.json* could look like the following one setting the password in a config file.

```json
"Plans": [
        {
            "Name": "Replace the string 'PLACEHOLDER_PASSWORD' in the file 'config.ini' with 'MyPa55w0rd' in "
            "Tasks": [
                {
                    "Type": "ReplaceInFileTask",
                    "Preferences": {
                        "FilePath": "/opt/mytool/config.ini",
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
  * `-p [PLAN_ID]` (*PLAN_ID* can be found using the `list` command)
* `list`: Show a list of all available plans and their task as defined in the *plans.json* file
* `help`: Show a list of avilable commands
* `clear`: Clear all output from the console
* `exit`: Quit the program

### Available tasks

The general structure of a task looks as follows. It consists out of a *type*, and (optional) *desciption* and *preferences*. The *preferences* are individuell to every task.

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
    "Args": [
      "Hello",
      "World"
    ]
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

* `Command`: The actiont to perform *start*/*stop*
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

## Build

To build the *exe* in a a way that it requests administrator privileges (embed a *manifest*) use the following commands:

```shell
go generate
go build
```