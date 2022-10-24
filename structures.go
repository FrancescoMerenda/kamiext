package kamiext

type Response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}
type Extension struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	Authors     string   `json:"author"`
	Source      string   `json:"source"`
	Executable  string   `json:"executable"`
	Methods     []string `json:"methods"`
	Logo        string   `json:"logo"`
	UpdatesUrl  string   `json:"updatesurl"`
	Updater     string   `json:"updater"`
	Domain      string   `json:"domain"`
	Version     string   `json:"version"`
	DataVersion string   `json:"dataversion"`
}
type Comic struct {
	Title         string    `json:"title"`
	Id            string    `json:"id"`
	Url           string    `json:"url"`
	SourceId      string    `json:"sourceid"`
	SourceName    string    `json:"sourcename"`
	ComicType     string    `json:"comictype"`
	ReaderMode    string    `json:"readermode"`
	ReleaseStatus string    `json:"releasestatus"`
	Summary       string    `json:"summary"`
	Cover         string    `json:"cover"`
	Authors       []string  `json:"authors"`
	Artists       []string  `json:"artists"`
	Altnames      []string  `json:"altnames"`
	Genres        []string  `json:"genres"`
	Year          string    `json:"year"`
	Trackers      []Tracker `json:"trackers"`
	Chapters      []Chapter `json:"chapters"`
}
type SearchComic struct {
	Title  string `json:"title"`
	Id     string `json:"id"`
	Handle string `json:"handle"`
	Cover  string `json:"cover"`
}
type Tracker struct {
	Tracker  string `json:"tracker"`
	Tracking bool   `json:"tracking"`
	Status   string `json:"status"`
	Id       string `json:"id"`
	Rating   string `json:"rating"`
}
type Chapter struct {
	Name         string   `json:"name"`
	Number               string   `json:"number"` //should "Number" be a float nstead?
	Timestamp    int      `json:"timestamp"`
	Handle       string   `json:"handle"`
	Pages        []string `json:"pages"`
	Local        bool     `json:"local"`
	Bookmark     bool     `json:"bookmark"`
	LastPageRead int      `json:"lastpageread"`
	Read         bool     `json:"read"`
}
