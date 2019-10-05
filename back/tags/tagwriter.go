package tags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/falcinspire/scriptblock/back/output"
	"github.com/falcinspire/scriptblock/front/location"
)

// WriteAllTagsToFiles writes all of the tags and locations provided into the output directory.
func WriteAllTagsToFiles(tags map[string]LocationList, output output.OutputDirectory) {
	for tagInformal, locations := range tags {
		tag := location.TagLocationFromInformal(tagInformal)
		WriteTagToFile(tag, locations, output)
	}
}

// minecraftTagFormat is the struct used to json deserialization.
type minecraftTagFormat struct {
	Values []string `json:"values"`
}

// WriteTagToFile writes a single tag and its locations to the output directory.
func WriteTagToFile(tag *location.TagLocation, locations LocationList, output output.OutputDirectory) {
	tagFile := fmt.Sprintf("%s.json", tag.Identity)
	parentPath := filepath.Join(output, tag.Namespace, "tags/functions/")
	filePath := filepath.Join(parentPath, tagFile)
	os.MkdirAll(parentPath, os.ModePerm)
	data, _ := json.Marshal(minecraftTagFormat{locations})
	ioutil.WriteFile(filePath, data, 0666) // TODO No idea what 0666 is. Cite it from the stackoverflow post and try to figure it out
}
