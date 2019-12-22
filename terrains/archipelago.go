package terrains

func MakeArchipelago() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint16(2, 3),
		CombineUint16(2, 7),
		CombineUint16(3, 5),
		CombineUint16(3, 9),
		CombineUint16(4, 2),
		CombineUint16(5, 7),
		CombineUint16(6, 2),
		CombineUint16(6, 5),
		CombineUint16(6, 9),
		CombineUint16(7, 6),
		CombineUint16(8, 3),
		CombineUint16(9, 8),
	}
	MarkIslands(terrain, islands)
	return terrain
}
