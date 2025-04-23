package services

import (
	"backend/database"
	"backend/model/server"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type ServerManager struct {
	client  *client.Client
	db      *database.DB
	Servers map[string]string
}

func NewServerManager(db *database.DB) *ServerManager {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	sm := ServerManager{
		cli,
		db,
		nil,
	}

	containers, err := cli.ContainerList(context.TODO(), containertypes.ListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "giorgioservermanager",
			Value: "1",
		}),
	})
	//if no container, return with an empty map. Is the first time of the user??
	if err != nil {
		return &sm
	}

	// Should i check if the string is ulid or not?
	// Is a bit useless but idk
	server_map := make(map[string]string)
	for _, container := range containers {
		server_id := container.Labels["giorgioserverid"]
		server_map[server_id] = container.ID
	}

	return &ServerManager{
		cli,
		db,
		server_map,
	}
}

func (sm *ServerManager) Close(ctx context.Context) {
	if err := sm.client.Close(); err != nil {
		panic(err)
	}
}

func (sm *ServerManager) PullImageBasedOn(ctx context.Context, servertype server.SeverType, version string) error {
	switch servertype {
	case "VANILLA":
		java_version, err := javaVersionFromManifest(strings.ToLower(version))
		if err != nil {
			panic(err)
		}
		tag := "java" + java_version
		if err := sm.pullImage(ctx, tag); err != nil {
			panic(err)
		}
	}
	return nil
}

func (sm *ServerManager) pullImage(ctx context.Context, containerTag string) error {
	//Is async and reader must be read to complete the task!
	reader, err := sm.client.ImagePull(ctx, fmt.Sprintf("itzg/minecraft-server:%s", containerTag), image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	//TODO: find a way to save this logs somewhere instead of stdout
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}

	return nil
}

type Manifest struct {
	Latest struct {
		Release  string
		Snapshot string
	}
	Versions []struct {
		ID  string `json:"id"`
		Url string `json:"url"`
	}
}

type VersionManifest struct {
	JavaVersion struct {
		Component    string `json:"component,omitempty"`
		MajorVersion int    `json:"majorVersion"`
	} `json:"javaVersion"`
}

// TODO: OPTIMIZE THIS FUNCTION
func javaVersionFromManifest(version string) (string, error) {
	res, err := regexp.MatchString(`\d+\.\d+\.\d+`, version)
	if err != nil {
		return "", err
	}
	if !res {
		panic(fmt.Sprintf("the version passed must be of the form \\d.\\d.\\d but get: %s", version))
	}

	resp, err := http.Get("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	manifest := new(Manifest)
	if err := json.NewDecoder(resp.Body).Decode(manifest); err != nil {
		return "", err
	}

	for _, v := range manifest.Versions {
		if !strings.EqualFold(strings.ToLower(v.ID), strings.ToLower(version)) {
			continue
		}

		resp, err := http.Get(v.Url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		var version_manifest VersionManifest
		if err := json.NewDecoder(resp.Body).Decode(&version_manifest); err != nil {
			return "", err
		}

		return strconv.Itoa(version_manifest.JavaVersion.MajorVersion), nil
	}
	return "", err
}

func (sm *ServerManager) CreateContainer(ctx context.Context) error {
	return nil
}
