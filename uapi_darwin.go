// +build darwin freebsd openbsd

/* SPDX-License-Identifier: GPL-2.0
 *
 * Copyright (C) 2017-2018 WireGuard LLC. All Rights Reserved.
 */

package main

import (
  "os"
  "os/user"
  "strconv"
)

var uapiSocketUmask = 0007

func setUapiSocketPermissions(socketPath string) (err error) {
  var (
    uid, gid int
    group    *user.Group
  )
  uid = os.Getuid()

  group, err = user.LookupGroup("staff")
  if err != nil {
    return err
  }

  gid, err = strconv.Atoi(group.Gid)
  if err != nil {
    return err
  }

  // change group to staff
  err = os.Chown(socketPath, uid, gid)
  return err
}
