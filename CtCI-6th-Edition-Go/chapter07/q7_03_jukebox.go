package chapter07

import "fmt"

type Song struct {
	Title       string
	Artist      string
	DurationSec int
}

func NewSong(title, artist string, durationSec int) *Song {
	if title == "" || artist == "" {
		panic("title and artist required")
	}
	if durationSec < 1 {
		panic("duration must be positive")
	}
	return &Song{Title: title, Artist: artist, DurationSec: durationSec}
}

type Cd struct {
	Name  string
	Songs []*Song
}

func NewCd(name string, songs []*Song) *Cd {
	if name == "" {
		panic("name required")
	}
	if len(songs) == 0 {
		panic("CD must have at least one song")
	}
	return &Cd{Name: name, Songs: append([]*Song{}, songs...)}
}

type Jukebox struct {
	cds         []*Cd
	currentCd   *int
	currentSong *int
}

func NewJukebox() *Jukebox { return &Jukebox{} }

func (j *Jukebox) AddCd(cd *Cd) {
	if cd == nil {
		panic("cd is nil")
	}
	j.cds = append(j.cds, cd)
}

func (j *Jukebox) SelectCd(index int) {
	if index < 0 || index >= len(j.cds) {
		panic("invalid CD index")
	}
	j.currentCd = &index
	zero := 0
	j.currentSong = &zero
}

func (j *Jukebox) SelectSong(index int) *Song {
	if j.currentCd == nil {
		panic("no CD selected")
	}
	cd := j.cds[*j.currentCd]
	if index < 0 || index >= len(cd.Songs) {
		panic("invalid song index")
	}
	j.currentSong = &index
	return cd.Songs[index]
}

func (j *Jukebox) Play() *Song {
	if j.currentCd == nil || j.currentSong == nil {
		panic("nothing selected")
	}
	return j.cds[*j.currentCd].Songs[*j.currentSong]
}

func (j *Jukebox) NextSong() *Song {
	if j.currentCd == nil || j.currentSong == nil {
		return nil
	}
	cd := j.cds[*j.currentCd]
	if *j.currentSong+1 >= len(cd.Songs) {
		return nil
	}
	next := *j.currentSong + 1
	j.currentSong = &next
	return cd.Songs[next]
}

func RunQ703() {
	box := NewJukebox()
	box.AddCd(NewCd("Greatest", []*Song{NewSong("A", "Band", 180), NewSong("B", "Band", 200)}))
	box.SelectCd(0)
	fmt.Println(box.Play().Title)
	fmt.Println(box.NextSong().Title)
}
