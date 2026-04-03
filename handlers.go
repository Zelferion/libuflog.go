// Copyright (c) 2026 Serhii Yeriemieiev
// Licensed under the MIT License. See LICENSE file in the project root.
package libuflog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/zelferion/libuflog.go/formatting"
)

var fileOnce sync.Once

func FileLogging(l *Logger, msg Message) {
	const logDir = "logs"
	const latestPath = logDir + "/latest.log"

	fileOnce.Do(func() {
		os.MkdirAll(logDir, 0o755)

		if f, err := os.Open(latestPath); err == nil {
			scanner := bufio.NewScanner(f)
			if scanner.Scan() {
				var header struct {
					Created string `json:"created"`
				}
				if json.Unmarshal([]byte(scanner.Text()), &header) == nil && header.Created != "" {
					f.Close()
					os.Rename(latestPath, logDir+"/"+header.Created+".log")
					return
				}
			}
			f.Close()
		}
	})

	file, err := os.OpenFile(latestPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return
	}
	defer file.Close()

	info, _ := file.Stat()
	if info.Size() == 0 {
		header, _ := json.Marshal(map[string]string{"created": time.Now().Format("2006-01-02T15-04-05")})
		fmt.Fprintln(file, string(header))
	}

	entry, err := json.Marshal(map[string]string{
		"time":    time.Now().Format(time.RFC3339),
		"type":    msg.GetMessageType(),
		"message": msg.GetRawMessage(),
	})
	if err != nil {
		return
	}
	fmt.Fprintln(file, string(entry))
}

func ColorfulLogging(l *Logger, msg Message) {
	t := formatting.EqualPadding(l.Formatting.FormatType(msg.GetMessageType(), msg.GetTypeStyle()...))
	msg.SetFormattedMessage(l.Formatting.FormatMessage(msg.GetFormattedMessage()))
	message := msg.GetFormattedMessage()
	var caller string
	if l.Caller {
		caller = l.Formatting.FormatCaller(msg.GetCaller(), formatting.Italic, formatting.Cyan)
	}
	time := l.Formatting.FormatTime(time.Now().Format("15:04:05"), formatting.Gray)
	if !l.Caller {
		fmt.Printf("%s %s%s\n", time, t, message)
	} else {
		fmt.Printf("%s %s %s %s\n", time, t, caller, message)
	}
}

func JSONLogging(l *Logger, msg Message) {
	entry := map[string]string{
		"time":    time.Now().Format(time.RFC3339),
		"type":    msg.GetMessageType(),
		"message": msg.GetRawMessage(),
	}

	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSONLogging: failed to marshal: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
