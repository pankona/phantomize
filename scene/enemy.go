package scene

type enemy struct {
	eID     string
	isAlive bool
}

type enemies []enemy

func (e enemy) action() {
	// TODO: implement
}

func (e enemy) dead() {
	// TODO: implement
	e.isAlive = false
}

type enemyconfig struct {
	enemyType string
	ftp       int64 // frame to pop
	pos       position
}

type enemyconfigs []enemyconfig

func createGoblin() *enemy {
	return nil
}

func jsonToEnemyConfig(json string) enemyconfigs {
	return nil
}

func (e enemyconfigs) initEnemies() enemies {

	// TODO: json unmarshall and generate enemyconfigs []enemyconfig
	enemies := make(enemies, 0)
	for _, v := range e {
		switch v.enemyType {
		case "goblin":
			enemies = append(enemies, *createGoblin())
		}
	}
	return enemies
}

func (e enemies) getEnemiesToPop(currentFrame int64) enemies {
	// TODO: implement
	return nil
}

func (e enemies) spawn() {
	// TODO: implement
}
