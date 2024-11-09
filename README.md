# Yet another Traveller NPC generator

This is a simple NPC generator for the Traveller RPG. 
It follows rules described in [this article](https://greatdungeonnorth.blogspot.com/2020/02/stock-in-trade-typical-traveller-npcs.html). 
I'll briefly describe the rules here.

## Ability scores

Instead of randomly generating ability scores, this generator uses several standard arrays:

| Citizen Category | Average Score  |      Characteristic Array |
|------------------|:--------------:|--------------------------:|
| Below Average    |       6        |          8, 7, 6, 6, 5, 4 |
| Average          |       7        |          9, 8, 7, 7, 6, 5 |
| Above Average    |       8        |         10, 9, 8, 8, 7, 6 |
| Exceptional      |       9        |        11, 10, 9, 9, 8, 7 |

## Skills

According to the previous experience of the NPC, the generator will assign a number of skill points according to this table, "Average Skill Levels by Term":

<table><caption>Average Skill Levels by Experience</caption> 
<thead>
<tr> <th rowspan="2">Experience</th> <th colspan="4">Number of Skills by Skill Level</th> </tr>
<tr> <th>3</th> <th>2</th> <th>1</th> <th>0</th> </tr>
</thead> 
<tbody>
<tr> <td>Recruit</td> <td>0</td> <td>0</td> <td>0</td> <td>4</td> </tr>
<tr> <td>Rookie</td> <td>0</td> <td>0</td> <td>2</td> <td>4</td> </tr>
</tbody> 
<tbody>
<tr> <td>Intermediate</td> <td>0</td> <td>1</td> <td>2</td> <td>4</td> </tr>
<tr> <td>Regular</td> <td>0</td> <td>2</td> <td>2</td> <td>5</td> </tr>
</tbody>
<tbody>
<tr> <td>Veteran</td> <td>0</td> <td>3</td> <td>2</td> <td>5</td> </tr>
<tr> <td>Elite</td> <td>1</td> <td>2</td> <td>3</td> <td>6</td> </tr>
</tbody> 
</table>

## Build & Run

To build the project execute the following command:

```bash
make build
```

This will generate this binary `out/generate-npc`. 

To run the project execute the following command:

```bash
./out/generate-npc --help
```

These are the options for the command:

<!--
Usage:
  generate-npc [OPTIONS]

Application Options:
  -c, --category=[0|1|2|3]                                                                                                   Citizen Category: 0-Below average, 1-Average, 2-Above Average, 3-Exceptional (default: 1)
  -e, --experience=[0|1|2|3|4|5]                                                                                             Experience: 0-Recruit, 1-Rookie, 2-Intermediate, 3-Regular, 4-Veteran, 5-Elite (default: 3)
  -r, --role=[pilot|navigator|engineer|steward|medic|marine|gunner|scout|technician|leader|diplomat|entertainer|trader|thug] Crew role in a starship
  -g, --gender=[female|male|unspecified]                                                                                     Gender of the NPC (default: unspecified)
  -d, --debug                                                                                                                Enable debug mode

Help Options:
  -h, --help                                                                                                                 Show this help message
-->

| Option | Long Option  | Description                                                                                                                                                     | Default     |
|--------|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| -c     | --category   | Citizen Category: 0-Below average, 1-Average, 2-Above Average, 3-Exceptional                                                                                    | 1           |
| -e     | --experience | Experience: 0-Recruit, 1-Rookie, 2-Intermediate, 3-Regular, 4-Veteran, 5-Elite                                                                                  | 3           |
| -r     | --role       | Crew role in a starship: pilot, navigator , engineer , steward , medic , marine , gunner , scout , technician , leader , diplomat , entertainer , trader , thug | required    |
| -g     |  --gender    | Gender of the NPC: female, male, unspecified                                                                                                                    | unspecified |
| -d     | --debug      | Enable debug mode                                                                                                                                               | false       |
| -h     | --help       | Show the help message                                                                                                                                           |             |

<img src="demo/demo.gif" alt="Demo" />
