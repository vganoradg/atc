package cliserver

import (
	"net/http"
	"path/filepath"
)

func (s *Server) Download(w http.ResponseWriter, r *http.Request) {
	if s.cliDownloadsDir == "" {
		http.Error(w, "cli downloads directory not configured", http.StatusNotFound)
		return
	}
	platform := r.URL.Query().Get("platform")

	arch := r.URL.Query().Get("arch")

	switch platform {
	case "windows", "darwin", "linux":
	default:
		http.Error(w, "invalid platform", http.StatusBadRequest)
		return
	}

	switch arch {
	case "amd64":
	case "i386":
		http.Error(w, "too few bits", http.StatusPaymentRequired)
		return
	default:
		http.Error(w, "invalid architecture", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=fly")

	http.ServeFile(w, r, filepath.Join(s.cliDownloadsDir, platform, arch, "fly"))
}
