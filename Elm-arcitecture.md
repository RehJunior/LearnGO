# Elm-Architektur in Go

In der Elm-Architektur, die in Go mit dem BubbleTea-Framework implementiert wird, haben die `Init`, `Update` und `View` Funktionen folgende Rollen:

## Init

Die `Init`-Funktion setzt den Anfangszustand der Anwendung und führt den ersten Befehl aus. In Ihrem Code gibt die `Init`-Funktion den `checkServer`-Befehl zurück. Dieser Befehl sendet eine HTTP-Anfrage an einen Server und gibt das Ergebnis als eine Nachricht zurück. Diese Nachricht kann entweder den Statuscode der HTTP-Antwort oder einen Fehler enthalten, wenn bei der Anfrage ein Problem aufgetreten ist.

## Update

Die `Update`-Funktion nimmt Nachrichten entgegen und aktualisiert den Zustand der Anwendung entsprechend. In Ihrem Code überprüft die `Update`-Funktion die Art der empfangenen Nachricht. Wenn die Nachricht den Statuscode einer HTTP-Antwort enthält, wird dieser Statuscode im Modell gespeichert und das Programm beendet. Wenn die Nachricht einen Fehler enthält, wird dieser Fehler im Modell gespeichert und das Programm ebenfalls beendet.

## View

Die `View`-Funktion stellt den aktuellen Zustand der Anwendung dar. In Ihrem Code überprüft die `View`-Funktion, ob im Modell ein Fehler gespeichert ist. Wenn ein Fehler vorliegt, gibt sie eine Fehlermeldung zurück. Wenn kein Fehler vorliegt, gibt sie eine Statusmeldung zurück, die den Status der HTTP-Anfrage darstellt. Diese Statusmeldung enthält die URL, die überprüft wird, und den Statuscode der HTTP-Antwort, wenn einer vorhanden ist.
