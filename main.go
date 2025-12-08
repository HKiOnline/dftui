package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/hkionline/dftui/services"
	"github.com/hkionline/dftui/ui"
)

var (
	port    = flag.String("port", "2222", "Port to listen on")
	hostKey = flag.String("host-key", "", "Path to host key (default: ~/.dftui/id_rsa)")
)

func main() {
	flag.Parse()

	// Initialize backend service (stub for now)
	backend := services.NewStubBackend()

	// Determine host key path
	keyPath := *hostKey
	if keyPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Failed to get user home directory:", err)
		}
		keyPath = filepath.Join(home, ".dftui", "id_rsa")
	}

	// Ensure .dftui directory exists
	keyDir := filepath.Dir(keyPath)
	if err := os.MkdirAll(keyDir, 0700); err != nil {
		log.Fatal("Failed to create host key directory:", err)
	}

	// Create SSH server with Wish
	// The bubbletea middleware creates a new Bubble Tea program for each SSH session
	s, err := wish.NewServer(
		wish.WithAddress(":"+*port),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			// Bubble Tea middleware - creates TUI for each session
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				// Extract username from SSH session (task 2.2)
				username := s.User()

				// Create new model for this user session
				m := ui.NewModel(username, backend)

				// Return model with alt screen buffer (clears screen on start/exit)
				return m, []tea.ProgramOption{
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				}
			}),
			// Logging middleware for debugging
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	// Graceful shutdown handling (task 2.4)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		addr := lipgloss.NewStyle().Bold(true).Render(s.Addr)
		log.Printf("Starting SSH server on %s", addr)
		log.Printf("Connect with: ssh localhost -p %s", *port)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	// Wait for shutdown signal
	<-done
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown server:", err)
	}

	log.Println("Server stopped")
}
