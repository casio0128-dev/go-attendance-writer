package main

type Config struct {
	TargetFile	string `toml:"target-file"`
	BackupDirectory	string	`toml:"backup-directory"`
	SpecifiedTime	string	`toml:"specified-time"`
}