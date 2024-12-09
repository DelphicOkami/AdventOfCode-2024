package days

import (
	"fmt"
	"strconv"
	"strings"
)

func DayNinePart1(input []string) int {
	disk := Parse9Input(input)
	disk.Defrag()
	return disk.GetChecksum()
}

func DayNinePart2(input []string) int {
	disk := Parse9Input(input)
	disk.ChunkDefrag()
	return disk.GetChecksum()
}

func Parse9Input(input []string) Disk9 {
	cmap := strings.Join(input, "")
	out := Disk9{
		CompressedMap: cmap,
	}
	fullMap := make([]int, 0)
	for fi := 0; fi < len(cmap); fi += 2 {
		index := fi / 2
		fileLen, _ := strconv.Atoi(string(cmap[fi]))
		spaceLen := 0
		if len(cmap) > fi+1 {
			spaceLen, _ = strconv.Atoi(string(cmap[fi+1]))
		}

		for l := 0; l < fileLen; l++ {
			fullMap = append(fullMap, index)
		}
		for l := 0; l < spaceLen; l++ {
			fullMap = append(fullMap, -1)
		}
	}

	out.FullMap = make([]int, len(fullMap))
	out.OriginalFullMap = make([]int, len(fullMap))
	copy(out.FullMap, fullMap)
	copy(out.OriginalFullMap, fullMap)
	return out
}

type Disk9 struct {
	CompressedMap   string
	FullMap         []int
	OriginalFullMap []int
	nextFreeSpace   int
}

func (d *Disk9) GetNextFreeSpace() (int, error) {
	for i := d.nextFreeSpace; i <= len(d.FullMap); i++ {
		if i == len(d.FullMap) {
			return -1, fmt.Errorf("no freespace left on disk")
		}
		if d.FullMap[i] == int(-1) {
			d.nextFreeSpace = i
			break
		}
	}
	return d.nextFreeSpace, nil
}

func (d *Disk9) GetFileStart(fileID int) (int, error) {
	for i, id := range d.FullMap {
		if id == fileID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("file %d not on disk", fileID)
}

func (d *Disk9) GetFileLen(fileID int) (int, error) {
	if len(d.CompressedMap) < fileID*2 {
		return -1, fmt.Errorf("file not on disk")
	}
	fl, _ := strconv.Atoi(string(d.CompressedMap[fileID*2]))
	return fl, nil
}

func (d *Disk9) GetNextFreeOfSizeSpace(size int) (int, error) {
	fs := 0
	starti := -1
	for i, fi := range d.FullMap {
		if i == len(d.FullMap) {
			return -1, fmt.Errorf("no freespace left on disk")
		}
		if fi != -1 {
			fs = 0
			starti = -1
			continue
		}
		if starti == -1 {
			starti = i
		}
		fs++
		if fs >= size {
			return starti, nil
		}
	}
	return -1, fmt.Errorf("no frespace big enough")
}

func (d *Disk9) Defrag() {
	copy(d.FullMap, d.OriginalFullMap)
	d.nextFreeSpace = 0
	for i := len(d.OriginalFullMap) - 1; i >= 0; i-- {
		if d.OriginalFullMap[i] == -1 {
			continue
		}
		fsi, err := d.GetNextFreeSpace()
		if err != nil {
			break
		}
		if fsi >= i {
			break
		}
		d.FullMap[fsi] = d.OriginalFullMap[i]
		d.FullMap[i] = -1
	}
}

func (d *Disk9) ChunkDefrag() {
	copy(d.FullMap, d.OriginalFullMap)
	startI := len(d.CompressedMap) / 2
	if len(d.CompressedMap)%2 == 1 {
		startI = (len(d.CompressedMap) - 1) / 2
	}
	for fileID := startI; fileID >= 0; fileID-- {
		fileLength, _ := d.GetFileLen(fileID)
		spaceStart, err := d.GetNextFreeOfSizeSpace(fileLength)
        if err != nil { continue }
		fileStart, _ := d.GetFileStart(fileID)
		if fileStart <= spaceStart { continue }
		for move := 0; move < fileLength; move++ {
			d.FullMap[spaceStart + move] = fileID
			d.FullMap[fileStart + move] = -1
		}
	}
}

func (d *Disk9) GetChecksum() int {
	out := 0
	for i, id := range d.FullMap {
		if id == -1 {
			continue
		}
		out += id * i
	}
	return out
}
