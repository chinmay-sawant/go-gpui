package collector

// procNames are the process names the fixture draws from. They are plain
// names, not real command lines, so the fixture never looks like host data.
var procNames = []string{
	"chrome", "firefox", "code", "node", "python3", "go", "gopls",
	"systemd", "sshd", "cron", "rsyslogd", "NetworkManager", "cupsd",
	"postgres", "nginx", "redis", "docker", "containerd", "podman",
	"bash", "zsh", "tmux", "vim", "slack", "discord", "spotify", "steam",
	"pipewire", "pulseaudio", "Xorg", "gnome-shell", "plasmashell",
}

// procUsers are the fixture user names.
var procUsers = []string{
	"root", "chinmay", "postgres", "www-data", "systemd+", "redis", "nobody",
}
