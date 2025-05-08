# Context

Now: often software uses long-running processes (here in the form of goroutines). If something goes wrong with thing that started these, it gets really hard to debug. In this case we should use `context`.
In this scenario we'll start a web server, make a request to it, and cancel the request before the data is retrieved - and we'll have to make sure the cancellation is done cleanly.

---

*a bunch of stuff here i should write about that i haven't gotten around to yet...* 

--- 

We'll now write helper test methods on our object `Store`. This is new, but I guess it makes sense. So you will make the `SpyStore` have a *testing.T field, so you don't have to pass that every time you want to run this helper.

---

The handler function/server/whatever that I think gets requests, must take the context from those request (which i think are automatically created by Go), and then pass those into the functions it is calling. So if the handler gets a request to get something from a db, it should take a ctx from request and then do something like `data := store.FetchFromDB(ctx, id)`.

---

NOTE: I have to redo this lesson I think, because a lot of it was pretty confusing, perhaps because I didn't really understand the http stuff. I will probably delete the code and do it again. If I want to look at the first version I did of this stuff, I'll have to go back and look at an older coommit.
