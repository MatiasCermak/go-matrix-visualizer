package handlers

import (
	"encoding/json"
	"matrix-visualizer/backend/pkg/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Map where matrix operations will be stored
var matrixOperationMap map[string]func(int, int) [][]uint = map[string]func(int, int) [][]uint{
	"spiralFibonacci": generateSpiralFibonacciMatrix,
}

func calcFibonacci(n int, ch chan (uint)) {
	var firstNumber uint = 0
	var secondNumber uint = 1
	ch <- firstNumber
	ch <- secondNumber
	for i := 2; i < n; i++ {
		firstNumber, secondNumber = secondNumber, firstNumber+secondNumber
		ch <- secondNumber
	}
}

func generateSpiralFibonacciMatrix(rows, cols int) [][]uint {
	matrix := make([][]uint, rows)
	for i := range matrix {
		matrix[i] = make([]uint, cols)
	}

	row, col := 0, 0
	intChannel := make(chan (uint))
	go calcFibonacci(rows*cols, intChannel)

	for i := 0; i < (rows+1)/2 && i < (cols+1)/2; i++ {
		// Traverse right
		for j := col; j < cols-i; j++ {
			matrix[row][j] = <-intChannel
		}

		// Traverse down
		for j := row + 1; j < rows-i; j++ {
			matrix[j][cols-i-1] = <-intChannel
		}

		// Traverse left
		for j := cols - i - 2; j >= col && rows-i-1 != row; j-- {
			matrix[rows-i-1][j] = <-intChannel
		}

		// Traverse up
		for j := rows - i - 2; j > row && cols-i-1 != col; j-- {
			matrix[j][col] = <-intChannel
		}

		row++
		col++
	}

	return matrix
}

func handleError(w http.ResponseWriter, response domain.Response, message string) {
	w.WriteHeader(http.StatusBadRequest)
	response.Message = message
	response.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	json.NewEncoder(w).Encode(response)
}

func GenerateMatrix(w http.ResponseWriter, r *http.Request) {
	var response domain.Response
	pathVariables := mux.Vars(r)

	rowsParam := r.URL.Query().Get("rows")
	colsParam := r.URL.Query().Get("cols")
	if rowsParam == "" || colsParam == "" {
		handleError(w, response, "Missing 'rows' or 'cols' parameter")
		return
	}

	rows, err := strconv.Atoi(rowsParam)
	if err != nil || rows < 1 {
		handleError(w, response, "Invalid 'rows' parameter")
		return
	}

	cols, err := strconv.Atoi(colsParam)
	if err != nil || cols < 1 {
		handleError(w, response, "Invalid 'cols' parameter")
		return
	}

	//Matrix type cannot be nil because of gorilla mux pattern matching
	matrixType := pathVariables["matrixOperation"]
	matrixOperationFunction, exists := matrixOperationMap[matrixType]
	if !exists {
		handleError(w, response, "Invalid matrix operation")
		return
	}

	response.Rows = matrixOperationFunction(rows, cols)

	response.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
