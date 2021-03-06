/*
 * Copyright (C) 2014 Cloudius Systems, Ltd.
 *
 * This work is open source software, licensed under the terms of the
 * BSD license as described in the LICENSE file in the top-level directory.
 */

package instance

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func ListInstances() error {
	header := fmt.Sprintf("%-20s %-10s %-10s %-15s", "Name", "Platform", "Status", "Image")
	fmt.Println(header)
	var homepath string
	if runtime.GOOS == "windows" {
		homepath = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
	} else {
		homepath = os.Getenv("HOME")
	}
	rootDir := filepath.Join(homepath, ".capstan", "instances")
	platforms, _ := ioutil.ReadDir(rootDir)
	for _, platform := range platforms {
		if platform.IsDir() {
			platformDir := filepath.Join(rootDir, platform.Name())
			instances, _ := ioutil.ReadDir(platformDir)
			for _, instance := range instances {
				if instance.IsDir() {
					instanceDir := filepath.Join(platformDir, instance.Name())
					printInstance(instance.Name(), platform.Name(), instanceDir)
				}
			}
		}
	}

	return nil
}

func printInstance(name, platform, dir string) error {
	fmt.Printf("%-20s %-10s %-10s %-15s\n", name, platform, "Running", "")
	return nil
}
