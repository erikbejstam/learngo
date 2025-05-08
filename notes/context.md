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

What about `context.Value()`? In short, don't use it. You can pass values in the context. But you lose type safety. Om du vill ha values, skicka dem bara som vanligt i funktionen because then they'll have types. It's possible to argue that you can use Value for stuff like trace id. "Value is for maintainers, not for users".

---

NOTE: I have to redo this lesson I think, because a lot of it was pretty confusing, perhaps because I didn't really understand the http stuff. I will probably delete the code and do it again. If I want to look at the first version I did of this stuff, I'll have to go back and look at an older coommit.

## Doing it all again now

I think it makes more sense now. I don't have a firm grip on the httptest package though. But we'll see. 

---

Why do we create our own mock ResponseWriter? Because the one httptest provides has no way of checking, apparently, if it has been written to. And we want to check that.

.... Later we remove the Cancel method from Store and from our SpyStore, since it is now the context's responsability. So we are not checking I guess for if it was cancelled. I'm not sure why we can't do that (I guess because it is not the responsability of Store, but still...). It doesn't "register" that it's been "cancelled". Instead it just should return data. So let's check that.

---

Ok! I think I roughly understand this pretty good now. Of course I should delve more deeply into context, mux and stuff. But I could *almost* write this code by myself.