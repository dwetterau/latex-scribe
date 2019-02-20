package recognize

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockClient struct {
	mock.Mock
}

var _ client = mockClient{}

func (m mockClient) post(url string, body string) ([]byte, error) {
	args := m.Called(url)
	return args.Get(0).([]byte), args.Error(1)
}

func TestToLatex(t *testing.T) {
	// Load in the image
	b, err := ioutil.ReadFile("./double_integral_test.jpg")
	require.NoError(t, err)

	mockResp, err := ioutil.ReadFile("./double_integral_test_resp.json")
	require.NoError(t, err)
	c := mockClient{}
	c.On("post", latexEndpoint).Return(mockResp, nil)
	r := recognizerImpl{c: c}
	text, err := r.ToLatex(b)

	require.NoError(t, err)
	require.Equal(t, text, "\\int x d x")
}
