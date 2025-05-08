# Select

Here we will build a lol function that GETs two different websites, and returns the URL of the one of them that returned first.

Testing HTTP stuff is so common that the Go standard lib has tools for it. We don't want tests that are slow or flaky. We also want to test edge cases.

The `net/http/httptest` package provides mock servers.

---

`defer` is good to use. Often for freeing up resources. It is mostly used for readability.

---

Now we're realizing we should be using goroutines. When you want to "listen" to multiple routines and see which one returns first, you can use a `select` statement.

I'm slightly weirded out by the syntax here. But I guess the select statement runs the functions in the cases, that is multiple instances of `ping`. For some reason, it returns a channel. That's' how do it i guess. If you were to return struct{} then that wouldn't have anything to do with the goroutine, and so you wouldn't be able to do that - you wouldn't get the info you want.

So if you return a channel, the select statement waits for the instance that finishes and returns a channel, and then runs that case. And when you don't care about *what is in the channel*, you can make a channel of `struct{}`, because that is the smallest data type.

NOTE: Now i will try to explain the different possible cases in a select statement. 1. You "listen" for a value to come from the channel. 2. You listen for a channel being ready to receive a certain value. 3. You listen for the channel to close.
They looks like this respectively:

1. `case value := <-ch:`
2. `case ch <- value:`
3. `case <-ch:`

--- 

Next we're adding a timeout. This is very common in select statements. In fact time.After returns a channel in the same way. We add this, and then let's make it configurable so we don't have to wait so long.

We add a time.Duration parameter in our Racer function. But let's not rush to add parameters in the tests. 

The solve here is to make a wrapper `Racer` that wraps `ConfigurableRacer` (that has the old code.). So the real code just uses Racer, and the timeout variable is automateically supplied. But when you want to go more in depth and supply it yourself, you directly access the ConfigurableRacer.