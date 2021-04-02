package validation

import (
	"github.com/xeipuuv/gojsonschema"
)

type Error struct {
	Field       string                 `json:"field"`
	Type        string                 `json:"type"`
	Context     string                 `json:"context"`
	Code        string                 `json:"code"`
	Description string                 `json:"description"`
	Value       interface{}            `json:"value"`
	Details     map[string]interface{} `json:"details"`
}

func (e *Error) Validate(schemapath string, object []byte) (*Result, error) {
	gojsonschema.Locale = e

	schemaLoader := gojsonschema.NewReferenceLoader(schemapath)
	documentLoader := gojsonschema.NewBytesLoader(object)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	return &Result{result}, err
}

// False returns a format-string for "false" schema validation errors
func (e Error) False() string {
	return "False always fails validation"
}

// Required returns a format-string for "required" schema validation errors
func (e Error) Required() string {
	return `{{.property}} is required`
}

// InvalidType returns a format-string for "invalid type" schema validation errors
func (e Error) InvalidType() string {
	return `Invalid type. Expected: {{.expected}}, given: {{.given}}`
}

// NumberAnyOf returns a format-string for "anyOf" schema validation errors
func (e Error) NumberAnyOf() string {
	return `Must validate at least one schema (anyOf)`
}

// NumberOneOf returns a format-string for "oneOf" schema validation errors
func (e Error) NumberOneOf() string {
	return `Must validate one and only one schema (oneOf)`
}

// NumberAllOf returns a format-string for "allOf" schema validation errors
func (e Error) NumberAllOf() string {
	return `Must validate all the schemas (allOf)`
}

// NumberNot returns a format-string to format a NumberNotError
func (e Error) NumberNot() string {
	return `Must not validate the schema (not)`
}

// MissingDependency returns a format-string for "missing dependency" schema validation errors
func (e Error) MissingDependency() string {
	return `Has a dependency on {{.dependency}}`
}

// Internal returns a format-string for internal errors
func (e Error) Internal() string {
	return `Internae Error {{.error}}`
}

// Const returns a format-string to format a ConstError
func (e Error) Const() string {
	return `{{.field}} does not match: {{.allowed}}`
}

// Enum returns a format-string to format an EnumError
func (e Error) Enum() string {
	return `{{.field}} must be one of the following: {{.allowed}}`
}

// ArrayNoAdditionalItems returns a format-string to format an ArrayNoAdditionalItemsError
func (e Error) ArrayNoAdditionalItems() string {
	return `No additional items allowed on array`
}

// ArrayNotEnoughItems returns a format-string to format an error for arrays having not enough items to match positional list of schema
func (e Error) ArrayNotEnoughItems() string {
	return `Not enough items on array to match positional list of schema`
}

// ArrayMinItems returns a format-string to format an ArrayMinItemsError
func (e Error) ArrayMinItems() string {
	return `Array must have at least {{.min}} items`
}

// ArrayMaxItems returns a format-string to format an ArrayMaxItemsError
func (e Error) ArrayMaxItems() string {
	return `Array must have at most {{.max}} items`
}

// Unique returns a format-string  to format an ItemsMustBeUniqueError
func (e Error) Unique() string {
	return `{{.type}} items[{{.i}},{{.j}}] must be unique`
}

// ArrayContains returns a format-string to format an ArrayContainsError
func (e Error) ArrayContains() string {
	return `At least one of the items must match`
}

// ArrayMinProperties returns a format-string to format an ArrayMinPropertiesError
func (e Error) ArrayMinProperties() string {
	return `Must have at least {{.min}} properties`
}

// ArrayMaxProperties returns a format-string to format an ArrayMaxPropertiesError
func (e Error) ArrayMaxProperties() string {
	return `Must have at most {{.max}} properties`
}

// AdditionalPropertyNotAllowed returns a format-string to format an AdditionalPropertyNotAllowedError
func (e Error) AdditionalPropertyNotAllowed() string {
	return `Additional property {{.property}} is not allowed`
}

// InvalidPropertyPattern returns a format-string to format an InvalidPropertyPatternError
func (e Error) InvalidPropertyPattern() string {
	return `Property "{{.property}}" does not match pattern {{.pattern}}`
}

// InvalidPropertyName returns a format-string to format an InvalidPropertyNameError
func (e Error) InvalidPropertyName() string {
	return `Property name of "{{.property}}" does not match`
}

// StringGTE returns a format-string to format an StringLengthGTEError
func (e Error) StringGTE() string {
	return `String length must be greater than or equal to {{.min}}`
}

// StringLTE returns a format-string to format an StringLengthLTEError
func (e Error) StringLTE() string {
	return `String length must be less than or equal to {{.max}}`
}

// DoesNotMatchPattern returns a format-string to format an DoesNotMatchPatternError
func (e Error) DoesNotMatchPattern() string {
	return `Does not match pattern '{{.pattern}}'`
}

// DoesNotMatchFormat returns a format-string to format an DoesNotMatchFormatError
func (e Error) DoesNotMatchFormat() string {
	return `Does not match format '{{.format}}'`
}

// MultipleOf returns a format-string to format an MultipleOfError
func (e Error) MultipleOf() string {
	return `Must be a multiple of {{.multiple}}`
}

// NumberGTE returns the format string to format a NumberGTEError
func (e Error) NumberGTE() string {
	return `Must be greater than or equal to {{.min}}`
}

// NumberGT returns the format string to format a NumberGTError
func (e Error) NumberGT() string {
	return `Must be greater than {{.min}}`
}

// NumberLTE returns the format string to format a NumberLTEError
func (e Error) NumberLTE() string {
	return `Must be less than or equal to {{.max}}`
}

// NumberLT returns the format string to format a NumberLTError
func (e Error) NumberLT() string {
	return `Must be less than {{.max}}`
}

// Schema validators

// RegexPattern returns a format-string to format a regex-pattern error
func (e Error) RegexPattern() string {
	return `Invalid regex pattern '{{.pattern}}'`
}

// GreaterThanZero returns a format-string to format an error where a number must be greater than zero
func (e Error) GreaterThanZero() string {
	return `{{.number}} must be strictly greater than 0`
}

// MustBeOfA returns a format-string to format an error where a value is of the wrong type
func (e Error) MustBeOfA() string {
	return `{{.x}} must be of a {{.y}}`
}

// MustBeOfAn returns a format-string to format an error where a value is of the wrong type
func (e Error) MustBeOfAn() string {
	return `{{.x}} must be of an {{.y}}`
}

// CannotBeUsedWithout returns a format-string to format a "cannot be used without" error
func (e Error) CannotBeUsedWithout() string {
	return `{{.x}} cannot be used without {{.y}}`
}

// CannotBeGT returns a format-string to format an error where a value are greater than allowed
func (e Error) CannotBeGT() string {
	return `{{.x}} cannot be greater than {{.y}}`
}

// MustBeOfType returns a format-string to format an error where a value does not match the required type
func (e Error) MustBeOfType() string {
	return `{{.key}} must be of type {{.type}}`
}

// MustBeValidRegex returns a format-string to format an error where a regex is invalid
func (e Error) MustBeValidRegex() string {
	return `{{.key}} must be a valid regex`
}

// MustBeValidFormat returns a format-string to format an error where a value does not match the expected format
func (e Error) MustBeValidFormat() string {
	return `{{.key}} must be a valid format {{.given}}`
}

// MustBeGTEZero returns a format-string to format an error where a value must be greater or equal than 0
func (e Error) MustBeGTEZero() string {
	return `{{.key}} must be greater than or equal to 0`
}

// KeyCannotBeGreaterThan returns a format-string to format an error where a value is greater than the maximum  allowed
func (e Error) KeyCannotBeGreaterThan() string {
	return `{{.key}} cannot be greater than {{.y}}`
}

// KeyItemsMustBeOfType returns a format-string to format an error where a key is of the wrong type
func (e Error) KeyItemsMustBeOfType() string {
	return `{{.key}} items must be {{.type}}`
}

// KeyItemsMustBeUnique returns a format-string to format an error where keys are not unique
func (e Error) KeyItemsMustBeUnique() string {
	return `{{.key}} items must be unique`
}

// ReferenceMustBeCanonical returns a format-string to format a "reference must be canonical" error
func (e Error) ReferenceMustBeCanonical() string {
	return `Reference {{.reference}} must be canonical`
}

// NotAValidType returns a format-string to format an invalid type error
func (e Error) NotAValidType() string {
	return `has a primitive type that is NOT VALID -- given: {{.given}} Expected valid values are:{{.expected}}`
}

// Duplicated returns a format-string to format an error where types are duplicated
func (e Error) Duplicated() string {
	return `{{.type}} type is duplicated`
}

// HttpBadStatus returns a format-string for errors when loading a schema using HTTP
func (e Error) HttpBadStatus() string {
	return `Could not read schema from HTTP, response status is {{.status}}`
}

// ErrorFormat returns a format string for errors
// Replacement options: field, description, context, value
func (e Error) ErrorFormat() string {
	return `{{.field}}: {{.description}}`
}

// ParseError returns a format-string for JSON parsing errors
func (e Error) ParseError() string {
	return `Expected: {{.expected}}, given: Invalid JSON`
}

// ConditionThen returns a format-string for ConditionThenError errors
// If/Else
func (e Error) ConditionThen() string {
	return `Must validate "then" as "if" was valid`
}

// ConditionElse returns a format-string for ConditionElseError errors
func (e Error) ConditionElse() string {
	return `Must validate "else" as "if" was not valid`
}

// constants
const (
	STRING_NUMBER                     = "number"
	STRING_ARRAY_OF_STRINGS           = "array of strings"
	STRING_ARRAY_OF_SCHEMAS           = "array of schemas"
	STRING_SCHEMA                     = "valid schema"
	STRING_SCHEMA_OR_ARRAY_OF_STRINGS = "schema or array of strings"
	STRING_PROPERTIES                 = "properties"
	STRING_DEPENDENCY                 = "dependency"
	STRING_PROPERTY                   = "property"
	STRING_UNDEFINED                  = "undefined"
	STRING_CONTEXT_ROOT               = "(root)"
	STRING_ROOT_SCHEMA_PROPERTY       = "(root)"
)
