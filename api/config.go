package api

type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`
	MongoDb struct {
		Url      string `json:"url"`
		Database string `json:"database"`
	} `json:"mongoDb"`
	Context struct {
		Timeout int `json:"timeout"`
	} `json:"context"`
}
