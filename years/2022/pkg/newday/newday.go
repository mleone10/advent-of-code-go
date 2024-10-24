package newday

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed day.go.template
var dayTemplate string

//go:embed day_test.go.template
var dayTestTemplate string

var tDay = template.Must(template.New("day").Parse(dayTemplate))
var tDayTest = template.Must(template.New("dayTest").Parse(dayTestTemplate))

type templateValues struct {
	DayNum int
}

// TODO: Implement tests for helper client (will need to mock os.Mkdir and os.Create)
func InitializeDay(dayNum int) error {
	err := createDayDirectory(dayNum)
	if err != nil {
		return err
	}

	tvs := templateValues{DayNum: dayNum}

	err = writeTemplateToFile(tDay, tvs, fmt.Sprintf("src/day%02[1]d/day%02[1]d.go", dayNum))
	if err != nil {
		return err
	}
	err = writeTemplateToFile(tDayTest, tvs, fmt.Sprintf("src/day%02[1]d/day%02[1]d_test.go", dayNum))
	if err != nil {
		return err
	}

	_, err = createFile(fmt.Sprintf("src/day%02[1]d/input.txt", dayNum))
	if err != nil {
		return err
	}

	_, err = createFile(fmt.Sprintf("src/day%02[1]d/test_input.txt", dayNum))
	if err != nil {
		return err
	}

	return nil
}

func createDayDirectory(dayNum int) error {
	err := os.Mkdir(fmt.Sprintf("src/day%02d", dayNum), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory for day %d: %w", dayNum, err)
	}

	return nil
}

func writeTemplateToFile(t *template.Template, tvs templateValues, filePath string) error {
	file, err := createFile(filePath)

	err = t.Execute(file, tvs)
	if err != nil {
		return fmt.Errorf("failed to write template data to file: %w", err)
	}

	return nil
}

func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create empty file: %w", err)
	}
	return file, nil
}
