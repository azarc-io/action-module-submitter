package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	v1 "github.com/azarc-io/verathread-global/api/v1"
	ga "github.com/sethvargo/go-githubactions"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

const (
	readme = "readme.md"
)

func main() {
	action := ga.New()
	if action.Getenv("GITHUB_REF_TYPE") != "tag" {
		action.Fatalf("module scrape can only be used on push tag")
	}

	workspace := action.Getenv("GITHUB_WORKSPACE")

	request := &v1.AddModuleRequest{
		Module: &v1.Module{
			Config: &v1.ModuleConfig{},
			Repo:   action.Getenv("GITHUB_REPOSITORY"),
			Readme: readFile(action, readme),
		},
	}
	if _, err := fmt.Sscanf(action.Getenv("GITHUB_REF"), "refs/tags/%s", request.Module.Version); err != nil {
		action.Fatalf("getting version: %s", err.Error())
	}
	parseYaml(action, fmt.Sprintf("%s/module.yaml", workspace), &request.Module.Config)

	sparksRoot := fmt.Sprintf("%s/sparks", workspace)
	files, err := ioutil.ReadDir(sparksRoot)
	if err != nil {
		action.Fatalf("listing sparks: %s", err.Error())
	}

	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}

		sparkRoot := fmt.Sprintf("%s/%s", sparksRoot, dir.Name())
		spark := v1.Spark{Readme: readFile(action, fmt.Sprintf("%s/%s", sparkRoot, readme))}
		parseYaml(action, fmt.Sprintf("%s/%s", sparkRoot, "spark.yaml"), &spark.Config)
		loadSchema(action, fmt.Sprintf("%s/%s", sparkRoot, "input_schema.json"), &spark.InputSchema)
		request.Sparks = append(request.Sparks, &spark)
	}

	buf := &bytes.Buffer{}
	if err = json.NewEncoder(buf).Encode(request); err != nil {
		action.Fatalf("could not encode module: %s", err.Error())
	}
	resp, err := http.Post("https://auth-events.cloud.azarc.dev/api/v1/module", "application/json", buf)
	if err != nil {
		action.Fatalf("could not add module: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		action.Fatalf("received %d from add module request", resp.StatusCode)
	}
	action.Infof("scraped and submitted for module [repo]: %s, [version]: %s, [sparks]: %d",
		request.Module.Repo, request.Module.Version, len(request.Sparks))

	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//conn, err := grpc.DialContext(ctx, ":443", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	action.Fatalf("could not dail verathread: %s", err.Error())
	//}
	//fmt.Println("connected")
	//client := v1.NewModuleSupportClient(conn)
	//if _, err = client.AddModule(ctx, &v1.AddModuleRequest{Module: module}); err != nil {
	//	action.Fatalf("could not upload module: %s", err.Error())
	//}
}

func readFile(action *ga.Action, file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		action.Fatalf("reading file: %s", err.Error())
	}
	return data
}

func parseYaml(action *ga.Action, file string, v interface{}) {
	data := readFile(action, file)
	if err := yaml.Unmarshal(data, v); err != nil {
		action.Fatalf("unmarshal file: %s", err.Error())
	}
}

func loadSchema(action *ga.Action, file string, v *[]byte) {
	*v = readFile(action, file)
	if _, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(*v)); err != nil {
		action.Fatalf("loading schema: %s", err.Error())
	}
}
