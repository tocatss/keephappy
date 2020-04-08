package functor

// https://github.com/uber-go/guide/blob/master/style.md#functional-options

type user struct {
	name     string
	sex      string
	password string
	imgURL   string
}

// Bad
func NewUser() *user {
	return &user{}
}

func NewUserWithName(name string) *user {
	return &user{}
}

func NewUserWithImgURL(imgURL string) *user {
	return &user{}
}

// Good
// Add option.
type option struct {
	name     string
	sex      string
	password string
	imgURL   string
}

type setOption func(*option)

func (f setOption) apply(opt *option) {
	f(opt)
}

type Option interface {
	apply(opt *option)
}

func WithName(name string) Option {
	return setOption(
		func(opt *option) {
			opt.name = name
		},
	)
}

func WithSex(sex string) Option {
	return setOption(
		func(opt *option) {
			opt.sex = sex
		},
	)
}

func WithPassword(password string) Option {
	return setOption(
		func(opt *option) {
			opt.password = password
		},
	)
}

func WithImgURL(imgURL string) Option {
	return setOption(
		func(opt *option) {
			opt.imgURL = imgURL
		},
	)
}

func NewUserWithOptions(opts ...Option) *user {
	defaultOption := &option{
		name:     "default",
		sex:      "default",
		password: "default",
		imgURL:   "default",
	}
	for _, o := range opts {
		o.apply(defaultOption)
	}
	return &user{
		name:     defaultOption.name,
		sex:      defaultOption.sex,
		password: defaultOption.password,
		imgURL:   defaultOption.imgURL,
	}
}
