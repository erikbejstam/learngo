### Structs, Methods and Interfaces

Spontant skulle man kunna tänka sig att man vill döpa funktionerna för t.ex. area speciellt så att det är t.ex. RectangleArea, så att den fungerar som väntat för användaren. Men en bättre metod är att definiera objekt som kallas ***strukturer***.

Vill ju i detta fall ha *flera* area-funktioner t.ex. Då använder vi ***metoder***, dvs funktioner kopplade till strukturer.

---

Skapar så småning om här ett interface. Interfaces implementeras implicit i Go, vilket innebär att sålänge en struct har de metoder som ges interfacet, så *är* strukten av den typen. Rectangle är en Shape.

---

När vi börjar få tillräckligt många test vi vill köra, t.ex. testa varje area-funktion för massa olika shapes, kan det vara klokt att ha ett Test Table.

Om man har tester i ett test table kan man fortfarande köra dem individuellt genom att i terminalen skriva `go test -run <path till funk typ>`.