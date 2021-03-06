# Sandpiles

Imagine an infinite 2-D checkerboard on which you place piles of coins of various heights on some of the squares (at most 1 pile per square). Consider the following toppling operation:

```
topple(r, c): if square (r, c) has  4 coins on it, move 1 coin from (r, c) to each of the 4 neighbors of (r, c) (diagonal neighbors don’t count, only north, south, east, and west). If square (r, c) has < 4 coins, do nothing.
```

A configuration of coins is said to be stable if all squares have < 4 coins on them. If we repeatedly topple until we can’t topple any more, we’ll end up at a stable configuration. Somewhat surprisingly, you can show than the order of the topples that you do won’t a↵ect the final stable configuration that you end up in, given a particular starting point.

For example, if you start with a pile of 10000 coins on a single square, and no coins elsewhere, you will end up with the following configuration:

![Imgur](http://i.imgur.com/K1l0IZb.png)
