# Explaining to myself the http package

Ok, so I've been tripping up on the http package a little bit. I doesn't feel as straight forward as I would like. I'm going to talk about the *handler* stuff first.

So I can see the need for a `handler` interface - that just makes everything more clear and clean. A handler is something that implements a function `ServeHTTP` that takes a `ResponseWriter` and a `*Response`. That is to say: it takes a little box in which it will write what will be the HTTP response, and a pointer to the Go representation of the HTTP request it has received. 
My first intuition was: OK great, then you should just do something like 

```
{
    type HelloHandler struct {}

    func (h *HelloHandler) ServeHTTP(rw ResponseWriter, req *Request) {
        // do the thing
    }
}
```

That would be the most straight forward way to me. But for some reason that I'm not yet experienced enough to understand, this is not the way to do it. They designed it so it would be more "elegant" I guess, and shorter to write.

So atleast in the basic case your handler isn't an *object* (which I think of structs as being), but a function. In Go, functions are *first class citizens*, meaning they are basicallu ***values***. So if you alias your function as a type, you can define methods on it. Remember, a struct is just a type with "sub-fields" in it. So that is almost just like a map, that has methods on it. In this way, it doesn't seem far-fetched to think of simple types as having methods on it. Ok this was digressing a bit. 

I'm just going to continue to write myself off here. So, so far I've been thinking that it is just structs that implement interfaces, but it doesn't have to be that way. You can have a type implementing an interface. And therefore also a function. What does that do? It saves you the boilerplate code of creating an exampleHandler struct that implements Handler. You just write your function that is supposed to be handling the request. You just make sure it is following the rules/convention (have parameters ResponseWriter and \*Request). Then you kind of wrap that in a function that says it is returning a `HandlerFunc`, which is the function you defined. That makes it much shorter for you to implement the Handler interface. It is still kind of convoluted for me, but that tells Go/everyone "ok, this function I wrote is now officially a HANDLER!". Nice - everything clear. In a very asbtract sense, this should be all we need to satisfy the Handler interface, but it's just that we've defined the interface to have a ServeHTTP function. It feels weird to have our handler implement a function that *just calls itself*. But that's what we've said, and so the solution is to just make our handler "call itself via ServeHTTP". Or something. 
Maybe it is easier to explain like this: Handler is just an interface that requires you to have a ServeHTTP method on your object to satisfy it. If you had a helloHandler struct, you'd have to implement ServeHTTP, not weird at all. Then Go wants to make it easier for you to just write your own handler function without using a custom struct and stuff. So you have the HandlerFunc type that just "automatically" assigns the ServeHTTP to your method. Now, what can it do? Since you already designed the logic, it should just call your method. It might feel weird, but in the larger context that is the only reasonable thing here (... since that's just the design of the Handler interface).

En till forklaring: You can think of your own function as device with roung pegs. You want to PLUG IT IN to the wall (the http server). But the wall only accepts square pegs. Then you just use an *adapter* in between the two. This is what the HandlerFunc is. You just make sure you can plug into the adapter (your function returns a Handlerfunc), and you just get the added functionality of being able to plug in to the wall (because the HandlerFunc has added the ServerHTTP function for you).