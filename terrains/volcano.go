package terrains

func MakeVolcano() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint8(2, 8),
		CombineUint8(3, 3),
		CombineUint8(5, 5),
		CombineUint8(5, 9),
		CombineUint8(6, 3),
		CombineUint8(7, 6),
		CombineUint8(8, 9),
		CombineUint8(9, 5),
	}
	MarkIslands(terrain, islands)
	return terrain
}

