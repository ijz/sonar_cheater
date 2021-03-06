package terrains

import "fmt"

type Terrain uint8
const (
	ArchipelagoTerrain Terrain = 0
	VolcanoTerrain Terrain = 1
	NorthPoleTerrain Terrain = 2
	OpenWatersTerrain Terrain = 3
)

var TerrainDict = map[Terrain]func() []uint8 {
	ArchipelagoTerrain: MakeArchipelago,
	VolcanoTerrain: MakeVolcano,
	NorthPoleTerrain: MakeNorthPole,
	OpenWatersTerrain: MakeOpenWaters,
}

var TerrainNameDict = map[Terrain]string {
	ArchipelagoTerrain: "Archipelago",
	VolcanoTerrain: "Volcano",
	NorthPoleTerrain: "North Pole",
	OpenWatersTerrain: "Open Waters",
}

func InitializeEmptyTerrain() []uint8 {
	terrain := make([]uint8, 100, 100)
	return terrain
}

func GetValueAt(terrain []uint8, r uint8, c uint8) uint8 {
	index := r * 10 + c
	return terrain[index]
}

func SetValueAt(terrain []uint8, r uint8, c uint8, val uint8) {
	index := r * 10 + c
	terrain[index] = val
}

func SplitUint16(i uint16) (uint8, uint8) {
	var mask uint16 = 0x00FF
	c := uint8(i & mask)
	r := uint8((i & ^mask) >> 8)
	return r, c
}

func CombineUint8(r uint8, c uint8) uint16 {
	var i uint16 = 0
	i = i | uint16(r)
	i = i << 8
	i = i | uint16(c)
	return i
}

func Int8sFromString(s string) (int8, int8) {
	colS := s[0]
	rowS := s[1]
	var row, col int8
	if '-' == colS {
		col = -1
	} else {
		col = int8(colS - 'A')
	}
	if '-' == rowS {
		row = -1
	} else {
		row = int8(rowS - '1')
		if '1' == rowS && 3 == len(s) && '0' == s[2] {
			row = 9
		}
	}
	return row, col
}

func StringUint16(i uint16) string {
	r, c := SplitUint16(i)
	return fmt.Sprintf("%c%d",c + 'A', r + 1)
}

func MarkIslands(terrain []uint8, islands []uint16) {
	for _, i := range islands {
		x, y := SplitUint16(i)
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

func IsIsland(terrain []uint8, r uint8, c uint8) bool {
	return 1 == GetValueAt(terrain, r, c)
}
