package orm

type option struct {
	addr         string
	username     string
	password     string
	dbname       string
	maxOpenConns int
	idleConns    int
	charset      string
}

type Option func(op *option)

func WithCharset(charset string) Option {
	return func(o *option) {
		o.charset = charset
	}
}

func WithAddr(addr string) Option {
	return func(o *option) {
		o.addr = addr
	}
}
func WithUsername(username string) Option {
	return func(o *option) {
		o.username = username
	}
}
func WithPassword(password string) Option {
	return func(o *option) {
		o.password = password
	}
}
func WithMaxOpenConns(maxOpenConns int) Option {
	return func(o *option) {
		o.maxOpenConns = maxOpenConns
	}
}

func WithDBname(dbname string) Option {
	return func(o *option) {
		o.dbname = dbname
	}
}

func WithIdleConns(idleConns int) Option {
	return func(o *option) {
		o.idleConns = idleConns
	}
}

func defaultOption() option {
	return option{
		maxOpenConns: 20,
		idleConns:    20,
		charset:      "utf8",
	}
}
