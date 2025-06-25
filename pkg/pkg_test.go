package pkg

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ExampleDBusCall shows calling methods with arguments and results.
func ExampleDBusCall() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "org.freedesktop.DBus"
	var nameOwner string
	err := DBusCall(ctx, "org.freedesktop.DBus", "/org/freedesktop/DBus", "org.freedesktop.DBus", "GetNameOwner", Args{query}, &nameOwner)
	if err != nil {
		log.Println("Failed", err)
		return
	}

	fmt.Println(nameOwner)
	// Output: org.freedesktop.DBus
}

// ExampleDBusCallIgnis shows calling methods provided by ignis.
func ExampleDBusCallIgnis() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var windowNames []string
	err := DBusCallIgnis(ctx, "ListWindows", Args{}, &windowNames)
	if err != nil {
		log.Println("Failed:", err)
		return
	}

	log.Println("Window names:", windowNames)
	fmt.Println("Done")
	// Output: Done
}

// ExampleListWindows shows calling APIs from this package.
func ExampleListWindows() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	windowNames, err := ListWindows(ctx)
	if err != nil {
		log.Println("Failed:", err)
		return
	}

	log.Println("Window names:", windowNames)
}
