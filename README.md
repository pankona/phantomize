
kanban (powered by zenhub) here -> <a href="https://app.zenhub.com/workspace/o/pankona/phantomize/boards?repos=90622123"><img src="https://raw.githubusercontent.com/ZenHubIO/support/master/zenhub-badge.png"></a>

# phantomize

a game like free place tower diffence written in Go

## how to build

* place this sources under GOPATH like:

`$GOPATH/src/github.com/pankona/phantomize`

* install dependencies with [glide](https://github.com/Masterminds/glide). (if you don't have glide, install it in advance.)

`$ glide install`

* build phantomize itself

`$ go build -tags=release`

* executable `phantomize` will be out.

## how to play

purpose of this game is defeating all enemies using 3 types of player unit.

* enemies will be spawned automatically. they will move to player unit to destroy them.
* player unit also moves to enemies to kill them all automatically.

you can summon players with following way:
* choose a unit from control panel (control panel to choose unit is placed on right-bottom side)
* click on field to summon the choosen unit
  * up to 2 units are available to summon in parallel.
  * every summoning spends money a little. remaining money is displayed on left-bottom side.

## tips

* if an enemy is down, it gains your money a little.
* summoned unit can be "recalled". recall gains money that is used for summoning the unit.
  * to recall, click the summoned unit.
  * then press "RECALL" button on unit icon on charactor information, that is placed on center-bottom.

## limitation

* only linux is supported. this doesn't work on osx or windows.

# credits

sounds: [maoudamashii.jokersounds.com](http://maoudamashii.jokersounds.com/)
