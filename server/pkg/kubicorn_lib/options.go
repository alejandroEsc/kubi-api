// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubicornlib

import (
    "fmt"
    "log"

    "github.com/kris-nova/kubicorn/state/fs"
    "github.com/kris-nova/kubicorn/state/jsonfs"
    gg "github.com/tcnksm/go-gitconfig"
    "errors"
    "github.com/kris-nova/kubicorn/state/git"
)

type Options struct {
    StateStore     string
    StateStorePath string
    Name           string
    CloudId        string
    Set            string
    AwsProfile     string
    GitRemote      string
}

func (options Options) NewStateStore() (ClusterStorer, error) {
    var stateStore ClusterStorer

    switch options.StateStore {
    case "fs":
        log.Print("Selected [fs] state store")
        stateStore = fs.NewFileSystemStore(&fs.FileSystemStoreOptions{
            BasePath:    options.StateStorePath,
            ClusterName: options.Name,
        })
    case "git":
        log.Print("Selected [git] state store")
        if options.GitRemote == "" {
            return nil, errors.New("Empty GitRemote url. Must specify the link to the remote git repo.")
        }
        user, _ := gg.Global("user.name")
        email, _ := gg.Email()

        stateStore = git.NewJSONGitStore(&git.JSONGitStoreOptions{
            BasePath:    options.StateStorePath,
            ClusterName: options.Name,
            CommitConfig: &git.JSONGitCommitConfig{
                Name:   user,
                Email:  email,
                Remote: options.GitRemote,
            },
        })
    case "jsonfs":
        log.Print("Selected [jsonfs] state store")
        stateStore = jsonfs.NewJSONFileSystemStore(&jsonfs.JSONFileSystemStoreOptions{
            BasePath:    options.StateStorePath,
            ClusterName: options.Name,
        })
    default:
        return nil, fmt.Errorf("State store [%s] has an invalid type [%s].", options.Name, options.StateStore)
    }

    return stateStore, nil
}
