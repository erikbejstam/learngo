### Arrays and Slices

Något speciellt och viktigt i Go är ju att arrays har fixerade storlekar. Det gör att man oftaste använder **slices**. 

Istället för att skriva den klassiska loopen, använd range. Lite snyggare:
`for index, number := range numbers`

En fet grej är att det finns ett coverage tool, som kollar hur mycket av kärnkoden som "kollas av" av testerna. Såklart kanon om man kan ha 100% coverage. Man kör verktyget genom `go test -cover <path>`.

Anm. I Go om man vill att en funktion ska kunna ta emot ett variabelt antal parametrar skriver man `Sum(numbers ...int)` t.ex.

Anm. När vi använder slices kan vi inte jämföra dem på samma sätt som tidigare med vanlig !=. Nu måste vi ta till ett bibliotek istället: `reflect`. DeepEqual, som man ofta använder, är inte type-safe ska sägas. Försiktig!

---

`make` låter en skapa en slice med en starting capacity. Man specificerar också vilken typ som slicen ska innehålla. 

---

Här provar vi också att slicea en slice. Vi ville ha den sk svansen av slicearna, dvs alla element utom det första. Då kan man skriva `tail := numbers[1:]`.

--

Ny grej som visas är att assigna en funktion till en variabel. Detta är lite krångligt att wrapa huvudet runt tycker jag. Men det handlar om att begränsa scopet för funktionen, och därför göra koden/APIt man skriver ha mindre "surface area". Finns fler syften som för mig är lite mer diffusa. Men t.ex. kan man ju "passa runt" funktionen som en vanlig variabel. Man kan också använda higher-order functions.

