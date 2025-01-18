package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	procKeybdEvent       = user32.NewProc("keybd_event")

	running      = true  // Variable pour contrôler l'exécution
	lastClipboard string // Dernière valeur connue du presse-papier
	exitChan     = make(chan bool)
)

func main() {
	// Lancer l'icône dans la barre des tâches dans un thread séparé
	go func() {
		systray.Run(onReady, onExit)
	}()

	fmt.Println("Programme en cours d'exécution... Copie automatiquement le texte sélectionné.")

loop:
	for {
		select {
		case <-exitChan:
			break loop // Quitter la boucle principale
		default:
			if running {
				// Simuler Ctrl+C pour copier la sélection active
				copySelection()

				// Lire le contenu du presse-papier
				text, err := clipboard.ReadAll()
				if err != nil {
					fmt.Println("Erreur lors de la lecture du presse-papier :", err)
					time.Sleep(1 * time.Second)
					continue
				}

				// Vérifier si le contenu a changé
				if text != lastClipboard {
					lastClipboard = text
					fmt.Println("Texte copié :", text)
				}
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Programme terminé.")
}

func copySelection() {
	// Simule Ctrl+C pour copier la sélection active
	if !isKeyPressed(0x11) { // 0x11 est le code de la touche Ctrl
		keybdEvent(0x11, 0, 0)          // Appuie sur Ctrl
		keybdEvent(0x43, 0, 0)          // Appuie sur 'C'
		keybdEvent(0x43, 0, 0x0002)     // Relâche 'C'
		keybdEvent(0x11, 0, 0x0002)     // Relâche Ctrl
	}
}

func isKeyPressed(vk int) bool {
	// Vérifie si une touche est enfoncée en appelant GetAsyncKeyState
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

func keybdEvent(bVk, bScan, dwFlags int) {
	// Appelle keybd_event pour simuler l'appui sur une touche
	procKeybdEvent.Call(uintptr(bVk), uintptr(bScan), uintptr(dwFlags), 0)
}

func onReady() {
	// Charger une icône (utiliser une icône en format .ico)
	systray.SetIcon(getIcon())
	systray.SetTitle("Clipboard Copier")
	systray.SetTooltip("Surveille automatiquement les sélections et copie.")

	// Ajouter uniquement le bouton "Quitter"
	quitMenu := systray.AddMenuItem("Quitter", "Quitter l'application")

	// Gérer le clic sur "Quitter"
	go func() {
		<-quitMenu.ClickedCh
		fmt.Println("Quitter l'application demandé.")
		exitChan <- true // Envoyer le signal de quitter
		systray.Quit()   // Supprimer l'icône de la barre des tâches
	}()
}

func onExit() {
	// Nettoyage avant la fermeture
	fmt.Println("Nettoyage avant fermeture...")
	os.Exit(0)
}

func getIcon() []byte {
	// Charger le fichier .ico
	icon, err := os.ReadFile("cp.ico")
	if err != nil {
		fmt.Println("Erreur lors du chargement de l'icône :", err)
		return []byte{}
	}
	return icon
}