### Maps

Vi ska igen leka med att göra en egen typ, så vi gör en liten wrapper Dictionary runt map.

Börjar också hantera lite mer avancerade errors. Nu gör vi errors till const, och då måste vi dessutom skapa en custom typ för dem. Hur gör vi custom errors? Jo, vi implementerar Error()-funktionen för dem.

```
{
    const ErrNotFound = DictionaryError("some string..")

    type DictionaryError string

    func (e DictionaryError) Error() string {
        return string(e)
    }
}
```

---

Ok so basically this is continuing the TDD process, and creating more custom structs, types and errors. 

So basically there are different ways of creating errors. The way to determine if something is an error is if it implements the `Error` interface. It does that if it has a method called Error that returns a string, that is all. 
Either you can do errors.New("the error string") or you can create a custom error type that is aliasing a string, and then implement Error() for it, like above.

I think this is all.