package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//Version holds version in uint64
type Version uint64

func ParseVersion(version string) (Version, error) {
	decode, err := VersionDecode(version)
	if err != nil {
		return 0, err
	}

	var value uint64

	decode[0] = (decode[0] << 48) & 0xffff000000000000
	decode[1] = (decode[1] << 32) & 0x0000ffff00000000
	decode[2] = decode[2] & 0x00000000ffffffff

	value = decode[0] | decode[1] | decode[2]

	return Version(value), nil
}

//VersionEncode convert Version to string represantation
func VersionEncode(major, minor, patch uint64) string {
	return strconv.FormatUint(major, 10) + "." +
		strconv.FormatUint(minor, 10) + "." +
		strconv.FormatUint(patch, 10)
}

//VersionDecode gets an major.minor.patch and returns array of parsed version
func VersionDecode(value string) ([]uint64, error) {
	versionSegments := strings.Split(value, ".")

	if len(versionSegments) != 3 {
		return nil, fmt.Errorf("Version should have 3 parts, got %d", len(versionSegments))
	}

	major, err := strconv.ParseUint(versionSegments[0], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Major part of Version is not parrsable")
	}

	minor, err := strconv.ParseUint(versionSegments[1], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Minor part of Version is not parrsable")
	}

	patch, err := strconv.ParseUint(versionSegments[2], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Patch part of Version is not parrsable")
	}

	return []uint64{major, minor, patch}, nil
}

//MarshalJSON for type Version
func (v Version) MarshalJSON() ([]byte, error) {
	value := uint64(v)

	major := (value & 0xffff000000000000) >> 48
	minor := (value & 0x0000ffff00000000) >> 32
	patch := (value & 0x00000000ffffffff)

	version := VersionEncode(major, minor, patch)

	return json.Marshal(version)
}

//UnmarshalJSON for type Version
func (v *Version) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Version should be a string, got %s", data)
	}

	version, err := VersionDecode(s)

	if err != nil {
		return err
	}

	var value uint64

	version[0] = (version[0] << 48) & 0xffff000000000000
	version[1] = (version[1] << 32) & 0x0000ffff00000000
	version[2] = version[2] & 0x00000000ffffffff

	value = version[0] | version[1] | version[2]

	*v = Version(value)
	return nil
}

func VersionAdd(version Version, major, minor, patch uint64) Version {
	value := uint64(version)

	major = major + ((value & 0xffff000000000000) >> 48)
	minor = minor + ((value & 0x0000ffff00000000) >> 32)
	patch = patch + (value & 0x00000000ffffffff)

	major = (major << 48) & 0xffff000000000000
	minor = (minor << 32) & 0x0000ffff00000000
	patch = patch & 0x00000000ffffffff

	value = major | minor | patch

	return Version(value)
}

func MaskVersion(version Version, major, minor, patch uint64) Version {
	value := uint64(version)

	v1 := (value & 0xffff000000000000) >> 48
	v2 := (value & 0x0000ffff00000000) >> 32
	v3 := value & 0x00000000ffffffff

	if major == 0 {
		v1 = 0
	}

	if minor == 0 {
		v2 = 0
	}

	if patch == 0 {
		v3 = 0
	}

	v1 = (v1 << 48) & 0xffff000000000000
	v2 = (v2 << 32) & 0x0000ffff00000000
	v3 = v3 & 0x00000000ffffffff

	value = v1 | v2 | v3

	return Version(value)
}