package printage

import (
	"fmt"
	"strconv"

	"lem-in/utils"
)

func M(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}

func shortIndex(Map map[int]int) int {
	i := Map[0]
	index := 0
	for t, len := range Map {
		if len < i {
			i = len
			index = t
		}
	}
	return index
}

func Sendants(colony *utils.AntFarm) {
	fmt.Println(colony)

	antGroups := make([][]string, len(utils.Paths))
	antId := 1
	mapp := M(utils.Paths)

	for antId <= utils.Ants {
		minPath := shortIndex(mapp)
		antGroups[minPath] = append(antGroups[minPath], "L"+strconv.Itoa(antId))
		antId++
		mapp[minPath]++
	}
	control_trafic(antGroups,colony)
}

func control_trafic(antGroups [][]string,colony *utils.AntFarm ) {
	trafic := make(map[string]int)
	Emptyroom := make(map[string]bool)
	finished := 0
	for finished != utils.Ants {
		for i := 0; i < len(utils.Paths); i++ {
			Emptyroom[colony.End.Name] = false
			for currentStep := 0; currentStep < len(antGroups[i]); currentStep++ {
				ant := antGroups[i][currentStep]
				nextroom := utils.Paths[i][trafic[ant]+1]
				if !Emptyroom[nextroom] {
					fmt.Printf("%v-%v ", ant, nextroom)
					Emptyroom[nextroom] = true
					Emptyroom[utils.Paths[i][trafic[ant]]] = false
					if nextroom == colony.End.Name {
						finished++
						delete(trafic, ant)
						antGroups[i] = append(antGroups[i][:currentStep], antGroups[i][currentStep+1:]...)
						currentStep--
						Emptyroom[colony.End.Name] = true
						continue
					}
					trafic[ant]++
				}
			}
		}
		fmt.Println()
	}
}
