package terrains

func MakeOpenWaters() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint16(2, 7),
		CombineUint16(3, 8),
		CombineUint16(5, 4),
		CombineUint16(7, 9),
		CombineUint16(8, 2),
		CombineUint16(8, 3),
	}
	MarkIslands(terrain, islands)
	return terrain
}
