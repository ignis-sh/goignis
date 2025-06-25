package pkg

import (
	"context"
	"fmt"

	dbus "github.com/godbus/dbus/v5"
)

type (
	Args = []any
)

// DBusCall connects to the session bus and call a method specified by the destination, object path, interface and method name.
//
//   - param "args" is a slice of arguments, or nil if no arguments needed
//   - param "retvalues" is a series of pointers to return values, or left empty if no return values expected
func DBusCall(ctx context.Context, dest string, objectPath dbus.ObjectPath, interfaceName, methodName string, args Args, retvalues ...any) (err error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		err = fmt.Errorf("failed to connect to session bus: %w", err)
		return
	}
	defer conn.Close()

	method := interfaceName + "." + methodName
	err = conn.Object(dest, objectPath).CallWithContext(ctx, method, 0, args...).Store(retvalues...)
	if err != nil {
		err = fmt.Errorf("failed to call method `%s`: %w", method, err)
		return
	}

	return
}
