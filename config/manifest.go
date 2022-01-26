package config

import (
	"encoding/json"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type (
	AssetList map[string][]string
	Config    struct {
		// DevHost webpack-dev-server host:port
		// DevHost string
		// FsPath filesystem path to public webpack dir
		FsPath string
		// WebPath http path to public webpack dir
		// WebPath string
		// Plugin webpack plugin to use, can be stats or manifest
		// Plugin string
		// IgnoreMissing ignore assets missing on manifest or fail on them
		IgnoreMissing bool
		// Verbose - show more info
		Verbose bool
		// IsDev - true to use webpack-serve or webpack-dev-server, false to use filesystem and manifest.json
		IsDev bool
	}
	assetResponse map[string]string
)
type manifest struct {
	AssetHelper func(string) (template.HTML, error)
}

var Manifest manifest

// Init Set current environment and preload manifest
func (manifest) Init(dev bool, fsPath string) {
	Manifest.AssetHelper = GetAssetHelper(&Config{
		FsPath:        fsPath,
		IgnoreMissing: true,
		Verbose:       true,
		IsDev:         dev,
	})
}

// Read webpack-manifest-plugin format manifest
func ReadManifest(path string) (AssetList, error) {
	//log.Println("read:", path+"/manifest.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return unmarshalManifest(data)
}

func unmarshalManifest(data []byte) (AssetList, error) {
	response := make(assetResponse, 0)
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, errors.Wrap(err, "go-webpack: Error unmarshaling manifest file")
	}

	assets := make(AssetList, len(response))
	for key, value := range response {
		//log.Println("found asset", key, value)
		if !strings.HasSuffix(value, ".map") {
			assets[key] = []string{value}
		}
	}
	return assets, nil
}

func GetAssetHelper(conf *Config) func(string) (template.HTML, error) {
	preloadedAssets := AssetList{}

	var err error
	if conf.IsDev {
		// Try to preload manifest, so we can show an error if webpack-dev-server is not running
		_, err = ReadManifest(conf.FsPath)
		if err != nil {
			log.Println(err)
		}
		//if err != nil {
		//return ErrorFunction(err)
		//}
	} else {
		preloadedAssets, err = ReadManifest(conf.FsPath)
		// we won't ever re-check assets in this case.  this should be a hard error.
		if err != nil {
			return ErrorFunction(err)
		}
	}

	return createAssetHelper(conf, preloadedAssets)
}

func createAssetHelper(conf *Config, preloadedAssets AssetList) func(string) (template.HTML, error) {
	return func(key string) (template.HTML, error) {
		var err error

		var assets AssetList
		if conf.IsDev {
			assets, err = ReadManifest(conf.FsPath)
			if err != nil {
				return template.HTML(""), err
			}
		} else {
			assets = preloadedAssets
		}

		parts := strings.Split(key, ".")
		kind := parts[len(parts)-1]
		//log.Println("showing assets:", key, parts, kind)
		v, ok := assets[key]
		if !ok {
			return displayAssetError(conf, key, assets)
		}

		buf := []string{}
		for _, s := range v {
			// if strings.HasSuffix(s, "."+kind) {
			buf = append(buf, AssetTag(kind, s))
			// } else {
			// 	log.Println("skip asset", s, ": bad type")
			// }
		}
		return template.HTML(strings.Join(buf, "\n")), nil
	}
}

// ErrorFunction returns a template function that returns a fixed error message
func ErrorFunction(err error) func(string) (template.HTML, error) {
	log.Println("go-webpack: error:", err)
	return func(string) (template.HTML, error) {
		return template.HTML(""), err
	}
}

func displayAssetError(conf *Config, key string, assets AssetList) (template.HTML, error) {
	message := "go-webpack: Asset file '" + key + "' not found in manifest"
	if conf.Verbose {
		log.Printf("%s. Manifest contents:", message)
		for k, a := range assets {
			log.Printf("%s: %s", k, a)
		}
	}
	if conf.IgnoreMissing {
		return template.HTML(""), nil
	}
	return template.HTML(""), errors.New(message)
}

func LinkTag(url string) string {
	return `<link type="text/css" rel="stylesheet" href="` + html.EscapeString(url) + `"></link>`
}

// ScriptTag make js script tag from url
func ScriptTag(url string) string {
	return `<script type="text/javascript" src="` + html.EscapeString(url) + `"></script>`
}

// AssetTag make js or css tag from url
func AssetTag(kind, url string) string {
	var buf string
	if kind == "css" {
		buf = LinkTag(url)
	} else if kind == "js" {
		buf = ScriptTag(url)
	} else {
		log.Println("go-webpack: unsupported asset kind: " + kind)
		buf = ""
	}
	return buf
}
