//+build ignore

package main

import (
	"os"
	"text/template"
)

const cronTemplate = `
# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of the month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * * command to execute
  0 8 * * 1-5 {{ .GoPath }}/bin/daily-hard-reset
  05 17 * * 1-5 {{ .GoPath }}/bin/daily-hard-reset
`

func main() {
	parsedTemplate, err := template.New("cron").Parse(cronTemplate)
	if err != nil {
		panic(err)
	}
	goPath, ok := os.LookupEnv("GOPATH")
	if !ok {
		panic("Could not find $GOPATH")
	}
	err = parsedTemplate.Execute(os.Stdout, struct {
		GoPath string
	}{
		GoPath: goPath,
	})
	if err != nil {
		panic(err)
	}
}
