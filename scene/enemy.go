package scene

import (
	"fmt"

	"github.com/pankona/phantomize/scene/config"
)

func (g *game) initUnits(json string) {
	// TODO: load from json file

	// TODO: unitpopTimeTable should be sorted by popTime
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 3 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 4 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 5 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 7 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 7 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 8 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 9 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 10 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 15 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 15 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 15 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 20 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 20 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 20 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 25 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 25 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 35 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 35 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 35 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 42 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 42 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 42 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 49 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 49 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 49 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * fps, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	units := make(map[string]uniter)
	for i, v := range g.unitPopTimeTable {
		g.unitPopTimeTable[i].unitID = fmt.Sprintf("e%d", i)
		units[v.unitID] = newUnit(v.unitID, v.unittype, g)
	}
	g.uniters = units
}
