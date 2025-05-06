# Reflection

We'll now look reflection, which apparently means "the ability of a program to examine its own structure, particularly types - a form of meta-programming".

---

`Reflect` is a little abstract I would say. 

But apparently you can turn a struct into a slice of fields like so: `val := reflect.ValueOf(x)`. Then you can do `val.Field(0)` to get the specific field.

---

Note: if you have a variable that you want to compare to multiple values, it is generally best to use a switch statement for readability. 

---

Ok so this lesson has been so uninteresting I think. I'm not sure why but it feels pretty contrived. I'm sure I'll have use for reflection at some point, but right now it just feels "besides the point". I went through the motions, but I doubt I will remember much from this lesson. But I shoulld probably revisit it sometime.

Note: I did not finish this lesson. I stopped at the "channel implementation". Revisit the lesson when ready.