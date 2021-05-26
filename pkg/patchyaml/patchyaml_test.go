package patchyaml

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v1"
)

func TestPatchData(t *testing.T) {
	Convey("test for patch file", t, func() {
		// source file
		var sourceObj = map[interface{}]interface{}{}
		sourceFileBytes, err := ioutil.ReadFile("files/source.yml")
		if err != nil {
			panic(err)
		}
		yaml.Unmarshal(sourceFileBytes, &sourceObj)

		// patch file
		var patchObj = map[interface{}]interface{}{}
		patchFileBytes, err := ioutil.ReadFile("files/patch.yml")
		if err != nil {
			panic(err)
		}
		yaml.Unmarshal(patchFileBytes, &patchObj)

		// patched file
		finnalData := PatchData(patchObj, sourceObj)
		finnalBytes, err := yaml.Marshal(finnalData)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("files/finnal.yml", finnalBytes, os.ModePerm); err != nil {
			panic(err)
		}
	})
}
