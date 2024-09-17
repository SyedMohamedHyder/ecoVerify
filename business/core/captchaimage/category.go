package captchaimage

import "fmt"

// Set of possible categories for a captcha image.
var (
	CaptchaImageCategoryEcoFriendly    = CaptchaImageCategory{"ECOFRIENDLY"}
	CaptchaImageCategoryNonEcoFriendly = CaptchaImageCategory{"NONECOFRIENDLY"}
)

// Map of known categories.
var categories = map[string]CaptchaImageCategory{
	CaptchaImageCategoryEcoFriendly.name:    CaptchaImageCategoryEcoFriendly,
	CaptchaImageCategoryNonEcoFriendly.name: CaptchaImageCategoryNonEcoFriendly,
}

// CaptchaImageCategory defines category of images used in the CAPTCHA.
type CaptchaImageCategory struct {
	name string
}

// ParseCaptchaImageCategory parses the string value and returns a CaptchaImageCategory if one exists.
func ParseCaptchaImageCategory(value string) (CaptchaImageCategory, error) {
	category, exists := categories[value]
	if !exists {
		return CaptchaImageCategory{}, fmt.Errorf("invalid role %q", value)
	}

	return category, nil
}

// MustParseCaptchaImageCategory parses the string value and returns a CaptchaImageCategory if one exists. If
// an error occurs the function panics.
func MustParseCaptchaImageCategory(value string) CaptchaImageCategory {
	category, err := ParseCaptchaImageCategory(value)
	if err != nil {
		panic(err)
	}

	return category
}

// Name returns the name of the category.
func (c CaptchaImageCategory) Name() string {
	return c.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (c *CaptchaImageCategory) UnmarshalText(data []byte) error {
	category, err := ParseCaptchaImageCategory(string(data))
	if err != nil {
		return err
	}

	c.name = category.name
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (c CaptchaImageCategory) MarshalText() ([]byte, error) {
	return []byte(c.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (c CaptchaImageCategory) Equal(c2 CaptchaImageCategory) bool {
	return c.name == c2.name
}
