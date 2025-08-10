package pkg

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	dbus "github.com/godbus/dbus/v5"
)

const (
	IgnisDest       string          = "com.github.linkfrg.ignis"
	IgnisObjectPath dbus.ObjectPath = "/com/github/linkfrg/ignis"
	IgnisInterface  string          = "com.github.linkfrg.ignis"
)

// DBusCallIgnis calls methods provided by interface "com.github.linkfrg.ignis".
//
// For methods from other interfaces or objects, use [DBusCall] instead.
func DBusCallIgnis(ctx context.Context, methodName string, args Args, retvalues ...any) (err error) {
	err = DBusCall(ctx, IgnisDest, IgnisObjectPath, IgnisInterface, methodName, args, retvalues...)
	return
}

// IgnisSystemInfo simply executes command "ignis systeminfo".
func IgnisSystemInfo(ctx context.Context) (err error) {
	cmd := exec.CommandContext(ctx, "ignis", "systeminfo")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run command: %w", err)
		return
	}
	return
}

// InitIgnis executes command "ignis init".
//
// If daemon is true, the child process is spawned and leaked, without stdin/stdout/stderr piped.
func InitIgnis(ctx context.Context, configPath string, daemon bool) (pid int, err error) {
	cmd := exec.CommandContext(ctx, "ignis", "init")
	if len(configPath) != 0 {
		cmd.Args = append(cmd.Args, "-c", configPath)
	}
	if daemon {
		err = cmd.Start()
		pid = cmd.Process.Pid
	} else {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}
	if err != nil {
		err = fmt.Errorf("failed to run command: %w", err)
		return
	}
	return
}

func QuitIgnis(ctx context.Context) (err error) {
	err = DBusCallIgnis(ctx, "Quit", nil)
	return
}

func ReloadIgnis(ctx context.Context) (err error) {
	err = DBusCallIgnis(ctx, "Reload", nil)
	return
}

func OpenInspector(ctx context.Context) (err error) {
	err = DBusCallIgnis(ctx, "Inspector", nil)
	return
}

func ListWindows(ctx context.Context) (windows []string, err error) {
	err = DBusCallIgnis(ctx, "ListWindows", nil, &windows)
	return
}

func ToggleWindow(ctx context.Context, windowName string) (found bool, err error) {
	err = DBusCallIgnis(ctx, "ToggleWindow", Args{windowName}, &found)
	return
}

func OpenWindow(ctx context.Context, windowName string) (found bool, err error) {
	err = DBusCallIgnis(ctx, "OpenWindow", Args{windowName}, &found)
	return
}

func CloseWindow(ctx context.Context, windowName string) (found bool, err error) {
	err = DBusCallIgnis(ctx, "CloseWindow", Args{windowName}, &found)
	return
}

func ListCommands(ctx context.Context) (commands []string, err error) {
	err = DBusCallIgnis(ctx, "ListCommands", nil, &commands)
	return
}

func RunCommand(ctx context.Context, commandName string, commandArgs []string) (found bool, output string, err error) {
	err = DBusCallIgnis(ctx, "RunCommand", Args{commandName, commandArgs}, &found, &output)
	return
}
