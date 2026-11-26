package mongodb

import (
	"testing"

	"github.com/stretchr/testify/require"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/server/common/persistence/visibility/store/query"
	"go.temporal.io/server/common/searchattribute/sadefs"
)

func TestMongoFieldName_RootRunID(t *testing.T) {
	field, err := mongoFieldName(sadefs.RootRunID)
	require.NoError(t, err)
	require.Equal(t, "root_run_id", field)
}

func TestIsOpenExecutionStatus(t *testing.T) {
	require.True(t, isOpenExecutionStatus(enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING))
	require.True(t, isOpenExecutionStatus(enumspb.WORKFLOW_EXECUTION_STATUS_PAUSED))
	require.False(t, isOpenExecutionStatus(enumspb.WORKFLOW_EXECUTION_STATUS_COMPLETED))
	require.False(t, isOpenExecutionStatus(enumspb.WORKFLOW_EXECUTION_STATUS_TERMINATED))
}

func TestStripTransientSearchAttributes_KeepsPauseInfoWhilePaused(t *testing.T) {
	s := &visibilityStore{}
	doc := &visibilityExecutionDocument{
		Status:              int32(enumspb.WORKFLOW_EXECUTION_STATUS_PAUSED),
		SearchAttributesMap: map[string]any{sadefs.TemporalPauseInfo: []string{"x"}},
	}
	removed, err := s.stripTransientSearchAttributes(doc)
	require.NoError(t, err)
	require.False(t, removed)
	require.Contains(t, doc.SearchAttributesMap, sadefs.TemporalPauseInfo)
}

func TestEncodeGroupValuePayload_MissingKeywordIsEmptyString(t *testing.T) {
	col := query.NewSAColumn(sadefs.TemporalNamespaceDivision, sadefs.TemporalNamespaceDivision, enumspb.INDEXED_VALUE_TYPE_KEYWORD)
	payload, err := encodeGroupValuePayload(col, nil)
	require.NoError(t, err)
	decoded, err := sadefs.DecodeValue(payload, enumspb.INDEXED_VALUE_TYPE_KEYWORD, false)
	require.NoError(t, err)
	require.Equal(t, "", decoded)
}
