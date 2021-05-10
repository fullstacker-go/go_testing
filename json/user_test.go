package json

import "testing"

func BenchmarkFiletoJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiletoJSON("MOCK_DATA.json")
	}
}

func BenchmarkFiletoJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiletoJsonMarshal("MOCK_DATA.json")
	}
}
