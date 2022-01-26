package config

import (
	"encoding/json"
	"flag"
	"path/filepath"

	"github.com/gofiber/template/html"
	"github.com/markbates/pkger"
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

func (v *ViewConfig) Load(path string) {
	path = MakeDir(filepath.Join(path, v.Template.Path))
	v.Template.TemplateEngine = html.NewFileSystem(pkger.Dir(path), v.Template.Extension)

	is_dev := flag.Bool("dev", false, "development mode")
	flag.Parse()

	Manifest.Init(*is_dev, "./public/mix-manifest.json")
	v.Template.TemplateEngine.AddFunc("asset", Manifest.AssetHelper)
	v.Template.TemplateEngine.AddFunc("gjsonGet", gjson.Get)
	v.Template.TemplateEngine.AddFunc("sjsonSet", sjson.Set)
	v.Template.TemplateEngine.AddFunc("unmarshal", func(str string) (map[string]any, error) {
		obj := map[string]any{}
		json.Unmarshal([]byte(str), &obj)
		return obj, nil
	})

}
