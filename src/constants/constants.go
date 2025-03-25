package constants

const (
	RunningServerPort = "Running Server on port : %v"
)

const (
	PANCardRegex     = `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	PasswordRegex    = `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	UppercaseRegex   = `[A-Z]`
	DigitRegex       = `\d`
	SpecialCharRegex = `[@$!%*?&]`
	LowercaseRegex   = `[a-z]`
)

const (
	FieldPassword        = "Password"
	FieldConfirmPassword = "ConfirmPassword"
	FieldPanCard         = "PanCard"
	FieldStrongPassword  = "strongPassword"
	FieldPhoneNumber     = "PhoneNumber"
)

// Migration success Message
const (
	MsgDBMigrationSuccess = "Database migration completed successfully!"
)

const (
	Postgres = "postgres"
	Yaml     = "yaml"
)

const (
	DSNString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s"
)
