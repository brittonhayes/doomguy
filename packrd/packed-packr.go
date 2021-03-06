// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "333291d33cfffaf2f92784ca0f4d393f"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"125d889b9d306beab68b733d631f999e": "1f8b08000000000000ffe2e50acacf2f5170cecfcd4dcc4b29e6e54a4848a84ccccde1e5aaae56284acc4b4f55d073ce4d2956d0adade5e5d28529b452a8ae56d083721440520a0a2ea9c5c94599052599f97910692401b012b099a950f50909092091e0d2245dbc9607972651c5f2ea6a149b01010000ffff528e134ff8000000",
		"63d55de4b5f1d89d0f1606a418723c2b": "1f8b08000000000000ff4a4848e0e5aaae56484c4a2a4a2d53b0303050d073492d4e2eca2c28c9cccf53a8ade5e502ab01abca4c53d0f34d2d494c2eca2cc94c56d0adadfd30bf67be829656506249665eba969695427535b28adadaeaead4bc14902918da43837c20264cdffc7e473fc890d4b2ccd4720c4340ea30cd71492d4bcdc92f482d0299f1684e2fc408b828c4146465984684a726156796a442bd3141414b0b2a02d30c5380a415100000ffff1cc2723930010000",
		"b2b340d9641fe1122e4f857ea796f117": "1f8b08000000000000ff4a4848a84ccccde1e50a2d4e4c4fb5e2e5aaae56d003b3156a6b79b912121278b900010000ffff4f1b4d8224000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("Templates", "./templates")
		b.SetResolver("game.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "63d55de4b5f1d89d0f1606a418723c2b"})
		b.SetResolver("help.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "125d889b9d306beab68b733d631f999e"})
		b.SetResolver("usage.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "b2b340d9641fe1122e4f857ea796f117"})
	}()
	return nil
}()
