//go:generate npm --prefix ../../web run build && go build -o ../../release/. -ldflags "-s -w" .

package main

import "github.com/MrOldOwl/react_server/internal/app"

func main() {
	app.Run()
}
