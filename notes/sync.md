# Sync

Let's create a counter that's safe to use concurrently. First, let's make an unsafe one.

---

Ok now let's do the safe one.

Note: we'll use `WaitGroup`. This is how we use it, atleast here:
1. Create the wg
2. Call wg.Add() with the number of goroutines it should wait for.
3. then you start the goroutine. inside it you do `wg.Done()` to indicate to the wg that the goroutine has finished.
4. Outside of the goroutines you have a `wg.Wait()` that tells the wg to wait for all wg.Done calls.

---

The test *might not* fail, but the point is, this is not safe for concurrency right now. Let's add some type of lock. This is done with `sync.Mutex` ("mutually exclusive"). What we basically have to do is, as soon as the function is called, lock the object, and defer unlocking. There we go!

Note: sometimes people embed the mutex directly in the object, but this is bad practice since it exposes it, so that users can lock and unlock as they please.

Note: you can use `go vet` to find some suspicious code not caught by compiler. In our case, it tells us that our assertCounter function is copying the mutex, which is not good. I guess it is because you "think" you're getting the mutex, but it isn't being updated since it it a copy of the original. You should pass a reference to it.

---

So, back to the pointer passing. Let the function receive a pointer to the Counter. But also create a counter constructor, and use that instead of creating the counter and THEN pass an explicit reference.

--- 

When to use channels vs mutex? Use mutex when you just want to *manage state*. Use channels when you want to *pass ownership of data*.

---

All in all, pretty nice and easy lesson. 