package mirrorfile

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type MirrorEntry string

func parseContent(reader io.Reader) ([]MirrorEntry, error) {
	mirrors := []MirrorEntry{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		entry := parseLine(line)
		if entry != nil {
			mirrors = append(mirrors, *entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("parse content: %w", err)
	}

	return mirrors, nil
}

func parseLine(line string) *MirrorEntry {

	split := strings.SplitN(line, "=", 2)
	if len(split) != 2 {
		return nil
	}

	split[0] = strings.TrimSpace(split[0])
	split[1] = strings.TrimSpace(split[1])

	if split[0] == "Server" && split[1] != "" {
		return (*MirrorEntry)(&split[1])
	}

	return nil
}
