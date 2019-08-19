package tags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/front/location"
)

type MinecraftTagFormat struct {
	Values []string `json:"values"`
}

func WriteTags(tags map[string]LocationList, output evaluator.OutputDirectory) {
	for tagInformal, locations := range tags {
		tag := location.TagLocationFromInformal(tagInformal)
		TranslateTag(tag, locations, output)
	}
}

func TranslateTag(tag *location.TagLocation, locations LocationList, output evaluator.OutputDirectory) {
	tagFile := fmt.Sprintf("%s.json", tag.Identity)
	parentPath := filepath.Join(output, tag.Namespace, "tags/functions/")
	filePath := filepath.Join(parentPath, tagFile)
	os.MkdirAll(parentPath, os.ModePerm)
	data, _ := json.Marshal(MinecraftTagFormat{locations})
	ioutil.WriteFile(filePath, data, 0666) // TODO I have no idea what 0666 is. Cite it from the stackoverflow post and try to figure it out
}
