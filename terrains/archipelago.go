package terrains

func MakeArchipelago() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint8(2, 3),
		CombineUint8(2, 7),
		CombineUint8(3, 5),
		CombineUint8(3, 9),
		CombineUint8(4, 2),
		CombineUint8(5, 7),
		CombineUint8(6, 2),
		CombineUint8(6, 5),
		CombineUint8(6, 9),
		CombineUint8(7, 6),
		CombineUint8(8, 3),
		CombineUint8(9, 8),
	}
	MarkIslands(terrain, islands)
	return terrain
}
