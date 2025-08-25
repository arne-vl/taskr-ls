package taskr

type Config struct {
	Tasks                 map[string]Task
	Environments          map[string]Environment
	HasDefaultEnvironment bool
}

type Task struct {
	Name  string
	Run   string
	Desc  string
	Alias []string
	Needs []string
}

type Environment struct {
	Name      string
	FileName  string
	IsDefault bool
}

func Analyze(doc string) Config {
	config := Config{}

	return config
}
