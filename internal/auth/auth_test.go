package auth

import (
	"net/http"
	"testing"
)



func TestAuth(t *testing.T){
	testHeader1 := http.Header{
		"Authorization": []string{"ApiKey TestKey"},
	}
	testHeader2 := http.Header{
		"Authorization": []string{"ApiKey TestKey2"},
	}
	cases := []struct {
		testHeader http.Header
		expectingKey string
		

	}{
		{
			testHeader: testHeader1,
			expectingKey: "TestKey",
			
		},
		{
			testHeader: testHeader2,
			expectingKey: "TestKey2",
			
		},
	}

	for _, c := range cases {
		apiKey, err := GetAPIKey(c.testHeader)
		if err != nil {
			t.Errorf("%s: Err when expecting %s", err, c.expectingKey)
			return
		}
		if apiKey != c.expectingKey {
			t.Errorf("%s: was found but expecting:%s", apiKey, c.expectingKey)
			return
		}
	}
}