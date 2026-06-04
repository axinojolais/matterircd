package mattermost

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestChannelMessageType(t *testing.T) {
	tests := []struct {
		name     string
		settings map[string]bool
		message  string
		want     string
	}{
		{
			name:    "default here mention is notice",
			message: "hello @here",
			want:    "notice",
		},
		{
			name: "all and here can be regular messages",
			settings: map[string]bool{
				"mattermost.DisableHereAllNotices": true,
			},
			message: "hello @all",
			want:    "",
		},
		{
			name: "channel mention stays notice with new option",
			settings: map[string]bool{
				"mattermost.DisableHereAllNotices": true,
			},
			message: "hello @channel",
			want:    "notice",
		},
		{
			name: "disable default mentions keeps legacy behavior",
			settings: map[string]bool{
				"mattermost.DisableDefaultMentions": true,
			},
			message: "hello @channel",
			want:    "",
		},
		{
			name:    "regular messages stay regular",
			message: "hello team",
			want:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := viper.New()
			for key, value := range tc.settings {
				v.Set(key, value)
			}

			m := &Mattermost{v: v}
			assert.Equal(t, tc.want, m.channelMessageType(tc.message))
		})
	}
}
