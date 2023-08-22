package chunker

//TODO: make generic later
//TODO: unit-test
func Chunk(workerNum int, data []string) [][]string {
	totalNum := len(data)
	if totalNum <= 0 {
		return nil
	}
	if totalNum < workerNum {
		workerNum = totalNum
	}
	result := make([][]string, workerNum)
	var chunkSize int = len(data)/workerNum
	var count int;
	for i := 0; i < workerNum; i++ {
		result[i] = data[count:count+chunkSize]
		count += chunkSize
	}
	return result
}