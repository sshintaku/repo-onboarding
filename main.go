package main

import (
	"encoding/json"
	"fmt"

	CloudType "github.com/sshintaku/cloud_types"
	"github.com/sshintaku/prisma_session"
)

func main() {
	var params = prisma_session.ReadParameters()
	session := prisma_session.Session{ApiUrl: params.ApiUrl}
	session.CreateSession()
	var repo CloudType.RepoProperties
	repo.Type = "github"
	repo.Data = append(repo.Data, "sshintaku/prisma_session")
	result, _ := json.Marshal(repo)
	fmt.Println(string(result))
	session.RepoOnboaarding(repo)
}
