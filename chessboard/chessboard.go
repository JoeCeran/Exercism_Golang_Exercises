package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var squareOccupied int
	for _, square := range cb[file] {
		if square == true {
			squareOccupied = squareOccupied + 1
		}
	}
	return squareOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var squareOccupied int
	file := [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, fileRank := range file{
		for squareRank, squareFile := range cb[fileRank] {
			if squareFile == true && squareRank == rank - 1{
				squareOccupied = squareOccupied + 1
			}
		}
	}
	return squareOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	file := [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	var returnI int
	for _, fileRank := range file {
		for _, _ = range cb[fileRank] {
			returnI = returnI + 1
		}
	}
	return returnI
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var squareOccupied int
	file := [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, fileRank := range file{
		for _, squareFile := range cb[fileRank] {
			if squareFile == true{
				squareOccupied = squareOccupied + 1
			}
		}
	}
	return squareOccupied
}
