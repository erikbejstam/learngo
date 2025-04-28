### Integers

Tidigare användes %q för att vi ville printa ut strängar. Nu använder vi %d som i ints.

En bra grej angående dokumentation är att skriva **Testable Examples**. Jag tror att läggs till som exempel när man kompilerar en godoc-dokumentation. De kollas vid kompilering så att de alltid är up-to-date. Om en kommentar på formen `Output: x` är tillagd körs testet, annars inte. Ibland passar inte unittester för funktioner, och det kan man illustrera med att inte ha med denna kommentar. 

