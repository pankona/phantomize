package scene

import (
	"fmt"

	"github.com/pankona/phantomize/scene/config"
)

func (g *game) initUnits(json string) {
	// TODO: load from json file

	// TODO: unitpopTimeTable should be sorted by popTime
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 3 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 7 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 10 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 10 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 14 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 14 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 18 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 18 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 22 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 22 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 22 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 27 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 27 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 32 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 32 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 37 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 40 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 42 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 42 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 47 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 49 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 49 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 52 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 54 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy1", popTime: 57 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4}})
	g.unitPopTimeTable = append(g.unitPopTimeTable, &unitPopTime{unittype: "enemy2", popTime: 60 * framePerSec, initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3}})

	units := make(map[string]uniter)
	for i, v := range g.unitPopTimeTable {
		g.unitPopTimeTable[i].unitID = fmt.Sprintf("e%d", i)
		units[v.unitID] = newUnit(v.unitID, v.unittype, g)
	}
	g.uniters = units
}
