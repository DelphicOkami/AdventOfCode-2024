package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NineSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunNineSuite(t *testing.T) {
	suite.Run(t, new(NineSuite))
}

func (suite *NineSuite) SetupTest() {
	suite.ProvidedInput = []string{"2333133121414131402"}
}

func (suite *NineSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 1928, days.DayNinePart1(suite.ProvidedInput))
}

func (suite *NineSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 2858, days.DayNinePart2(suite.ProvidedInput))
}

func (suite *NineSuite) TestInputParsing() {
	expected := []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}
	assert.Equal(suite.T(), expected, days.Parse9Input(suite.ProvidedInput).FullMap)
}

func (suite *NineSuite) TestGetNextFreeSpace() {
	disk := days.Parse9Input(suite.ProvidedInput)
	fs, err := disk.GetNextFreeSpace()
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)
	fs, err = disk.GetNextFreeSpace()
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)

	disk.FullMap[2] = 7
	fs, err = disk.GetNextFreeSpace()
	assert.Equal(suite.T(), 3, fs)
	assert.Nil(suite.T(), err)

	disk.FullMap = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fs, err = disk.GetNextFreeSpace()
	assert.Equal(suite.T(), -1, fs)
	assert.NotNil(suite.T(), err)

}

func (suite *NineSuite) TestDefrag() {
	defragged := []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	disk := days.Parse9Input(suite.ProvidedInput)

	disk.Defrag()
	assert.Equal(suite.T(), defragged, disk.FullMap)
	defragged = []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}

	disk.ChunkDefrag()
	assert.Equal(suite.T(), defragged, disk.FullMap)
}

func (suite *NineSuite) TestChecksum() {
	disk := days.Parse9Input(suite.ProvidedInput)

	disk.Defrag()
	assert.Equal(suite.T(), 1928, disk.GetChecksum())
}

func (suite *NineSuite) TestGetFreespaceOfSize() {
	disk := days.Parse9Input(suite.ProvidedInput)
	fs, err := disk.GetNextFreeOfSizeSpace(1)
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)
	fs, err = disk.GetNextFreeOfSizeSpace(2)
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)
	fs, err = disk.GetNextFreeOfSizeSpace(3)
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)
	disk.FullMap[2] = 7
	fs, err = disk.GetNextFreeOfSizeSpace(3)
	assert.Equal(suite.T(), 8, fs)
	assert.Nil(suite.T(), err)
	_, err = disk.GetNextFreeOfSizeSpace(4)
	assert.NotNil(suite.T(), err)

}

func (suite *NineSuite) TestGetFileStarts() {
	disk := days.Parse9Input(suite.ProvidedInput)
	fs, err := disk.GetFileStart(0)
	assert.Equal(suite.T(), 0, fs)
	assert.Nil(suite.T(), err)
	fs, err = disk.GetFileStart(2)
	assert.Equal(suite.T(), 11, fs)
	assert.Nil(suite.T(), err)
	_, err = disk.GetFileStart(11)
	assert.NotNil(suite.T(), err)
}

func (suite *NineSuite) TestGetFileLen() {
	disk := days.Parse9Input(suite.ProvidedInput)
	fs, err := disk.GetFileLen(0)
	assert.Equal(suite.T(), 2, fs)
	assert.Nil(suite.T(), err)
	fs, err = disk.GetFileLen(2)
	assert.Equal(suite.T(), 1, fs)
	assert.Nil(suite.T(), err)
	_, err = disk.GetFileLen(11)
	assert.NotNil(suite.T(), err)
}
