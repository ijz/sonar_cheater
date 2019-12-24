package terrains

func MakeNorthPole() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint8(2, 6),
		CombineUint8(4, 4),
		CombineUint8(5, 7),
		CombineUint8(6, 6),
		CombineUint8(7, 5),
		CombineUint8(9, 2),
		CombineUint8(10, 8),
	}
	MarkIslands(terrain, islands)
	return terrain
}
