package pgcode

// Constants for every PostgreSQL SQLSTATE code.
// When multiple codes share a condition name, the primary code keeps the clean name
// and aliases are suffixed with their class: e.g. ModifyingSQLDataNotPermittedClass38.
const (
	// Class 00 — Successful Completion
	SuccessfulCompletion = "00000"

	// Class 01 — Warning
	Warning                          = "01000"
	NullValueEliminatedInSetFunction = "01003"
	StringDataRightTruncation        = "01004"
	PrivilegeNotRevoked              = "01006"
	PrivilegeNotGranted              = "01007"
	ImplicitZeroBitPadding           = "01008"
	DynamicResultSetsReturned        = "0100C"
	DeprecatedFeature                = "01P01"

	// Class 02 — No Data (this is also a warning class per the SQL standard)
	NoData                                = "02000"
	NoAdditionalDynamicResultSetsReturned = "02001"

	// Class 03 — SQL Statement Not Yet Complete
	SQLStatementNotYetComplete = "03000"

	// Class 08 — Connection Exception
	ConnectionException                           = "08000"
	SqlclientUnableToEstablishSqlconnection       = "08001"
	ConnectionDoesNotExist                        = "08003"
	SqlserverRejectedEstablishmentOfSqlconnection = "08004"
	ConnectionFailure                             = "08006"
	TransactionResolutionUnknown                  = "08007"
	ProtocolViolation                             = "08P01"

	// Class 09 — Triggered Action Exception
	TriggeredActionException = "09000"

	// Class 0A — Feature Not Supported
	FeatureNotSupported = "0A000"

	// Class 0B — Invalid Transaction Initiation
	InvalidTransactionInitiation = "0B000"

	// Class 0F — Locator Exception
	LocatorException            = "0F000"
	InvalidLocatorSpecification = "0F001"

	// Class 0L — Invalid Grantor
	InvalidGrantor        = "0L000"
	InvalidGrantOperation = "0LP01"

	// Class 0P — Invalid Role Specification
	InvalidRoleSpecification = "0P000"

	// Class 0Z — Diagnostics Exception
	DiagnosticsException                           = "0Z000"
	StackedDiagnosticsAccessedWithoutActiveHandler = "0Z002"

	// Class 20 — Case Not Found
	CaseNotFound = "20000"

	// Class 21 — Cardinality Violation
	CardinalityViolation = "21000"

	// Class 22 — Data Exception
	DataException                             = "22000"
	StringDataRightTruncationClass22          = "22001"
	NullValueNoIndicatorParameter             = "22002"
	NumericValueOutOfRange                    = "22003"
	NullValueNotAllowed                       = "22004"
	ErrorInAssignment                         = "22005"
	InvalidDatetimeFormat                     = "22007"
	DatetimeFieldOverflow                     = "22008"
	InvalidTimeZoneDisplacementValue          = "22009"
	EscapeCharacterConflict                   = "2200B"
	InvalidUseOfEscapeCharacter               = "2200C"
	InvalidEscapeOctet                        = "2200D"
	ZeroLengthCharacterString                 = "2200F"
	MostSpecificTypeMismatch                  = "2200G"
	SequenceGeneratorLimitExceeded            = "2200H"
	NotAnXMLDocument                          = "2200L"
	InvalidXMLDocument                        = "2200M"
	InvalidXMLContent                         = "2200N"
	InvalidXMLComment                         = "2200S"
	InvalidXMLProcessingInstruction           = "2200T"
	InvalidIndicatorParameterValue            = "22010"
	SubstringError                            = "22011"
	DivisionByZero                            = "22012"
	InvalidPrecedingOrFollowingSize           = "22013"
	InvalidArgumentForNtileFunction           = "22014"
	IntervalFieldOverflow                     = "22015"
	InvalidArgumentForNthValueFunction        = "22016"
	InvalidCharacterValueForCast              = "22018"
	InvalidEscapeCharacter                    = "22019"
	InvalidRegularExpression                  = "2201B"
	InvalidArgumentForLogarithm               = "2201E"
	InvalidArgumentForPowerFunction           = "2201F"
	InvalidArgumentForWidthBucketFunction     = "2201G"
	InvalidRowCountInLimitClause              = "2201W"
	InvalidRowCountInResultOffsetClause       = "2201X"
	CharacterNotInRepertoire                  = "22021"
	IndicatorOverflow                         = "22022"
	InvalidParameterValue                     = "22023"
	UnterminatedCString                       = "22024"
	InvalidEscapeSequence                     = "22025"
	StringDataLengthMismatch                  = "22026"
	TrimError                                 = "22027"
	ArraySubscriptError                       = "2202E"
	InvalidTablesampleRepeat                  = "2202G"
	InvalidTablesampleArgument                = "2202H"
	DuplicateJSONObjectKeyValue               = "22030"
	InvalidArgumentForSQLJSONDatetimeFunction = "22031"
	InvalidJSONText                           = "22032"
	InvalidSQLJSONSubscript                   = "22033"
	MoreThanOneSQLJSONItem                    = "22034"
	NoSQLJSONItem                             = "22035"
	NonNumericSQLJSONItem                     = "22036"
	NonUniqueKeysInAJSONObject                = "22037"
	SingletonSQLJSONItemRequired              = "22038"
	SQLJSONArrayNotFound                      = "22039"
	SQLJSONMemberNotFound                     = "2203A"
	SQLJSONNumberNotFound                     = "2203B"
	SQLJSONObjectNotFound                     = "2203C"
	TooManyJSONArrayElements                  = "2203D"
	TooManyJSONObjectMembers                  = "2203E"
	SQLJSONScalarRequired                     = "2203F"
	SQLJSONItemCannotBeCastToTargetType       = "2203G" // since PG15
	FloatingPointException                    = "22P01"
	InvalidTextRepresentation                 = "22P02"
	InvalidBinaryRepresentation               = "22P03"
	BadCopyFileFormat                         = "22P04"
	UntranslatableCharacter                   = "22P05"
	NonstandardUseOfEscapeCharacter           = "22P06"

	// Class 23 — Integrity Constraint Violation
	IntegrityConstraintViolation = "23000"
	RestrictViolation            = "23001"
	NotNullViolation             = "23502"
	ForeignKeyViolation          = "23503"
	UniqueViolation              = "23505"
	CheckViolation               = "23514"
	ExclusionViolation           = "23P01"

	// Class 24 — Invalid Cursor State
	InvalidCursorState = "24000"

	// Class 25 — Invalid Transaction State
	InvalidTransactionState                         = "25000"
	ActiveSQLTransaction                            = "25001"
	BranchTransactionAlreadyActive                  = "25002"
	InappropriateAccessModeForBranchTransaction     = "25003"
	InappropriateIsolationLevelForBranchTransaction = "25004"
	NoActiveSQLTransactionForBranchTransaction      = "25005"
	ReadOnlySQLTransaction                          = "25006"
	SchemaAndDataStatementMixingNotSupported        = "25007"
	HeldCursorRequiresSameIsolationLevel            = "25008"
	NoActiveSQLTransaction                          = "25P01"
	InFailedSQLTransaction                          = "25P02"
	IdleInTransactionSessionTimeout                 = "25P03"
	TransactionTimeout                              = "25P04" // since PG17

	// Class 26 — Invalid SQL Statement Name
	InvalidSQLStatementName = "26000"

	// Class 27 — Triggered Data Change Violation
	TriggeredDataChangeViolation = "27000"

	// Class 28 — Invalid Authorization Specification
	InvalidAuthorizationSpecification = "28000"
	InvalidPassword                   = "28P01"

	// Class 2B — Dependent Privilege Descriptors Still Exist
	DependentPrivilegeDescriptorsStillExist = "2B000"
	DependentObjectsStillExist              = "2BP01"

	// Class 2D — Invalid Transaction Termination
	InvalidTransactionTermination = "2D000"

	// Class 2F — SQL Routine Exception
	SQLRoutineException               = "2F000"
	ModifyingSQLDataNotPermitted      = "2F002"
	ProhibitedSQLStatementAttempted   = "2F003"
	ReadingSQLDataNotPermitted        = "2F004"
	FunctionExecutedNoReturnStatement = "2F005"

	// Class 34 — Invalid Cursor Name
	InvalidCursorName = "34000"

	// Class 38 — External Routine Exception
	ExternalRoutineException               = "38000"
	ContainingSQLNotPermitted              = "38001"
	ModifyingSQLDataNotPermittedClass38    = "38002"
	ProhibitedSQLStatementAttemptedClass38 = "38003"
	ReadingSQLDataNotPermittedClass38      = "38004"

	// Class 39 — External Routine Invocation Exception
	ExternalRoutineInvocationException = "39000"
	InvalidSQLStateReturned            = "39001"
	NullValueNotAllowedClass39         = "39004"
	TriggerProtocolViolated            = "39P01"
	SRFProtocolViolated                = "39P02"
	EventTriggerProtocolViolated       = "39P03"

	// Class 3B — Savepoint Exception
	SavepointException            = "3B000"
	InvalidSavepointSpecification = "3B001"

	// Class 3D — Invalid Catalog Name
	InvalidCatalogName = "3D000"

	// Class 3F — Invalid Schema Name
	InvalidSchemaName = "3F000"

	// Class 40 — Transaction Rollback
	TransactionRollback                     = "40000"
	SerializationFailure                    = "40001"
	TransactionIntegrityConstraintViolation = "40002"
	StatementCompletionUnknown              = "40003"
	DeadlockDetected                        = "40P01"

	// Class 42 — Syntax Error or Access Rule Violation
	SyntaxErrorOrAccessRuleViolation   = "42000"
	InsufficientPrivilege              = "42501"
	SyntaxError                        = "42601"
	InvalidName                        = "42602"
	InvalidColumnDefinition            = "42611"
	NameTooLong                        = "42622"
	DuplicateColumn                    = "42701"
	AmbiguousColumn                    = "42702"
	UndefinedColumn                    = "42703"
	UndefinedObject                    = "42704"
	DuplicateObject                    = "42710"
	DuplicateAlias                     = "42712"
	DuplicateFunction                  = "42723"
	AmbiguousFunction                  = "42725"
	GroupingError                      = "42803"
	DatatypeMismatch                   = "42804"
	WrongObjectType                    = "42809"
	InvalidForeignKey                  = "42830"
	CannotCoerce                       = "42846"
	UndefinedFunction                  = "42883"
	GeneratedAlways                    = "428C9"
	ReservedName                       = "42939"
	UndefinedTable                     = "42P01"
	UndefinedParameter                 = "42P02"
	DuplicateCursor                    = "42P03"
	DuplicateDatabase                  = "42P04"
	DuplicatePreparedStatement         = "42P05"
	DuplicateSchema                    = "42P06"
	DuplicateTable                     = "42P07"
	AmbiguousParameter                 = "42P08"
	AmbiguousAlias                     = "42P09"
	InvalidColumnReference             = "42P10"
	InvalidCursorDefinition            = "42P11"
	InvalidDatabaseDefinition          = "42P12"
	InvalidFunctionDefinition          = "42P13"
	InvalidPreparedStatementDefinition = "42P14"
	InvalidSchemaDefinition            = "42P15"
	InvalidTableDefinition             = "42P16"
	InvalidObjectDefinition            = "42P17"
	IndeterminateDatatype              = "42P18"
	InvalidRecursion                   = "42P19"
	WindowingError                     = "42P20"
	CollationMismatch                  = "42P21"
	IndeterminateCollation             = "42P22"

	// Class 44 — WITH CHECK OPTION Violation
	WithCheckOptionViolation = "44000"

	// Class 53 — Insufficient Resources
	InsufficientResources      = "53000"
	DiskFull                   = "53100"
	OutOfMemory                = "53200"
	TooManyConnections         = "53300"
	ConfigurationLimitExceeded = "53400"

	// Class 54 — Program Limit Exceeded
	ProgramLimitExceeded = "54000"
	StatementTooComplex  = "54001"
	TooManyColumns       = "54011"
	TooManyArguments     = "54023"

	// Class 55 — Object Not In Prerequisite State
	ObjectNotInPrerequisiteState = "55000"
	ObjectInUse                  = "55006"
	CantChangeRuntimeParam       = "55P02"
	LockNotAvailable             = "55P03"
	UnsafeNewEnumValueUsage      = "55P04"

	// Class 57 — Operator Intervention
	OperatorIntervention = "57000"
	QueryCanceled        = "57014"
	AdminShutdown        = "57P01"
	CrashShutdown        = "57P02"
	CannotConnectNow     = "57P03"
	DatabaseDropped      = "57P04"
	IdleSessionTimeout   = "57P05"

	// Class 58 — System Error (errors external to PostgreSQL itself)
	SystemError     = "58000"
	IOError         = "58030"
	UndefinedFile   = "58P01"
	DuplicateFile   = "58P02"
	FileNameTooLong = "58P03" // since PG18

	// Class 72 — Snapshot Failure
	SnapshotTooOld = "72000"

	// Class F0 — Configuration File Error
	ConfigFileError = "F0000"
	LockFileExists  = "F0001"

	// Class HV — Foreign Data Wrapper Error (SQL/MED)
	FDWError                             = "HV000"
	FDWOutOfMemory                       = "HV001"
	FDWDynamicParameterValueNeeded       = "HV002"
	FDWInvalidDataType                   = "HV004"
	FDWColumnNameNotFound                = "HV005"
	FDWInvalidDataTypeDescriptors        = "HV006"
	FDWInvalidColumnName                 = "HV007"
	FDWInvalidColumnNumber               = "HV008"
	FDWInvalidUseOfNullPointer           = "HV009"
	FDWInvalidStringFormat               = "HV00A"
	FDWInvalidHandle                     = "HV00B"
	FDWInvalidOptionIndex                = "HV00C"
	FDWInvalidOptionName                 = "HV00D"
	FDWOptionNameNotFound                = "HV00J"
	FDWReplyHandle                       = "HV00K"
	FDWUnableToCreateExecution           = "HV00L"
	FDWUnableToCreateReply               = "HV00M"
	FDWUnableToEstablishConnection       = "HV00N"
	FDWNoSchemas                         = "HV00P"
	FDWSchemaNotFound                    = "HV00Q"
	FDWTableNotFound                     = "HV00R"
	FDWFunctionSequenceError             = "HV010"
	FDWTooManyHandles                    = "HV014"
	FDWInconsistentDescriptorInformation = "HV021"
	FDWInvalidAttributeValue             = "HV024"
	FDWInvalidStringLengthOrBufferLength = "HV090"
	FDWInvalidDescriptorFieldIdentifier  = "HV091"

	// Class P0 — PL/pgSQL Error
	PLpgSQLError   = "P0000"
	RaiseException = "P0001"
	NoDataFound    = "P0002"
	TooManyRows    = "P0003"
	AssertFailure  = "P0004"

	// Class XX — Internal Error
	InternalError  = "XX000"
	DataCorrupted  = "XX001"
	IndexCorrupted = "XX002"

	// Class 10 — XQuery Error
	InvalidArgumentForXquery = "10608" // since PG18
)

// IsSuccessfulCompletion returns true if code belongs to
// Class 00 — Successful Completion.
func IsSuccessfulCompletion(code string) bool {
	switch code {
	case SuccessfulCompletion:
		return true
	}
	return false
}

// IsWarning returns true if code belongs to
// Class 01 — Warning.
func IsWarning(code string) bool {
	switch code {
	case Warning, NullValueEliminatedInSetFunction, StringDataRightTruncation, PrivilegeNotRevoked, PrivilegeNotGranted, ImplicitZeroBitPadding, DynamicResultSetsReturned, DeprecatedFeature:
		return true
	}
	return false
}

// IsNoData returns true if code belongs to
// Class 02 — No Data (this is also a warning class per the SQL standard).
func IsNoData(code string) bool {
	switch code {
	case NoData, NoAdditionalDynamicResultSetsReturned:
		return true
	}
	return false
}

// IsSqlStatementNotYetComplete returns true if code belongs to
// Class 03 — SQL Statement Not Yet Complete.
func IsSqlStatementNotYetComplete(code string) bool {
	switch code {
	case SQLStatementNotYetComplete:
		return true
	}
	return false
}

// IsConnectionException returns true if code belongs to
// Class 08 — Connection Exception.
func IsConnectionException(code string) bool {
	switch code {
	case ConnectionException, SqlclientUnableToEstablishSqlconnection, ConnectionDoesNotExist, SqlserverRejectedEstablishmentOfSqlconnection, ConnectionFailure, TransactionResolutionUnknown, ProtocolViolation:
		return true
	}
	return false
}

// IsTriggeredActionException returns true if code belongs to
// Class 09 — Triggered Action Exception.
func IsTriggeredActionException(code string) bool {
	switch code {
	case TriggeredActionException:
		return true
	}
	return false
}

// IsFeatureNotSupported returns true if code belongs to
// Class 0A — Feature Not Supported.
func IsFeatureNotSupported(code string) bool {
	switch code {
	case FeatureNotSupported:
		return true
	}
	return false
}

// IsInvalidTransactionInitiation returns true if code belongs to
// Class 0B — Invalid Transaction Initiation.
func IsInvalidTransactionInitiation(code string) bool {
	switch code {
	case InvalidTransactionInitiation:
		return true
	}
	return false
}

// IsLocatorException returns true if code belongs to
// Class 0F — Locator Exception.
func IsLocatorException(code string) bool {
	switch code {
	case LocatorException, InvalidLocatorSpecification:
		return true
	}
	return false
}

// IsInvalidGrantor returns true if code belongs to
// Class 0L — Invalid Grantor.
func IsInvalidGrantor(code string) bool {
	switch code {
	case InvalidGrantor, InvalidGrantOperation:
		return true
	}
	return false
}

// IsInvalidRoleSpecification returns true if code belongs to
// Class 0P — Invalid Role Specification.
func IsInvalidRoleSpecification(code string) bool {
	switch code {
	case InvalidRoleSpecification:
		return true
	}
	return false
}

// IsDiagnosticsException returns true if code belongs to
// Class 0Z — Diagnostics Exception.
func IsDiagnosticsException(code string) bool {
	switch code {
	case DiagnosticsException, StackedDiagnosticsAccessedWithoutActiveHandler:
		return true
	}
	return false
}

// IsCaseNotFound returns true if code belongs to
// Class 20 — Case Not Found.
func IsCaseNotFound(code string) bool {
	switch code {
	case CaseNotFound:
		return true
	}
	return false
}

// IsCardinalityViolation returns true if code belongs to
// Class 21 — Cardinality Violation.
func IsCardinalityViolation(code string) bool {
	switch code {
	case CardinalityViolation:
		return true
	}
	return false
}

// IsDataException returns true if code belongs to
// Class 22 — Data Exception.
func IsDataException(code string) bool {
	switch code {
	case DataException, StringDataRightTruncationClass22, NullValueNoIndicatorParameter, NumericValueOutOfRange, NullValueNotAllowed, ErrorInAssignment, InvalidDatetimeFormat, DatetimeFieldOverflow, InvalidTimeZoneDisplacementValue, EscapeCharacterConflict, InvalidUseOfEscapeCharacter, InvalidEscapeOctet, ZeroLengthCharacterString, MostSpecificTypeMismatch, SequenceGeneratorLimitExceeded, NotAnXMLDocument, InvalidXMLDocument, InvalidXMLContent, InvalidXMLComment, InvalidXMLProcessingInstruction, InvalidIndicatorParameterValue, SubstringError, DivisionByZero, InvalidPrecedingOrFollowingSize, InvalidArgumentForNtileFunction, IntervalFieldOverflow, InvalidArgumentForNthValueFunction, InvalidCharacterValueForCast, InvalidEscapeCharacter, InvalidRegularExpression, InvalidArgumentForLogarithm, InvalidArgumentForPowerFunction, InvalidArgumentForWidthBucketFunction, InvalidRowCountInLimitClause, InvalidRowCountInResultOffsetClause, CharacterNotInRepertoire, IndicatorOverflow, InvalidParameterValue, UnterminatedCString, InvalidEscapeSequence, StringDataLengthMismatch, TrimError, ArraySubscriptError, InvalidTablesampleRepeat, InvalidTablesampleArgument, DuplicateJSONObjectKeyValue, InvalidArgumentForSQLJSONDatetimeFunction, InvalidJSONText, InvalidSQLJSONSubscript, MoreThanOneSQLJSONItem, NoSQLJSONItem, NonNumericSQLJSONItem, NonUniqueKeysInAJSONObject, SingletonSQLJSONItemRequired, SQLJSONArrayNotFound, SQLJSONMemberNotFound, SQLJSONNumberNotFound, SQLJSONObjectNotFound, TooManyJSONArrayElements, TooManyJSONObjectMembers, SQLJSONScalarRequired, SQLJSONItemCannotBeCastToTargetType, FloatingPointException, InvalidTextRepresentation, InvalidBinaryRepresentation, BadCopyFileFormat, UntranslatableCharacter, NonstandardUseOfEscapeCharacter:
		return true
	}
	return false
}

// IsIntegrityConstraintViolation returns true if code belongs to
// Class 23 — Integrity Constraint Violation.
func IsIntegrityConstraintViolation(code string) bool {
	switch code {
	case IntegrityConstraintViolation, RestrictViolation, NotNullViolation, ForeignKeyViolation, UniqueViolation, CheckViolation, ExclusionViolation:
		return true
	}
	return false
}

// IsInvalidCursorState returns true if code belongs to
// Class 24 — Invalid Cursor State.
func IsInvalidCursorState(code string) bool {
	switch code {
	case InvalidCursorState:
		return true
	}
	return false
}

// IsInvalidTransactionState returns true if code belongs to
// Class 25 — Invalid Transaction State.
func IsInvalidTransactionState(code string) bool {
	switch code {
	case InvalidTransactionState, ActiveSQLTransaction, BranchTransactionAlreadyActive, InappropriateAccessModeForBranchTransaction, InappropriateIsolationLevelForBranchTransaction, NoActiveSQLTransactionForBranchTransaction, ReadOnlySQLTransaction, SchemaAndDataStatementMixingNotSupported, HeldCursorRequiresSameIsolationLevel, NoActiveSQLTransaction, InFailedSQLTransaction, IdleInTransactionSessionTimeout, TransactionTimeout:
		return true
	}
	return false
}

// IsInvalidSqlStatementName returns true if code belongs to
// Class 26 — Invalid SQL Statement Name.
func IsInvalidSqlStatementName(code string) bool {
	switch code {
	case InvalidSQLStatementName:
		return true
	}
	return false
}

// IsTriggeredDataChangeViolation returns true if code belongs to
// Class 27 — Triggered Data Change Violation.
func IsTriggeredDataChangeViolation(code string) bool {
	switch code {
	case TriggeredDataChangeViolation:
		return true
	}
	return false
}

// IsInvalidAuthorizationSpecification returns true if code belongs to
// Class 28 — Invalid Authorization Specification.
func IsInvalidAuthorizationSpecification(code string) bool {
	switch code {
	case InvalidAuthorizationSpecification, InvalidPassword:
		return true
	}
	return false
}

// IsDependentPrivilegeDescriptorsStillExist returns true if code belongs to
// Class 2B — Dependent Privilege Descriptors Still Exist.
func IsDependentPrivilegeDescriptorsStillExist(code string) bool {
	switch code {
	case DependentPrivilegeDescriptorsStillExist, DependentObjectsStillExist:
		return true
	}
	return false
}

// IsInvalidTransactionTermination returns true if code belongs to
// Class 2D — Invalid Transaction Termination.
func IsInvalidTransactionTermination(code string) bool {
	switch code {
	case InvalidTransactionTermination:
		return true
	}
	return false
}

// IsSqlRoutineException returns true if code belongs to
// Class 2F — SQL Routine Exception.
func IsSqlRoutineException(code string) bool {
	switch code {
	case SQLRoutineException, ModifyingSQLDataNotPermitted, ProhibitedSQLStatementAttempted, ReadingSQLDataNotPermitted, FunctionExecutedNoReturnStatement:
		return true
	}
	return false
}

// IsInvalidCursorName returns true if code belongs to
// Class 34 — Invalid Cursor Name.
func IsInvalidCursorName(code string) bool {
	switch code {
	case InvalidCursorName:
		return true
	}
	return false
}

// IsExternalRoutineException returns true if code belongs to
// Class 38 — External Routine Exception.
func IsExternalRoutineException(code string) bool {
	switch code {
	case ExternalRoutineException, ContainingSQLNotPermitted, ModifyingSQLDataNotPermittedClass38, ProhibitedSQLStatementAttemptedClass38, ReadingSQLDataNotPermittedClass38:
		return true
	}
	return false
}

// IsExternalRoutineInvocationException returns true if code belongs to
// Class 39 — External Routine Invocation Exception.
func IsExternalRoutineInvocationException(code string) bool {
	switch code {
	case ExternalRoutineInvocationException, InvalidSQLStateReturned, NullValueNotAllowedClass39, TriggerProtocolViolated, SRFProtocolViolated, EventTriggerProtocolViolated:
		return true
	}
	return false
}

// IsSavepointException returns true if code belongs to
// Class 3B — Savepoint Exception.
func IsSavepointException(code string) bool {
	switch code {
	case SavepointException, InvalidSavepointSpecification:
		return true
	}
	return false
}

// IsInvalidCatalogName returns true if code belongs to
// Class 3D — Invalid Catalog Name.
func IsInvalidCatalogName(code string) bool {
	switch code {
	case InvalidCatalogName:
		return true
	}
	return false
}

// IsInvalidSchemaName returns true if code belongs to
// Class 3F — Invalid Schema Name.
func IsInvalidSchemaName(code string) bool {
	switch code {
	case InvalidSchemaName:
		return true
	}
	return false
}

// IsTransactionRollback returns true if code belongs to
// Class 40 — Transaction Rollback.
func IsTransactionRollback(code string) bool {
	switch code {
	case TransactionRollback, SerializationFailure, TransactionIntegrityConstraintViolation, StatementCompletionUnknown, DeadlockDetected:
		return true
	}
	return false
}

// IsSyntaxErrorOrAccessRuleViolation returns true if code belongs to
// Class 42 — Syntax Error or Access Rule Violation.
func IsSyntaxErrorOrAccessRuleViolation(code string) bool {
	switch code {
	case SyntaxErrorOrAccessRuleViolation, InsufficientPrivilege, SyntaxError, InvalidName, InvalidColumnDefinition, NameTooLong, DuplicateColumn, AmbiguousColumn, UndefinedColumn, UndefinedObject, DuplicateObject, DuplicateAlias, DuplicateFunction, AmbiguousFunction, GroupingError, DatatypeMismatch, WrongObjectType, InvalidForeignKey, CannotCoerce, UndefinedFunction, GeneratedAlways, ReservedName, UndefinedTable, UndefinedParameter, DuplicateCursor, DuplicateDatabase, DuplicatePreparedStatement, DuplicateSchema, DuplicateTable, AmbiguousParameter, AmbiguousAlias, InvalidColumnReference, InvalidCursorDefinition, InvalidDatabaseDefinition, InvalidFunctionDefinition, InvalidPreparedStatementDefinition, InvalidSchemaDefinition, InvalidTableDefinition, InvalidObjectDefinition, IndeterminateDatatype, InvalidRecursion, WindowingError, CollationMismatch, IndeterminateCollation:
		return true
	}
	return false
}

// IsWithCheckOptionViolation returns true if code belongs to
// Class 44 — WITH CHECK OPTION Violation.
func IsWithCheckOptionViolation(code string) bool {
	switch code {
	case WithCheckOptionViolation:
		return true
	}
	return false
}

// IsInsufficientResources returns true if code belongs to
// Class 53 — Insufficient Resources.
func IsInsufficientResources(code string) bool {
	switch code {
	case InsufficientResources, DiskFull, OutOfMemory, TooManyConnections, ConfigurationLimitExceeded:
		return true
	}
	return false
}

// IsProgramLimitExceeded returns true if code belongs to
// Class 54 — Program Limit Exceeded.
func IsProgramLimitExceeded(code string) bool {
	switch code {
	case ProgramLimitExceeded, StatementTooComplex, TooManyColumns, TooManyArguments:
		return true
	}
	return false
}

// IsObjectNotInPrerequisiteState returns true if code belongs to
// Class 55 — Object Not In Prerequisite State.
func IsObjectNotInPrerequisiteState(code string) bool {
	switch code {
	case ObjectNotInPrerequisiteState, ObjectInUse, CantChangeRuntimeParam, LockNotAvailable, UnsafeNewEnumValueUsage:
		return true
	}
	return false
}

// IsOperatorIntervention returns true if code belongs to
// Class 57 — Operator Intervention.
func IsOperatorIntervention(code string) bool {
	switch code {
	case OperatorIntervention, QueryCanceled, AdminShutdown, CrashShutdown, CannotConnectNow, DatabaseDropped, IdleSessionTimeout:
		return true
	}
	return false
}

// IsSystemError returns true if code belongs to
// Class 58 — System Error (errors external to PostgreSQL itself).
func IsSystemError(code string) bool {
	switch code {
	case SystemError, IOError, UndefinedFile, DuplicateFile, FileNameTooLong:
		return true
	}
	return false
}

// IsSnapshotFailure returns true if code belongs to
// Class 72 — Snapshot Failure.
func IsSnapshotFailure(code string) bool {
	switch code {
	case SnapshotTooOld:
		return true
	}
	return false
}

// IsConfigurationFileError returns true if code belongs to
// Class F0 — Configuration File Error.
func IsConfigurationFileError(code string) bool {
	switch code {
	case ConfigFileError, LockFileExists:
		return true
	}
	return false
}

// IsForeignDataWrapperError returns true if code belongs to
// Class HV — Foreign Data Wrapper Error (SQL/MED).
func IsForeignDataWrapperError(code string) bool {
	switch code {
	case FDWError, FDWOutOfMemory, FDWDynamicParameterValueNeeded, FDWInvalidDataType, FDWColumnNameNotFound, FDWInvalidDataTypeDescriptors, FDWInvalidColumnName, FDWInvalidColumnNumber, FDWInvalidUseOfNullPointer, FDWInvalidStringFormat, FDWInvalidHandle, FDWInvalidOptionIndex, FDWInvalidOptionName, FDWOptionNameNotFound, FDWReplyHandle, FDWUnableToCreateExecution, FDWUnableToCreateReply, FDWUnableToEstablishConnection, FDWNoSchemas, FDWSchemaNotFound, FDWTableNotFound, FDWFunctionSequenceError, FDWTooManyHandles, FDWInconsistentDescriptorInformation, FDWInvalidAttributeValue, FDWInvalidStringLengthOrBufferLength, FDWInvalidDescriptorFieldIdentifier:
		return true
	}
	return false
}

// IsPl returns true if code belongs to
// Class P0 — PL/pgSQL Error.
func IsPl(code string) bool {
	switch code {
	case PLpgSQLError, RaiseException, NoDataFound, TooManyRows, AssertFailure:
		return true
	}
	return false
}

// IsInternalError returns true if code belongs to
// Class XX — Internal Error.
func IsInternalError(code string) bool {
	switch code {
	case InternalError, DataCorrupted, IndexCorrupted:
		return true
	}
	return false
}

// IsXqueryError returns true if code belongs to
// Class 10 — XQuery Error.
func IsXqueryError(code string) bool {
	switch code {
	case InvalidArgumentForXquery:
		return true
	}
	return false
}

// Name returns the constant name for a PostgreSQL error code,
// or an empty string if the code is not recognised.
func Name(code string) string {
	switch code {
	case SuccessfulCompletion:
		return "SuccessfulCompletion"
	case Warning:
		return "Warning"
	case NullValueEliminatedInSetFunction:
		return "NullValueEliminatedInSetFunction"
	case StringDataRightTruncation:
		return "StringDataRightTruncation"
	case PrivilegeNotRevoked:
		return "PrivilegeNotRevoked"
	case PrivilegeNotGranted:
		return "PrivilegeNotGranted"
	case ImplicitZeroBitPadding:
		return "ImplicitZeroBitPadding"
	case DynamicResultSetsReturned:
		return "DynamicResultSetsReturned"
	case DeprecatedFeature:
		return "DeprecatedFeature"
	case NoData:
		return "NoData"
	case NoAdditionalDynamicResultSetsReturned:
		return "NoAdditionalDynamicResultSetsReturned"
	case SQLStatementNotYetComplete:
		return "SQLStatementNotYetComplete"
	case ConnectionException:
		return "ConnectionException"
	case SqlclientUnableToEstablishSqlconnection:
		return "SqlclientUnableToEstablishSqlconnection"
	case ConnectionDoesNotExist:
		return "ConnectionDoesNotExist"
	case SqlserverRejectedEstablishmentOfSqlconnection:
		return "SqlserverRejectedEstablishmentOfSqlconnection"
	case ConnectionFailure:
		return "ConnectionFailure"
	case TransactionResolutionUnknown:
		return "TransactionResolutionUnknown"
	case ProtocolViolation:
		return "ProtocolViolation"
	case TriggeredActionException:
		return "TriggeredActionException"
	case FeatureNotSupported:
		return "FeatureNotSupported"
	case InvalidTransactionInitiation:
		return "InvalidTransactionInitiation"
	case LocatorException:
		return "LocatorException"
	case InvalidLocatorSpecification:
		return "InvalidLocatorSpecification"
	case InvalidGrantor:
		return "InvalidGrantor"
	case InvalidGrantOperation:
		return "InvalidGrantOperation"
	case InvalidRoleSpecification:
		return "InvalidRoleSpecification"
	case DiagnosticsException:
		return "DiagnosticsException"
	case StackedDiagnosticsAccessedWithoutActiveHandler:
		return "StackedDiagnosticsAccessedWithoutActiveHandler"
	case CaseNotFound:
		return "CaseNotFound"
	case CardinalityViolation:
		return "CardinalityViolation"
	case DataException:
		return "DataException"
	case StringDataRightTruncationClass22:
		return "StringDataRightTruncation"
	case NullValueNoIndicatorParameter:
		return "NullValueNoIndicatorParameter"
	case NumericValueOutOfRange:
		return "NumericValueOutOfRange"
	case NullValueNotAllowed:
		return "NullValueNotAllowed"
	case ErrorInAssignment:
		return "ErrorInAssignment"
	case InvalidDatetimeFormat:
		return "InvalidDatetimeFormat"
	case DatetimeFieldOverflow:
		return "DatetimeFieldOverflow"
	case InvalidTimeZoneDisplacementValue:
		return "InvalidTimeZoneDisplacementValue"
	case EscapeCharacterConflict:
		return "EscapeCharacterConflict"
	case InvalidUseOfEscapeCharacter:
		return "InvalidUseOfEscapeCharacter"
	case InvalidEscapeOctet:
		return "InvalidEscapeOctet"
	case ZeroLengthCharacterString:
		return "ZeroLengthCharacterString"
	case MostSpecificTypeMismatch:
		return "MostSpecificTypeMismatch"
	case SequenceGeneratorLimitExceeded:
		return "SequenceGeneratorLimitExceeded"
	case NotAnXMLDocument:
		return "NotAnXMLDocument"
	case InvalidXMLDocument:
		return "InvalidXMLDocument"
	case InvalidXMLContent:
		return "InvalidXMLContent"
	case InvalidXMLComment:
		return "InvalidXMLComment"
	case InvalidXMLProcessingInstruction:
		return "InvalidXMLProcessingInstruction"
	case InvalidIndicatorParameterValue:
		return "InvalidIndicatorParameterValue"
	case SubstringError:
		return "SubstringError"
	case DivisionByZero:
		return "DivisionByZero"
	case InvalidPrecedingOrFollowingSize:
		return "InvalidPrecedingOrFollowingSize"
	case InvalidArgumentForNtileFunction:
		return "InvalidArgumentForNtileFunction"
	case IntervalFieldOverflow:
		return "IntervalFieldOverflow"
	case InvalidArgumentForNthValueFunction:
		return "InvalidArgumentForNthValueFunction"
	case InvalidCharacterValueForCast:
		return "InvalidCharacterValueForCast"
	case InvalidEscapeCharacter:
		return "InvalidEscapeCharacter"
	case InvalidRegularExpression:
		return "InvalidRegularExpression"
	case InvalidArgumentForLogarithm:
		return "InvalidArgumentForLogarithm"
	case InvalidArgumentForPowerFunction:
		return "InvalidArgumentForPowerFunction"
	case InvalidArgumentForWidthBucketFunction:
		return "InvalidArgumentForWidthBucketFunction"
	case InvalidRowCountInLimitClause:
		return "InvalidRowCountInLimitClause"
	case InvalidRowCountInResultOffsetClause:
		return "InvalidRowCountInResultOffsetClause"
	case CharacterNotInRepertoire:
		return "CharacterNotInRepertoire"
	case IndicatorOverflow:
		return "IndicatorOverflow"
	case InvalidParameterValue:
		return "InvalidParameterValue"
	case UnterminatedCString:
		return "UnterminatedCString"
	case InvalidEscapeSequence:
		return "InvalidEscapeSequence"
	case StringDataLengthMismatch:
		return "StringDataLengthMismatch"
	case TrimError:
		return "TrimError"
	case ArraySubscriptError:
		return "ArraySubscriptError"
	case InvalidTablesampleRepeat:
		return "InvalidTablesampleRepeat"
	case InvalidTablesampleArgument:
		return "InvalidTablesampleArgument"
	case DuplicateJSONObjectKeyValue:
		return "DuplicateJSONObjectKeyValue"
	case InvalidArgumentForSQLJSONDatetimeFunction:
		return "InvalidArgumentForSQLJSONDatetimeFunction"
	case InvalidJSONText:
		return "InvalidJSONText"
	case InvalidSQLJSONSubscript:
		return "InvalidSQLJSONSubscript"
	case MoreThanOneSQLJSONItem:
		return "MoreThanOneSQLJSONItem"
	case NoSQLJSONItem:
		return "NoSQLJSONItem"
	case NonNumericSQLJSONItem:
		return "NonNumericSQLJSONItem"
	case NonUniqueKeysInAJSONObject:
		return "NonUniqueKeysInAJSONObject"
	case SingletonSQLJSONItemRequired:
		return "SingletonSQLJSONItemRequired"
	case SQLJSONArrayNotFound:
		return "SQLJSONArrayNotFound"
	case SQLJSONMemberNotFound:
		return "SQLJSONMemberNotFound"
	case SQLJSONNumberNotFound:
		return "SQLJSONNumberNotFound"
	case SQLJSONObjectNotFound:
		return "SQLJSONObjectNotFound"
	case TooManyJSONArrayElements:
		return "TooManyJSONArrayElements"
	case TooManyJSONObjectMembers:
		return "TooManyJSONObjectMembers"
	case SQLJSONScalarRequired:
		return "SQLJSONScalarRequired"
	case SQLJSONItemCannotBeCastToTargetType:
		return "SQLJSONItemCannotBeCastToTargetType"
	case FloatingPointException:
		return "FloatingPointException"
	case InvalidTextRepresentation:
		return "InvalidTextRepresentation"
	case InvalidBinaryRepresentation:
		return "InvalidBinaryRepresentation"
	case BadCopyFileFormat:
		return "BadCopyFileFormat"
	case UntranslatableCharacter:
		return "UntranslatableCharacter"
	case NonstandardUseOfEscapeCharacter:
		return "NonstandardUseOfEscapeCharacter"
	case IntegrityConstraintViolation:
		return "IntegrityConstraintViolation"
	case RestrictViolation:
		return "RestrictViolation"
	case NotNullViolation:
		return "NotNullViolation"
	case ForeignKeyViolation:
		return "ForeignKeyViolation"
	case UniqueViolation:
		return "UniqueViolation"
	case CheckViolation:
		return "CheckViolation"
	case ExclusionViolation:
		return "ExclusionViolation"
	case InvalidCursorState:
		return "InvalidCursorState"
	case InvalidTransactionState:
		return "InvalidTransactionState"
	case ActiveSQLTransaction:
		return "ActiveSQLTransaction"
	case BranchTransactionAlreadyActive:
		return "BranchTransactionAlreadyActive"
	case InappropriateAccessModeForBranchTransaction:
		return "InappropriateAccessModeForBranchTransaction"
	case InappropriateIsolationLevelForBranchTransaction:
		return "InappropriateIsolationLevelForBranchTransaction"
	case NoActiveSQLTransactionForBranchTransaction:
		return "NoActiveSQLTransactionForBranchTransaction"
	case ReadOnlySQLTransaction:
		return "ReadOnlySQLTransaction"
	case SchemaAndDataStatementMixingNotSupported:
		return "SchemaAndDataStatementMixingNotSupported"
	case HeldCursorRequiresSameIsolationLevel:
		return "HeldCursorRequiresSameIsolationLevel"
	case NoActiveSQLTransaction:
		return "NoActiveSQLTransaction"
	case InFailedSQLTransaction:
		return "InFailedSQLTransaction"
	case IdleInTransactionSessionTimeout:
		return "IdleInTransactionSessionTimeout"
	case TransactionTimeout:
		return "TransactionTimeout"
	case InvalidSQLStatementName:
		return "InvalidSQLStatementName"
	case TriggeredDataChangeViolation:
		return "TriggeredDataChangeViolation"
	case InvalidAuthorizationSpecification:
		return "InvalidAuthorizationSpecification"
	case InvalidPassword:
		return "InvalidPassword"
	case DependentPrivilegeDescriptorsStillExist:
		return "DependentPrivilegeDescriptorsStillExist"
	case DependentObjectsStillExist:
		return "DependentObjectsStillExist"
	case InvalidTransactionTermination:
		return "InvalidTransactionTermination"
	case SQLRoutineException:
		return "SQLRoutineException"
	case ModifyingSQLDataNotPermitted:
		return "ModifyingSQLDataNotPermitted"
	case ProhibitedSQLStatementAttempted:
		return "ProhibitedSQLStatementAttempted"
	case ReadingSQLDataNotPermitted:
		return "ReadingSQLDataNotPermitted"
	case FunctionExecutedNoReturnStatement:
		return "FunctionExecutedNoReturnStatement"
	case InvalidCursorName:
		return "InvalidCursorName"
	case ExternalRoutineException:
		return "ExternalRoutineException"
	case ContainingSQLNotPermitted:
		return "ContainingSQLNotPermitted"
	case ModifyingSQLDataNotPermittedClass38:
		return "ModifyingSQLDataNotPermitted"
	case ProhibitedSQLStatementAttemptedClass38:
		return "ProhibitedSQLStatementAttempted"
	case ReadingSQLDataNotPermittedClass38:
		return "ReadingSQLDataNotPermitted"
	case ExternalRoutineInvocationException:
		return "ExternalRoutineInvocationException"
	case InvalidSQLStateReturned:
		return "InvalidSQLStateReturned"
	case NullValueNotAllowedClass39:
		return "NullValueNotAllowed"
	case TriggerProtocolViolated:
		return "TriggerProtocolViolated"
	case SRFProtocolViolated:
		return "SRFProtocolViolated"
	case EventTriggerProtocolViolated:
		return "EventTriggerProtocolViolated"
	case SavepointException:
		return "SavepointException"
	case InvalidSavepointSpecification:
		return "InvalidSavepointSpecification"
	case InvalidCatalogName:
		return "InvalidCatalogName"
	case InvalidSchemaName:
		return "InvalidSchemaName"
	case TransactionRollback:
		return "TransactionRollback"
	case SerializationFailure:
		return "SerializationFailure"
	case TransactionIntegrityConstraintViolation:
		return "TransactionIntegrityConstraintViolation"
	case StatementCompletionUnknown:
		return "StatementCompletionUnknown"
	case DeadlockDetected:
		return "DeadlockDetected"
	case SyntaxErrorOrAccessRuleViolation:
		return "SyntaxErrorOrAccessRuleViolation"
	case InsufficientPrivilege:
		return "InsufficientPrivilege"
	case SyntaxError:
		return "SyntaxError"
	case InvalidName:
		return "InvalidName"
	case InvalidColumnDefinition:
		return "InvalidColumnDefinition"
	case NameTooLong:
		return "NameTooLong"
	case DuplicateColumn:
		return "DuplicateColumn"
	case AmbiguousColumn:
		return "AmbiguousColumn"
	case UndefinedColumn:
		return "UndefinedColumn"
	case UndefinedObject:
		return "UndefinedObject"
	case DuplicateObject:
		return "DuplicateObject"
	case DuplicateAlias:
		return "DuplicateAlias"
	case DuplicateFunction:
		return "DuplicateFunction"
	case AmbiguousFunction:
		return "AmbiguousFunction"
	case GroupingError:
		return "GroupingError"
	case DatatypeMismatch:
		return "DatatypeMismatch"
	case WrongObjectType:
		return "WrongObjectType"
	case InvalidForeignKey:
		return "InvalidForeignKey"
	case CannotCoerce:
		return "CannotCoerce"
	case UndefinedFunction:
		return "UndefinedFunction"
	case GeneratedAlways:
		return "GeneratedAlways"
	case ReservedName:
		return "ReservedName"
	case UndefinedTable:
		return "UndefinedTable"
	case UndefinedParameter:
		return "UndefinedParameter"
	case DuplicateCursor:
		return "DuplicateCursor"
	case DuplicateDatabase:
		return "DuplicateDatabase"
	case DuplicatePreparedStatement:
		return "DuplicatePreparedStatement"
	case DuplicateSchema:
		return "DuplicateSchema"
	case DuplicateTable:
		return "DuplicateTable"
	case AmbiguousParameter:
		return "AmbiguousParameter"
	case AmbiguousAlias:
		return "AmbiguousAlias"
	case InvalidColumnReference:
		return "InvalidColumnReference"
	case InvalidCursorDefinition:
		return "InvalidCursorDefinition"
	case InvalidDatabaseDefinition:
		return "InvalidDatabaseDefinition"
	case InvalidFunctionDefinition:
		return "InvalidFunctionDefinition"
	case InvalidPreparedStatementDefinition:
		return "InvalidPreparedStatementDefinition"
	case InvalidSchemaDefinition:
		return "InvalidSchemaDefinition"
	case InvalidTableDefinition:
		return "InvalidTableDefinition"
	case InvalidObjectDefinition:
		return "InvalidObjectDefinition"
	case IndeterminateDatatype:
		return "IndeterminateDatatype"
	case InvalidRecursion:
		return "InvalidRecursion"
	case WindowingError:
		return "WindowingError"
	case CollationMismatch:
		return "CollationMismatch"
	case IndeterminateCollation:
		return "IndeterminateCollation"
	case WithCheckOptionViolation:
		return "WithCheckOptionViolation"
	case InsufficientResources:
		return "InsufficientResources"
	case DiskFull:
		return "DiskFull"
	case OutOfMemory:
		return "OutOfMemory"
	case TooManyConnections:
		return "TooManyConnections"
	case ConfigurationLimitExceeded:
		return "ConfigurationLimitExceeded"
	case ProgramLimitExceeded:
		return "ProgramLimitExceeded"
	case StatementTooComplex:
		return "StatementTooComplex"
	case TooManyColumns:
		return "TooManyColumns"
	case TooManyArguments:
		return "TooManyArguments"
	case ObjectNotInPrerequisiteState:
		return "ObjectNotInPrerequisiteState"
	case ObjectInUse:
		return "ObjectInUse"
	case CantChangeRuntimeParam:
		return "CantChangeRuntimeParam"
	case LockNotAvailable:
		return "LockNotAvailable"
	case UnsafeNewEnumValueUsage:
		return "UnsafeNewEnumValueUsage"
	case OperatorIntervention:
		return "OperatorIntervention"
	case QueryCanceled:
		return "QueryCanceled"
	case AdminShutdown:
		return "AdminShutdown"
	case CrashShutdown:
		return "CrashShutdown"
	case CannotConnectNow:
		return "CannotConnectNow"
	case DatabaseDropped:
		return "DatabaseDropped"
	case IdleSessionTimeout:
		return "IdleSessionTimeout"
	case SystemError:
		return "SystemError"
	case IOError:
		return "IOError"
	case UndefinedFile:
		return "UndefinedFile"
	case DuplicateFile:
		return "DuplicateFile"
	case FileNameTooLong:
		return "FileNameTooLong"
	case SnapshotTooOld:
		return "SnapshotTooOld"
	case ConfigFileError:
		return "ConfigFileError"
	case LockFileExists:
		return "LockFileExists"
	case FDWError:
		return "FDWError"
	case FDWOutOfMemory:
		return "FDWOutOfMemory"
	case FDWDynamicParameterValueNeeded:
		return "FDWDynamicParameterValueNeeded"
	case FDWInvalidDataType:
		return "FDWInvalidDataType"
	case FDWColumnNameNotFound:
		return "FDWColumnNameNotFound"
	case FDWInvalidDataTypeDescriptors:
		return "FDWInvalidDataTypeDescriptors"
	case FDWInvalidColumnName:
		return "FDWInvalidColumnName"
	case FDWInvalidColumnNumber:
		return "FDWInvalidColumnNumber"
	case FDWInvalidUseOfNullPointer:
		return "FDWInvalidUseOfNullPointer"
	case FDWInvalidStringFormat:
		return "FDWInvalidStringFormat"
	case FDWInvalidHandle:
		return "FDWInvalidHandle"
	case FDWInvalidOptionIndex:
		return "FDWInvalidOptionIndex"
	case FDWInvalidOptionName:
		return "FDWInvalidOptionName"
	case FDWOptionNameNotFound:
		return "FDWOptionNameNotFound"
	case FDWReplyHandle:
		return "FDWReplyHandle"
	case FDWUnableToCreateExecution:
		return "FDWUnableToCreateExecution"
	case FDWUnableToCreateReply:
		return "FDWUnableToCreateReply"
	case FDWUnableToEstablishConnection:
		return "FDWUnableToEstablishConnection"
	case FDWNoSchemas:
		return "FDWNoSchemas"
	case FDWSchemaNotFound:
		return "FDWSchemaNotFound"
	case FDWTableNotFound:
		return "FDWTableNotFound"
	case FDWFunctionSequenceError:
		return "FDWFunctionSequenceError"
	case FDWTooManyHandles:
		return "FDWTooManyHandles"
	case FDWInconsistentDescriptorInformation:
		return "FDWInconsistentDescriptorInformation"
	case FDWInvalidAttributeValue:
		return "FDWInvalidAttributeValue"
	case FDWInvalidStringLengthOrBufferLength:
		return "FDWInvalidStringLengthOrBufferLength"
	case FDWInvalidDescriptorFieldIdentifier:
		return "FDWInvalidDescriptorFieldIdentifier"
	case PLpgSQLError:
		return "PLpgSQLError"
	case RaiseException:
		return "RaiseException"
	case NoDataFound:
		return "NoDataFound"
	case TooManyRows:
		return "TooManyRows"
	case AssertFailure:
		return "AssertFailure"
	case InternalError:
		return "InternalError"
	case DataCorrupted:
		return "DataCorrupted"
	case IndexCorrupted:
		return "IndexCorrupted"
	case InvalidArgumentForXquery:
		return "InvalidArgumentForXquery"
	}
	return ""
}
