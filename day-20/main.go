package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Orientation struct {
	Grid        [][]rune
	TopRow      string
	BottomRow   string
	LeftColumn  string
	RightColumn string
}

type Tile struct {
	ID                 int
	CurrentOrientation int
	Grid               [][]rune
	Orientations       []Orientation
	Sides              map[string]bool
}

func main() {
	tiles, err := readTiles("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	seaMonster, err := readSeaMonster("sea_monster.txt")
	if err != nil {
		log.Fatal(err)
	}

	filledGrid, value := partOne(tiles)
	fmt.Println(value)
	fmt.Println(partTwo(filledGrid, seaMonster))
}

func partOne(tiles []Tile) ([][]Tile, int) {
	for tileIndex, tile := range tiles {
		tiles[tileIndex].Orientations = determineOrientations(tile)
		tiles[tileIndex].Sides = findSides(tiles[tileIndex])
	}
	tileSize := int(math.Sqrt(float64(len(tiles))))
	grid := make([][]Tile, tileSize)
	for j := range grid {
		grid[j] = make([]Tile, tileSize)
	}
	filledGrid, isFilled := fillGrid(grid, tiles)
	if isFilled {
		return filledGrid, filledGrid[0][0].ID * filledGrid[0][len(grid)-1].ID * filledGrid[len(grid)-1][0].ID * filledGrid[len(grid)-1][len(grid)-1].ID
	}
	return [][]Tile{}, -1
}

func determineOrientations(tile Tile) []Orientation {
	orientations := []Orientation{}
	for i := 0; i < 8; i++ {
		newTile := tile
		newTile.Grid = orientGrid(newTile.Grid, i)
		topRow, bottomRow, leftColumn, rightColumn := "", "", "", ""
		for _, value := range newTile.Grid[0] {
			topRow += string(value)
		}
		for _, value := range newTile.Grid[len(newTile.Grid)-1] {
			bottomRow += string(value)
		}
		for j := range newTile.Grid[0] {
			leftColumn += string(newTile.Grid[j][0])
		}
		for k := range newTile.Grid[0] {
			rightColumn += string(newTile.Grid[k][len(newTile.Grid)-1])
		}
		orientations = append(orientations, Orientation{Grid: newTile.Grid, TopRow: topRow, BottomRow: bottomRow, LeftColumn: leftColumn, RightColumn: rightColumn})
	}
	return orientations
}

func orientGrid(grid [][]rune, value int) [][]rune {
	newGrid := grid
	for i := 0; i < value/4; i++ {
		newGrid = reflectGrid(newGrid)
	}
	for j := 0; j < value%4; j++ {
		newGrid = rotateGrid(newGrid)
	}
	return newGrid
}

func rotateGrid(grid [][]rune) [][]rune {
	newGrid := [][]rune{}
	for i := range grid[0] {
		row := []rune{}
		for j := range grid {
			row = append(row, grid[len(grid)-j-1][i])
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

func reflectGrid(grid [][]rune) [][]rune {
	newGrid := [][]rune{}
	for i := range grid {
		row := []rune{}
		for j := range grid[0] {
			row = append(row, grid[i][len(grid[0])-j-1])
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

func findSides(tile Tile) map[string]bool {
	sides := map[string]bool{}
	for _, orientation := range tile.Orientations {
		sides[orientation.TopRow] = true
		sides[orientation.BottomRow] = true
		sides[orientation.LeftColumn] = true
		sides[orientation.RightColumn] = true
	}
	return sides
}

func fillGrid(grid [][]Tile, tiles []Tile) ([][]Tile, bool) {
	if len(tiles) == 0 {
		return grid, true
	}
	nextMissingTileRow, nextMissingTileColumn := findNextMissingTile(grid)
	for tileIndex, tile := range tiles {
		for i := 0; i < 8; i++ {
			newTile := tile
			if tileFits(grid, newTile, nextMissingTileRow, nextMissingTileColumn, i) {
				newTile.CurrentOrientation = i
				newGrid := make([][]Tile, len(grid))
				for j := range grid {
					newGrid[j] = make([]Tile, len(grid))
					copy(newGrid[j], grid[j])
				}
				newGrid[nextMissingTileRow][nextMissingTileColumn] = newTile
				remainingTiles := []Tile{}
				remainingTiles = append(remainingTiles, tiles[:tileIndex]...)
				remainingTiles = append(remainingTiles, tiles[tileIndex+1:]...)
				completeGrid, isFilled := fillGrid(newGrid, remainingTiles)
				if isFilled {
					return completeGrid, true
				}
			}
		}
	}
	return nil, false
}

func findNextMissingTile(grid [][]Tile) (int, int) {
	for i, row := range grid {
		for j, tile := range row {
			if tile.Grid == nil {
				return i, j
			}
		}
	}
	return -1, -1
}

func tileFits(grid [][]Tile, tile Tile, row, column, orientation int) bool {
	if row > 0 {
		aboveTile := grid[row-1][column]
		if aboveTile.Orientations[aboveTile.CurrentOrientation].BottomRow != tile.Orientations[orientation].TopRow {
			return false
		}
	}
	if column > 0 {
		leftTile := grid[row][column-1]
		if leftTile.Orientations[leftTile.CurrentOrientation].RightColumn != tile.Orientations[orientation].LeftColumn {
			return false
		}
	}
	return true
}

func partTwo(filledGrid [][]Tile, seaMonster [][]rune) int {
	newGrid := [][]rune{}
	for _, row := range filledGrid {
		newRows := make([][]rune, len(filledGrid[0][0].Grid)-2)
		for _, tile := range row {
			reducedGrid := cropGrid(tile.Orientations[tile.CurrentOrientation].Grid)
			for reducedGridIndex, reducedGridRow := range reducedGrid {
				newRows[reducedGridIndex] = append(newRows[reducedGridIndex], reducedGridRow...)
			}
		}
		newGrid = append(newGrid, newRows...)
	}

	horizontalSeaMonsters := [][][]rune{}
	for i := 0; i < 4; i++ {
		horizontalSeaMonsters = append(horizontalSeaMonsters, orientGrid(seaMonster, i*2))
	}
	verticalSeaMonsters := [][][]rune{}
	for i := 0; i < 4; i++ {
		verticalSeaMonsters = append(verticalSeaMonsters, orientGrid(seaMonster, i*2+1))
	}

	seaMonsterIndices := map[[2]int]bool{}
	for i := 0; i < len(newGrid)-len(horizontalSeaMonsters[0]); i++ {
		for j := 0; j < len(newGrid[0])-len(horizontalSeaMonsters[0][0]); j++ {
			seaMonsterIndices = findSeaMonsters(i, j, newGrid, horizontalSeaMonsters, seaMonsterIndices)
		}
	}
	for i := 0; i < len(newGrid)-len(verticalSeaMonsters[0]); i++ {
		for j := 0; j < len(newGrid[0])-len(verticalSeaMonsters[0][0]); j++ {
			seaMonsterIndices = findSeaMonsters(i, j, newGrid, verticalSeaMonsters, seaMonsterIndices)
		}
	}

	notSeaMonsterCount := 0
	for rowindex, row := range newGrid {
		for columnIndex := range row {
			if _, ok := seaMonsterIndices[[2]int{rowindex, columnIndex}]; newGrid[rowindex][columnIndex] == '#' && !ok {
				notSeaMonsterCount++
			}
		}
	}
	return notSeaMonsterCount
}

func cropGrid(grid [][]rune) [][]rune {
	newGrid := [][]rune{}
	for i := 1; i < len(grid)-1; i++ {
		row := []rune{}
		for j := 1; j < len(grid)-1; j++ {
			row = append(row, grid[i][j])
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

func findSeaMonsters(row, column int, grid [][]rune, seaMonsters [][][]rune, seaMonsterMap map[[2]int]bool) map[[2]int]bool {
CheckSeaMonster:
	for _, seaMonster := range seaMonsters {
		positions := map[[2]int]bool{}
		for smRowIndex, smRow := range seaMonster {
			for smColumnIndex, pos := range smRow {
				if pos == '#' {
					if grid[smRowIndex+row][smColumnIndex+column] != '#' {
						continue CheckSeaMonster
					}
					positions[[2]int{smRowIndex + row, smColumnIndex + column}] = true
				}
			}
		}
		for key, value := range positions {
			seaMonsterMap[key] = value
		}
	}
	return seaMonsterMap
}

func readTiles(fileName string) ([]Tile, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	tiles := []Tile{}
	newTile := Tile{Grid: [][]rune{}}
	for scanner.Scan() {
		if scanner.Text() == "" {
			tiles = append(tiles, newTile)
			newTile = Tile{}
		} else if scanner.Text()[0] == 'T' {
			idString := strings.TrimSuffix(strings.TrimPrefix(scanner.Text(), "Tile "), ":")
			id, err := strconv.Atoi(idString)
			if err != nil {
				return nil, err
			}

			newTile.ID = id
		} else {
			row := []rune{}
			for _, character := range scanner.Text() {
				row = append(row, character)
			}
			newTile.Grid = append(newTile.Grid, row)
		}
	}
	tiles = append(tiles, newTile)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tiles, nil
}

func readSeaMonster(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	monster := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		for _, character := range scanner.Text() {
			row = append(row, character)
		}
		monster = append(monster, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return monster, nil
}
