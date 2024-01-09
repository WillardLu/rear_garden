// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package main

import "gin_http_server/internal/myhttpserver"

func main() {
	myhttpserver.CreateHttpServer()
	return
}
