package facade

import (
	"fmt"
	"io"
)

const (
	defaultMovieVolume = 5
	defaultDimLevel    = 10
)

type PopcornPopper struct {
	out io.Writer
}

func NewPopcornPopper(out io.Writer) *PopcornPopper {
	return &PopcornPopper{out: out}
}

func (p *PopcornPopper) On() {
	fmt.Fprintln(p.out, "ポップコーンメーカーの電源を入れます")
}

func (p *PopcornPopper) Pop() {
	fmt.Fprintln(p.out, "ポップコーンを作ります！")
}

func (p *PopcornPopper) Off() {
	fmt.Fprintln(p.out, "ポップコーンメーカーの電源を切ります")
}

type TheaterLights struct {
	out io.Writer
}

func NewTheaterLights(out io.Writer) *TheaterLights {
	return &TheaterLights{out: out}
}

func (l *TheaterLights) Dim(level int) {
	fmt.Fprintf(l.out, "シアターの照明を%d%%に暗くします\n", level)
}

func (l *TheaterLights) On() {
	fmt.Fprintln(l.out, "シアターの照明を明るくします")
}

type Screen struct {
	out io.Writer
}

func NewScreen(out io.Writer) *Screen {
	return &Screen{out: out}
}

func (s *Screen) Down() {
	fmt.Fprintln(s.out, "シアタースクリーンを下ろします")
}

func (s *Screen) Up() {
	fmt.Fprintln(s.out, "シアタースクリーンを上げます")
}

type Projector struct {
	out io.Writer
}

func NewProjector(out io.Writer) *Projector {
	return &Projector{out: out}
}

func (p *Projector) On() {
	fmt.Fprintln(p.out, "プロジェクターの電源を入れます")
}

func (p *Projector) WideScreenMode() {
	fmt.Fprintln(p.out, "プロジェクターをワイドスクリーンモード(16:9)に設定します")
}

func (p *Projector) Off() {
	fmt.Fprintln(p.out, "プロジェクターの電源を切ります")
}

type StreamingPlayer struct {
	out io.Writer
}

func NewStreamingPlayer(out io.Writer) *StreamingPlayer {
	return &StreamingPlayer{out: out}
}

func (sp *StreamingPlayer) On() {
	fmt.Fprintln(sp.out, "ストリーミングプレーヤーの電源を入れます")
}

func (sp *StreamingPlayer) Play(movie string) {
	fmt.Fprintf(sp.out, "ストリーミングプレーヤーで「%s」を再生します\n", movie)
}

func (sp *StreamingPlayer) Stop() {
	fmt.Fprintln(sp.out, "ストリーミングプレーヤーを停止します")
}

func (sp *StreamingPlayer) Off() {
	fmt.Fprintln(sp.out, "ストリーミングプレーヤーの電源を切ります")
}

type Amplifier struct {
	out io.Writer
}

func NewAmplifier(out io.Writer) *Amplifier {
	return &Amplifier{out: out}
}

func (a *Amplifier) On() {
	fmt.Fprintln(a.out, "アンプの電源を入れます")
}

func (a *Amplifier) SetStreamingPlayer(_ *StreamingPlayer) {
	fmt.Fprintln(a.out, "アンプの入力をストリーミングプレーヤーに設定します")
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Fprintln(a.out, "アンプをサラウンド音声に設定します")
}

func (a *Amplifier) SetVolume(level int) {
	fmt.Fprintf(a.out, "アンプの音量を%dに設定します\n", level)
}

func (a *Amplifier) Off() {
	fmt.Fprintln(a.out, "アンプの電源を切ります")
}

type HomeTheaterFacade struct {
	amp       *Amplifier
	player    *StreamingPlayer
	projector *Projector
	lights    *TheaterLights
	screen    *Screen
	popper    *PopcornPopper
}

func NewHomeTheaterFacade(
	amp *Amplifier,
	player *StreamingPlayer,
	projector *Projector,
	lights *TheaterLights,
	screen *Screen,
	popper *PopcornPopper,
) *HomeTheaterFacade {
	if amp == nil || player == nil || projector == nil || lights == nil || screen == nil || popper == nil {
		panic("home theater facade requires all subsystem components")
	}

	return &HomeTheaterFacade{
		amp:       amp,
		player:    player,
		projector: projector,
		lights:    lights,
		screen:    screen,
		popper:    popper,
	}
}

func (f *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Fprintln(f.popper.out, "映画を見る準備をしています...")
	f.popper.On()
	f.popper.Pop()
	f.lights.Dim(defaultDimLevel)
	f.screen.Down()
	f.projector.On()
	f.projector.WideScreenMode()
	f.amp.On()
	f.amp.SetStreamingPlayer(f.player)
	f.amp.SetSurroundSound()
	f.amp.SetVolume(defaultMovieVolume)
	f.player.On()
	f.player.Play(movie)
}

func (f *HomeTheaterFacade) EndMovie() {
	fmt.Fprintln(f.popper.out, "\n映画館をシャットダウンしています...")
	f.popper.Off()
	f.lights.On()
	f.screen.Up()
	f.projector.Off()
	f.amp.Off()
	f.player.Stop()
	f.player.Off()
}
