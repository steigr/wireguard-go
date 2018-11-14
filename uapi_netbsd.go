// +build darwin freebsd openbsd

/* SPDX-License-Identifier: GPL-2.0
 *
 * Copyright (C) 2017-2018 WireGuard LLC. All Rights Reserved.
 */

package main

var uapiSocketUmask = 0077

func setUapiSocketPermissions(string socketPath) error {
  return nil
}
