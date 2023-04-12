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

	options := ListOptions{
		Page:      0,
		SortBy:    "",
		SortOrder: "",
	}
	include := IncludeGroups()
	status := StatusFilter(Status("OPEN"), LessThanOrEqualTo)

	_, err = client.SearchTicketsSideLoad("", &options, include, status)
	require.NoError(t, err)
}
