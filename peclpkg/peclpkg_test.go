package peclpkg_test

import (
	"testing"

	"github.com/NiR-/notpecl/peclpkg"
	"github.com/go-test/deep"
)

type loadTC struct {
	file        string
	expected    peclpkg.Package
	expectedErr error
}

func TestLoadPackageXMLFromFile(t *testing.T) {
	testcases := map[string]loadTC{
		"successfully load package.xml for redis ext": {
			file: "testdata/package-redis.xml",
			expected: peclpkg.Package{
				Name:        "redis",
				Summary:     "PHP extension for interfacing with Redis",
				Description: "This extension provides an API for communicating with Redis servers.",
				PublishDate: "2019-11-11",
				PublishTime: "07:36:41",
				Version: peclpkg.Version{
					Release: "5.1.1",
					API:     "5.1.0",
				},
				Stability: peclpkg.PackageStability{
					Release: peclpkg.Stable,
					API:     peclpkg.Stable,
				},
				License: peclpkg.License{
					Name: "PHP",
					URI:  "http://www.php.net/license",
				},
				Dependencies: peclpkg.Dependencies{
					Required: peclpkg.RequiredDependencies{
						PHP: peclpkg.PHPConstraint{
							Min: "7.0.0",
							Max: "7.9.99",
						},
					},
				},
				ExtSrcRelease: peclpkg.ExtSrcRelease{
					ConfigureOptions: []peclpkg.ConfigureOption{
						{
							Name:    "enable-redis-igbinary",
							Default: "no",
							Prompt:  "enable igbinary serializer support?",
						},
						{
							Name:    "enable-redis-lzf",
							Default: "no",
							Prompt:  "enable lzf compression support?",
						},
						{
							Name:    "enable-redis-zstd",
							Default: "no",
							Prompt:  "enable zstd compression support?",
						},
					},
				},
				Changelog: peclpkg.Changelog{
					Releases: []peclpkg.Release{
						{
							Stability: peclpkg.PackageStability{
								Release: peclpkg.Stable,
								API:     peclpkg.Stable,
							},
							Version: peclpkg.Version{
								Release: "5.1.1",
								API:     "5.1.0",
							},
							Date: "2019-11-11",
							Notes: `
phpredis 5.1.1

This release contains only bugfix for unix-socket connection.

* Fix fail to connect to redis through unix socket [2bae8010, 9f4ededa] (Pavlo Yatsukhnenko, Michael Grunder)
* Documentation improvements (@fitztrev)
   `,
						},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		out, err := peclpkg.LoadPackageXMLFromFile(tc.file)
		if tc.expectedErr != nil {
			if err == nil {
				t.Fatalf("Expected error: %v\nGot: <nil>", tc.expectedErr)
			}
			if tc.expectedErr.Error() != err.Error() {
				t.Fatalf("Expected error: %v\nGot: %v", tc.expectedErr, err)
			}
			return
		}
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Match only the 1st release in the changelog as testing more
		// releases wouldn't bring anything more but would make this file
		// quite big.
		out.Changelog.Releases = out.Changelog.Releases[:1]

		if diff := deep.Equal(out, tc.expected); diff != nil {
			t.Fatal(diff)
		}
	}
}
