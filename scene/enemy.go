package scene

import "github.com/pankona/phantomize/scene/config"

func (g *game) initUnits(json string) {
	// TODO: load from json file
	units := make(map[string]uniter)
	units["e1"] = newUnit("e1", "enemy1", g)
	units["e2"] = newUnit("e2", "enemy1", g)
	units["e3"] = newUnit("e3", "enemy2", g)

	// TODO: unitpopTimeTable should be sorted by popTime
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e1",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e2",
			popTime:         4 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e3",
			popTime:         5 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3},
		})

	g.uniters = units
}
