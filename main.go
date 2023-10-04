package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
	"github.com/tealeg/xlsx"
)

type Config struct {
	DBHost     string `toml:"db_host"`
	DBPort     int    `toml:"db_port"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
	DBName     string `toml:"db_name"`
	TableName  string `toml:"table_name"`
}

var (
	xlsxFilePathFlag string
)

func init() {
	flag.StringVar(&xlsxFilePathFlag, "f", "", "Path to XLSX file")
	flag.Parse()
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(&config)
	return config, err
}

func main() {
	xlsxFilePath := xlsxFilePathFlag
	if xlsxFilePath == "" {
		log.Fatal("No XLSX file specified. Usage: ./XLSXtoMySQL-Importer -f <file_path>")
	}

	config, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	logMulti := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logMulti)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	xlFile, err := xlsx.OpenFile(xlsxFilePath)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}

			values := make([]interface{}, 0)
			columns := make([]string, 0)

			for i, cell := range row.Cells {
				cellValue := strings.TrimSpace(cell.String())

				var value interface{}
				if cellValue == "" {
					value = nil
				} else {
					cellValue = strings.ReplaceAll(cellValue, "'", "''")
					value = cellValue
				}

				values = append(values, value)
				columns = append(columns, xlFile.Sheets[0].Rows[0].Cells[i].String())
			}

			columnsStr := strings.Join(columns, ", ")
			placeholders := strings.Repeat("?, ", len(columns)-1) + "?"
			query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", config.TableName, columnsStr, placeholders)

			// Use nil where needed to insert empty values as NULL
			_, err := db.Exec(query, values...)
			if err != nil {
				log.Printf("Error: %v", err)
				log.Printf("Failed to insert data SQL: %s, Values: %+v", query, values)
			}
		}
	}

	log.Println("Data transfer completed. Please check logfile.log file.")
}
