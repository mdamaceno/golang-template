package view

import (
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"

    "gopkg.in/yaml.v3"
    "github.com/gin-contrib/multitemplate"
)

var viewConfigFile = filepath.Join("config", "views.yml")
var defaultLayout = "application"

func SetViews() multitemplate.Renderer {
    r := multitemplate.NewRenderer()

    vcfg := make(map[string]interface{})

    data, err := ioutil.ReadFile(viewConfigFile)

    if err != nil {
        log.Fatalf("Failed to read YAML file: %v", err)
    }

    err = yaml.Unmarshal(data, &vcfg)

    if err != nil {
        log.Fatalf("Failed to unmarshal YAML: %v", err)
    }

    for m, values := range vcfg {
        for _, v := range values.([]interface{}) {
            s := strings.Split(v.(string), ":")

            if len(s) < 2 {
                r.AddFromFiles(s[0], filepath.Join(m, "view", s[0] + ".html"))
            } else {
                l := s[1]

                if l == "" {
                    l = defaultLayout
                }

                r.AddFromFiles(
                    s[0],
                    filepath.Join(m, "view", "layouts", l + ".html"),
                    filepath.Join(m, "view", s[0] + ".html"),
                )
            }
        }
    }

    return r
}
