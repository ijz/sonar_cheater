package terrains

import "fmt"

func InitializeEmptyTerrain() []uint8 {
	terrain := make([]uint8, 100, 100)
	return terrain
}

func GetValueAt(terrain []uint8, r uint8, c uint8)uint8 {
	index := r * 10 + c
	return terrain[index]
}

func SetValueAt(terrain []uint8, r uint8, c uint8, val uint8) {
	index := r * 10 + c
	terrain[index] = val
}

func SplitUint32(i uint16) (uint8, uint8) {
	var mask uint16 = 0x00FF
	c := uint8(i & mask)
	r := uint8((i & ^mask) >> 8)
	return r, c
}

func CombineUint16(r uint8, c uint8) uint16 {
	var i uint16 = 0
	i = i | uint16(r)
	i = i << 8
	i = i | uint16(c)
	return i
}

func MarkIslands(terrain []uint8, islands []uint16) {
	for _, i := range islands {
		x, y := SplitUint32(i)
		SetValueAt(terrain, x - 1, y - 1, 1)
	}
}

func PrintTerrain(terrain []uint8) {
	fmt.Println("\tA\tB\tC\tD\tE\tF\tG\tH\tI\tJ\t")
	for x := 0; x < 10; x++ {
		fmt.Printf("%d\t", x + 1)
		for y := 0; y < 10; y++ {
			fmt.Printf("%d\t", GetValueAt(terrain, uint8(x), uint8(y)))
		}
		fmt.Println()
	}
}
