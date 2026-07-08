package fileops

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func GetFloatFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000 , errors.New("FAIled to parse stored value")
	}

	valueText := string(data)
	value ,  err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return  1000 , errors.New("failed to parse stored balance value")
	}

	return value , nil
}
func WriteFloatToFile(value float64, fileName string){
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
