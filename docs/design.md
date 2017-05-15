# design 

## scene transition

![scene transition](http://www.plantuml.com/plantuml/svg/RP11Si8m34NtFeMMfGmNOD6Bb2wKe2IQ1ZjZov3S7bjs8D1cuOc_lkIl9omZGMZ94eX76rZOMMC6FXDqpadZPcE-Ftz0luFdVf335qZkCtg-w1Uo4OuWE1dzrMBI1ngdf0_k1k9W6d4nAgtrV_lMeOrdwjZsoH04lPY7a5jxl52gO4o3dokSB7P2_j5hrBO-vod4KL9NLjsLlAaAR4UANm1xPFjQtSInDCl9Vd07awShknXomfxWXW7QaT3JUMSAfnHUrrTc3RQ_BtBA43BvgkqLXRB8bxBsUjjIL8nrnaPFePe59cubUjfLKnswvZy2pSdB5h1nKNIP4dzZsR-3DKgoa2iHolsF4L2UU2j1pV16DMlvwJi0)

## game scene

### outline

![outline1](./images/outline1.png)

### 1. field

* stocked units can place anywehere on this pane
* enemies spawn on right edge of screen, and toword "you"
* "you" must be placed to start game. in other word, game starts at placing "you".
* ally units (that are summoned by player) will fight automatically.
* if ally units are tapped, they will back to 3. stocked units
  * ready to re-summon

### 2. resources

* money and number of units ( placed / generated )
* money is used to generate units
* increase when enemies are killed

### 3. stocked units

* units can be generated at "unit generation panel"
* generated units will be listed here
* tap on a generated unit, change mode into "summon mode"
  * then tap on 1. field, units will be placed there (summon)
* unit that is in generating will also be listed here
  * if tap them, it means canceling of generating
  * money will be back if unit generation is canceled

### 4. unit generation panel

* list of units they can start to generate
* unit generation takes short time.
* when unit generation completed, it will be listed on 3. stocked units
