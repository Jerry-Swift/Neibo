package common

import (
	"bufio"
	"github.com/projectdiscovery/stringsutil"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func ReadFile(filePath string) (urls []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Error("Oops,Cannot open the target file...\n")
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines []string
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
			return lines, err
		}
		if strings.TrimSpace(str) == "" {
			continue
		}
		if stringsutil.HasSuffixAny(str, "\r\n", "\n", "\r") {
			str = stringsutil.TrimSuffixAny(str, "\r\n", "\n", "\r")
		}
		//  ******打印日志
		logrus.Debug("url is:", str)
		lines = append(lines, str)
	}
	return lines, err

}
