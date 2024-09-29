# Mindless Stitchcraft (2024)

I like to knit, crochet, cross-stitch, make friendship bracelets/macrame as
its

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

If you don't want to use Docker, you can use these steps:

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

### Knitting: Zigzag

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

This first one has a motif that repeats many times

```
mindless-stitchcraft 10 "v--"

Output:
v-vv-vv-vv
-v--v--v--
-vv-vv-vv- # --> ...
--v--v--v- # <--
vv-vv-vv-v # --> second row is stitched on the reverse side
v--v--v--v # <-- first row is stitched from right to left
```

Changing the fabric width can result in a very different pattern. E.g. here's
a motif that lines up with 

```
mindless-stitchcraft 9 "v--"

-vv-vv-vv
--v--v--v
```

#### Zigzag Math

Some details for the mathematically inclined:

- Let $w$ be the given fabric width
- Let $m$ be the length of the motif in stitches
- Let's temporarily ignore the zig-zag nature of flat knitting. Assume every row
is knit in the same direction for now.
- The pattern will repeat after $h$ rows. This means the pattern will be
short when the motif evenly divides the fabric width (or vice-versa). And when
the two lengths are coprime, there are exactly $m$ rows.

$$h=\frac{m}{\text{gcd}(m, w)}$$

- Interestingly, the area is the least common multiple of the two lengths.

$$A=\frac{wm}{\text{gcd}(m, w)} = \text{lcm}(m, w)$$

- As for the number of motif repeats, divide $\frac{a}{m}$ and get:

$$\frac{w}{\text{gcd}(m,w)}$$

- Now, let's account for the zig-zag stitching order. In general, the stitches
on the front will look different on the back due to swapping knits and purls.
So this means we have to _double_ the number of rows/area/motif repeats.

### Friendship Bracelets: Repeat

```
mindless-stitchcraft knit-zigzag STRAND_COUNT MOTIF
```

Where:

| Argument | Description |
| --- | --- |
| `STRAND_COUNT` | The number of strands in the friendship bracelet |
| `MOTIF` | A string of knots (see below) that represents the pattern |

Knots (and some useful properties of each one)

| Symbol | Knot Type | Visible color? | Swaps strands? |
| ------ | --------- | --- | --- |
|  `\`   | Forward knot | Left strand | Yes |
|  `/`   | Backward knot | Right strand | Yes |
|  `>`   | Forward-backward knot | Left strand | No |
|  `<`   | Backward-forward knot | Right strand | No |
