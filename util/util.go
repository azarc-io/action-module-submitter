package util

import (
	"encoding/base64"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"os"
)

const maxImageSize = 500000

type Logger interface {
	Fatalf(msg string, args ...any)
}

func LoadImage(log Logger, file string) string {
	fp, err := os.Open(file)
	if err != nil {
		log.Fatalf("opening file: %s", err)
	}

	cfg, tp, err := image.DecodeConfig(fp)
	if err != nil {
		//if errors.Is(err, image.ErrFormat) {
		// TODO add vector support
		//}
		log.Fatalf("error loading %s: %s", file, err)
	}

	if cfg.Width*cfg.Height > maxImageSize {
		log.Fatalf("error loading %s's HxW cannot exceed %d", file, maxImageSize)
	}

	return fmt.Sprintf("data:image/%s;base64, %s", tp, base64.StdEncoding.EncodeToString(ReadFile(log, file)))
}

func LoadDirs(log Logger, path string, cb func(path, name string)) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("listing %s: %s", path, err.Error())
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		cb(fmt.Sprintf("%s/%s", path, dir.Name()), dir.Name())
	}
}

func ReadFile(log Logger, file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("reading file: %s", err.Error())
	}
	return data
}

func ParseYaml(log Logger, file string, v interface{}) {
	data := ReadFile(log, file)
	if err := yaml.Unmarshal(data, v); err != nil {
		log.Fatalf("unmarshal file: %s", err.Error())
	}
}

func LoadSchema(log Logger, file string) []byte {
	v := ReadFile(log, file)
	if _, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(v)); err != nil {
		log.Fatalf("loading schema: %s", err.Error())
	}
	return v
}
