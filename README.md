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

<table><caption>Average Skill Levels by Term</caption> <thead>
<tr> <th rowspan="2">Experience</th><th rowspan="2">Terms</th> <th colspan="4">Number of Skills by Skill Level</th> </tr>
<tr> <th>3</th> <th>2</th> <th>1</th> <th>0</th> </tr>
</thead> <tbody>
<tr> <td>Recruit</td><td>0</td> <td>0</td> <td>0</td> <td>0</td> <td>4</td> </tr>
<tr> <td>Rookie</td> <td>1</td> <td>0</td> <td>0</td> <td>2</td> <td>4</td> </tr>
</tbody> <tbody>
<tr> <td>Intermediate</td><td>2</td> <td>0</td> <td>1</td> <td>2</td> <td>4</td> </tr>
<tr> <td>Regular</td><td>3</td> <td>0</td> <td>2</td> <td>2</td> <td>5</td> </tr>
</tbody><tbody>
<tr> <td>Veteran</td><td>4</td> <td>0</td> <td>3</td> <td>2</td> <td>5</td> </tr>
<tr> <td>Elite</td><td>5</td> <td>1</td> <td>2</td> <td>3</td> <td>6</td> </tr>
</tbody> </table>
