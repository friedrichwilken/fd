package path

import (
	"fmt"
	"log"
	"strings"
)

type Path struct {
	Segments  []string
	seperator string
}

func New(path, sep string) Path {
	p := strings.Split(path, sep)

	// remove empty segments
	segs := []string{}
	for _, s := range p {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			segs = append(segs, s)
		}
	}

	return Path{
		Segments:  segs,
		seperator: sep,
	}
}

func (p *Path) String() string {
	str := ""
	for _, s := range p.Segments {
		str = fmt.Sprintf("%s%s%s", str, p.seperator, s)
	}

	return str
}

func (p *Path) GoToParent() error {
	if !(len(p.Segments) > 0) {
		return fmt.Errorf("no parent folder")
	}

    log.Printf("before: %v", len(p.Segments))
    p.Segments = p.Segments[:len(p.Segments)-1]
    log.Printf("after: %v", len(p.Segments))

	return nil
}

func (p *Path) GoToSub(s string) {
	p.Segments = append(p.Segments, s)
}

func (p Path) Len() int {
	return len(p.Segments)
}
