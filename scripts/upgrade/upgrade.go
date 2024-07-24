package main

import (
	"log"
	"os/exec"
)

func Panic(who error, where string) {
	log.Println("Upgrade failed!")

	log.Println("Please, check 'Error' or undo changes with: make upgrade-reset")

	log.Panicf("Error: (%s): [%s]", where, who)
}

func UpgradeGo() {
	if err := exec.Command("make", "upgrade-go").Run(); err != nil {
		Panic(err, "UpgradeGo")
	}
}

func UpgradeNode() {
	if err := exec.Command("make", "upgrade-node").Run(); err != nil {
		Panic(err, "UpgradeNode")
	}
}

func UpgradeTools() {
	if err := exec.Command("make", "install-tools").Run(); err != nil {
		Panic(err, "UpgradeTools")
	}
}

func RunLint() {
	if err := exec.Command("make", "lint").Run(); err != nil {
		Panic(err, "RunLint")
	}
}

func RunTest() {
	if err := exec.Command("make", "test-unit").Run(); err != nil {
		Panic(err, "RunTest")
	}
}

func Commit() {
	if err := exec.Command("git", "add", ".", "--update").Run(); err != nil {
		Panic(err, "Commit")
	}

	if err := exec.Command("git", "commit", "-m", "chore(deps): upgrade").Run(); err != nil {
		Panic(err, "Commit")
	}
}

func main() {
	log.Println("Starting upgrades...")
	RunLint()

	log.Println("Upgrading Go")
	UpgradeGo()

	log.Println("Upgrading Node")
	UpgradeNode()

	log.Println("Upgrading Tools")
	UpgradeTools()

	log.Println("Running Lint")
	RunLint()

	log.Println("Running Test")
	RunTest()

	log.Println("Commit changes")
	Commit()

	log.Println("Upgrade completed!")
}
