package main
// Исправить позже
import (
    "fmt"
    "time"
    "github.com/caarlos0/env"
)

type config struct {
    Home        string        `env:"HOME"`
    Port        int           `env:"PORT" envDefault:"8080"`
    IsProduction bool          `env:"PRODUCTION,required"`
    Hosts       []string      `env:"HOSTS" envSeparator:":"`
    Duration   time.Duration `env:"DURATION"`
    TempFolder  string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp"`
}

/*
PRODUCTION=true HOSTS="host1:host2:host3" DURATION=1s go run 1.go
{Home:/your/home Port:8080 IsProduction:true Hosts:[host1 host2 host3] Duration:1s}
 */

func main() {
    cfg := config{}
    if err := env.Parse(cfg); err!= nil {
        fmt.Println(err)
        return
    }
	fmt.Println(cfg)
}