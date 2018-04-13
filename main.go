package main

import (
  "flag"
  "os"
  "fmt"

  "github.com/google/logger"
  "github.com/jinzhu/configor"

)

const logPath = "./logs/backup.log"
const configPath = "./etc/config.yml"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

var Config = struct {
        Mode string `required:"true" default:"cron"`
        Format string `default:"/service_name/date"`
        LogLevel string `default:"info"`
        LogPath string `default:"./logs/backups.log"`
        Verbose bool `default:"false"`
        Scheduler string `default:"00 23 * * *"`

        Notifications []struct {
                  id  string
        }

        Storages []struct {
                  id  string
        }

        Services []struct {
                  id  string
        }
}{}

func main() {
  flag.Parse()

  cf, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
  if err != nil {
    logger.Fatalf("Failed to open configuration file: %v", err)
  }
  defer cf.Close()

  defer configor.Load(&Config, configPath)
  fmt.Printf("config: %#v", Config)

  lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
  if err != nil {
    logger.Fatalf("Failed to open log file: %v", err)
  }

  defer lf.Close()

  defer logger.Init("LoggerExample", *verbose, true, lf).Close()

  logger.Info("I'm about to do something!")
  // if err := doSomething(); err != nil {
  //   logger.Errorf("Error running doSomething: %v", err)
  // }
}
