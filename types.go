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
	Chapters      []Chapter `json:"chapters"`
}

type SearchComic struct {
	Title  string `json:"title"`
	Id     string `json:"id"`
	Handle string `json:"handle"`
	Cover  string `json:"cover"`
}
type Chapter struct {
	Name         string   `json:"name"`
	Number       string   `json:"number"`
	Timestamp    int      `json:"timestamp"`
	Handle       string   `json:"handle"`
	Pages        []string `json:"pages"`
}