package path_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/friedrichwilken/fd/pkg/path"
)

func Test_ParseSegment(t *testing.T) {
	//arrange
	wantSegment := "thrid"
	wantString := fmt.Sprintf("/first/second/%s/fourth", wantSegment)
	p := path.New(wantString, "/")

	//act
	actualSegnent := p.Segments[2]
	actualString := p.String()

	//assert
	assert.Equal(t, wantSegment, actualSegnent, "segments should stay intact")
	assert.Equal(t, wantString, actualString, "whole string should stay the same")
}

func Test_RemoveEmptySegments(t *testing.T) {
	//arrange
	wantSegment := "only"
	wantLen := 1
	input := fmt.Sprintf("//%s/", wantSegment)
	p := path.New(input, "/")

	//act
	actualSegment := p.Segments[0]
	actualLen := p.Len()

	//assert
	assert.Equal(t, wantLen, actualLen, "should have the right amount of segments")
	assert.Equal(t, wantSegment, actualSegment, "should contain the only segment")
}

func Test_GoToParent(t *testing.T) {
	//arrange
	wantString := "/one/two/three"
	wantLen := 3
	input := fmt.Sprintf("%s/%s", wantString, "last")
	p := path.New(input, "/")

	//act
    p.GoToParent()
	actualString := p.String()
	actualLen := p.Len()

	//assert
	assert.Equal(t, wantString, actualString, "should return the right string")
	assert.Equal(t, wantLen, actualLen, "should have the right amount of segments")
}

func Test_GoToSub(t *testing.T) {
	//arrange
    input := "/one/two/three"
    wantSegment := "last"
    wantString := fmt.Sprintf("%s/%s", input, wantSegment) 
	wantLen := 4
	p := path.New(input, "/")

	//act
    p.GoToSub(wantSegment)
	actualString := p.String()
	actualLen := p.Len()

	//assert
	assert.Equal(t, wantString, actualString, "should return the right string")
	assert.Equal(t, wantLen, actualLen, "should have the right amount of segments")
}
