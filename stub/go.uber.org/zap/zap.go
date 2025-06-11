package zap

type Logger struct{}

type AtomicLevel struct{}

type Config struct{ Level AtomicLevel }

func NewNop() *Logger { return &Logger{} }

func ParseAtomicLevel(l string) (AtomicLevel, error) { return AtomicLevel{}, nil }

func NewProductionConfig() Config { return Config{} }

func (c Config) Build(opts ...interface{}) (*Logger, error) { return &Logger{}, nil }

func (l *Logger) Info(msg string, fields ...interface{})  {}
func (l *Logger) Debug(msg string, fields ...interface{}) {}
func (l *Logger) Warn(msg string, fields ...interface{})  {}
func (l *Logger) Error(msg string, fields ...interface{}) {}
func (l *Logger) Fatal(msg string, fields ...interface{}) {}
