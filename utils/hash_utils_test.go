package utils

import (
	"testing"
)

func TestGenerateHashFromPassword(t *testing.T) {
	const password = "password"
	hashPassword, err := GenerateHashFromPassword(password)

	if err != nil {
		t.Errorf("utils.GenerateHashFromPassword. \nError: %v+", err)
	}

	if len(hashPassword) == 0 {
		t.Errorf("utils.GenerateHashFromPassword. \nError: Password length is %v+", len(hashPassword))
	}
}

func TestCompareHashWithPassword(t *testing.T) {
	tests := []struct {
		name         string
		password     string
		hashPassword string
		want         bool
	}{
		{
			name:         "First password test",
			password:     "password",
			hashPassword: "$argon2id$v=19$m=65536,t=3,p=2$S5gYgNoUv+3VWDVVw2gc8A$Ewqb9u7hpu7fRQZC5jdcrP3x+1pZ7zq59AUMKt75oQ8",
			want:         true,
		},
		{
			name:         "Second password test",
			password:     "password1",
			hashPassword: "$argon2id$v=19$m=65536,t=3,p=2$S5gYgNoUv+3VWDVVw2gc8A$Ewqb9u7hpu7fRQZC5jdcrP3x+1pZ7zq59AUMKt75oQ8",
			want:         false,
		},
		{
			name:         "Third password test",
			password:     "password2",
			hashPassword: "$argon2id$v=19$m=65536,t=3,p=2$8wisnm4qg5h2BMC1IFTFXA$rigIEwvbjHM7ldBiTO7AX9sbZ8EtCR8YAFtxKmSL3vo",
			want:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match, err := CompareHashWithPassword(test.password, test.hashPassword)

			if err != nil {
				t.Errorf("utils.CompareHashWithPassword, \nError: %v+", err)
			}

			if match != test.want {
				t.Errorf("utils.CompareHashWithPassword, \nError: match = %v+ | want = %v+", match, test.want)
			}
		})
	}

}
