# Property-based tests

I've implemented the "roman numeral kata" in the "roman" package. But not really using TDD. That's what we'll be trying now.

Remember: *"thin vertical slices of useful functinality, then iterate..."*

---

Ok, so always SCALE BACK. You don't need to be "intelligent". Just try it out. First write the `ConvertToRoman` that takes an int and outputs a string. Doesn't need to be more complicated than that. You'll refactor later.

Ok, after a while we have three test cases, the numbers 1, 2, 3. We realize that we're just adding `"I"` for everyone one of them. We know that we might have to do other stuff later. But for now, let's do a stringbuilder that adds an I in a for-loop.

... And we're continuing to write stupid code etc etc...

But then, I realize I think I thought in the same way that the book shows. You iterate starting from the number, down to 0. And you look at what you have. Then you subtract what you're adding to them stringbuilder from your arabic. 

If you add V for 5 and IV for 4 and I for everything else, works for 1-8. I wouldn't have thought of that actually. That's neat. And this actually works for up to 39.

``` 
	for arabic > 0 {
		switch {
		case arabic > 8:
			sb.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			sb.WriteString("V")
			arabic -= 5
		case arabic > 3:
			sb.WriteString("IV")
			arabic -= 4
		default:
			sb.WriteString("I")
			arabic -= 1
		}
	}
```
... but it's not totally sustainable is it.

--- 

Now comes a lesson: Some times you should think extra about switch statements. This comes from OOP. When you manually encode behavior, that's not very OO. You should think about if you shouldn't really abstract that away into interfaces, classes and methods. You decouple the data from the behavior.

So what do we get then? Well, we see that in each case in the switch statement, we have the same logic. If we define a type `RomanNumeral` that has fields Value and Symbol, and we create a slice of all the roman numerals (so far). Then we can in our function iterate through our list of numerals checking if our arabic is >= than the value we're currently looking at. If so, subtract that numeral.Value from our arabic, and write numeral.Symbol to our stringbuilder. Et voila.

--- 

Now let's do the rest of the symbols. Let's see if it holds up still. It does. Turns out this was a pretty neat solution. In my own solution I didn't treat IV as a separate case, but that forced me to be more complex in other places.

---

Now let's get into parsing roman numerals.

After building a dumb switch statement, we can see that we need to, in the same way as before, *iterate* through the thing and build *something* gradually. Let's start by just counting the characters in the roman numeral.

---

So there is some new stuff in this solution. I don't think I've every used a while-loop in Go so far. What you do is: you go through the list of numerals. For each numerals, enter a while loop with the condition "does the roman have the prefix numeral.Symbol"? If so, add numeral.Value to total, and remove the prefix from roman.

--- 

Next, I'm trying to implement a recursive solution to the problem. Took a little bit of thinking. But remember: work out the base case first. In this case, it is when roman only has one numeral left (len(roman) == len(numeral.Symbol)). Then you return the value. At the bottom of your function I just did `return arabic + Recursive(newRoman)`.

Skrev benchmarktester for them both. Of course Recursive is slower by about four times.

## Property based testing

We've come to realize the *domain* of *Roman numerals* have certain rules. So far, we've only tested concrete examples (example based tests), but what if we could concretize the rules we've found, and test them against the code? So what rules have we found?

- can't have more than 3 consecutive symbols
- only I, X and C can be subtractors
- convert and then deconverting using our functions should return us to the original number.

We can write tests for these.

--- 

We're using `quick.Check` that takes an anon function. We realize then that a flaw in our code: it can't handle negative ints or numerals over 3999. This suggests to use that maybe `int` isn't the right type? We can print out with `t.Log` btw if we want to see what values aren't making it. Then you run the test with `go test -v` (verbose flag i guess).
Using `unit16` gets us closer to what we want at least.

---

## Wrapping up

### TDD practice

*"Did this thought of writing code that converts 1984 to MCMLXXXIV feel intimidating. It did to me too. The trick is always 'take small steps'. Not in any part of this process did we take any huge step. Always have this approach."*

The skill is knowing *how* to split up work, which comes with practice, and specifically TDD practice.


### Property based tests

Excellent tool if you can figure out the rules for your domain. The functionality you need is in the stdlib. These types of tests are potentially niche though.

