package spoty

import (
	"github.com/MrMohebi/tel-link-to-file/common"
	"os/exec"
	"strings"
)

func DownloadAndSave(link string) {

	cmd := exec.Command("spotdl", "download", link, "--output", "\"./music/{artist} - {title}.{output-ext}\"")
	output, err := cmd.Output()
	common.IsErr(err)
	println(strings.TrimSuffix(string(output), "\n"))

}
