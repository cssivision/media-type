package mediaType

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	t.Run("invalid type", func(t *testing.T) {
		mt1 := &MediaType{}
		_, err := mt1.Format()
		require.NotNil(t, err)

		mt2 := &MediaType{
			Type: "-vsdfvsd",
		}
		_, err = mt2.Format()
		require.NotNil(t, err)
	})

	t.Run("invalid subtype", func(t *testing.T) {
		mt1 := &MediaType{
			Type:    "application",
			SubType: "",
		}
		_, err := mt1.Format()
		require.NotNil(t, err)

		mt2 := &MediaType{
			Type:    "application",
			SubType: "vdfsvsd*",
		}
		_, err = mt2.Format()
		require.NotNil(t, err)
	})

	t.Run("invalid suffix", func(t *testing.T) {
		mt1 := &MediaType{
			Type:    "application",
			SubType: "json",
			Suffix:  "vsdvdfs*|",
		}
		_, err := mt1.Format()
		require.NotNil(t, err)
	})

	t.Run("valid media type", func(t *testing.T) {
		mt1 := &MediaType{
			Type:    "application",
			SubType: "json",
			Suffix:  "xml",
		}
		str, err := mt1.Format()
		require.Nil(t, err)
		assert.Equal(t, str, "application/json+xml")
	})
}

func TestParse(t *testing.T) {
	t.Run("invalid parse string", func(t *testing.T) {
		_, err := Parse("")
		require.NotNil(t, err)
		_, err = Parse("vdv")
		require.NotNil(t, err)
		_, err = Parse("application/json*")
		require.NotNil(t, err)
		_, err = Parse("application/json+xml*")
		require.NotNil(t, err)
	})

	t.Run("valid parse string", func(t *testing.T) {
		mt, err := Parse("application/json+xml")
		require.Nil(t, err)
		assert.Equal(t, mt.Type, "application")
		assert.Equal(t, mt.SubType, "json")
		assert.Equal(t, mt.Suffix, "xml")
	})
}
