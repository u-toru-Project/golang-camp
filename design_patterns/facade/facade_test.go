package facade

import (
	"bytes"
	"strings"
	"testing"
)

func TestHomeTheaterFacadeWatchMovieAndEndMovie(t *testing.T) {
	var out bytes.Buffer

	homeTheater := NewHomeTheaterFacade(
		NewAmplifier(&out),
		NewStreamingPlayer(&out),
		NewProjector(&out),
		NewTheaterLights(&out),
		NewScreen(&out),
		NewPopcornPopper(&out),
	)

	homeTheater.WatchMovie("レイダース/失われたアーク《聖櫃》")
	homeTheater.EndMovie()

	expectedLines := []string{
		"映画を見る準備をしています...",
		"ポップコーンメーカーの電源を入れます",
		"ポップコーンを作ります！",
		"シアターの照明を10%に暗くします",
		"シアタースクリーンを下ろします",
		"プロジェクターの電源を入れます",
		"プロジェクターをワイドスクリーンモード(16:9)に設定します",
		"アンプの電源を入れます",
		"アンプの入力をストリーミングプレーヤーに設定します",
		"アンプをサラウンド音声に設定します",
		"アンプの音量を5に設定します",
		"ストリーミングプレーヤーの電源を入れます",
		"ストリーミングプレーヤーで「レイダース/失われたアーク《聖櫃》」を再生します",
		"",
		"映画館をシャットダウンしています...",
		"ポップコーンメーカーの電源を切ります",
		"シアターの照明を明るくします",
		"シアタースクリーンを上げます",
		"プロジェクターの電源を切ります",
		"アンプの電源を切ります",
		"ストリーミングプレーヤーを停止します",
		"ストリーミングプレーヤーの電源を切ります",
	}

	gotLines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	if len(gotLines) != len(expectedLines) {
		t.Fatalf("unexpected number of lines: want %d, got %d\noutput:\n%s", len(expectedLines), len(gotLines), out.String())
	}

	for i, expected := range expectedLines {
		if gotLines[i] != expected {
			t.Fatalf("line %d mismatch: want %q, got %q", i+1, expected, gotLines[i])
		}
	}
}

func TestNewHomeTheaterFacadePanicsWhenSubsystemIsMissing(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when subsystem is missing")
		}
	}()

	var out bytes.Buffer
	NewHomeTheaterFacade(
		nil,
		NewStreamingPlayer(&out),
		NewProjector(&out),
		NewTheaterLights(&out),
		NewScreen(&out),
		NewPopcornPopper(&out),
	)
}
