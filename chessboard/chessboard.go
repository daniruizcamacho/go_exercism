package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	rankValues, exists := cb[rank]

	if exists == false {
		return count
	}

	for _, rankValue := range rankValues {
		if rankValue {
			count += 1
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
	for _, rankValues := range cb {
		for fileIndex, fileValue := range rankValues {
			if fileIndex != file-1 {
				continue
			}

			if fileValue {
				count += 1
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	rows := len(cb)
	columns := len(cb["A"])

	return rows * columns
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, rankValues := range cb {
		for _, fileValue := range rankValues {
			if fileValue {
				count += 1
			}
		}
	}

	return count
}
