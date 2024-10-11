# Mindless Stitchcraft (2024)

I like to knit, crochet, cross-stitch, make friendship bracelets/macrame,
etc. These are relaxing hands-on activities.

I prefer patterns that are simple to remember yet look visually interesting.
Surprisingly, you can do a lot with patterns that involve the same motif
of stitches over and over (e.g. "knit 2, purl 3") depending on the 
fabric width and other factors.

This program lets me explore and preview patterns before putting them to 
fabric. This way I can rule out patterns that would be too long to
remember easily, or uninspiring designs.

This repo also is a way for me to practice some things I've learned about
recently:

- Golang
- Unit testing
- Docker

## Docker

This repo is intended to be run in a Docker container for consistent 
environment.

### Usage

```
docker container run --rm -it ptrgags/mindless-stitchcraft COMMAND <ARGS>
```

The commands and options are detailed in the [Pattern Types](#pattern-types)
section below.

### Building the container

```
docker image build -f docker/Dockerfile -t ptrgags/mindless-stitchcraft .
```

## Running Manually

If you don't want to use Docker and have Golang installed, you can run

```
go run . COMMAND <ARGS>
```

Or build it as an excutable

```
go build -o mindless-stitchcraft main.go
./mindless-stitchcraft COMMAND <ARGS>
```

## Pattern Types

Below is a list of the pattern types currently available in this repo, and
a description of the options.

### Knitting: Zigzag (2024)

This method takes a simple motif of knits `v` and purls `-` and repeats
them over and over (knitting flat). At the end of the row, flip the fabric
over and keep going.

This method is inspired by something I saw in passing. I was reading the Bridges
paper, ["Predicting Planned Pooling Patterns"](https://archive.bridgesmathart.org/2024/bridges2024-361.html#gsc.tab=0) 
and found a reference to [Sequence Knitting](https://ceceliacampochiaro.com/sequence-knitting/)
by Cecilia Campochiaro. I haven't actually read that book, but just from the
main page and other videos about the technique (see 
[This YouTube video](https://youtu.be/uxrPQXZmRlQ?si=porLJlSq0GwmEFJr) for 
example) I got a rough understanding of the technique. I began exploring the
math on my own. I think this pattern corresponds to the "serpentine" method.

Usage:

```
mindless-stitchcraft knit-zigzag FABRIC_WIDTH MOTIF
```

Where

| Argument | Description |
| --- | --- |
| `FABRIC_WIDTH` | How many stitches wide is the fabric? |
| `MOTIF` | A string of knits (`v`) and purls (`-`) listed in the order you stitch them. E.g. `vv---v--` means "k2 p3 k1 p2" |

The output is a chart which shows how the front of the work will look when stitched. Stitch from the
bottom right and zigzag.

Examples:

This first example has a motif that repeats many times:

```
mindless-stitchcraft knit-zigzag 10 "v--"

Output:
v-vv-vv-vv
-v--v--v--
-vv-vv-vv- # --> ...
--v--v--v- # <--
vv-vv-vv-v # --> second row is stitched on the reverse side
v--v--v--v # <-- first row is stitched from right to left
```

Changing the fabric width can result in a very different pattern. E.g. here's
a motif of length 3 that evenly divides 9. This pattern repeats after only
2 rows!

```
mindless-stitchcraft knit-zigzag 9 "v--"

-vv-vv-vv
--v--v--v
```

### Knitting: Sync (2024)

On closer inspection of one of the sample images from [Sequence Knitting](https://ceceliacampochiaro.com/sequence-knitting/), I realized that the technique shown was simpler:
at the start of a new row, restart at the beginning of the motif instead
of continuing. 

This will produce 2-row patterns in general (one on each side of the fabric).
To add more variation, I allow specifying multiple motifs. The program will
cycle through these from row to row.

I call this mode `knit-sync` because the act of restarting the periodic pattern
at the end of the row reminds me of oscillator sync on synthesizers.

Usage:

```
mindless-stitchcraft knit-sync FABRIC_WIDTH MOTIF [MOTIF ...]
```

| Argument | Description |
| --- | --- |
| `FABRIC_WIDTH` | How many stitches wide is the fabric? |
| `MOTIF` | A string of knits (`v`) and purls (`-`) listed in the order you stitch them. E.g. `vv---v--` means "k2 p3 k1 p2" |

Examples:

Simple example with a single motif:

```
mindless-stitchcraft knit-sync 10 "v--"

-vv-vv-vv-  --> motif (reverse)
v--v--v--v  <-- motif
```

An example that uses multiple motifs to make a more interesting pattern:

```
mindless-stitchcraft knit-sync 10 "vvv--" "vvv--" "v----"

-vvvv-vvvv  --> third motif (reverse)
--vvv--vvv  <-- second motif
---vv---vv  --> first motif (reverse)
----v----v  <-- third motif
---vv---vv  --> second motif (reverse)
--vvv--vvv  <-- first motif
```

### Friendship Bracelets: Repeat (2024)

I took the concept of repeating a motif and applied it to friendship
bracelets. 

Friendship bracelets take parallel strands of embroidery floss and
tie pairs of strands together in staggered rows like this:

```
x x x x
 x x x
x x x x
 x x x
```

Taking a motif of knots `/ > \`, we can fit them into the
slots from left to right like this:

```
/ > \ /
 > \ /
> \ / >
 \ / >
\ / > \
 / > \
```

Usage:

```
mindless-stitchcraft bracelet-repeat STRAND_COUNT MOTIF
```

Where:

| Argument | Description |
| --- | --- |
| `STRAND_LABELS` | A string of Unicode characters that represents the colors of each strand. E.g. `ABCD` represents 4 strands labeled A, B, C, D. Labels can be repeated (e.g. `ABCCBA`) to indicate multiple strands of the same color |
| `MOTIF` | A string of knots (see below) that represents the pattern |

To make a short pattern, have the length of a motif divide the length of
two rows of knots. Since strands are knotted in pairs and every other row
is staggered, this means the motif should be $\text{len(STRANDS)} - 1$
for the shortest patterns. That said, any non-zero length will work!

Knots (and some useful properties of each one)

| Symbol | Knot Type | Visible color? | Swaps strands? |
| ------ | --------- | --- | --- |
|  `\`   | Forward knot | Left strand | Yes |
|  `/`   | Backward knot | Right strand | Yes |
|  `>`   | Forward-backward knot | Left strand | No |
|  `<`   | Backward-forward knot | Right strand | No |



The output is two versions of the pattern. The uncolored pattern lists the knots you make, while the colored pattern gives a preview of what
the result will look like.

Example: the classic chevron pattern:

```
mindless-stitchcraft bracelet-repeat '.ahBBha.' '\\//\//'

Uncolored pattern:
\ \ / /
 \ / / 
Colored pattern:
. a h B B h a .
| | | | | | | |
 .   h   h   . 
a  .   h   .  a
 a   .   .   a 
B  a   .   a  B
 B   a   a   B 
h  B   a   B  h
 h   B   B   h 
.  h   B   h  .
 .   h   h   . 
a  .   h   .  a
 a   .   .   a 
B  a   .   a  B
 B   a   a   B 
h  B   a   B  h
 h   B   B   h 
.  h   B   h  .
| | | | | | | |
. a h B B h a .
```

One neat trick I realized by accident:
Using emoji helps the distinct colors stand out
(albeit the spacing gets wonky).

```
mindless-stitchcraft bracelet-repeat '🔴 🟢 🔵 🟡 🟡 🔵 🟢 🔴' '\\//\//'

Uncolored pattern:
\ \ / /
 \ / / 
Colored pattern:
🔴 🟢 🔵 🟡 🟡 🔵 🟢 🔴
| | | | | | | |
 🔴   🔵   🔵   🔴 
🟢  🔴   🔵   🔴  🟢
 🟢   🔴   🔴   🟢 
🟡  🟢   🔴   🟢  🟡
 🟡   🟢   🟢   🟡 
🔵  🟡   🟢   🟡  🔵
 🔵   🟡   🟡   🔵 
🔴  🔵   🟡   🔵  🔴
 🔴   🔵   🔵   🔴 
🟢  🔴   🔵   🔴  🟢
 🟢   🔴   🔴   🟢 
🟡  🟢   🔴   🟢  🟡
 🟡   🟢   🟢   🟡 
🔵  🟡   🟢   🟡  🔵
 🔵   🟡   🟡   🔵 
🔴  🔵   🟡   🔵  🔴
| | | | | | | |
🔴 🟢 🔵 🟡 🟡 🔵 🟢 🔴
```