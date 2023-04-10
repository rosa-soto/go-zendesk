package zendesk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchTicketsSideLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	v := ListOptions{
		Page:      0,
		SortBy:    "",
		SortOrder: "",
	}
	x := IncludeGroups()
	y := StatusFilter(Status("OPEN"), LessThanOrEqualTo)

	_, err = client.SearchTicketsSideLoad("", &v, x, y)
	require.NoError(t, err)
}
