# The game design of Below the Keep

This file describes the game: the style, a run, the rules, the screen, and the content. Each fact cites its decision in `docs/decisions.md`.

An item with a `P-` id is a proposal of the agent. The owner accepted this file and each proposal in it on 2026-10-16 (D-71).

The numbers of the proposals are a first guess. A later gate for the balance can change them.

## 1. The style (D-36)

- The game is a classic roguelike. It has one hero, a top-down tile map, and turns.
- The world is classic fantasy: a dungeon below a ruined keep. It has goblins, potions, and an artifact at the bottom.
- The text is short, and it has a light tone with dry humor.

## 2. A run (D-37, D-40, D-50, D-67)

- A run has 5 floors. Each floor is harder than the floor before it.
- The player sees the whole floor from the start: each tile, each item, and each monster (D-67).
- On each floor, the hero walks to the items, fights the monsters, and goes down the stairs.
- The run ends with a win when the hero picks up the artifact on floor 5.
- The run ends with a loss when the hero dies.
- The game has no save, so a run ends when the player quits the game (D-50).
- Nothing carries over from one run to the next.

**P-1. The start and the end of a floor.** The hero starts each floor on a random open tile. The stairs down are on the open tile with the longest path from the start. On floor 5, the artifact takes the place of the stairs, and the Warden stands next to it.

## 3. The turn (D-38, D-55, D-69)

- On each turn, the hero acts one time. Then each monster acts one time.
- All actors have the same speed.
- An actor moves in 4 directions: up, down, left, and right (D-69).
- A move into a monster is an attack.
- A move into a wall is a rule outcome, not an error (D-55). The game shows a message.
- The hero has these actions: move, attack, wait, pick up, drop, equip, use a potion or a scroll, and go down the stairs (D-38).

**P-2. A blocked move takes no turn.** When the hero walks into a wall, the monsters do not act.

**P-3. The order of the monsters.** The monsters act in the order of the monster list of the floor. A new monster goes at the end of the list.

## 4. The stats (D-39, D-41, D-54)

Each actor has five stats:

| Stat | The job of the stat |
|---|---|
| hit points | When the hit points reach 0, the actor dies. |
| strength | A bonus to each damage roll. |
| accuracy | The base of the hit chance. |
| armor | It lowers the damage of each hit. |
| evasion | It lowers the hit chance of an attacker. |

The weapon of the hero gives the damage range (D-41). The equipped armor gives the armor value.

**P-4. The damage range of a monster.** A monster has no weapon. Its kind holds its damage range and its armor value.

**P-5. The formulas.** The play source (D-53) makes both rolls.

- The hit chance, in percent, is the accuracy of the attacker minus the evasion of the defender. The chance is min 5 and max 95.
- The damage is a roll in the damage range, plus the strength of the attacker, minus the armor of the defender. The damage of a hit is min 1.

**P-6. The hero at the start.** Hit points 20, strength 0, accuracy 80, armor 0, and evasion 10. The hero starts with a dagger and leather armor, both equipped, and one healing potion.

## 5. Levels and experience points (D-40, D-41)

- Only kills give experience points. Each monster kind gives a fixed number, and a stronger kind gives more.
- At each new level, a menu opens. The player picks one stat. The game stores no unspent points.
- A full run gives about 10 levels.

**P-7. The level menu.**

- Hit points: +5 to the max hit points, and +5 to the current hit points.
- Strength: +1.
- Evasion: +5.

**P-8. The level curve.** The hero reaches level n + 1 at 5 × n × (n + 1) experience points in total. So level 2 needs 10 points, level 3 needs 30, and level 11 needs 550. The panel shows the points and the next goal, for example "XP 41/60".

## 6. The recovery of hit points (D-39)

- The hero gets back hit points over time, and a potion heals at once.
- The hero needs no food.

**P-9. The regen.** The hero gets back 1 hit point every 10 turns. Monsters do not regen.

## 7. Items (D-38, D-39, D-53, D-54)

- The map source puts items on each floor.
- Some monsters drop an item when they die. The play source makes the drop roll.

**P-10. The pack.** The pack holds max 10 items, with the letters a to j. The equipped weapon and the equipped armor stay in the pack, with a mark. When the pack is full, the hero can not pick up an item.

**P-11. The item kinds.**

| Item | Glyph | Effect | First floor |
|---|---|---|---|
| dagger | `)` | damage 1 to 4 | 1 |
| short sword | `)` | damage 2 to 6 | 1 |
| mace | `)` | damage 3 to 7 | 2 |
| long sword | `)` | damage 3 to 9 | 3 |
| war axe | `)` | damage 4 to 11 | 4 |
| leather armor | `[` | armor 1 | 1 |
| chain mail | `[` | armor 2 | 2 |
| plate armor | `[` | armor 4 | 4 |
| healing potion | `!` | heals 10 hit points | 1 |
| greater healing potion | `!` | heals 25 hit points | 3 |
| teleport scroll | `?` | moves the hero to a random open tile | 2 |
| fire scroll | `?` | 8 damage to each monster within 3 tiles | 3 |
| the Keepstone (the artifact) | `♦` | when the hero picks it up, the run ends with a win | 5 |

**P-12. The items on a floor.** Each floor gets 5 items. Each item comes from the kinds of that floor or of a floor above it.

## 8. Monsters (D-38, D-54)

- A monster walks at random until it sees the hero. Then it walks to the hero and attacks.

**P-13. The sight of a monster.** A monster sees the hero when the hero is max 8 tiles away, and a straight line between them crosses no wall. The player sees the whole floor (D-67), but a monster does not.

**P-14. The chase.** On each turn, the monster takes the first step of a shortest path to the hero. A breadth-first search over the open tiles, in 4 directions, finds the path. When another monster blocks the step, the monster waits.

**P-15. The wander.** The monster moves to a random open tile next to it. When it has no open tile next to it, it waits.

**P-16. The monster kinds.** Two new kinds come on each floor. The Warden guards the artifact on floor 5.

| Monster | Glyph | Floor | Hit points | Damage | Accuracy | Armor | Evasion | XP | Drop chance |
|---|---|---|---|---|---|---|---|---|---|
| rat | `r` | 1 | 4 | 1 to 2 | 60 | 0 | 10 | 2 | 0% |
| kobold | `k` | 1 | 6 | 1 to 3 | 65 | 0 | 10 | 4 | 20% |
| goblin | `g` | 2 | 10 | 1 to 4 | 70 | 1 | 10 | 6 | 20% |
| cave bat | `b` | 2 | 6 | 1 to 3 | 75 | 0 | 30 | 6 | 0% |
| orc | `o` | 3 | 16 | 2 to 6 | 70 | 2 | 5 | 12 | 25% |
| giant spider | `s` | 3 | 12 | 2 to 5 | 80 | 1 | 20 | 12 | 0% |
| skeleton | `z` | 4 | 20 | 3 to 7 | 75 | 3 | 5 | 18 | 25% |
| troll | `T` | 4 | 30 | 3 to 8 | 65 | 2 | 0 | 22 | 25% |
| ogre | `O` | 5 | 36 | 4 to 10 | 65 | 3 | 0 | 25 | 30% |
| wraith | `W` | 5 | 24 | 3 to 8 | 85 | 1 | 30 | 25 | 0% |
| the Warden | `&` | 5 | 60 | 5 to 12 | 80 | 4 | 10 | 50 | 0% |

**P-17. The monsters on a floor.** Floor 1 has 6 monsters, and each next floor has 2 more. Each monster comes from the kinds of that floor or of the floor above it.

## 9. The map (D-47, D-53, D-68)

- The whole map is always visible, and it has a fixed size (D-47).
- The map source makes each floor as a cave, with a cellular automaton (D-68).
- The code checks that each open part of the cave connects (D-68).

**P-18. The size of the map.** The map is 80 tiles wide and 30 tiles high. Each border tile is a wall.

**P-19. The cave maker.**

1. Make each inner tile a wall, with a chance of 45%.
2. Do 5 rounds of this rule: a tile becomes a wall when 5 or more of its 8 neighbors are walls. Else it becomes open.
3. Find the biggest open area that connects in 4 directions. Fill each other open tile with wall.
4. When the open area is less than 40% of the map, make the floor again with the next numbers of the map source.

## 10. The screen (D-43 to D-49)

The layout (D-46): the map on the left, the panel on the right, and the messages at the bottom.

```
┌──────────────── the map: 80 x 30 ─────────────────┬─ the panel: 24 ─┐
│                                                   │ Lvl 3  XP 41/60 │
│        #########                                  │ HP  12/20       │
│      ###·····r·###                                │ Str 1   Acc 80  │
│      #·····@·····##                               │ Arm 1   Eva 15  │
│      ##··!······>#                                │ Floor 2  T 312  │
│        ##########                                 │ a) dagger *     │
│                                                   │ b) leather *    │
├───────────────────────────────────────────────────┴─────────────────┤
│ The rat bites you. It looks pleased with itself.                    │
│ You hit the rat.                                                    │
└─────────────────────────────────────────────────────────────────────┘
```

**P-20. The size of the window.** One border column on each side and between the parts: 1 + 80 + 1 + 24 + 1 = 107 columns. The rows: 1 + 30 + 1 + 4 + 1 = 37. So the minimum window is 107 columns by 37 rows. When the window is smaller, the game pauses and shows the size that it needs (D-47).

**P-21. The panel.** The level, the experience points, the hit points, the five stats, the floor, the turn, and the pack. The panel has room for the 10 lines of the pack.

**P-22. The message lines.** 4 lines, with the newest line at the bottom.

**P-23. The glyph list (D-43, D-48).**

| Thing | Glyph | Width class |
|---|---|---|
| hero | `@` | Na |
| open floor | `·` (U+00B7) | A |
| wall | `#` | Na |
| stairs down | `>` | Na |
| artifact | `♦` (U+2666) | N |
| items | `)` `[` `!` `?` | Na |
| monsters | the letters of P-16, and `&` | Na |
| borders | `─ │ ┌ ┐ └ ┘ ├ ┤ ┬ ┴` (U+2500 block) | A |

Per the Unicode 13.0 data in Python (read 2026-10-16), "Na" and "N" are one column. "A" is ambiguous: an option of the terminal can show it two columns wide. D-48 accepts this risk.

**P-24. The colors (D-44, D-49).** Each cell gets a background color, so the game looks the same on each theme. The numbers are indexes of the 256-color palette.

| Thing | Foreground | Background |
|---|---|---|
| each cell | none | 234 (very dark grey) |
| wall | 244 (grey) | 234 |
| open floor | 238 (dim grey) | 234 |
| hero | 231 (white) | 234 |
| stairs and artifact | 220 (gold) | 234 |
| items | 229 (pale yellow) | 234 |
| monsters | a color in each kind | 234 |
| text | 252 (light grey) | 234 |
| low hit points (25% or less) | 196 (red) | 234 |

**P-25. The keys (D-38, D-69).**

| Key | Action |
|---|---|
| the arrow keys, or `h` `j` `k` `l` | move, or attack |
| `.` | wait one turn |
| `g` | pick up |
| `d`, then a letter | drop |
| `e`, then a letter | equip |
| `u`, then a letter | use a potion or a scroll |
| `>` | go down the stairs |
| `?` | help |
| `Q` | quit, after a yes or no question |
| `1`, `2`, `3` | a choice in the level menu |
| Esc | cancel a letter choice |

**P-26. The screens.** The title screen (the name, a new run, and quit), the play screen, the help screen, the level menu as a box over the map, and the end screen. The end screen shows the win or the cause of death, the floor, the level, the turns, and the seed (D-53).

## 11. The random sources (D-52, D-53)

- One seed makes two sources of `math/rand/v2`: the map source and the play source.
- The map source makes the cave of each floor, the items on it, and the monsters on it. So one seed gives the same floors.
- The play source makes each roll during play: the hits, the damage, the drops, and the wander.
- A command-line flag sets the seed. With no flag, the game picks a random seed.

## 12. The content (D-48, D-51, D-57, D-61)

- The content is in JSON files, and `embed` puts them in the binary (D-51).
- The files are in `internal/`, in the folder of the package that embeds them (D-57). The owner writes them, and the guard blocks the agent (D-61).
- The content holds the monster kinds, the item kinds, and the text of the messages.
- A test checks each glyph of the content against the glyph list (D-48).

**P-27. The tone of the text.** Short lines with dry humor, for example:

- "The rat bites you. It looks pleased with itself."
- "You drink the potion. It tastes of old socks, but it works."
- "You die on floor 3. The goblins take your boots."

## 13. Not in the first version

These parts are out of the first version. A later gate can add one, after an owner decision.

- A save and a load (D-50).
- Ranged attacks and spells (D-38).
- Hunger (D-39).
- A score list and unlocks (D-37).
- Fog, and a line of sight for the player (D-67).

**P-28. More parts out of the first version.** Doors, traps, shops, and gold.
