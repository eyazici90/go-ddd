package aggregate

import (
	"github.com/google/uuid"
)

type Version string

func NewVersion() Version { return Version(uuid.New().String()) }

func (v Version) String() string { return string(v) }

func ToVersion(version string) Version { return Version(version) }
