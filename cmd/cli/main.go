////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2022 The HeyApple Authors.
//
// Use of this source code is governed by the GNU Affero General
// Public License as published by the Free Software Foundation,
// either version 3 of the License, or any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
//
////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/and-rad/heyapple/internal/conv"
)

var (
	flagConvert = flag.String("convert", "", "convert food data in specified file to a format usd by HeyApple internally")
	flagFrom    = flag.String("from", "", "the input data format for conversion")
	flagOut     = flag.String("out", "", "the output file name")
)

type cli struct {
	cwd string
}

func (c *cli) convert(in, out, format string) {
	path, err := filepath.Abs(in)
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	foods := []conv.Food{}
	switch format {
	case "usda":
		foods, err = conv.FromUSDA(data)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	data, err = json.Marshal(foods)
	if err != nil {
		log.Fatal(err.Error())
	}

	path, err = filepath.Abs(out)
	if err != nil {
		log.Fatal(err.Error())
	}

	data = bytes.ReplaceAll(data, []byte("},{"), []byte("},\n\t{"))
	data = bytes.ReplaceAll(data, []byte("[{"), []byte("[\n\t{"))
	data = bytes.ReplaceAll(data, []byte("}]"), []byte("}\n]"))

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	flag.Parse()

	exe, _ := os.Executable()
	cli := cli{cwd: filepath.Dir(exe)}

	if *flagConvert != "" {
		cli.convert(*flagConvert, *flagOut, *flagFrom)
	}
}
