package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
											 {"Name" : "Arjesh", "Wins" : 10},
											{"Name" : "RJS", "Wins" : 10}
]`)
		store := FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []Player{
			{"Arjesh", 10},
			{"RJS", 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
											 {"Name" : "Arjesh", "Wins" : 10},
											{"Name" : "RJS", "Wins" : 30}
]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("RJS")
		want := 30
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
