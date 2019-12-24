package terrains

func MakeOpenWaters() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint8(2, 7),
		CombineUint8(3, 8),
		CombineUint8(5, 4),
		CombineUint8(7, 9),
		CombineUint8(8, 2),
		CombineUint8(8, 3),
	}
	MarkIslands(terrain, islands)
	return terrain
}
