package terrains

func MakeNorthPole() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint16(2, 6),
		CombineUint16(4, 4),
		CombineUint16(5, 7),
		CombineUint16(6, 6),
		CombineUint16(7, 5),
		CombineUint16(9, 2),
		CombineUint16(10, 8),
	}
	MarkIslands(terrain, islands)
	return terrain
}
