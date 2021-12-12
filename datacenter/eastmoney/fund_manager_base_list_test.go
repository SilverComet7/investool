package eastmoney

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFundManagerBaseList(t *testing.T) {
	data, err := _em.FundMangerBaseList(_ctx, "", "YIELDSE")
	require.Nil(t, err)
	require.NotEmpty(t, data)
	t.Logf("data:%+v", data)
	_, err = json.Marshal(data)
	require.Nil(t, err)
}
