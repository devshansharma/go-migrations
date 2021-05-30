package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

func main() {
	nameFlag := flag.String("name", "", "Name of the migration")
	flag.Parse()

	if *nameFlag == "" {
		fmt.Println("Please provide 'name' for migration")
		os.Exit(1)
	}

	*nameFlag = strings.ToLower(*nameFlag)
	*nameFlag = strings.ReplaceAll(*nameFlag, " ", "_")

	data := make(map[string]string)
	data["Name"] = *nameFlag
	data["Version"] = time.Now().Format("20060102150405")

	var out bytes.Buffer
	t := template.Must(template.New("").Parse(`
		package migrations

		import (
			"database/sql"
		)

		func init() {
			migrator.AddMigration(&Migration{
				Version: "{{.Version}}",
				Name:    "{{.Name}}",
				Up:      mig_{{.Version}}_{{.Name}}_up,
				Down:    mig_{{.Version}}_{{.Name}}_down,
			})
		}

		func mig_{{.Version}}_{{.Name}}_up(tx *sql.Tx) error {
			return nil
		}

		func mig_{{.Version}}_{{.Name}}_down(tx *sql.Tx) error {
			return nil
		}
	`))
	if err := t.Execute(&out, data); err != nil {
		fmt.Println("Error while writing output:", err.Error())
		os.Exit(1)
	}

	filePath := fmt.Sprintf("./migrations/mig_%s_%s.go", data["Version"], data["Name"])
	fmt.Println("Writing to file:", filePath)

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		fmt.Println("Error while writing to file:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Created migration file")
}
