# Concurrency

Note about anonymous functions: These have lexical scope. This means that all the variables available where the function is declared are available inside the function, which is nice. So in this case, the `url` variable is available inside the func.

--- 

Note: we get a `fatal error` because several goroutines are trying to write to the map at once. This is called a ***race condition***. To spot race conditions, run tests with a `-race` flag. 

So how do we solve this? We use **channels**. Let each routine put a value into the channel, then take the values from the channel one by one.

---

Yeah that is pretty much it for this lesson. A good quote for Donald Knuth is:

>Premature optimization is the root of all evil.

Worth pondering!