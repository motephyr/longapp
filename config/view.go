package config

import (
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"net/http"

	"github.com/gofiber/template/html"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TemplateConfig struct {
	TemplateEngine *html.Engine
	Path           string `yaml:"path" env-default:"resources/views"`
	Extension      string `yaml:"extension" env-default:".html"`
}
type ViewConfig struct {
	Template TemplateConfig `yaml:"template"`
}

var EmbedDirStaticView embed.FS

func (v *ViewConfig) Load(path string) {
	// viewPath := MakeDir(filepath.Join(path, v.Template.Path))

	contentFS, _ := fs.Sub(EmbedDirStaticView, "resources/views")

	v.Template.TemplateEngine = html.NewFileSystem(http.FS(contentFS), v.Template.Extension)

	is_dev := flag.Bool("dev", false, "development mode")
	flag.Parse()

	Manifest.Init(*is_dev, path+"/public/mix-manifest.json")

	v.Template.TemplateEngine.AddFunc("asset", Manifest.AssetHelper)
	v.Template.TemplateEngine.AddFunc("gjsonGet", gjson.Get)
	v.Template.TemplateEngine.AddFunc("sjsonSet", sjson.Set)
	v.Template.TemplateEngine.AddFunc("unmarshal", func(str string) (map[string]any, error) {
		obj := map[string]any{}
		json.Unmarshal([]byte(str), &obj)
		return obj, nil
	})

}
