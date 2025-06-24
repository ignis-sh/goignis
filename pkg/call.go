package pkg

import (
	"context"
	"fmt"

	dbus "github.com/godbus/dbus/v5"
)

type (
	Args = []any
)

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
