package storage

import (
	"bytes"
	"slices"
	"strings"

	"github.com/adrg/frontmatter"
)

type Byte struct {
	Id      string
	Title   string
	Slug    string
	Tags    []string
	Content string
}

func (b Byte) HasTag(tag string) bool {
	return slices.Contains(b.Tags, tag)
}

func slugify(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}

func ParseByte(id string, content []byte) (Byte, error) {
	var matter struct {
		Title string   `yaml:"title"`
		Slug  string   `yaml:"slug"`
		Tags  []string `yaml:"tags"`
	}

	data, err := frontmatter.Parse(bytes.NewReader(content), &matter)
	if err != nil {
		return Byte{}, err
	}

	slug := matter.Slug
	if slug == "" {
		slug = slugify(matter.Title)
	}

	return Byte{
		Id:      id,
		Title:   matter.Title,
		Slug:    slug,
		Tags:    matter.Tags,
		Content: string(data),
	}, nil
}
