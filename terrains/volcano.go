package terrains

func MakeVolcano() []uint8 {
	terrain := InitializeEmptyTerrain()
	islands := []uint16{
		CombineUint16(2, 8),
		CombineUint16(3, 3),
		CombineUint16(5, 5),
		CombineUint16(5, 9),
		CombineUint16(6, 3),
		CombineUint16(7, 6),
		CombineUint16(8, 9),
		CombineUint16(9, 5),
	}
	MarkIslands(terrain, islands)
	return terrain
}

