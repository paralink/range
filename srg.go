package srg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	hostnamePattern = regexp.MustCompile(`^([a-zA-Z]+)(\d+)$`)
	portPattern     = regexp.MustCompile(`^(\d+)$`)
)

// ParseRange - converts a given server range expression to a list of hostnames
//
// TODO:
//   - unify error handling
//   - more UT
//   - UT for error cases
func ParseRange(svrexp string) ([]string, error) {
	servers := []string{}

	if len(svrexp) == 0 {
		return servers, nil
	}

	ranges := strings.Split(svrexp, ",")

	for _, r := range ranges {
		parts := strings.Split(r, "~")

		switch len(parts) {
		case 1:
			servers = append(servers, parts[0])
		case 2:
			fstPrefix, fst, err := split(parts[0])
			if err != nil {
				return nil, err
			}

			sndPrefix, snd, err := split(parts[1])
			if err != nil {
				return nil, err
			}

			if fstPrefix != sndPrefix {
				return nil, fmt.Errorf("invalid range: %s", r)
			}

			fstNum, err := strconv.Atoi(fst)
			if err != nil {
				return nil, err
			}
			sndNum, err := strconv.Atoi(snd)
			if err != nil {
				return nil, err
			}

			if fstNum > sndNum {
				return nil, fmt.Errorf("server range is reversed: %s", r)
			}

			for i := fstNum; i <= sndNum; i++ {
				servers = append(servers, fmt.Sprintf("%s%d", fstPrefix, i))
			}
		default:
			return nil, fmt.Errorf("invalid server range: %s", r)
		}
	}

	return servers, nil
}

func split(token string) (string, string, error) {
	match := hostnamePattern.FindStringSubmatch(token)
	if len(match) == 3 {
		return match[1], match[2], nil
	}

	match = portPattern.FindStringSubmatch(token)
	if len(match) == 2 {
		return "", match[0], nil
	}

	return "", "", fmt.Errorf("invalid range '%s'", token)
}
