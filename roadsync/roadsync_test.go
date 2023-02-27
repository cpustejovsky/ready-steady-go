package roadsync_test

import (
	"github.com/cpustejovsky/ready-steady-go/roadsync"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 5
	got := roadsync.Add(2, 3)
	require.Equal(t, want, got)
}
