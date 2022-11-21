package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid, err := readGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(grid))
	fmt.Println(partTwo(grid))
}

func partOne(grid [][]string) int {
	threeDimensionalGrid := [][][]string{grid}
	for i := 1; i <= 6; i++ {
		threeDimensionalGrid = runCycle3D(threeDimensionalGrid)
	}
	return countAllActiveCubes3D(threeDimensionalGrid)
}

func runCycle3D(grid [][][]string) [][][]string {
	newGrid := [][][]string{}
	for i := -1; i <= len(grid); i++ {
		gridSlice := [][]string{}
		for j := -1; j <= len(grid[0]); j++ {
			row := []string{}
			for k := -1; k <= len(grid[0][0]); k++ {
				activeNearbyCubes := countNearbyActiveCubes3D(i, j, k, grid)
				currentStatus := "#"
				if i == -1 || i == len(grid) || j == -1 || j == len(grid[0]) || k == -1 || k == len(grid[0][0]) || grid[i][j][k] == "." {
					currentStatus = "."
				}
				row = append(row, determineNewStatus(currentStatus, activeNearbyCubes))
			}
			gridSlice = append(gridSlice, row)
		}
		newGrid = append(newGrid, gridSlice)
	}
	return newGrid
}

func countNearbyActiveCubes3D(x, y, z int, grid [][][]string) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				} else if i+x < 0 || i+x >= len(grid) || j+y < 0 || j+y >= len(grid[0]) || k+z < 0 || k+z >= len(grid[0][0]) {
					continue
				}
				if grid[i+x][j+y][k+z] == "#" {
					count++
				}
			}
		}
	}
	return count
}

func determineNewStatus(currentStatus string, activeNearbyCubes int) string {
	if currentStatus == "#" && (activeNearbyCubes == 2 || activeNearbyCubes == 3) {
		return "#"
	} else if currentStatus == "#" {
		return "."
	} else if currentStatus == "." && activeNearbyCubes == 3 {
		return "#"
	} else {
		return "."
	}
}

func countAllActiveCubes3D(grid [][][]string) int {
	total := 0
	for _, gridSlice := range grid {
		for _, row := range gridSlice {
			for _, cube := range row {
				if cube == "#" {
					total++
				}
			}
		}
	}
	return total
}

func partTwo(grid [][]string) int {
	gridSlice := [][][]string{grid}
	fourDimensionalGrid := [][][][]string{gridSlice}
	for i := 1; i <= 6; i++ {
		fourDimensionalGrid = runCycle4D(fourDimensionalGrid)
	}
	return countAllActiveCubes4D(fourDimensionalGrid)
}

func runCycle4D(grid [][][][]string) [][][][]string {
	newGrid := [][][][]string{}
	for i := -1; i <= len(grid); i++ {
		grid3D := [][][]string{}
		for j := -1; j <= len(grid[0]); j++ {
			gridSlice := [][]string{}
			for k := -1; k <= len(grid[0][0]); k++ {
				row := []string{}
				for l := -1; l <= len(grid[0][0][0]); l++ {
					activeNearbyCubes := countNearbyActiveCubes4D(i, j, k, l, grid)
					currentStatus := "#"
					if i == -1 || i == len(grid) || j == -1 || j == len(grid[0]) || k == -1 || k == len(grid[0][0]) || l == -1 || l == len(grid[0][0][0]) || grid[i][j][k][l] == "." {
						currentStatus = "."
					}
					row = append(row, determineNewStatus(currentStatus, activeNearbyCubes))
				}
				gridSlice = append(gridSlice, row)
			}
			grid3D = append(grid3D, gridSlice)
		}
		newGrid = append(newGrid, grid3D)
	}
	return newGrid
}

func countNearbyActiveCubes4D(x, y, z, w int, grid [][][][]string) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					} else if i+x < 0 || i+x >= len(grid) || j+y < 0 || j+y >= len(grid[0]) || k+z < 0 || k+z >= len(grid[0][0]) || l+w < 0 || l+w >= len(grid[0][0][0]) {
						continue
					}
					if grid[i+x][j+y][k+z][l+w] == "#" {
						count++
					}
				}
			}
		}
	}
	return count
}

func countAllActiveCubes4D(grid [][][][]string) int {
	total := 0
	for _, grid3D := range grid {
		for _, gridSlice := range grid3D {
			for _, row := range gridSlice {
				for _, cube := range row {
					if cube == "#" {
						total++
					}
				}
			}
		}
	}
	return total
}

func readGrid(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
